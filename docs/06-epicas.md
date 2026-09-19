# Épicas: MesaBank MVP

Épicas derivadas del PRD (`02-prd.md`), agrupadas por bloque funcional. Cada
épica lista los requisitos funcionales (RF-xx) que cubre, criterios de
aceptación de alto nivel y dependencias con otras épicas.

## Resumen

| ID | Título | RF cubiertos | Fase del roadmap | Dependencias |
|----|--------|--------------|-------------------|---------------|
| EP-01 | Autenticación y Sesión | RF-01 a RF-04 | Fase 1 | Ninguna |
| EP-02 | Gestión de Partidas | RF-05 a RF-11 | Fase 1 | EP-01 |
| EP-03 | Configuración de Partida | RF-12 a RF-18 | Fase 1 | EP-02 |
| EP-04 | Operaciones Financieras | RF-19 a RF-27 | Fase 2 y 3 | EP-03 |
| EP-05 | Tiempo Real y Turnos | RF-28 a RF-31 | Fase 3 | EP-02, EP-04 |
| EP-06 | Ranking y Puntuación | RF-32 a RF-34 | Fase 5 | EP-04 |
| EP-07 | Administración | (sin RF numerado, ver PRD §2) | Transversal, baja prioridad | EP-01 |

> **Nota:** RF-15 (inflación) está formalmente listado en EP-03 porque el PRD
> lo incluye en la configuración de partida, pero **su lógica funcional queda
> fuera del MVP** (decisión del PO, 2026-09-18): se requiere un POC con una
> ronda de juego real antes de definir el mecanismo. En el MVP, el campo se
> guarda como valor informativo/reservado, sin efecto sobre precios, tasa de
> cambio ni saldos. Ver detalle en `07-historias-usuario.md`, sección "Fuera
> del MVP".

---

## EP-01: Autenticación y Sesión

**Objetivo:** Permitir que un usuario se registre, inicie sesión y mantenga
su sesión activa entre reconexiones, con soporte para acceso como invitado.

**Requisitos funcionales cubiertos:** RF-01, RF-02, RF-03, RF-04.

**Criterios de aceptación de alto nivel:**
- Un usuario puede registrarse con usuario y PIN, y luego iniciar sesión con
  esas mismas credenciales.
- Si la conexión se pierde y se recupera, el usuario retoma la partida donde
  quedó sin tener que re-autenticarse manualmente (dentro de la vigencia de
  su sesión).
- Un invitado puede unirse a una partida sin registro previo, pero no puede
  crear partidas (rol Banco) ni aparecer en el ranking global.

**Dependencias:** Ninguna (épica base).

---

## EP-02: Gestión de Partidas

**Objetivo:** Permitir crear, descubrir y unirse a partidas, asignar roles y
administrar el ciclo de vida de una partida (reinicio, expulsión).

**Requisitos funcionales cubiertos:** RF-05, RF-06, RF-07, RF-08, RF-09,
RF-10, RF-11.

**Criterios de aceptación de alto nivel:**
- El creador de una partida recibe automáticamente el rol Banco.
- Otros usuarios pueden unirse por código, QR o link, y el sistema soporta
  múltiples partidas activas simultáneamente sin interferencia entre ellas.
- El banco puede expulsar a un jugador en bancarrota.
- El reinicio de una partida requiere al menos tres confirmaciones antes de
  ejecutarse (CA-04).

**Dependencias:** EP-01 (requiere usuario autenticado o invitado).

---

## EP-03: Configuración de Partida

**Objetivo:** Dar al banco control total sobre los parámetros económicos de
la partida antes y durante su desarrollo: masa monetaria, moneda, comisiones
e intereses.

**Requisitos funcionales cubiertos:** RF-12, RF-13, RF-14, RF-15 (reservado,
fuera del MVP funcional), RF-16, RF-17, RF-18.

**Criterios de aceptación de alto nivel:**
- El banco define la masa monetaria total (digital + papel) y la asignación
  inicial por jugador antes de iniciar la partida.
- El banco define moneda, tasa de cambio, fee de transferencia (default 0) e
  interés (por ronda o por tiempo) como parte de la configuración inicial.
- El banco define una deuda máxima como múltiplo del valor inicial por
  jugador.
- El campo de inflación se puede guardar en la configuración pero no altera
  ningún cálculo del sistema en el MVP (ver nota en el resumen).

**Dependencias:** EP-02 (la configuración pertenece a una partida ya creada).

---

## EP-04: Operaciones Financieras

**Objetivo:** Ejecutar todas las operaciones de dinero dentro de una
partida —transferencias, consignaciones, préstamos, intereses, liquidación—
de forma segura, idempotente y auditable, y dar al banco visibilidad total
mediante el arqueo de caja.

**Requisitos funcionales cubiertos:** RF-19, RF-20, RF-21, RF-22, RF-23,
RF-24, RF-25, RF-26, RF-27.

**Criterios de aceptación de alto nivel:**
- Una transferencia P2P con la misma `idempotency_key` reintentada nunca
  duplica el débito/crédito (ADR-002).
- Una consignación reduce el papel moneda declarado del jugador y aumenta su
  saldo digital.
- El banco puede otorgar préstamos con interés y plazo, y el sistema calcula
  automáticamente el pago de intereses según la configuración de la partida.
- El arqueo de caja en tiempo real es una función exclusiva del banco: **el
  banco declara manualmente el papel moneda en caja y el estimado por cada
  jugador** (decisión del PO, 2026-09-18); no existe una acción de "declarar
  papel" para el jugador. El sistema compara lo declarado contra la masa
  total configurada y señala si cuadra o no.
- Cada jugador puede ver su propio historial de ingresos y egresos; el banco
  tiene acceso a la auditoría completa de todas las operaciones de la
  partida.
- La liquidación de un jugador en bancarrota paga a los acreedores de forma
  inmediata y lo deja fuera de la partida (CA-03).

**Dependencias:** EP-03 (necesita la configuración económica de la partida).

---

## EP-05: Tiempo Real y Turnos

**Objetivo:** Sincronizar el estado de la partida entre todos los
dispositivos conectados en tiempo real y gestionar el flujo de turnos que
dispara el cálculo de intereses por ronda.

**Requisitos funcionales cubiertos:** RF-28, RF-29, RF-30, RF-31.

**Criterios de aceptación de alto nivel:**
- Cuando ocurre una operación financiera, todos los dispositivos conectados
  a la partida ven el cambio de saldo sin necesidad de refrescar (CA-02).
- El sistema permite asignar y avanzar turnos desde la app.
- El interés por ronda se calcula automáticamente cuando el turno regresa al
  jugador; el interés por tiempo se calcula cada N minutos configurados.

**Dependencias:** EP-02 (partida activa), EP-04 (las operaciones que se
sincronizan y el cálculo de intereses).

---

## EP-06: Ranking y Puntuación

**Objetivo:** Calcular una puntuación final por partida y exponer un
leaderboard global para usuarios registrados.

**Requisitos funcionales cubiertos:** RF-32, RF-33, RF-34.

**Criterios de aceptación de alto nivel:**
- Al finalizar una partida, cada jugador registrado recibe una puntuación
  igual a su patrimonio neto (digital + papel declarado − deuda) en el
  momento de la liquidación (assumption #3).
- El ranking global es un **leaderboard de las mejores puntuaciones de
  partidas individuales** (decisión del PO, 2026-09-18): no se acumula ni se
  promedia el desempeño histórico de un mismo usuario entre partidas.
- Los invitados no aparecen en el ranking global.

**Dependencias:** EP-04 (necesita el cálculo de patrimonio neto y el cierre
de partida).

---

## EP-07: Administración

**Objetivo:** Dar al Admin general una interfaz mínima para gestionar
parámetros globales de la aplicación (monedas soportadas, valores por
defecto), sin ser prioridad para el MVP (assumption #5).

**Requisitos funcionales cubiertos:** Ninguno con numeración RF explícita en
el PRD; se deriva de la tabla de roles (PRD §2, rol "Admin general").

**Criterios de aceptación de alto nivel:**
- Un Admin general puede ver y editar la lista de monedas soportadas y sus
  parámetros por defecto.
- Esta funcionalidad no bloquea ninguna otra épica del MVP: una partida
  puede crearse y jugarse completa sin que exista aún un panel de admin.

**Dependencias:** EP-01 (requiere un rol de Admin autenticado). Épica
transversal, de baja prioridad — ver MoSCoW en `07-historias-usuario.md`.
