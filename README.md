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

## Convenciones

- Idioma: español.
- Formato: Markdown.
- Ubicación: repositorio Git dedicado (`mesabank-cash-castor-app`; el nombre
  del repositorio no refleja el nombre del producto —**CashCastor**— por
  decisión explícita de mantenerlo así).
