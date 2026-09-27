# CashCastor — Documentación BMAD

Documentación inicial (MVP) para **CashCastor**, una aplicación web que digitaliza la
gestión de la caja del banco en juegos de mesa económicos como Monopoly o Tío Rico.

> **¿Sos dev nuevo en esto?** Leé en orden: esta sección ("Dónde encontrar
> todo"), después "Cómo empezar a desarrollar", y después "Flujo de trabajo
> (Git)" antes de tu primer PR.

## Dónde encontrar todo

| Carpeta/archivo | Qué es | Cuándo mirarlo |
|---|---|---|
| [`docs/`](docs/) | Documentación BMAD: brief, PRD, arquitectura, ADRs, roadmap, épicas, historias de usuario (con criterios de aceptación en Gherkin), OpenAPI, plan de pruebas, backlog de Sprint 0, reglamento de los juegos | Para entender **qué** hay que construir, **por qué**, y con qué reglas de negocio/criterios de aceptación |
| [`backlog/tareas.md`](backlog/tareas.md) | Todas las épicas/features/fixes con su ID y estado actual (`Por hacer`/`En progreso`/`En revisión`/`Hecho`) | Para ver **qué está pendiente** y actualizar el estado de lo que estés haciendo |
| [`backlog/asignaciones.md`](backlog/asignaciones.md) | Quién está trabajando en cada ID, y el frente (backend/frontend/DevOps) de cada dev en el sprint actual | Para saber **quién hace qué** ahora mismo |
| [`backend/`](backend/) | Código del backend (Go + Gin) | Trabajo de backend |
| [`frontend/`](frontend/) | Código del frontend (Vite + React + TypeScript) | Trabajo de frontend |
| [`supuestos-y-preguntas-abiertas.md`](supuestos-y-preguntas-abiertas.md) | Supuestos ya tomados y preguntas de negocio todavía sin resolver | Si algo del comportamiento esperado no está claro, revisá acá antes de asumir |
| [`.github/workflows/quality-checks.yml`](.github/workflows/quality-checks.yml) | Los checks que corren en cada PR | Para entender por qué un PR está en rojo |
| [`CHANGELOG.md`](CHANGELOG.md) | Historial de decisiones y cambios de alcance | Para ver el porqué de una decisión ya tomada |

**Regla general:** `docs/` es planeación (cambia poco, es el diseño
acordado); `backlog/` es el estado real del trabajo (cambia todo el
tiempo). Si algo en `backlog/` contradice `docs/`, `docs/` manda salvo que
se haya discutido y actualizado explícitamente.

## Cómo empezar a desarrollar

Prerrequisitos: Go 1.22+, Node 20+, Docker (para construir las imágenes).

### Backend

```bash
cd backend
go run .                 # levanta el servidor en :8080 (GET /healthz)
# o, para probar la imagen de despliegue:
docker build -t cashcastor-api .
docker run -p 8080:8080 cashcastor-api
```

Es un scaffold mínimo: todavía no conecta a Postgres/Redis (eso es
`T-04`/`T-05` en `backlog/tareas.md`, sin hacer aún).

### Frontend

```bash
cd frontend
npm install
npm run dev               # servidor de desarrollo con hot-reload
npm run build              # build de producción a dist/
# o, para probar la imagen de despliegue (sirve el build vía Nginx):
docker build -t cashcastor-web .
docker run -p 8081:80 cashcastor-web
```

### API

El contrato de la API (endpoints, schemas, ejemplos) está en
[`docs/08-openapi.yaml`](docs/08-openapi.yaml). El backend real todavía no
implementa esos endpoints (ver estado en `backlog/tareas.md`). Para verla
como documentación navegable, abrí `docs/openapi.html` en el navegador
(generado con Redoc; regenerarlo con `npx @redocly/cli build-docs
docs/08-openapi.yaml -o docs/openapi.html` después de tocar la spec).

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
        │   ├── ADR-007-kubernetes-single-iac.md
        │   └── ADR-008-gorm-migraciones.md
        ├── 05-roadmap.md
        ├── 06-epicas.md
        ├── 07-historias-usuario.md
        ├── 08-openapi.yaml
        ├── 09-plan-de-pruebas.md
        ├── 10-backlog-sprint-0.md
        └── 11-reglamento-juegos.md

## Alcance de esta fase

- **Fase anterior:** documentación (Brief, PRD, Arquitectura, ADRs, Roadmap).
- **Fase siguiente:** desglose técnico para los devs (épicas, historias de
  usuario, OpenAPI, plan de pruebas, backlog de sprint 0).
- **Fase actual:** arranque de Sprint 0 — scaffold de backend/frontend,
  git-flow y CI configurados, equipo confirmado y trabajando (ver
  `backlog/`).

## Roles

- **Product Owner / Arquitecto / DevOps:** Erick Salamanca.
- **Equipo de desarrollo (todos fullstack):** Harri, Cris (frente backend
  en el sprint actual), Aleja, Andrety (frente frontend en el sprint
  actual). El frente es el foco del sprint, no una limitación de skill —
  ver el detalle y las tareas asignadas en
  [`backlog/asignaciones.md`](backlog/asignaciones.md).

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
- `OpenAPI lint` — corre `@redocly/cli lint` sobre `docs/08-openapi.yaml`
  en todos los PRs (rápido, no depende de qué carpeta cambió).
- `Backend build` — compila y lintea (`golangci-lint`) `backend/`.
- `Frontend build` — compila, lintea, formatea y testea `frontend/`
  (`npm run lint/format:check/test:run/build`).
- `Backend Docker image` / `Frontend Docker image` — construyen las
  imágenes de `backend/Dockerfile` y `frontend/Dockerfile`.
- Los últimos cuatro solo ejecutan el build/lint/test real si el PR
  modificó esa carpeta; si no hay cambios ahí, reportan éxito sin gastar
  tiempo de CI.
- Aún no hay pruebas de integración/e2e reales (ver `09-plan-de-pruebas.md`);
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
