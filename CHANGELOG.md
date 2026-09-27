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
- `README.md`: el árbol de estructura y "Convenciones" referían al repo como
  `cashcastor-docs`, mientras el repositorio real es `mesabank-cash-castor-app`.
  Se corrigió y se documentó explícitamente por qué el nombre del repo no
  coincide con el del producto.
- `supuestos-y-preguntas-abiertas.md`: las preguntas abiertas #3 (inflación)
  y #4 (espectadores) seguían listadas como sin resolver, pero ya habían
  sido decididas por el PO (ver "Decisiones de alcance" abajo). Se marcaron
  como resueltas con referencia a las historias correspondientes.
- `08-openapi.yaml`: los schemas `User` y `Game` no dejaban explícito que
  difieren intencionalmente del modelo de datos del PRD §5 (`User` omite
  `pin_hash` por seguridad; `Game.config` expone un objeto `GameConfig`
  estructurado en vez del `config_json` crudo del PRD). Se documentó la
  diferencia en la descripción de cada schema.

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

## [0.1.0-scaffold] - 2026-09-18

### Añadido
- Scaffold mínimo de `backend/` (Go + Gin, endpoint `GET /healthz`) y
  `frontend/` (Vite + React + TypeScript), con `Dockerfile` multi-stage para
  cada uno (T-01/T-02 del backlog de Sprint 0).
- Workflow `.github/workflows/quality-checks.yml`: quality gate de PRs hacia
  `develop` con cuatro checks (compilación de backend y frontend, build de
  ambas imágenes Docker), condicionados a qué carpeta cambió en el PR
  (T-13 del backlog de Sprint 0).
- `.github/CODEOWNERS` asignando a `@esalamancar` como revisor obligatorio
  de todo el repositorio.
- Rama `develop` como rama de trabajo, protegida: solo se integra vía PR con
  al menos una aprobación de code owner y los checks de calidad en verde.
- Acceso `Write` (equivalente a "developer") para `2alelagos-dot`, `LoBe4`
  y `nCrisz`.
- Documentado el flujo de trabajo de Git y los checks requeridos en
  `README.md`.

### Cambiado
- El repositorio pasó de privado a **público**: la API de branch
  protection/rulesets de GitHub rechaza ambas con 403 en repos privados de
  cuentas personales Free ("Upgrade to GitHub Pro or make this repository
  public"). Decisión del PO (2026-09-18): pasar a público en vez de pagar
  el upgrade, dado que el repo aún no tiene código de negocio.
- Rama por defecto cambiada de `main` a `develop`.
- Se quitaron las estimaciones de tiempo/esfuerzo de `10-backlog-sprint-0.md`
  (días-persona) y `07-historias-usuario.md` (S/M/L). Decisión del PO
  (2026-09-19): cuánto tarda una tarea depende de quién la haga
  (herramientas, experiencia), no es una propiedad fija de la tarea —
  distinto de estimar costo. Se conserva el orden de dependencia (backlog)
  y la prioridad MoSCoW (historias), que sí son decisiones de negocio.

## [Unreleased] (continuación)

### Añadido
- Carpeta `backlog/`: tablero de trabajo vivo, separado de `docs/`.
  - `backlog/tareas.md`: todas las épicas (`EP-xx`), features (`US-xxx`,
    reutilizando los IDs de `07-historias-usuario.md`) y las tareas de
    Sprint 0 (`T-xx`, de `10-backlog-sprint-0.md`), con su estado actual.
    Incluye una sección `FIX-xxx` vacía para fixes no planeados que surjan
    durante el desarrollo.
  - `backlog/asignaciones.md`: quién está trabajando en cada ID. Equipo de
    desarrollo confirmado: Harri, Aleja, Andrety, Cris (más Erick como PO/
    DevOps). Reparto inicial de Sprint 0 aplicado; las épicas de fase 1 en
    adelante quedan sin asignar hasta que arranque cada fase.
- `README.md` actualizado: equipo de desarrollo (antes "por definir"),
  árbol de estructura completo (`backend/`, `frontend/`, `backlog/`,
  `.github/`, que faltaban desde que se agregaron).

## [Unreleased] (continuación 2)

### Añadido
- Git-flow formal de 3 ramas: `develop → qa → prod`. Solo `develop` recibe
  código nuevo (vía PR de los devs); `qa` y `prod` son ramas de promoción
  (solo reciben merges desde la rama anterior en la cadena, nunca código
  directo). Las tres protegidas igual: sin push directo, sin force-push,
  sin borrado, PR obligatorio, 1 aprobación de code owner (`@esalamancar`,
  el único que revisa/aprueba), y los mismos checks de CI obligatorios.
- Check de CI `Branch name convention`: falla el PR si la rama origen no
  sigue el patrón `tipo/descripcion-en-kebab-case` (tipos: `feature`, `fix`,
  `chore`, `docs`, `refactor`, `test`, `hotfix`). Se agrega como check
  obligatorio en las tres rulesets.

### Cambiado
- Se retiró la rama `main` (su historia ya estaba completa dentro de
  `develop`; no se perdió nada). `develop` sigue siendo la rama por defecto.
- Se eliminó la rama `front`, creada sin seguir la convención de nombres
  (sin commits propios, no se perdió trabajo).

### Corregido / Limitación de plan encontrada
- La regla nativa de GitHub para forzar nombres de rama
  (`branch_name_pattern`) requiere GitHub Team/Enterprise — no está
  disponible en el plan Free ni siquiera en repos públicos (a diferencia de
  `pull_request`/`required_status_checks`, que sí funcionan en público).
  Se resolvió con el check de CI `Branch name convention` en vez de la
  regla nativa: mismo efecto práctico (bloquea el merge), sin necesidad de
  upgrade de plan.

## [Unreleased] (continuación 3)

### Cambiado
- `backlog/asignaciones.md`: reparto de Sprint 0 corregido por frente real
  (2026-09-27). Todos los devs son fullstack; para este sprint Harri y Cris
  se enfocan en backend, Aleja y Andrety en frontend, Erick en DevOps/
  infraestructura/arquitectura (además de PO). Reemplaza el reparto
  anterior, que tenía a Aleja y Andrety mezclados entre backend/frontend.

## [Unreleased] (continuación 4)

### Cambiado
- `README.md` reorganizado como punto de entrada real para el equipo:
  sección nueva "Dónde encontrar todo" (tabla docs/ vs backlog/ vs backend/
  vs frontend/, con cuándo mirar cada uno) y "Cómo empezar a desarrollar"
  (comandos para correr backend y frontend local y con Docker). Se agregó
  un banner al inicio guiando el orden de lectura para un dev nuevo, y se
  actualizó "Roles" con el frente (backend/frontend/DevOps) de cada dev.

## [Unreleased] (continuación 5)

### Corregido
- `backlog/asignaciones.md`: Sprint 0 (fundaciones arquitectónicas) las
  venía haciendo Erick directamente, no el equipo de devs — la asignación
  anterior los tenía puestos ahí por error. Se corrige: Sprint 0 completo
  queda a cargo de Erick, y los 4 devs pasan a Sprint 1 (Fase 1: Core
  Bancario). Se documenta el estado real pendiente de Sprint 0 (todo
  excepto T-13).

### Añadido
- Reparto de Sprint 1 (`EP-01`, `EP-02`, `EP-03`) en `backlog/asignaciones.md`:
  parejas backend+frontend por épica — Harri+Aleja en EP-01/EP-03, Cris+
  Andrety en EP-02.

## [Unreleased] (continuación 6)

### Añadido
- `docker-compose.yml` + `.env.example`: PostgreSQL y Redis para desarrollo
  local (T-04). Sintaxis validada (`docker compose config`); no se pudo
  levantar en este entorno por permisos de Docker del sandbox.
- Modelos GORM para las 7 entidades del PRD §5 (`backend/internal/models`)
  y flag de arranque `--migrations-only` en el backend, que corre
  `AutoMigrate` y termina (T-05).
- `ADR-008`: se decide GORM + `AutoMigrate` en vez de scripts SQL manuales
  (`golang-migrate`/`goose`), resolviendo la elección que estaba abierta
  entre `sqlc` y `GORM` en `03-arquitectura.md` §2.1.

### Cambiado
- `03-arquitectura.md` y `10-backlog-sprint-0.md` actualizados para
  reflejar ADR-008 (ya no mencionan `sqlc`/`golang-migrate`/`goose` como
  opciones abiertas).
