# Backlog Sprint 0: Fundaciones — CashCastor

Tareas técnicas de la Fase 0 del roadmap (`05-roadmap.md`, Semana 1-2),
ordenadas por dependencia. No son historias de usuario funcionales, sino
trabajo de habilitación para que las épicas de `06-epicas.md` puedan
empezar a construirse desde la Fase 1.

**Convención de estimación:** días-persona (1 persona, jornada completa).

## Resumen ordenado por dependencia

| # | Tarea | Estimación | Depende de |
|---|-------|------------|------------|
| T-01 | Setup del repositorio backend (Go) | 1 día | — |
| T-02 | Setup del repositorio frontend (React + Vite + PWA) | 1 día | — |
| T-03 | Definición y publicación de la OpenAPI spec | 0.5 día | `08-openapi.yaml` (ya generado) |
| T-04 | Setup de PostgreSQL y Redis en local/desarrollo (docker-compose) | 1 día | T-01 |
| T-05 | Modelo de datos y migraciones iniciales | 2 días | T-04 |
| T-06 | Esqueleto de capas backend (API/Servicios/Repositorios/WebSocket) | 2 días | T-01, T-05 |
| T-07 | Cliente API + WebSocket base en el frontend | 1.5 días | T-02, T-03 |
| T-08 | Coordinación con `single-iac`: convenciones de namespace, secrets y pipelines | 1 día (bloqueante externo) | — |
| T-09 | Manifiestos de Kubernetes (namespace, ConfigMaps, Secrets) siguiendo `single-iac` | 1.5 días | T-08 |
| T-10 | Pipeline de CI/CD base (build, test, lint, deploy) en `single-iac` | 2 días | T-01, T-02, T-08, T-09 |
| T-11 | Observabilidad base (logging estructurado, health checks, métricas) | 1.5 días | T-06 |
| T-12 | Documentación de entornos (local, dev, staging, prod) | 0.5 día | T-04, T-09 |
| T-13 | Quality gate de PRs en GitHub Actions (compilación + build de imágenes Docker) | 1 día | T-01, T-02 |

**Total estimado:** ~16.5 días-persona, ejecutable en paralelo entre 1-2
desarrolladores en las 2 semanas de la Fase 0.

---

## Detalle de tareas

### T-01: Setup del repositorio backend (Go)
- Inicializar módulo Go, estructura de carpetas por capas (`api/`,
  `service/`, `repository/`, `ws/`, `config/`).
- Configurar Gin, middleware base (recovery, CORS, logging).
- Configurar linting (`golangci-lint`) y formato (`gofmt`/`goimports`).
- Configurar `go test` con soporte para `-race` desde el inicio.
- **Estado (2026-09-18):** scaffold mínimo creado en `backend/` (`go.mod`,
  `main.go` con Gin y un endpoint `GET /healthz`, `Dockerfile` multi-stage
  verificado localmente). Pendiente: estructura de capas, middleware,
  `golangci-lint`.
- **Dependencias:** ninguna.

### T-02: Setup del repositorio frontend (React + Vite + PWA)
- Inicializar proyecto Vite + React + TypeScript.
- Instalar y configurar el plugin PWA (manifest, Service Worker básico de
  caché de assets estáticos, ADR-006).
- Configurar estado global (Zustand o Redux Toolkit, según ADR/arquitectura
  §2.2) con un store vacío de referencia.
- Configurar ESLint/Prettier y Vitest.
- **Estado (2026-09-18):** scaffold mínimo creado en `frontend/` (Vite +
  React + TypeScript, `Dockerfile` multi-stage que sirve el build vía Nginx,
  verificado localmente con `npm run build`). Pendiente: plugin PWA
  (manifest, Service Worker), store global, ESLint/Prettier, Vitest.
- **Dependencias:** ninguna.

### T-03: Definición y publicación de la OpenAPI spec
- Validar `docs/08-openapi.yaml` con `@redocly/cli lint` en el pipeline de
  CI (evitar que se rompa el contrato sin darse cuenta).
- Publicar la spec como documentación navegable (Redoc/Swagger UI) para el
  equipo, aunque sea en un job de CI o página estática interna.
- **Dependencias:** spec ya generada en `docs/08-openapi.yaml`.

### T-04: Setup de PostgreSQL y Redis en local/desarrollo
- `docker-compose.yml` con PostgreSQL y Redis para desarrollo local.
- Variables de entorno documentadas (`.env.example`).
- **Dependencias:** T-01 (para que el backend tenga dónde conectarse).

### T-05: Modelo de datos y migraciones iniciales
- Elegir herramienta de migraciones (`golang-migrate` o `goose`).
- Migraciones iniciales para las 7 entidades del PRD §5: `User`, `Game`,
  `GameConfig`, `Player`, `Transaction`, `Loan`, `AuditLog`, con índices y
  llaves foráneas correspondientes (`game_id`, `user_id`, `player_id`).
- Seed mínimo de datos de desarrollo (usuario de prueba, monedas por
  defecto para EP-07).
- **Dependencias:** T-04.

### T-06: Esqueleto de capas backend
- Handlers REST vacíos (o con respuesta *stub*) para los endpoints de
  `08-openapi.yaml`, organizados por los tags del spec (`auth`, `games`,
  `finanzas`, `tiempo-real`, `ranking`, `admin`).
- Middleware de autenticación JWT (validación de Bearer token, sin login
  real todavía si no está listo EP-01).
- Hub de WebSocket base (una conexión, un canal por partida, sin lógica de
  negocio aún).
- Capa de repositorios conectada a PostgreSQL (sqlc o GORM, según se decida
  en el equipo — la arquitectura menciona ambas opciones sin decidir).
- **Dependencias:** T-01, T-05.

### T-07: Cliente API + WebSocket base en el frontend
- Cliente HTTP tipado a partir de la OpenAPI spec (generado o escrito a
  mano) para consumir los endpoints.
- Cliente WebSocket nativo con reconexión automática (soporta US-003,
  US-080).
- **Dependencias:** T-02, T-03.

### T-08: Coordinación con `single-iac`
- Reunión con el equipo/dueño de `single-iac` para conocer: convención de
  nombres de namespace, formato de Secrets/ConfigMaps, estructura esperada
  de pipelines, y el agente de CI/CD existente (ADR-007).
- Documentar las convenciones encontradas (actualiza la pregunta abierta #8
  de `supuestos-y-preguntas-abiertas.md`; no se modifica ese archivo en esta
  entrega, pero es insumo directo para hacerlo después).
- **Dependencias:** ninguna, pero bloquea T-09 y T-10. Es la tarea de mayor
  riesgo de calendario por depender de un equipo externo: priorizarla desde
  el día 1 de la Fase 0.

### T-09: Manifiestos de Kubernetes
- Namespace `cashcastor` (arquitectura §2.4).
- ConfigMaps y Secrets para configuración de la app y credenciales de DB,
  siguiendo el formato acordado en T-08.
- Manifiestos base de `Deployment`/`Service` para `cashcastor-api` y
  `cashcastor-web` (sin CI/CD todavía, aplicables manualmente para pruebas).
- **Dependencias:** T-08.

### T-10: Pipeline de CI/CD base
- Pipeline en `single-iac` que ejecute: lint + test backend, lint + test
  frontend, build de imágenes, y despliegue a un entorno de desarrollo.
- Integrar el lint de la OpenAPI spec (T-03) como gate del pipeline.
- **Dependencias:** T-01, T-02, T-08, T-09.

### T-11: Observabilidad base
- Logging estructurado (JSON) en el backend, con correlación por
  `request_id` y, cuando aplique, `game_id`.
- Endpoints de `healthz`/`readyz` para probes de Kubernetes.
- Métricas básicas expuestas (Prometheus): latencia de endpoints,
  conexiones WebSocket activas, tasa de errores.
- **Dependencias:** T-06.

### T-12: Documentación de entornos
- Documentar cómo levantar el entorno local (docker-compose), y cómo se
  diferencian dev/staging/prod (namespaces, dominios, niveles de logging).
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
- **Estado (2026-09-18):** implementado; pendiente de verificar en el primer
  PR real hacia `develop`.
- **Dependencias:** T-01, T-02.

---

## Riesgos y notas

- **T-08 es la tarea de mayor riesgo:** las convenciones exactas de
  `single-iac` son una pregunta abierta (ver
  `supuestos-y-preguntas-abiertas.md`, pregunta #8). Si no se resuelve
  temprano, T-09 y T-10 —y por lo tanto todo el pipeline de despliegue—
  se retrasan. Se recomienda agendarla en los primeros 2 días de la Fase 0.
- La elección entre `sqlc` y `GORM` (mencionada como abierta en
  `03-arquitectura.md` §2.1) debe resolverse en T-06 antes de escribir el
  primer repositorio, para no reescribir código.
- La elección de `golang-migrate` vs. `goose` (T-05) es indistinta a nivel
  de producto; se recomienda la que el equipo de desarrollo ya conozca.
