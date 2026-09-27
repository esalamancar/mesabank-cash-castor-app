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
    ├── .github/
    │   ├── CODEOWNERS
    │   └── workflows/quality-checks.yml
    ├── backend/          # scaffold Go + Gin (Dockerfile incluido)
    ├── frontend/         # scaffold Vite + React + TS (Dockerfile incluido)
    ├── backlog/          # tablero de trabajo vivo (distinto de docs/, ver abajo)
    │   ├── tareas.md
    │   └── asignaciones.md
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
- **Equipo de desarrollo:** Harri, Aleja, Andrety, Cris. Ver reparto de
  tareas en [`backlog/asignaciones.md`](backlog/asignaciones.md).

## Flujo de trabajo (Git)

### Ramas

Solo existen tres ramas permanentes, en este orden de promoción:

```
develop  →  qa  →  prod
```

- **`develop`** es la única rama donde se integra código nuevo, y es la
  rama por defecto del repo. Todo feature/fix de los devs se mergea acá.
- **`qa`** y **`prod`** son ramas de **promoción**: nunca se codifica
  directo en ellas, solo reciben merges (vía PR) desde la rama anterior en
  la cadena (`develop` → `qa`, y `qa` → `prod`).
- Nadie hace push directo a ninguna de las tres — ni siquiera el PO. Las
  tres están protegidas con un Ruleset de GitHub: sin push directo, sin
  force-push, sin borrado. Todo cambio entra por **Pull Request**.

### Ramas de los devs

Cada dev crea su propia rama a partir de `develop` para trabajar, y abre un
PR hacia `develop` cuando esté listo. El nombre de la rama **debe** seguir
el patrón `tipo/descripcion-en-kebab-case`, con `tipo` uno de:
`feature`, `fix`, `chore`, `docs`, `refactor`, `test`, `hotfix`.

Ejemplos válidos: `feature/transferencias-p2p`, `fix/saldo-negativo-arqueo`,
`chore/actualizar-dependencias`. Nombres ambiguos como `front` o `test123`
**no** son válidos.

Esto se aplica con dos mecanismos (GitHub no permite forzar el nombre de
rama nativamente en el plan Free/público — esa regla concreta requiere
GitHub Team/Enterprise):
1. Convención documentada acá.
2. Check obligatorio de CI (`Branch name convention`, en
   `.github/workflows/quality-checks.yml`) que falla el PR si el nombre no
   cumple el patrón — bloquea el merge aunque alguien lo olvide.

### Revisión y aprobación

Cada PR (hacia `develop`, `qa` o `prod`) requiere al menos una aprobación
de code owner (`.github/CODEOWNERS`, actualmente `@esalamancar` para todo
el repositorio) — **Erick es el único que puede aprobar/mergear PRs.**
GitHub no permite auto-aprobar tu propio PR; los PRs de `@esalamancar`
(incluyendo las promociones `develop→qa` y `qa→prod`) se mergean vía el
bypass de Admin del Ruleset, no por auto-aprobación.

### Checks obligatorios antes de fusionar

`.github/workflows/quality-checks.yml`, corre en PRs hacia las tres ramas:

- `Branch name convention` — valida el nombre de la rama origen (ver
  arriba). Se omite si la rama origen es `develop`, `qa` o `prod` (PRs de
  promoción).
- `Backend build` — compila `backend/` con `go build`.
- `Frontend build` — compila `frontend/` con `npm run build`.
- `Backend Docker image` / `Frontend Docker image` — construyen las
  imágenes de `backend/Dockerfile` y `frontend/Dockerfile`.
- Los tres últimos solo ejecutan el build real si el PR modificó esa
  carpeta; si no hay cambios ahí, reportan éxito sin gastar tiempo de CI.
- Aún no hay pruebas unitarias ni de integración (ver `09-plan-de-pruebas.md`);
  estos checks son el gate mínimo mientras tanto.

### Repositorio público

Las reglas de rama (Ruleset) en GitHub requieren plan Pro para repos
privados en cuentas personales; se optó por hacer el repo público en vez
de pagar el upgrade (2026-09-18). No hay código de negocio real todavía,
solo scaffolding y documentación.

### Colaboradores

Acceso `Write` (equivalente a "developer" en GitHub): `2alelagos-dot`,
`LoBe4`, `nCrisz`.

## Convenciones

- Idioma: español.
- Formato: Markdown.
- Ubicación: repositorio Git dedicado (`mesabank-cash-castor-app`; el nombre
  del repositorio no refleja el nombre del producto —**CashCastor**— por
  decisión explícita de mantenerlo así).
