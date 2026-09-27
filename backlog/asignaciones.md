# Asignaciones

Quién está trabajando en cada ID de [`tareas.md`](./tareas.md). Se actualiza
a mano; si un ID no aparece acá, todavía no tiene dueño asignado.

## Equipo

Todos los devs son fullstack — el frente asignado abajo es dónde van a
enfocarse en el Sprint 0, no una limitación de skill.

| Dev | Rol | Frente en Sprint 0 |
|-----|-----|---------------------|
| Erick | Product Owner / DevOps / Infra / Arquitectura | DevOps, infraestructura, arquitectura |
| Harri | Desarrollo (fullstack) | Backend |
| Cris | Desarrollo (fullstack) | Backend |
| Aleja | Desarrollo (fullstack) | Frontend |
| Andrety | Desarrollo (fullstack) | Frontend |

## Sprint 0 (tareas técnicas `T-xx`)

Reparto por frente (2026-09-27): Harri y Cris en backend, Aleja y Andrety
en frontend, Erick en DevOps/infraestructura/arquitectura. Dentro de cada
frente se dividen las tareas según carga y dependencia, no necesariamente
una por persona.

| ID | Asignado a | Notas |
|----|------------|-------|
| T-01 | Harri | Backend. En progreso — falta estructura de capas y linting sobre el scaffold ya creado |
| T-02 | Aleja | Frontend. En progreso — falta plugin PWA y store sobre el scaffold ya creado |
| T-03 | Cris | Backend |
| T-04 | Erick | DevOps |
| T-05 | Cris | Backend. Depende de T-04 |
| T-06 | Harri | Backend. Depende de T-01, T-05 |
| T-07 | Andrety | Frontend. Depende de T-02, T-03 |
| T-08 | Erick | DevOps. Mayor riesgo de calendario (depende de un equipo externo); priorizar desde el día 1 |
| T-09 | Erick | DevOps. Depende de T-08 |
| T-10 | Erick | DevOps. Depende de T-01, T-02, T-08, T-09 |
| T-11 | Erick | DevOps. Depende de T-06 |
| T-12 | Erick | DevOps. Depende de T-04, T-09 |
| T-13 | Erick | DevOps. Hecho |

## Fase 1 en adelante (épicas y features)

Todavía sin asignar. Se reparte por épica completa (backend + frontend de
la misma persona) cuando arranque cada fase, para no fragmentar el trabajo
entre muchas personas. Se actualiza esta tabla cuando se decida.

| ID | Asignado a |
|----|------------|
| EP-01 (Autenticación y Sesión) | — |
| EP-02 (Gestión de Partidas) | — |
| EP-03 (Configuración de Partida) | — |
| EP-04 (Operaciones Financieras) | — |
| EP-05 (Tiempo Real y Turnos) | — |
| EP-06 (Ranking y Puntuación) | — |
| EP-07 (Administración) | — |
