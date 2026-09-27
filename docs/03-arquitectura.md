# Arquitectura: CashCastor

## 1. Visión General
Arquitectura de microservicio monolítico modular, desplegada en Kubernetes. El
backend en Go expone una API REST y un canal WebSocket. El frontend es una PWA
estática servida por Nginx. PostgreSQL es la fuente de verdad; Redis maneja
sesiones, caché y pub/sub para WebSockets.

    +-------------+     +-------------+     +-------------+
    |  Movil 1    |     |  Movil 2    |     |  Movil N    |
    |  (PWA)      |     |  (PWA)      |     |  (PWA)      |
    +------+------+     +------+------+     +------+------+
           |                   |                   |
           +-------------------+-------------------+
                               | HTTPS / WSS
                        +------v------+
                        |   Nginx     |
                        |  (Ingress)  |
                        +------+------+
                               |
                  +------------+------------+
                  |                         |
           +------v------+           +------v------+
           |  Frontend   |           |  Backend Go |
           |  (Estatico) |           |   (Gin)     |
           +-------------+           +------+------+
                                            |
                               +------------+------------+
                               |                         |
                        +------v------+           +------v------+
                        | PostgreSQL  |           |    Redis    |
                        | (Persist.)  |           | (Cache/WS)  |
                        +-------------+           +-------------+

## 2. Componentes

### 2.1 Backend (Go + Gin)
- Capa de API: Handlers REST, validación de entrada, middleware de
  autenticación (JWT), rate limiting.
- Capa de Servicios: Lógica de negocio: transferencias, préstamos,
  liquidación, cálculo de intereses, arqueo.
- Capa de Repositorios: Acceso a PostgreSQL vía GORM (ADR-008), Redis para
  idempotencia y pub/sub.
- Capa de WebSocket: Hub de conexiones por partida, broadcast de eventos
  (saldo actualizado, transacción, arqueo).

### 2.2 Frontend (React + Vite PWA)
- PWA: Service Worker para caché de assets estáticos y tolerancia a
  conexiones lentas.
- Estado: Zustand o Redux Toolkit para estado global de la partida.
- WebSocket: Cliente nativo WebSocket para actualizaciones en vivo.
- UI: Mobile-first, responsive, con componentes de billetera, transferencia,
  historial y panel de banco.

### 2.3 Base de Datos
- PostgreSQL: Tablas normalizadas para usuarios, partidas, jugadores,
  transacciones, préstamos, auditoría.
- Redis:
  - Sesiones de usuario (JWT blacklist/whitelist).
  - Idempotencia de transacciones (clave a resultado).
  - Pub/Sub para WebSocket (canal por partida).
  - Caché de saldos y arqueos (invalidación por evento).

### 2.4 Infraestructura (Kubernetes)

**Actualizado 2026-09-27** con lo que realmente aprovisionó `single-iac`
(rama `add-cashcastor-namespace`, pendiente de merge/apply — ver ADR-009).

- Namespace: `cashcastor` (agregado a `var.apps` en `01-cluster-resources`
  de `single-iac`: da namespace + quota + NetworkPolicy gratis, mismo
  patrón que el resto de apps de ese repo).
- Deployments/Services: `api` (Go) y `web` (Nginx + estáticos) — nombres
  cortos porque el namespace ya da el scope; es el patrón real de
  `single-iac`, no `cashcastor-api`/`cashcastor-web` como se bocetó
  originalmente.
- DNS: dos hosts (`cashcastor` y `cashcastor-api`), cada uno con su
  Ingress + TLS (cert-manager).
- Postgres y Redis: propios del namespace (no un patrón compartido entre
  apps), mismo criterio que otras apps de `single-iac`.
- ConfigMaps/Secrets: configuración de la app y credenciales de DB (ver
  `docs/architecture/secrets.md` en `single-iac`).
- CI/CD: **cada app construye, publica y despliega su propia imagen**
  (ADR-009); `single-iac` solo aplica los manifiestos (workflow genérico,
  ya cubre `cashcastor` sin cambios). El workflow de build+push+deploy
  vive en este repo, no en `single-iac`.
- Sin separación de infraestructura por entorno (dev/qa/prod): un único
  namespace `cashcastor`, mapeado a la rama `develop`. Decisión explícita
  de `single-iac` (no toda la organización usa infra separada por entorno,
  ADR-0002 de ese repo) y de este proyecto (triplicar infra ahora sería
  especulativo, sin lógica de negocio real todavía) — se revisa cuando
  haya necesidad real.

## 3. Decisiones de Diseño Clave

| Decisión | Razón |
|----------|-------|
| REST + WebSocket | REST para operaciones CRUD e idempotentes; WS para tiempo real sin polling. |
| Idempotencia en Redis | Evita duplicados por reintentos en conexiones inestables. |
| PostgreSQL + Redis | Postgres para consistencia; Redis para velocidad y pub/sub. |
| PWA estática servida por Nginx | No requiere SSR; tolera desconexiones con Service Worker. |
| JWT con PIN | Auth simple para el contexto de juego; sin OAuth. |
| Masa monetaria fija | Evita inflación artificial; el banco controla la emisión. |
| GORM con AutoMigrate (ADR-008) | Esquema definido en un solo lugar (modelos Go); sin scripts SQL manuales que se puedan desincronizar. |

## 4. Flujo de una Transferencia (Ejemplo)
1. Jugador A envía `POST /games/{code}/transfer` con `idempotency_key`.
2. API verifica idempotencia en Redis.
3. Servicio valida saldo suficiente, fee, y que ambos estén en la misma partida.
4. Transacción en PostgreSQL: débito A, crédito B, registro en Transaction y
   AuditLog.
5. Publica evento en Redis pub/sub, WebSocket hub, broadcast a todos los
   conectados a la partida.
6. Responde 200 con nuevo saldo.
