# Backlog Sprint 0: Fundaciones — CashCastor

Tareas técnicas de la Fase 0 del roadmap (`05-roadmap.md`, Semana 1-2),
ordenadas por dependencia. No son historias de usuario funcionales, sino
trabajo de habilitación para que las épicas de `06-epicas.md` puedan
empezar a construirse desde la Fase 1.

## Resumen ordenado por dependencia

| # | Tarea | Depende de |
|---|-------|------------|
| T-01 | Setup del repositorio backend (Go) | — |
| T-02 | Setup del repositorio frontend (React + Vite + PWA) | — |
| T-03 | Definición y publicación de la OpenAPI spec | `08-openapi.yaml` (ya generado) |
| T-04 | Setup de PostgreSQL y Redis en local/desarrollo (docker-compose) | T-01 |
| T-05 | Modelo de datos y migraciones iniciales | T-04 |
| T-06 | Esqueleto de capas backend (API/Servicios/Repositorios/WebSocket) | T-01, T-05 |
| T-07 | Cliente API + WebSocket base en el frontend | T-02, T-03 |
| T-08 | Coordinación con `single-iac`: convenciones de namespace, secrets y pipelines | — |
| T-09 | Manifiestos de Kubernetes (namespace, ConfigMaps, Secrets) siguiendo `single-iac` | T-08 |
| T-10 | Pipeline de CI/CD base (build, test, lint, deploy) en `single-iac` | T-01, T-02, T-08, T-09 |
| T-11 | Observabilidad base (logging estructurado, health checks, métricas) | T-06 |
| T-12 | Documentación de entornos (local, dev, staging, prod) | T-04, T-09 |
| T-13 | Quality gate de PRs en GitHub Actions (compilación + build de imágenes Docker) | T-01, T-02 |

No se asignan tiempos por tarea: cuánto tarda cada una depende de quién la
haga (herramientas, experiencia), no es una propiedad fija de la tarea. El
orden de dependencia arriba sí es fijo y es lo que determina qué se puede
paralelizar entre el equipo disponible.

---

## Detalle de tareas

### T-01: Setup del repositorio backend (Go)
- Inicializar módulo Go, estructura de carpetas por capas (`api/`,
  `service/`, `repository/`, `ws/`, `config/`).
- Configurar Gin, middleware base (recovery, CORS, logging).
- Configurar linting (`golangci-lint`) y formato (`gofmt`/`goimports`).
- Configurar `go test` con soporte para `-race` desde el inicio.
- **Estado (2026-09-27): Hecho.** Estructura de capas en `backend/internal/`
  (`api/`, `middleware/`, `ws/`, `db/`, `models/`); `golangci-lint`
  configurado (`.golangci.yml`) y corriendo en CI como parte del check
  `Backend build`. `go test -race` funcionará una vez existan tests (aún no
  hay lógica de negocio que probar).
- **Dependencias:** ninguna.

### T-02: Setup del repositorio frontend (React + Vite + PWA)
- Inicializar proyecto Vite + React + TypeScript.
- Instalar y configurar el plugin PWA (manifest, Service Worker básico de
  caché de assets estáticos, ADR-006).
- Configurar estado global (Zustand o Redux Toolkit, según ADR/arquitectura
  §2.2) con un store vacío de referencia.
- Configurar ESLint/Prettier y Vitest.
- **Estado (2026-09-27): Hecho.** `vite-plugin-pwa` configurado (manifest +
  Service Worker cacheando assets estáticos); store de referencia con
  Zustand (`src/store/useSessionStore.ts`); ESLint (flat config) +
  Prettier + Vitest con un test de humo, los cuatro corriendo en CI
  (`npm run lint/format:check/test:run/build`) como parte del check
  `Frontend build`. Validado con Node 20 (la versión de CI) además de
  local.
- **Dependencias:** ninguna.

### T-03: Definición y publicación de la OpenAPI spec
- Validar `docs/08-openapi.yaml` con `@redocly/cli lint` en el pipeline de
  CI (evitar que se rompa el contrato sin darse cuenta).
- Publicar la spec como documentación navegable (Redoc/Swagger UI) para el
  equipo, aunque sea en un job de CI o página estática interna.
- **Estado (2026-09-27): Hecho.** Check `OpenAPI lint` agregado como
  obligatorio en las 3 rulesets; `docs/openapi.html` (Redoc estático,
  autocontenido) generado y commiteado, se regenera a mano cuando cambia
  la spec (comando en `README.md`).
- **Dependencias:** spec ya generada en `docs/08-openapi.yaml`.

### T-04: Setup de PostgreSQL y Redis en local/desarrollo
- `docker-compose.yml` con PostgreSQL y Redis para desarrollo local.
- Variables de entorno documentadas (`.env.example`).
- **Dependencias:** T-01 (para que el backend tenga dónde conectarse).

### T-05: Modelo de datos y migraciones iniciales
- **Actualizado (ADR-008):** no se usa `golang-migrate`/`goose` ni SQL
  manual. El esquema se define como modelos GORM y se aplica con
  `AutoMigrate`, disparado por el propio binario del backend con el flag
  `--migrations-only` (corre las migraciones y termina, pensado para un Job
  de Kubernetes o paso de pipeline).
- Modelos GORM para las 7 entidades del PRD §5: `User`, `Game`,
  `GameConfig`, `Player`, `Transaction`, `Loan`, `AuditLog`, con índices y
  relaciones (`game_id`, `user_id`, `player_id`) declaradas vía tags GORM.
- Seed mínimo de datos de desarrollo (usuario de prueba, monedas por
  defecto para EP-07) — sigue pendiente de definir cómo se dispara.
- **Dependencias:** T-04.

### T-06: Esqueleto de capas backend
- Handlers REST vacíos (o con respuesta *stub*) para los endpoints de
  `08-openapi.yaml`, organizados por los tags del spec (`auth`, `games`,
  `finanzas`, `tiempo-real`, `ranking`, `admin`).
- Middleware de autenticación JWT (validación de Bearer token, sin login
  real todavía si no está listo EP-01).
- Hub de WebSocket base (una conexión, un canal por partida, sin lógica de
  negocio aún).
- Capa de repositorios conectada a PostgreSQL vía GORM (ADR-008).
- **Estado (2026-09-27): Hecho** (`backend/internal/api/router.go`) —
  las 28 operaciones de `08-openapi.yaml` están cableadas como stubs (501),
  con el middleware de auth (`internal/middleware`, placeholder estructural
  hasta que exista EP-01) aplicado donde el spec lo exige, y el hub de
  WebSocket (`internal/ws`) sirviendo `GET /games/{code}/ws` de verdad
  (upgrade + registro por partida, sin lógica de negocio). La capa de
  repositorios propiamente dicha (queries reales sobre los modelos GORM)
  se agrega historia por historia a partir de Sprint 1, no de una vez acá.
- **Dependencias:** T-01, T-05.

### T-07: Cliente API + WebSocket base en el frontend
- Cliente HTTP tipado a partir de la OpenAPI spec (generado o escrito a
  mano) para consumir los endpoints.
- Cliente WebSocket nativo con reconexión automática (soporta US-003,
  US-080).
- **Estado (2026-09-27): Hecho** (`frontend/src/api`, `frontend/src/ws`) —
  cliente HTTP tipado a mano (`types.ts`, `client.ts` con manejo de
  `ApiError`, `auth.ts`, `games.ts`) para los endpoints de EP-01/EP-02
  (alcance de Sprint 1); `GameSocket` con reconexión por backoff
  exponencial. No depende de T-03 en la práctica (se escribió a mano, no
  generado desde la spec), así que no bloquea a Aleja/Andrety.
- **Dependencias:** T-02, T-03.

### T-08: Coordinación con `single-iac`
- Reunión con el equipo/dueño de `single-iac` para conocer: convención de
  nombres de namespace, formato de Secrets/ConfigMaps, estructura esperada
  de pipelines, y el agente de CI/CD existente (ADR-007).
- Documentar las convenciones encontradas (actualiza la pregunta abierta #8
  de `supuestos-y-preguntas-abiertas.md`; no se modifica ese archivo en esta
  entrega, pero es insumo directo para hacerlo después).
- **Estado (2026-09-27): resuelto por el agente de `single-iac`** —
  namespace `cashcastor` agregado a `var.apps`, formato de Secrets/
  ConfigMaps y estructura de pipeline documentados en su ADR-0017. Ver
  ADR-009 de este repo para el resumen. Pendiente: revisión y merge de la
  rama `add-cashcastor-namespace` por el PO.
- **Dependencias:** ninguna, pero bloquea T-09 y T-10. Es la tarea de mayor
  riesgo de calendario por depender de un equipo externo: priorizarla desde
  el día 1 de la Fase 0.

### T-09: Manifiestos de Kubernetes
- Namespace `cashcastor` (arquitectura §2.4).
- ConfigMaps y Secrets para configuración de la app y credenciales de DB,
  siguiendo el formato acordado en T-08.
- Manifiestos base de `Deployment`/`Service` para `api` y `web` (sin CI/CD
  todavía, aplicables manualmente para pruebas).
- **Estado (2026-09-27): resuelto por el agente de `single-iac`** —
  Deployments/Services de `api` y `web`, Postgres+Redis propios, 2
  Ingress+TLS. Mismo lugar y estado que T-08 (rama sin mergear/aplicar).
- **Dependencias:** T-08.

### T-10: Pipeline de CI/CD base
- ~~Pipeline en `single-iac` que ejecute: lint + test backend, lint + test
  frontend, build de imágenes, y despliegue a un entorno de desarrollo.~~
- **Redefinido (ADR-009, 2026-09-27):** `single-iac` nunca construye ni
  despliega imágenes de otra app (incidente real de por medio); el
  workflow de build+push+deploy es de **este** repo, no de `single-iac`.
  Pendiente: crear ese workflow acá (Erick pidió que lo redacte el agente
  de `single-iac`, que conoce el patrón exacto, pero que el archivo viva
  en este repo).
- Integrar el lint de la OpenAPI spec (T-03, ya hecho como check propio)
  como gate del pipeline de deploy también.
- **Dependencias:** T-01, T-02, T-08, T-09.

### T-11: Observabilidad base
- Logging estructurado (JSON) en el backend, con correlación por
  `request_id` y, cuando aplique, `game_id`.
- Endpoints de `healthz`/`readyz` para probes de Kubernetes.
- Métricas básicas expuestas (Prometheus): latencia de endpoints,
  conexiones WebSocket activas, tasa de errores.
- **Estado (2026-09-27): Hecho** (`backend/internal/observability`) —
  logging JSON por request (`request_id` propio o del header
  `X-Request-Id`, más `game_code` cuando la ruta tiene `:code`); `GET
  /readyz` hace ping real a Postgres y responde 503 si falla (a diferencia
  de `/healthz`, que solo indica que el proceso está vivo); `GET /metrics`
  expone `http_requests_total`, `http_request_duration_seconds` y
  `websocket_connections_active` en formato Prometheus.
- **Dependencias:** T-06.

### T-12: Documentación de entornos
- Documentar cómo levantar el entorno local (docker-compose), y cómo se
  diferencian dev/staging/prod (namespaces, dominios, niveles de logging).
- **Estado (2026-09-27): parcialmente resuelto.** Entorno local: ver
  `README.md` (`docker-compose.yml`, `.env.example`). Entornos
  dev/staging/prod: `single-iac` **no separa infraestructura por entorno**
  (un namespace único por app, `cashcastor` mapeado a la rama `develop`) —
  decisión explícita documentada en ADR-009, no una omisión. Se revisa
  cuando haya necesidad real (más de un ambiente desplegado a la vez).
- **Dependencias:** T-04, T-09.

### T-13: Quality gate de PRs en GitHub Actions
- Workflow `.github/workflows/quality-checks.yml`, disparado en cada PR
  hacia `develop`. Cuatro checks requeridos: `Backend build` (`go build`),
  `Frontend build` (`npm run build`), `Backend Docker image` y
  `Frontend Docker image` (`docker build` de `backend/Dockerfile` y
  `frontend/Dockerfile`).
- Cada check usa un job de detección de cambios (`dorny/paths-filter`) para
  solo ejecutar la compilación/build real cuando el PR toca `backend/` o
  `frontend/` respectivamente; si no hay cambios en esa carpeta, el check
  igual reporta éxito (para no bloquear el merge indefinidamente), pero sin
  gastar tiempo de CI en algo que no cambió.
- Es un gate independiente y complementario a T-10 (pipeline de despliegue
  en `single-iac`): este corre en GitHub Actions sobre cada PR: valida que
  el código compile y las imágenes construyan antes de fusionar a `develop`;
  T-10 se encarga del build/despliegue real hacia los entornos.
- **Estado (2026-09-18):** implementado y verificado — los 4 checks
  pasaron en el PR #1 (scaffold inicial de backend/frontend).
- **Dependencias:** T-01, T-02.

---

## Riesgos y notas

- **T-08 es la tarea de mayor riesgo:** las convenciones exactas de
  `single-iac` son una pregunta abierta (ver
  `supuestos-y-preguntas-abiertas.md`, pregunta #8). Si no se resuelve
  temprano, T-09 y T-10 —y por lo tanto todo el pipeline de despliegue—
  se retrasan. Se recomienda agendarla en los primeros 2 días de la Fase 0.
- **Resuelto (ADR-008):** la elección entre `sqlc` y `GORM` (antes abierta
  en `03-arquitectura.md` §2.1) se decidió por GORM, con `AutoMigrate` en
  vez de `golang-migrate`/`goose` — ver ADR-008 para el detalle y el
  razonamiento.
