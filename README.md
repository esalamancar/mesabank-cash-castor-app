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

- **Rama de trabajo y por defecto: `develop`.** Todo el desarrollo se
  integra ahí; `main` queda reservada para releases (sin estrategia de
  release definida aún).
- **`develop` está protegida** mediante un Ruleset de GitHub (equivalente
  moderno a "branch protection"): no se permite push directo, force-push ni
  borrado de la rama. Todo cambio entra por **Pull Request**. Solo el rol
  Admin del repositorio (`@esalamancar`) puede saltarse la regla en caso de
  emergencia; los colaboradores con acceso `Write` no pueden.
- **Repositorio público:** las reglas de rama (Ruleset) en GitHub requieren
  plan Pro para repos privados en cuentas personales; se optó por hacer el
  repo público en vez de pagar el upgrade (2026-09-18). No hay código de
  negocio real todavía, solo scaffolding y documentación.
- **Revisión obligatoria:** cada PR requiere al menos una aprobación de
  code owner (`.github/CODEOWNERS`, actualmente `@esalamancar` para todo el
  repositorio). Nota: GitHub no permite auto-aprobar tu propio PR; los PRs
  de `@esalamancar` los aprueba/mergea vía el bypass de Admin.
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
