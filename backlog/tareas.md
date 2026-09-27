# Tablero de Tareas

Listado vivo de trabajo: épicas, features y fixes, cada uno con un ID. Se
actualiza a mano a medida que el equipo avanza — a diferencia de `docs/`
(que documenta el diseño y el alcance planeado), esta carpeta refleja el
**estado actual** del trabajo.

Quién está asignado a cada ID vive en [`asignaciones.md`](./asignaciones.md),
no aquí, para no duplicar esa información en dos lugares.

## Esquema de IDs

- **Épicas (`EP-xx`)** y **features (`US-xxx`)**: reutilizan los IDs que ya
  existen en `docs/06-epicas.md` y `docs/07-historias-usuario.md`. No se
  inventan IDs nuevos para trabajo que ya estaba planeado ahí.
- **Tareas técnicas de fundación (`T-xx`)**: reutilizan los IDs de
  `docs/10-backlog-sprint-0.md`.
- **Fixes (`FIX-xxx`)**: numeración nueva y consecutiva, propia de esta
  carpeta. Los fixes no se planean de antemano (surgen durante el
  desarrollo), así que no tienen equivalente en `docs/`.

## Estados

`Por hacer` → `En progreso` → `En revisión` (PR abierto) → `Hecho`

---

## Tareas técnicas de Sprint 0 (`T-xx`)

Detalle completo de cada una en `docs/10-backlog-sprint-0.md`.

| ID | Título | Estado |
|----|--------|--------|
| T-01 | Setup del repositorio backend (Go) | Hecho — estructura de capas (`internal/api,middleware,ws,db,models`), `golangci-lint` configurado y en CI |
| T-02 | Setup del repositorio frontend (React + Vite + PWA) | Hecho — PWA, store (Zustand), ESLint/Prettier, Vitest, todo en CI |
| T-03 | Definición y publicación de la OpenAPI spec | Por hacer |
| T-04 | Setup de PostgreSQL y Redis en local/desarrollo | En progreso — `docker-compose.yml` listo (sintaxis validada, no se pudo levantar en este entorno por permisos de Docker) |
| T-05 | Modelo de datos y migraciones iniciales | En progreso — modelos GORM de las 7 entidades y flag `--migrations-only` listos (ADR-008); falta seed de datos de desarrollo |
| T-06 | Esqueleto de capas backend | Hecho — 28 rutas del OpenAPI como stubs, middleware de auth (placeholder), hub de WebSocket base |
| T-07 | Cliente API + WebSocket base en el frontend | Por hacer |
| T-08 | Coordinación con `single-iac` | Por hacer |
| T-09 | Manifiestos de Kubernetes | Por hacer |
| T-10 | Pipeline de CI/CD base en `single-iac` | Por hacer |
| T-11 | Observabilidad base | Hecho — logging JSON con request_id/game_code, `/readyz` (chequea Postgres), `/metrics` (Prometheus) |
| T-12 | Documentación de entornos | Por hacer |
| T-13 | Quality gate de PRs en GitHub Actions | Hecho — 4 checks activos en `develop` |

---

## Épicas (`EP-xx`)

Detalle completo de cada una en `docs/06-epicas.md`.

| ID | Título | Estado |
|----|--------|--------|
| EP-01 | Autenticación y Sesión | Por hacer |
| EP-02 | Gestión de Partidas | Por hacer |
| EP-03 | Configuración de Partida | Por hacer |
| EP-04 | Operaciones Financieras | Por hacer |
| EP-05 | Tiempo Real y Turnos | Por hacer |
| EP-06 | Ranking y Puntuación | Por hacer |
| EP-07 | Administración | Por hacer |

## Features (`US-xxx`)

Detalle completo, criterios de aceptación en Gherkin y notas técnicas de
cada una en `docs/07-historias-usuario.md`. Las historias marcadas
"Fuera del MVP" (US-200 en adelante) no se incluyen aquí — ver esa sección
en `07-historias-usuario.md` si hace falta retomarlas en fase 2.

Sprint y talla (T-shirt: XS=1, S=2, M=3, L=5, XL=8) de EP-01 a EP-03
definidos en la planeación de Sprint 1; el reparto por tarea BE/FE, sus
dependencias y el orden están en [`asignaciones.md`](./asignaciones.md).
En Sprint 1, US-033 y US-035 cubren solo la configuración; su efecto llega
con US-050 y US-052. Las demás épicas se dimensionan en su planeación.

| ID | Épica | Prioridad | Título | Sprint | Talla BE | Talla FE | Estado |
|----|-------|-----------|--------|--------|----------|----------|--------|
| US-001 | EP-01 | Must | Registro de usuario | 1 | M | S | Por hacer |
| US-002 | EP-01 | Must | Login con usuario y PIN | 1 | M | M | Por hacer |
| US-003 | EP-01 | Must | Recuperación de sesión tras desconexión | 2 | S | M | Por hacer |
| US-004 | EP-01 | Must | Acceso como invitado | 1 | S | S | Por hacer |
| US-005 | EP-01 | Should | Cierre de sesión | 1 | S | XS | Por hacer |
| US-010 | EP-02 | Must | Crear partida | 1 | M | S | Por hacer |
| US-011 | EP-02 | Must | Generar código, QR y link de invitación | 1 | S | M | Por hacer |
| US-012 | EP-02 | Must | Unirse a partida por código, QR o link | 1 | L | M | Por hacer |
| US-013 | EP-02 | Must | Soporte de múltiples partidas simultáneas | 2 | L | — | Por hacer |
| US-014 | EP-02 | Must | Reinicio de partida con confirmaciones múltiples | 2 | M | M | Por hacer |
| US-015 | EP-02 | Must | Expulsión de jugador en bancarrota | 2 | S | S | Por hacer |
| US-016 | EP-02 | Should | Consultar información de la partida | 1 | S | M | Por hacer |
| US-017 | EP-02 | Must | Unirse a partida como espectador | 1 | S | S | Por hacer |
| US-030 | EP-03 | Must | Definir masa monetaria total | 1 | S | M | Por hacer |
| US-031 | EP-03 | Must | Definir asignación inicial por jugador | 1 | S | XS | Por hacer |
| US-032 | EP-03 | Must | Definir moneda y tasa de cambio | 1 | S | M | Por hacer |
| US-033 | EP-03 | Should | Definir fee de transferencia | 1 | XS | XS | Por hacer |
| US-034 | EP-03 | Must | Definir interés por ronda o por tiempo | 1 | XS | S | Por hacer |
| US-035 | EP-03 | Should | Definir deuda máxima como múltiplo del valor inicial | 1 | XS | XS | Por hacer |
| US-036 | EP-03 | Could | Guardar valor de inflación (informativo, sin efecto funcional) | 2 | XS | XS | Por hacer |
| US-050 | EP-04 | Must | Transferencia P2P idempotente | — | — | — | Por hacer |
| US-051 | EP-04 | Must | Consignación de papel moneda a cuenta digital | — | — | — | Por hacer |
| US-052 | EP-04 | Must | Préstamo del banco a jugador | — | — | — | Por hacer |
| US-053 | EP-04 | Must | Pago automático de intereses | — | — | — | Por hacer |
| US-054 | EP-04 | Must | Registro de deuda con el banco | — | — | — | Por hacer |
| US-055 | EP-04 | Must | Liquidación de bancarrota | — | — | — | Por hacer |
| US-056 | EP-04 | Must | Arqueo de caja en tiempo real (declaración del Banco) | — | — | — | Por hacer |
| US-057 | EP-04 | Must | Historial de ingresos y egresos por jugador | — | — | — | Por hacer |
| US-058 | EP-04 | Must | Auditoría completa para el Banco | — | — | — | Por hacer |
| US-080 | EP-05 | Must | Sincronización en tiempo real vía WebSocket | — | — | — | Por hacer |
| US-081 | EP-05 | Must | Asignación y avance de turnos | — | — | — | Por hacer |
| US-082 | EP-05 | Must | Cálculo de interés por ronda | — | — | — | Por hacer |
| US-083 | EP-05 | Must | Cálculo de interés por tiempo | — | — | — | Por hacer |
| US-084 | EP-05 | Must | Vista de espectador (solo saldos en vivo) | — | — | — | Por hacer |
| US-100 | EP-06 | Must | Cálculo de puntuación final de partida | — | — | — | Por hacer |
| US-101 | EP-06 | Must | Leaderboard global de mejores partidas | — | — | — | Por hacer |
| US-102 | EP-06 | Must | Exclusión de invitados del ranking | — | — | — | Por hacer |
| US-110 | EP-07 | Could | Gestión de monedas soportadas | — | — | — | Por hacer |
| US-111 | EP-07 | Could | Gestión de parámetros globales por defecto | — | — | — | Por hacer |

## Fixes (`FIX-xxx`)

Sin fixes registrados todavía. Se agregan aquí a medida que surgen durante
el desarrollo (no se planean de antemano), con ID consecutivo:

| ID | Título | Estado |
|----|--------|--------|
| — | — | — |
