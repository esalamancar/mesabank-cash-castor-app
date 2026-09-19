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
- Capa de Repositorios: Acceso a PostgreSQL (sqlc o GORM), Redis para
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
- Namespace: `cashcastor`.
- Deployments: `cashcastor-api` (Go), `cashcastor-web` (Nginx + estáticos).
- Services: ClusterIP para API y web.
- Ingress: Nginx Ingress con TLS (cert-manager).
- StatefulSets: PostgreSQL (o uso de operador), Redis (o servicio gestionado).
- ConfigMaps/Secrets: Configuración de la app, credenciales de DB.
- CI/CD: Pipelines definidos en `single-iac` (siguiendo reglas del repo).

## 3. Decisiones de Diseño Clave

| Decisión | Razón |
|----------|-------|
| REST + WebSocket | REST para operaciones CRUD e idempotentes; WS para tiempo real sin polling. |
| Idempotencia en Redis | Evita duplicados por reintentos en conexiones inestables. |
| PostgreSQL + Redis | Postgres para consistencia; Redis para velocidad y pub/sub. |
| PWA estática servida por Nginx | No requiere SSR; tolera desconexiones con Service Worker. |
| JWT con PIN | Auth simple para el contexto de juego; sin OAuth. |
| Masa monetaria fija | Evita inflación artificial; el banco controla la emisión. |

## 4. Flujo de una Transferencia (Ejemplo)
1. Jugador A envía `POST /games/{code}/transfer` con `idempotency_key`.
2. API verifica idempotencia en Redis.
3. Servicio valida saldo suficiente, fee, y que ambos estén en la misma partida.
4. Transacción en PostgreSQL: débito A, crédito B, registro en Transaction y
   AuditLog.
5. Publica evento en Redis pub/sub, WebSocket hub, broadcast a todos los
   conectados a la partida.
6. Responde 200 con nuevo saldo.
