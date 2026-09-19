# CashCastor — Documentación BMAD

Documentación inicial (MVP) para **CashCastor**, una aplicación web que digitaliza la
gestión de la caja del banco en juegos de mesa económicos como Monopoly o Tío Rico.

## Estructura

    mesabank-cash-castor-app/
    ├── README.md
    ├── CHANGELOG.md
    ├── LICENSE
    ├── .gitignore
    ├── supuestos-y-preguntas-abiertas.md
    └── docs/
        ├── 01-product-brief.md
        ├── 02-prd.md
        ├── 03-arquitectura.md
        ├── 04-adrs/
        │   ├── ADR-001-rest-websocket.md
        │   ├── ADR-002-idempotencia-redis.md
        │   ├── ADR-003-masa-monetaria-fija.md
        │   ├── ADR-004-postgres-redis.md
        │   ├── ADR-005-auth-pin-jwt.md
        │   ├── ADR-006-pwa-vite.md
        │   └── ADR-007-kubernetes-single-iac.md
        ├── 05-roadmap.md
        ├── 06-epicas.md
        ├── 07-historias-usuario.md
        ├── 08-openapi.yaml
        ├── 09-plan-de-pruebas.md
        ├── 10-backlog-sprint-0.md
        └── 11-reglamento-juegos.md

## Alcance de esta fase

- **Fase anterior:** documentación (Brief, PRD, Arquitectura, ADRs, Roadmap).
- **Fase actual:** desglose técnico para los devs (épicas, historias de
  usuario, OpenAPI, plan de pruebas, backlog de sprint 0).

## Roles

- **Product Owner / Arquitecto:** Erick Salamanca
- **Equipo de desarrollo:** por definir

## Flujo de trabajo (Git)

- **Rama de trabajo:** `develop`. Todo el desarrollo se integra ahí; `main`
  queda reservada para releases (sin estrategia de release definida aún).
- **`develop` está protegida:** no se permite push directo (salvo bypass de
  administrador de repositorio). Todo cambio entra por **Pull Request**.
- **Revisión obligatoria:** cada PR requiere al menos una aprobación de
  code owner (`.github/CODEOWNERS`, actualmente `@esalamancar` para todo el
  repositorio).
- **Checks obligatorios antes de fusionar** (`.github/workflows/quality-checks.yml`):
  - `Backend build` — compila `backend/` con `go build`.
  - `Frontend build` — compila `frontend/` con `npm run build`.
  - `Backend Docker image` / `Frontend Docker image` — construyen las
    imágenes de `backend/Dockerfile` y `frontend/Dockerfile`.
  - Cada check solo ejecuta el build real si el PR modificó esa carpeta;
    si no hay cambios ahí, reporta éxito sin gastar tiempo de CI.
  - Aún no hay pruebas unitarias ni de integración (ver `09-plan-de-pruebas.md`);
    estos checks son el gate mínimo mientras tanto.
- **Colaboradores con acceso `Write`** (equivalente a "developer" en
  GitHub): `2alelagos-dot`, `LoBe4`, `nCrisz`.

## Convenciones

- Idioma: español.
- Formato: Markdown.
- Ubicación: repositorio Git dedicado (`mesabank-cash-castor-app`; el nombre
  del repositorio no refleja el nombre del producto —**CashCastor**— por
  decisión explícita de mantenerlo así).
