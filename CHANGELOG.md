# Changelog

Todas las modificaciones notables de este proyecto serán documentadas en este archivo.

El formato se basa en Keep a Changelog y este proyecto se adhiere a Semantic Versioning.

## [Unreleased]

### Añadido
- Product Brief inicial.
- PRD MVP con requisitos funcionales y no funcionales.
- Documento de Arquitectura con diagrama y componentes.
- 7 ADRs fundacionales.
- Roadmap MVP en 6 fases.
- Documento de supuestos y preguntas abiertas.
- Épicas derivadas del PRD, una por bloque funcional (`06-epicas.md`).
- Historias de usuario del MVP (fases 1-5) con criterios de aceptación en
  Gherkin, priorización MoSCoW, estimación y sección explícita de historias
  fuera de alcance (`07-historias-usuario.md`).
- Especificación OpenAPI 3.1 de la API (`08-openapi.yaml`), validada con
  `@redocly/cli lint`.
- Plan de pruebas: estrategia unitaria/integración/e2e, casos derivados de
  CA-01 a CA-04, y pruebas dedicadas de concurrencia/idempotencia y de carga
  (`09-plan-de-pruebas.md`).
- Backlog técnico de Sprint 0 con tareas de fundación ordenadas por
  dependencia (`10-backlog-sprint-0.md`).
- Reglamento de referencia de Tío Rico (Rich Uncle) y Monopoly, con tabla de
  correspondencia con conceptos de CashCastor y una propuesta inicial de
  correspondencia específica para Tío Rico (borrador de fase 2)
  (`11-reglamento-juegos.md`).

### Corregido
- `supuestos-y-preguntas-abiertas.md`: la assumption #1 asumía que Tío Rico
  era funcionalmente equivalente a Monopoly. El reglamento oficial
  incorporado en `11-reglamento-juegos.md` mostró que no es así (Tío Rico es
  un juego de acciones y dividendos, sin propiedades/rentas/hipotecas). Se
  corrigió la assumption y se cerró la pregunta abierta #1 con la referencia
  al reglamento.
- `07-historias-usuario.md`: la fila US-200 ("Fuera del MVP") describía Tío
  Rico con mecánicas de Monopoly (hipotecas, rentas, subastas) y citaba la
  pregunta abierta #1 como pendiente. Se actualizó para reflejar las reglas
  reales (acciones, dividendos, cotizaciones) y su cierre.
- Se corrigió el nombre del producto en toda la documentación: la aplicación
  no se llama "MesaBank", se llama **CashCastor**. Se reemplazó en todos los
  archivos, incluyendo nombres técnicos derivados (namespace de Kubernetes,
  Deployments `*-api`/`*-web`, dominios de ejemplo del OpenAPI).

### Decisiones de alcance (PO, 2026-09-18)
- La lógica funcional de inflación (RF-15) queda fuera del MVP; se requiere
  un POC con una ronda de juego real antes de definir su mecanismo.
- El arqueo de papel moneda (RF-25) solo puede ser declarado por el rol
  Banco; no existe autodeclaración de papel por parte del jugador.
- El ranking global (RF-33) es un leaderboard de mejores partidas
  individuales, sin acumulación por usuario entre partidas.
- Los espectadores ven solo saldos en vivo, sin acceso al historial
  detallado de transacciones.

### Por definir
- Reglas específicas de Tío Rico (fase 2).
- Mecanismo funcional de inflación (fase 2, pendiente de POC).
- Convenciones exactas de `single-iac` (namespaces, secrets, pipelines).
- TTL exacto de JWT y estrategia completa de refresh.
