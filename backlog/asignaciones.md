# Asignaciones

Quién está trabajando en cada ID de [`tareas.md`](./tareas.md). Se actualiza
a mano; si un ID no aparece acá, todavía no tiene dueño asignado.

## Equipo

Todos los devs son fullstack — el frente asignado abajo es dónde van a
enfocarse en el sprint actual, no una limitación de skill.

| Dev | Rol | Frente en Sprint 1 |
|-----|-----|---------------------|
| Erick | Product Owner / DevOps / Infra / Arquitectura | DevOps, infraestructura, arquitectura (sigue cerrando Sprint 0) |
| Harri | Desarrollo (fullstack) | Backend — EP-01, EP-03 (+ US-005 FE) |
| Aleja | Desarrollo (fullstack) | Frontend — EP-01, EP-03 (+ US-034 BE) |
| Cris | Desarrollo (fullstack) | Backend — EP-02, US-033/035 (+ US-017 FE) |
| Andrety | Desarrollo (fullstack) | Frontend — EP-02, US-033/035 (+ US-011 BE) |

## Sprint 0 (tareas técnicas `T-xx`) — corregido 2026-09-27

**Corrección:** Sprint 0 son las fundaciones arquitectónicas (setup de
repos, migraciones, CI/CD, Kubernetes) y las venía haciendo Erick
directamente, no el equipo de devs. La versión anterior de esta tabla los
tenía asignados por error; se corrige acá. El detalle de qué falta de
Sprint 0 está al final de este archivo.

| ID | Asignado a | Estado |
|----|------------|--------|
| T-01 | Erick | Hecho |
| T-02 | Erick | Hecho |
| T-03 | Erick | Hecho |
| T-04 | Erick | En progreso — `docker-compose.yml` listo, sin verificar levantado (permiso de Docker) |
| T-05 | Erick | En progreso — modelos GORM y flag `--migrations-only` listos; falta seed de datos |
| T-06 | Erick | Hecho |
| T-07 | Erick | Hecho |
| T-08 | Erick | En revisión — resuelto por el agente de `single-iac`, pendiente merge/apply |
| T-09 | Erick | En revisión — mismo estado que T-08 |
| T-10 | Erick | En revisión — `deploy.yml` recibido de `single-iac`, sin probar, esperando tu aprobación de merge |
| T-11 | Erick | Hecho |
| T-12 | Erick | Parcial — local documentado; dev/qa/prod no aplica (ADR-009) |
| T-13 | Erick | Hecho |

## Sprint 1 (Fase 1: Core Bancario — `EP-01`, `EP-02`, `EP-03`) — semanas 3 a 5

Reparto **por tarea** (cada historia partida en BE y FE), con talla
T-shirt para medir carga. Reemplaza el reparto por épica anterior, que
dejaba a Harri/Aleja con 12 historias y a Cris/Andrety con 8.

### Escala de tallas

La talla mide **complejidad relativa**, no días (ver convención de
`docs/07-historias-usuario.md`). Para sumar carga se usa esta equivalencia:

| XS | S | M | L | XL |
|----|---|---|---|----|
| 1 | 2 | 3 | 5 | 8 |

### Reglas del reparto

- Todos son fullstack: cada dev tiene **una historia completa de punta a
  punta (★)**, tomando la tarea del otro frente. El intercambio es
  simétrico: los dos de backend toman algo de frontend y los dos de
  frontend toman algo de backend.
- EP-03 en este sprint es **solo configuración** (campo + validación). El
  efecto de fee (US-033) y deuda máxima (US-035) llega con US-050 y US-052
  en Sprint 2.
- Harri es dueño del paquete de validación de `GameConfig`; Cris solo lo
  invoca desde el handler de `POST /games` (US-010), para no pisarse en el
  mismo código.

### Tareas

| Historia | Tarea | Talla | Dueño | Depende de |
|----------|-------|-------|-------|------------|
| **EP-01** | | | | |
| US-001 Registro | BE: `POST /auth/register`, hash del PIN, validación 4-6 dígitos, 409 | M | Harri | — |
| | FE: pantalla de registro y validaciones | S | Aleja | T-07 (o mocks) |
| US-002 Login | BE: `POST /auth/login`, emisión de JWT, middleware de auth real (reemplaza el placeholder) | M | Harri | US-001 BE |
| | FE: pantalla de login, persistencia del token, rutas protegidas | M | Aleja | T-02 (store), T-07 |
| US-004 Invitado | BE: `POST /auth/guest`, JWT de invitado, rechazo al crear partida | S | Harri | US-002 BE |
| | FE: pantalla de nombre de invitado | S | Aleja | US-002 FE |
| US-005 Logout ★ (Should) | BE: blacklist de JWT en Redis, revisada en el middleware | S | Harri | US-002 BE, T-04 (Redis) |
| | FE: botón de logout y limpieza de estado | XS | **Harri ★** | US-002 FE |
| **EP-02** | | | | |
| US-010 Crear partida | BE: `POST /games`, rol Banco, persistir `GameConfig` | M | Cris | US-002 BE (arranca con el placeholder) |
| | FE: pantalla de crear partida (contenedor del formulario de config) | S | Andrety | T-07 (o mocks) |
| US-011 Código/QR/link ★ | BE: código único entre partidas activas, link de invitación | S | **Andrety ★** | — (función aislada) |
| | FE: mostrar código, generar QR, compartir link | M | Andrety | US-010 FE |
| US-012 Unirse | BE: `POST /games/{code}/join`, rol Jugador, asignación inicial (CA-01), errores | L | Cris | US-010 BE, US-011 BE, US-031 BE |
| | FE: unirse por código y ruta `/join/:code` (el QR abre el link, sin escáner dentro de la app) | M | Andrety | US-011 FE |
| US-016 Info de partida (Should) | BE: `GET /games/{code}` y `GET /games/{code}/players` sin datos sensibles | S | Cris | US-010 BE |
| | FE: vista de lobby con lista de jugadores | M | Andrety | US-012 FE |
| US-017 Espectador ★ | BE: rol Espectador y bloqueo de operaciones para ese rol | S | Cris | US-012 BE |
| | FE: elegir el rol al unirse y vista de solo lectura | S | **Cris ★** | US-012 FE, US-016 FE |
| **EP-03** | | | | |
| US-030 Masa monetaria | BE: validación, inmutabilidad, no iniciar sin config | S | Harri | US-010 BE |
| | FE: formulario base de configuración | M | Aleja | US-010 FE |
| US-031 Asignación inicial | BE: la suma de asignaciones no supera la masa total | S | Harri | US-030 BE |
| | FE: campo del formulario | XS | Aleja | US-030 FE |
| US-032 Moneda | BE: validar USD/COP/EUR o personalizada con tasa | S | Harri | US-030 BE |
| | FE: selector de moneda y helper de formato de montos para toda la app | M | Aleja | US-030 FE |
| US-034 Interés ★ | BE: validar tipo, tasa e intervalo | XS | **Aleja ★** | US-030 BE |
| | FE: campos condicionales según el tipo de interés | S | Aleja | US-030 FE |
| US-033 Fee (Should) | BE: validación, valor por defecto 0 | XS | Cris | US-030 BE |
| | FE: campo | XS | Andrety | US-030 FE |
| US-035 Deuda máxima (Should) | BE: validación del múltiplo | XS | Cris | US-030 BE |
| | FE: campo | XS | Andrety | US-030 FE |

### Carga por dev

| Dev | Puntos | Historia completa (★) |
|-----|--------|------------------------|
| Harri | 17 | US-005 |
| Aleja | 17 | US-034 |
| Cris | 16 | US-017 |
| Andrety | 15 | US-011 |
| **Total** | **65** | |

Es el primer sprint y no hay velocidad histórica: 65 puntos es una apuesta.
Se revisa a mitad del sprint. Si hay que recortar, lo primero que sale es
US-035, luego US-033 y luego US-005 (todas Should).

### Orden sugerido

1. **Semana 1:** Harri hace US-001 y US-002 BE (el middleware de auth real
   desbloquea a todos). Cris arranca US-010 BE sobre el placeholder de
   auth. Andrety hace US-011 BE y US-010 FE. Aleja hace las pantallas de
   auth contra mocks si T-07 no está listo.
2. **Semana 2:** Harri hace US-030, US-031 y US-032 BE. **US-031 BE tiene
   que estar antes de que Cris cierre US-012.** Cris hace US-012 y US-016
   BE. Aleja arma el formulario de config. Andrety hace US-011 FE y
   US-012 FE.
3. **Semana 3:** US-004, US-017, US-034 y las Should (US-005, US-033,
   US-035).

### Riesgos

- ~~El frontend depende de T-02 (store) y T-07 (cliente API)~~ **Resuelto
  (2026-09-27):** ambas listas (`frontend/src/store`,
  `frontend/src/api`, `frontend/src/ws`). Aleja y Andrety ya pueden
  consumirlas desde la semana 1, sin mocks.
- US-005 necesita Redis (T-04), que todavía no se ha verificado levantado.

Historias, criterios en Gherkin y notas técnicas:
`docs/07-historias-usuario.md`. Estado de cada tarea: `backlog/tareas.md`.

## Sprint 2 (movidas desde Sprint 1)

Sin dueño todavía; se asignan en la planeación de Sprint 2 junto con EP-04.

| Historia | Motivo | Talla (BE / FE) |
|----------|--------|------------------|
| US-003 Recuperación de sesión | Depende de US-080 (reconexión WebSocket, EP-05) | S / M |
| US-013 Partidas simultáneas | Depende de US-080 y de pruebas de carga | L / — |
| US-014 Reinicio con confirmaciones | El PO todavía no define el flujo; toca el arqueo (CA-04) | M / M |
| US-015 Expulsión por bancarrota | Depende de US-055 (EP-04) | S / S |
| US-036 Inflación informativa | Es Could y no tiene efecto funcional | XS / XS |

## Fase 2 en adelante (`EP-04` a `EP-07`)

Todavía sin asignar.

| ID | Asignado a |
|----|------------|
| EP-04 (Operaciones Financieras) | — |
| EP-05 (Tiempo Real y Turnos) | — |
| EP-06 (Ranking y Puntuación) | — |
| EP-07 (Administración) | — |

## Pendiente de Sprint 0 (a cargo de Erick)

Actualizado 2026-09-27. Hechas: `T-01`, `T-02`, `T-03`, `T-06`, `T-07`,
`T-11`, `T-13`. En progreso: `T-04`, `T-05`, `T-10`. En revisión (por
Erick, en `single-iac`): `T-08`, `T-09`. Parcial: `T-12`.

- **T-04** (Postgres/Redis local): `docker-compose.yml` + `.env.example`
  listos (sintaxis validada); sin verificar levantado por permisos de
  Docker en este entorno.
- **T-05** (migraciones): modelos GORM de las 7 entidades y flag
  `--migrations-only` listos (ADR-008); falta el seed de datos de
  desarrollo.
- **T-08 / T-09** (convenciones, manifiestos K8s): el agente de `single-iac`
  ya entregó todo en su rama `add-cashcastor-namespace` (commit
  `90baedb`) — **nada aplicado al cluster real**. Pendiente de Erick:
  revisar el diff, mergear en `single-iac`, y correr `terraform-apply.yml`
  a mano. Ver ADR-009 para el resumen y las dos desviaciones que propuso
  (nombres `api`/`web` en vez de `cashcastor-api`/`cashcastor-web`, y la
  redefinición de T-10).
- **T-10** (workflow de build+push+deploy): redefinido — vive en este
  repo, no en `single-iac` (ADR-009). El agente de `single-iac` va a
  redactar el contenido del workflow (conoce el patrón exacto de otras
  apps), pero el archivo se agrega acá. Todavía no llegó.
- **T-12** (documentación de entornos): local ya documentado en
  `README.md`; dev/qa/prod no aplica en este proyecto por ahora (decisión
  explícita, ADR-009), no es una tarea pendiente de verdad.
