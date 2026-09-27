# Asignaciones

Quién está trabajando en cada ID de [`tareas.md`](./tareas.md). Se actualiza
a mano; si un ID no aparece acá, todavía no tiene dueño asignado.

## Equipo

Todos los devs son fullstack — el frente asignado abajo es dónde van a
enfocarse en el sprint actual, no una limitación de skill.

| Dev | Rol | Frente en Sprint 1 |
|-----|-----|---------------------|
| Erick | Product Owner / DevOps / Infra / Arquitectura | DevOps, infraestructura, arquitectura (sigue cerrando Sprint 0) |
| Harri | Desarrollo (fullstack) | Backend — EP-01, EP-03 |
| Aleja | Desarrollo (fullstack) | Frontend — EP-01, EP-03 |
| Cris | Desarrollo (fullstack) | Backend — EP-02 |
| Andrety | Desarrollo (fullstack) | Frontend — EP-02 |

## Sprint 0 (tareas técnicas `T-xx`) — corregido 2026-09-27

**Corrección:** Sprint 0 son las fundaciones arquitectónicas (setup de
repos, migraciones, CI/CD, Kubernetes) y las venía haciendo Erick
directamente, no el equipo de devs. La versión anterior de esta tabla los
tenía asignados por error; se corrige acá. El detalle de qué falta de
Sprint 0 está al final de este archivo.

| ID | Asignado a | Estado |
|----|------------|--------|
| T-01 | Erick | En progreso — scaffold mínimo listo, falta estructura de capas y linting |
| T-02 | Erick | En progreso — scaffold mínimo listo, falta plugin PWA y store |
| T-03 | Erick | Por hacer |
| T-04 | Erick | Por hacer |
| T-05 | Erick | Por hacer — depende de T-04 |
| T-06 | Erick | Por hacer — depende de T-01, T-05 |
| T-07 | Erick | Por hacer — depende de T-02, T-03 |
| T-08 | Erick | Por hacer — mayor riesgo de calendario (depende de un equipo externo); prioridad |
| T-09 | Erick | Por hacer — depende de T-08 |
| T-10 | Erick | Por hacer — depende de T-01, T-02, T-08, T-09 |
| T-11 | Erick | Por hacer — depende de T-06 |
| T-12 | Erick | Por hacer — depende de T-04, T-09 |
| T-13 | Erick | Hecho |

## Sprint 1 (Fase 1: Core Bancario — `EP-01`, `EP-02`, `EP-03`)

Reparto por pareja backend+frontend sobre la misma épica, para minimizar
coordinación entre personas distintas en la misma pieza:

| Épica | Backend | Frontend | Notas |
|-------|---------|----------|-------|
| EP-01 (Autenticación y Sesión) | Harri | Aleja | Sin dependencias, puede arrancar ya |
| EP-02 (Gestión de Partidas) | Cris | Andrety | Depende de EP-01 (requiere usuario autenticado o invitado) |
| EP-03 (Configuración de Partida) | Harri | Aleja | Depende de EP-02 (la configuración pertenece a una partida ya creada) — arrancan con EP-01 mientras tanto |

Historias de cada épica, criterios de aceptación en Gherkin y notas
técnicas: `docs/07-historias-usuario.md`. Estado de cada historia:
`backlog/tareas.md`.

## Fase 2 en adelante (`EP-04` a `EP-07`)

Todavía sin asignar.

| ID | Asignado a |
|----|------------|
| EP-04 (Operaciones Financieras) | — |
| EP-05 (Tiempo Real y Turnos) | — |
| EP-06 (Ranking y Puntuación) | — |
| EP-07 (Administración) | — |

## Pendiente de Sprint 0 (a cargo de Erick)

Solo `T-13` está terminado. Todo lo demás sigue pendiente:

- **T-01 / T-02** (scaffold backend/frontend): el scaffold mínimo ya existe
  y compila, pero falta la estructura de capas del backend
  (`api/`, `service/`, `repository/`, `ws/`, `config/`) + `golangci-lint`,
  y del lado frontend falta el plugin PWA, el store global, ESLint/Prettier
  y Vitest.
- **T-03** (OpenAPI): falta agregar el lint (`@redocly/cli`) al pipeline de
  CI y publicar la spec como documentación navegable (Redoc/Swagger UI).
- **T-04** (Postgres/Redis local): sin empezar, `docker-compose.yml` +
  `.env.example` pendientes.
- **T-05** (migraciones): sin empezar, depende de T-04.
- **T-06** (esqueleto de capas backend, handlers stub, hub WS base): sin
  empezar, depende de T-01/T-05.
- **T-07** (cliente API + WebSocket del frontend): sin empezar, depende de
  T-02/T-03.
- **T-08** (coordinación con `single-iac`): sin empezar — es la tarea de
  mayor riesgo de calendario porque depende de un equipo externo.
- **T-09 / T-10** (manifiestos de Kubernetes, pipeline de CI/CD en
  `single-iac`): sin empezar, bloqueadas por T-08.
- **T-11** (observabilidad: logging estructurado, `readyz`, métricas):
  sin empezar, solo existe el `GET /healthz` básico del scaffold.
- **T-12** (documentación de entornos local/dev/staging/prod): sin
  empezar.
