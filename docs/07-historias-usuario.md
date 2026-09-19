# Historias de Usuario: CashCastor MVP

Historias derivadas de las épicas en `06-epicas.md`, cubriendo las fases 1 a 5
del roadmap (`05-roadmap.md`). Formato: `Como <rol>, quiero <acción>, para
<beneficio>`, con criterios de aceptación en Gherkin.

**Convenciones:**
- **Prioridad (MoSCoW):** Must / Should / Could / Won't (para esta iteración).
- **Estimación:** S (≤1 día), M (2-3 días), L (4-5 días), a nivel de una
  historia ya desglosada en tareas técnicas por el equipo.
- Las decisiones del PO tomadas el 2026-09-18 (inflación fuera del MVP, arqueo
  solo por el banco, ranking tipo leaderboard, espectador solo con saldos) se
  marcan explícitamente donde aplican.

---

## EP-01: Autenticación y Sesión

### US-001: Registro de usuario
- **Épica:** EP-01 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-01

Como jugador nuevo, quiero registrarme con un usuario y un PIN, para poder
identificarme en futuras partidas y aparecer en el ranking.

```gherkin
Escenario: Registro exitoso
  Dado que no existe un usuario con el nombre "erick99"
  Cuando envío mi nombre de usuario "erick99" y un PIN de 4 a 6 dígitos
  Entonces el sistema crea mi cuenta y me devuelve un token de sesión

Escenario: Usuario ya existe
  Dado que ya existe un usuario con el nombre "erick99"
  Cuando intento registrarme con ese mismo nombre
  Entonces el sistema rechaza la operación con un error de conflicto
```

- **Dependencias:** Ninguna.
- **Notas técnicas:** PIN se almacena con hash (bcrypt/argon2), nunca en
  texto plano (RNF-06). Validar longitud de PIN (4-6 dígitos) en backend.

---

### US-002: Login con usuario y PIN
- **Épica:** EP-01 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-02

Como usuario registrado, quiero iniciar sesión con mi usuario y PIN, para
acceder a mis partidas y mi historial.

```gherkin
Escenario: Login exitoso
  Dado un usuario registrado con usuario "erick99" y PIN válido
  Cuando envío esas credenciales al endpoint de login
  Entonces recibo un JWT válido y mis datos de perfil

Escenario: PIN incorrecto
  Dado un usuario registrado "erick99"
  Cuando envío un PIN incorrecto
  Entonces el sistema responde 401 sin revelar si el usuario existe
```

- **Dependencias:** US-001.
- **Notas técnicas:** Emitir JWT de corta duración (ADR-005). El TTL exacto y
  la estrategia de refresh completa quedan pendientes de definición (ver
  `supuestos-y-preguntas-abiertas.md`, pregunta abierta #7); no bloquea el
  desarrollo de esta historia, sí bloquea el detalle fino de US-003.

---

### US-003: Recuperación de sesión tras desconexión
- **Épica:** EP-01 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-03

Como jugador conectado a una partida, quiero que al perder y recuperar
conexión (o al reabrir la app) retome automáticamente donde quedé, para no
perder mi progreso ni tener que reconfigurar nada.

```gherkin
Escenario: Reconexión dentro de la partida activa
  Dado que tengo una sesión válida y estoy en la partida "ABCD"
  Cuando pierdo conexión y vuelvo a abrir la app antes de que expire mi sesión
  Entonces vuelvo automáticamente a la vista de la partida "ABCD" con mi
    saldo y estado actualizados

Escenario: Sesión expirada
  Dado que mi token de sesión expiró
  Cuando reabro la app
  Entonces se me pide iniciar sesión nuevamente
```

- **Dependencias:** US-002, US-080 (reconexión de WebSocket).
- **Notas técnicas:** El cliente debe persistir el token y el código de
  partida en almacenamiento local (Service Worker / localStorage) para
  reintentar la reconexión sin intervención del usuario.

---

### US-004: Acceso como invitado
- **Épica:** EP-01 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-04

Como persona sin cuenta, quiero unirme a una partida como invitado, para
jugar sin necesidad de registrarme.

```gherkin
Escenario: Unirse como invitado
  Dado un código de partida válido
  Cuando ingreso un nombre para mostrar sin registrarme
  Entonces se me asigna una sesión temporal y puedo unirme a la partida como
    Jugador o Espectador

Escenario: Invitado no puede crear partida
  Dado que estoy en sesión de invitado
  Cuando intento crear una partida
  Entonces el sistema rechaza la operación indicando que se requiere cuenta
    registrada
```

- **Dependencias:** Ninguna.
- **Notas técnicas:** La sesión de invitado también usa JWT (de vida corta),
  pero el usuario no se persiste como cuenta permanente. Un invitado no
  aparece en el ranking global (ver US-102).

---

### US-005: Cierre de sesión
- **Épica:** EP-01 | **Prioridad:** Should | **Estimación:** S | **RF:** complementaria, sin número en el PRD

Como usuario, quiero cerrar sesión explícitamente, para proteger mi cuenta
en un dispositivo compartido.

```gherkin
Escenario: Logout exitoso
  Dado que tengo una sesión activa
  Cuando selecciono "cerrar sesión"
  Entonces mi token se invalida y se me redirige a la pantalla de login
```

- **Dependencias:** US-002.
- **Notas técnicas:** Requiere blacklist de JWT en Redis (arquitectura §2.3).
  Historia complementaria no derivada de un RF explícito del PRD; se incluye
  para completar el ciclo de sesión.

---

## EP-02: Gestión de Partidas

### US-010: Crear partida
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-05, RF-08

Como usuario registrado, quiero crear una partida configurando su masa
monetaria inicial, para asumir el rol de Banco y empezar a jugar.

```gherkin
Escenario: Creación exitosa
  Dado que estoy autenticado como usuario registrado
  Cuando creo una partida con una configuración de masa monetaria válida
  Entonces se crea la partida, se me asigna automáticamente el rol Banco y
    recibo un código de sala único

Escenario: Invitado intenta crear partida
  Dado que estoy en sesión de invitado
  Cuando intento crear una partida
  Entonces la operación es rechazada (ver US-004)
```

- **Dependencias:** US-002 (o US-004 para el caso de rechazo).
- **Notas técnicas:** La configuración detallada de masa monetaria se cubre
  en EP-03; esta historia solo cubre la creación del registro `Game` y la
  asignación de rol.

---

### US-011: Generar código, QR y link de invitación
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-06

Como Banco, quiero que mi partida tenga un código corto, un QR y un link,
para invitar a otros jugadores fácilmente.

```gherkin
Escenario: Generación automática al crear la partida
  Dado que acabo de crear una partida
  Entonces el sistema genera un código alfanumérico único, una URL de QR y
    un link de invitación asociados a esa partida
```

- **Dependencias:** US-010.
- **Notas técnicas:** El código debe ser único entre partidas activas; puede
  reciclarse una vez la partida finaliza.

---

### US-012: Unirse a partida por código, QR o link
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-07

Como jugador, quiero unirme a una partida existente usando el código, QR o
link, para participar con el rol de Jugador.

```gherkin
Escenario: Unión exitosa como Jugador
  Dado un código de partida válido y activa
  Cuando me uno con ese código
  Entonces recibo el rol Jugador y mi asignación inicial de digital y papel
    configurada por el Banco (CA-01)

Escenario: Código inválido o partida finalizada
  Dado un código que no corresponde a ninguna partida activa
  Cuando intento unirme
  Entonces el sistema responde con un error claro
```

- **Dependencias:** US-010, US-011, US-030/US-031 (configuración de masa
  monetaria y asignación inicial ya definida).
- **Notas técnicas:** Ver CA-01 del PRD.

---

### US-013: Soporte de múltiples partidas simultáneas
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** L | **RF:** RF-09

Como sistema, debo soportar múltiples partidas activas al mismo tiempo sin
que las operaciones de una interfieran con otra, para que distintos grupos
de jugadores puedan usar la app en paralelo.

```gherkin
Escenario: Aislamiento entre partidas
  Dado dos partidas activas "ABCD" y "WXYZ"
  Cuando se realiza una transferencia en "ABCD"
  Entonces los saldos y eventos de "WXYZ" no se ven afectados ni reciben
    broadcast de esa transacción
```

- **Dependencias:** US-010, US-080 (hub de WebSocket particionado por
  partida).
- **Notas técnicas:** Historia de tipo "enabler" arquitectónico; se valida
  con pruebas de carga (ver `09-plan-de-pruebas.md`).

---

### US-014: Reinicio de partida con confirmaciones múltiples
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-10

Como Banco, quiero reiniciar la partida solo después de varias
confirmaciones, para evitar reinicios accidentales que borren todo el
progreso.

```gherkin
Escenario: Reinicio confirmado
  Dado que el Banco solicita reiniciar la partida
  Cuando se registran al menos tres confirmaciones (del Banco y/o jugadores,
    según defina el flujo de UI)
  Entonces la partida vuelve a su estado inicial de configuración y el
    arqueo de papel moneda coincide con la masa total original (CA-04)

Escenario: Reinicio sin confirmaciones suficientes
  Dado que solo se ha registrado una confirmación
  Cuando se intenta ejecutar el reinicio
  Entonces la operación se rechaza y se informa cuántas confirmaciones faltan
```

- **Dependencias:** US-010.
- **Notas técnicas:** Ver CA-04. El mecanismo exacto de quién puede confirmar
  (solo el Banco repitiendo confirmación, o Banco + jugadores) no está
  definido en el PRD; se propone como diseño mínimo un flujo de
  "solicitud + N confirmaciones" documentado en `08-openapi.yaml`. Confirmar
  con el PO antes de implementar si se requiere un flujo distinto.

---

### US-015: Expulsión de jugador en bancarrota
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-11

Como Banco, quiero expulsar a un jugador declarado en bancarrota, para que
la partida continúe con los jugadores restantes.

```gherkin
Escenario: Expulsión tras liquidación
  Dado un jugador marcado como bancarrota (is_bankrupt = true) tras la
    liquidación (US-055)
  Cuando el Banco confirma la expulsión
  Entonces el jugador pierde acceso a operar en la partida y se notifica a
    todos los conectados vía WebSocket
```

- **Dependencias:** US-055 (liquidación de bancarrota).
- **Notas técnicas:** Solo el rol Banco puede ejecutar esta acción (PRD §2).

---

### US-016: Consultar información de la partida
- **Épica:** EP-02 | **Prioridad:** Should | **Estimación:** S | **RF:** complementaria a RF-07

Como participante de una partida, quiero consultar su estado general
(jugadores conectados, código, estado), para saber si puedo unirme o seguir
jugando.

```gherkin
Escenario: Consulta de partida activa
  Dado un código de partida válido
  Cuando consulto su información
  Entonces recibo su estado, lista de jugadores y configuración pública
    (sin datos sensibles de otros usuarios)
```

- **Dependencias:** US-010.
- **Notas técnicas:** Corresponde a `GET /games/{code}` del PRD §6.

---

### US-017: Unirse a partida como espectador
- **Épica:** EP-02 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-07 (rol Espectador, PRD §2)

Como persona interesada en seguir la partida sin jugar, quiero unirme como
espectador, para ver el desarrollo sin poder operar.

```gherkin
Escenario: Unión como espectador
  Dado un código de partida válido
  Cuando me uno seleccionando el rol Espectador
  Entonces puedo ver la partida en modo solo lectura, sin balance propio ni
    posibilidad de transferir o consignar
```

- **Dependencias:** US-010, US-012.
- **Notas técnicas:** Ver US-084 para el alcance exacto de lo que un
  espectador puede visualizar (solo saldos en vivo, decisión del PO
  2026-09-18).

---

## EP-03: Configuración de Partida

### US-030: Definir masa monetaria total
- **Épica:** EP-03 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-12

Como Banco, quiero definir cuánto dinero digital y cuánto papel moneda
existen en total en la partida, para controlar la emisión y poder hacer
arqueo al final (ADR-003).

```gherkin
Escenario: Configuración válida
  Dado que estoy configurando una partida nueva
  Cuando defino el total de dinero digital y el total de papel moneda
  Entonces esos valores quedan fijos para toda la partida y no pueden
    incrementarse posteriormente

Escenario: Configuración incompleta
  Cuando intento iniciar la partida sin definir la masa monetaria
  Entonces el sistema me impide continuar
```

- **Dependencias:** US-010.
- **Notas técnicas:** Ver ADR-003 (masa monetaria fija).

---

### US-031: Definir asignación inicial por jugador
- **Épica:** EP-03 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-13

Como Banco, quiero definir cuánto dinero digital y papel recibe cada jugador
al unirse, para que el reparto inicial sea consistente.

```gherkin
Escenario: Asignación aplicada al unirse
  Dado una configuración con asignación inicial definida
  Cuando un jugador se une a la partida
  Entonces recibe exactamente esa asignación en digital y papel (CA-01)
```

- **Dependencias:** US-030.
- **Notas técnicas:** Debe validarse que la suma de asignaciones no exceda la
  masa monetaria total configurada.

---

### US-032: Definir moneda y tasa de cambio
- **Épica:** EP-03 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-14

Como Banco, quiero elegir la moneda de la partida (USD, COP, EUR o
personalizada) y su tasa de cambio, para adaptar la partida al contexto de
los jugadores.

```gherkin
Escenario: Selección de moneda estándar
  Cuando configuro la partida con moneda "COP"
  Entonces todos los montos de la partida se muestran y formatean en COP

Escenario: Moneda personalizada
  Cuando configuro una moneda personalizada con su símbolo y tasa de cambio
  Entonces el sistema la acepta y la usa de forma consistente en toda la
    partida
```

- **Dependencias:** US-030.
- **Notas técnicas:** Ninguna.

---

### US-033: Definir fee de transferencia
- **Épica:** EP-03 | **Prioridad:** Should | **Estimación:** S | **RF:** RF-16

Como Banco, quiero definir un fee (comisión) por transferencia, con valor
por defecto 0, para poder simular costos de transacción si lo deseo.

```gherkin
Escenario: Fee por defecto
  Dado que no configuro explícitamente un fee
  Entonces las transferencias se procesan sin comisión

Escenario: Fee configurado
  Dado un fee del 2% configurado
  Cuando un jugador transfiere 100 unidades
  Entonces el receptor recibe 98 y el fee se registra como ingreso del banco
```

- **Dependencias:** US-030.
- **Notas técnicas:** El destino exacto del fee cobrado (¿va a la caja del
  Banco?) se asume como ingreso del Banco por ser el único otro actor
  económico del sistema; confirmar con el PO si se requiere otro tratamiento
  antes de implementar en fase de desarrollo.

---

### US-034: Definir interés por ronda o por tiempo
- **Épica:** EP-03 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-17

Como Banco, quiero definir un porcentaje de interés y su periodicidad (por
ronda o por tiempo), para aplicar mecánicas de préstamos con interés.

```gherkin
Escenario: Interés por ronda
  Cuando configuro interés del 5% "por ronda"
  Entonces el sistema aplicará ese porcentaje cada vez que el turno regrese
    al jugador con deuda activa (ver US-082)

Escenario: Interés por tiempo
  Cuando configuro interés del 5% "por tiempo" cada 10 minutos
  Entonces el sistema aplicará ese porcentaje automáticamente cada 10
    minutos a las deudas activas (ver US-083)
```

- **Dependencias:** US-030.
- **Notas técnicas:** Ver EP-05 para el disparador técnico del cálculo.

---

### US-035: Definir deuda máxima como múltiplo del valor inicial
- **Épica:** EP-03 | **Prioridad:** Should | **Estimación:** S | **RF:** RF-18

Como Banco, quiero limitar la deuda máxima de un jugador a un múltiplo de su
valor inicial, para evitar endeudamiento ilimitado.

```gherkin
Escenario: Préstamo dentro del límite
  Dado un jugador con valor inicial 1500 y deuda máxima configurada en 2x
  Cuando solicita un préstamo que lo llevaría a una deuda de 2500
  Entonces el préstamo se aprueba

Escenario: Préstamo que excede el límite
  Dado el mismo jugador y configuración
  Cuando solicita un préstamo que lo llevaría a una deuda de 3500
  Entonces el sistema rechaza el préstamo indicando el límite máximo
```

- **Dependencias:** US-030, US-052 (préstamos).
- **Notas técnicas:** Ninguna.

---

### US-036: Guardar valor de inflación (informativo, sin efecto funcional)
- **Épica:** EP-03 | **Prioridad:** Could | **Estimación:** S | **RF:** RF-15 (alcance reducido)

Como Banco, quiero poder registrar un valor de inflación en la configuración
de la partida, para dejar constancia de la intención aunque el MVP todavía
no aplique ningún efecto automático.

```gherkin
Escenario: Guardar inflación sin efecto
  Cuando configuro un valor de inflación del 3%
  Entonces el valor se guarda y se muestra en el panel de configuración,
    pero ningún cálculo del sistema (precios, tasa de cambio, saldos,
    intereses) se ve afectado por él
```

- **Dependencias:** US-030.
- **Notas técnicas:** **Decisión del PO (2026-09-18): la lógica funcional de
  inflación queda fuera del MVP por completo.** Se requiere un POC jugando
  una ronda real antes de definir cómo debe comportarse (¿afecta precios de
  propiedades, tasa de cambio, saldos?). Ver sección "Fuera del MVP" más
  abajo (US-201). Esta historia es opcional (Could): si no aporta valor
  guardar un campo sin uso, puede omitirse del MVP sin impacto funcional.

---

## EP-04: Operaciones Financieras

### US-050: Transferencia P2P idempotente
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** L | **RF:** RF-19

Como jugador, quiero transferir dinero digital a otro jugador de forma
instantánea, para pagar rentas u otras obligaciones del juego sin usar papel
físico.

```gherkin
Escenario: Transferencia exitosa
  Dado que tengo saldo suficiente
  Cuando transfiero una cantidad válida a otro jugador de la misma partida
    con una idempotency_key nueva
  Entonces mi saldo se debita, el receptor se acredita, y ambos ven el
    cambio reflejado en tiempo real (CA-02)

Escenario: Reintento con la misma idempotency_key
  Dado que ya procesé una transferencia con la clave "abc-123"
  Cuando reenvío la misma solicitud con la clave "abc-123" (p. ej. por una
    reconexión de red)
  Entonces el sistema devuelve el mismo resultado sin duplicar el débito

Escenario: Saldo insuficiente
  Dado que mi saldo es menor al monto solicitado
  Cuando intento transferir
  Entonces la operación se rechaza y mi saldo no cambia

Escenario: Jugadores de partidas distintas
  Cuando intento transferir a un jugador que no pertenece a mi partida
  Entonces la operación se rechaza
```

- **Dependencias:** US-012, US-030.
- **Notas técnicas:** Ver ADR-002 (idempotencia en Redis, TTL 24h) y el flujo
  descrito en `03-arquitectura.md` §4. Requiere transacción atómica en
  PostgreSQL (débito + crédito + registro).

---

### US-051: Consignación de papel moneda a cuenta digital
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-20

Como jugador, quiero consignar papel moneda físico a mi cuenta digital, para
poder operar digitalmente con dinero que tenía en efectivo.

```gherkin
Escenario: Consignación exitosa
  Dado que declaro tener 500 en papel moneda
  Cuando registro una consignación de 200
  Entonces mi paper_balance baja en 200 y mi digital_balance sube en 200

Escenario: Consignación mayor al papel disponible
  Dado que mi paper_balance registrado es 100
  Cuando intento consignar 200
  Entonces la operación se rechaza
```

- **Dependencias:** US-030, US-031.
- **Notas técnicas:** Esta es la única forma en que cambia el `paper_balance`
  de un jugador en el MVP (junto con la asignación inicial), dado que no
  existe una acción de "declarar papel" independiente (decisión del PO,
  2026-09-18; ver US-056).

---

### US-052: Préstamo del banco a jugador
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-21

Como Banco, quiero otorgar un préstamo a un jugador con monto, interés y
plazo, para permitirle seguir jugando cuando no tiene liquidez.

```gherkin
Escenario: Préstamo otorgado
  Dado un jugador solicitando o el Banco iniciando un préstamo de 300 con
    interés 5% y plazo de 3 rondas
  Cuando el Banco confirma el préstamo
  Entonces el jugador recibe 300 en su saldo digital y se crea un registro
    Loan en estado activo
```

- **Dependencias:** US-030, US-035 (límite de deuda).
- **Notas técnicas:** Ninguna.

---

### US-053: Pago automático de intereses
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** L | **RF:** RF-22

Como jugador con un préstamo activo, quiero que el interés se calcule y
aplique automáticamente según la configuración de la partida, para no tener
que calcularlo manualmente.

```gherkin
Escenario: Interés aplicado automáticamente
  Dado un préstamo activo con interés configurado
  Cuando se cumple el disparador configurado (ronda o tiempo, ver US-082 y
    US-083)
  Entonces el interés correspondiente se suma a la deuda del jugador y se
    registra en su historial
```

- **Dependencias:** US-034, US-052, US-082, US-083.
- **Notas técnicas:** Historia de lógica de negocio compartida entre EP-04
  (regla de cálculo) y EP-05 (disparador de ronda/tiempo).

---

### US-054: Registro de deuda con el banco
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-23

Como Banco, quiero que la deuda de un jugador se refleje como saldo negativo
contable, para tener visibilidad clara de quién debe y cuánto.

```gherkin
Escenario: Deuda visible
  Dado un jugador con un préstamo activo y sin pagar
  Cuando consulto el panel de Banco
  Entonces veo su deuda total como parte de su información financiera
```

- **Dependencias:** US-052.
- **Notas técnicas:** Ninguna.

---

### US-055: Liquidación de bancarrota
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-24

Como Banco, quiero iniciar la liquidación de un jugador en bancarrota, para
que sus acreedores reciban pago inmediato y el jugador quede fuera de la
partida.

```gherkin
Escenario: Liquidación exitosa
  Dado un jugador cuyo patrimonio (digital + papel − deuda) es negativo o
    insuficiente para cubrir sus obligaciones
  Cuando el Banco inicia la liquidación
  Entonces los acreedores reciben el pago disponible de inmediato y el
    jugador queda marcado is_bankrupt = true (CA-03)
```

- **Dependencias:** US-052, US-054.
- **Notas técnicas:** Dispara US-015 (expulsión). El orden exacto de pago
  cuando hay múltiples acreedores y fondos insuficientes no está definido en
  el PRD; se propone prorrateo proporcional a la deuda como diseño mínimo,
  a confirmar con el PO antes de implementar.

---

### US-056: Arqueo de caja en tiempo real (declaración del Banco)
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-25

Como Banco, quiero declarar cuánto papel moneda tengo físicamente en caja y
cuánto estimo que tiene cada jugador, para comparar contra la masa monetaria
total configurada y verificar que todo cuadra.

```gherkin
Escenario: Arqueo cuadrado
  Dado una masa total de papel configurada en 10000
  Cuando el Banco declara 4000 en su caja y estima 6000 repartidos entre los
    jugadores
  Entonces el sistema confirma que el arqueo cuadra (4000 + 6000 = 10000)

Escenario: Arqueo descuadrado
  Dado la misma masa total de 10000
  Cuando el Banco declara 4000 en caja y estima 5000 en jugadores (suma 9000)
  Entonces el sistema señala una diferencia de 1000 y no la resuelve
    automáticamente
```

- **Dependencias:** US-030.
- **Notas técnicas:** **Decisión del PO (2026-09-18): solo el Banco declara
  el arqueo.** No existe una acción de "declarar mi papel" para el jugador;
  el `paper_balance` de cada jugador en el sistema solo cambia por acciones
  registradas en la app (asignación inicial, consignaciones — US-051). El
  arqueo compara la declaración del Banco (caja + estimado por jugador)
  contra la masa total configurada, y contra el `paper_balance` que el
  sistema tiene registrado para cada jugador, para ayudar a detectar
  discrepancias.

---

### US-057: Historial de ingresos y egresos por jugador
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-26

Como jugador, quiero ver mi propio historial de ingresos y egresos, para
llevar control de mis movimientos durante la partida.

```gherkin
Escenario: Consulta de historial propio
  Dado que he realizado transferencias, consignaciones y he recibido
    préstamos
  Cuando consulto mi historial
  Entonces veo cada movimiento con tipo, monto, contraparte y fecha
```

- **Dependencias:** US-050, US-051, US-052.
- **Notas técnicas:** Un jugador solo ve su propio historial, no el de otros
  (a diferencia del Banco, ver US-058).

---

### US-058: Auditoría completa para el Banco
- **Épica:** EP-04 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-27

Como Banco, quiero ver el registro completo de todas las operaciones de la
partida, para poder resolver disputas y verificar la integridad del juego.

```gherkin
Escenario: Consulta de auditoría
  Dado cualquier número de operaciones realizadas en la partida
  Cuando el Banco consulta la auditoría
  Entonces ve todas las transacciones, préstamos, consignaciones y eventos
    de liquidación/expulsión, con detalle y orden cronológico
```

- **Dependencias:** US-050, US-051, US-052, US-055.
- **Notas técnicas:** Corresponde a la tabla `AuditLog` del modelo de datos
  (PRD §5); cada operación relevante debe escribir un registro aquí, no solo
  en `Transaction`.

---

## EP-05: Tiempo Real y Turnos

### US-080: Sincronización en tiempo real vía WebSocket
- **Épica:** EP-05 | **Prioridad:** Must | **Estimación:** L | **RF:** RF-28

Como jugador o espectador conectado a una partida, quiero ver los cambios de
saldo y estado en tiempo real sin refrescar la app, para seguir el juego de
forma fluida.

```gherkin
Escenario: Broadcast de transacción
  Dado dos dispositivos conectados a la misma partida vía WebSocket
  Cuando ocurre una transferencia entre dos jugadores
  Entonces ambos dispositivos (y cualquier espectador conectado) reciben el
    evento y actualizan los saldos visibles sin recargar (CA-02)

Escenario: Reconexión de WebSocket
  Dado que un cliente pierde la conexión WebSocket
  Cuando se reconecta con un token de sesión válido
  Entonces vuelve a unirse al canal de su partida y recibe el estado
    actualizado
```

- **Dependencias:** US-012, US-050.
- **Notas técnicas:** Hub de conexiones por partida + Redis pub/sub
  (`03-arquitectura.md` §2.1, §2.3). Ver US-013 para aislamiento entre
  partidas.

---

### US-081: Asignación y avance de turnos
- **Épica:** EP-05 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-29

Como Banco, quiero gestionar de quién es el turno desde la app, para que el
sistema sepa cuándo aplicar el interés por ronda.

```gherkin
Escenario: Avanzar turno
  Dado un orden de jugadores definido y el turno actual en el jugador A
  Cuando se avanza el turno
  Entonces el turno pasa al siguiente jugador según el orden configurado y
    se notifica a todos los conectados
```

- **Dependencias:** US-012.
- **Notas técnicas:** El orden de jugadores se define en el momento de
  iniciar la partida o por orden de unión; no hay una regla de "orden de
  turno" definida en el PRD más allá de "asignación de turnos desde la app",
  se asume un orden simple (round-robin) como diseño mínimo.

---

### US-082: Cálculo de interés por ronda
- **Épica:** EP-05 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-30

Como jugador con deuda activa, quiero que el interés por ronda se calcule
automáticamente cuando el turno vuelve a mí, para que el juego siga las
reglas económicas configuradas sin cálculos manuales.

```gherkin
Escenario: Interés aplicado al volver el turno
  Dado un préstamo con interés "por ronda" del 5%
  Cuando el turno vuelve al jugador deudor
  Entonces su deuda se incrementa en 5% automáticamente y queda registrado
    en su historial
```

- **Dependencias:** US-034, US-053, US-081.
- **Notas técnicas:** Ninguna.

---

### US-083: Cálculo de interés por tiempo
- **Épica:** EP-05 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-31

Como jugador con deuda activa, quiero que el interés por tiempo se calcule
automáticamente cada N minutos configurados, para partidas donde no se sigue
un orden estricto de turnos.

```gherkin
Escenario: Interés aplicado por temporizador
  Dado un préstamo con interés "por tiempo" del 5% cada 10 minutos
  Cuando transcurren 10 minutos desde el último cálculo
  Entonces la deuda se incrementa en 5% automáticamente
```

- **Dependencias:** US-034, US-053.
- **Notas técnicas:** Requiere un proceso en background (scheduler) en el
  backend por partida activa; considerar el reinicio de pods (RNF-08) al
  diseñar la persistencia del último cálculo.

---

### US-084: Vista de espectador (solo saldos en vivo)
- **Épica:** EP-05 | **Prioridad:** Must | **Estimación:** S | **RF:** complementaria a RF-28 (rol Espectador, PRD §2)

Como espectador, quiero ver los saldos de los jugadores actualizarse en
tiempo real, para seguir el estado económico de la partida sin operar.

```gherkin
Escenario: Espectador ve saldos en vivo
  Dado que estoy conectado como espectador a una partida
  Cuando ocurre una transferencia entre jugadores
  Entonces veo los saldos actualizados en tiempo real

Escenario: Espectador sin acceso a historial detallado
  Dado que estoy conectado como espectador
  Cuando intento consultar el historial de transacciones de un jugador
  Entonces el sistema no me da acceso a ese detalle (solo a saldos actuales)
```

- **Dependencias:** US-017, US-080.
- **Notas técnicas:** **Decisión del PO (2026-09-18): el espectador ve solo
  saldos en vivo, sin acceso al historial detallado de transacciones.** Esto
  acota el brief ("ver saldos y movimientos en vivo") a actualizaciones de
  saldo en tiempo real, no a un log histórico consultable.

---

## EP-06: Ranking y Puntuación

### US-100: Cálculo de puntuación final de partida
- **Épica:** EP-06 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-32

Como jugador registrado, quiero que al finalizar la partida se calcule mi
puntuación, para saber cómo me fue.

```gherkin
Escenario: Cálculo al finalizar
  Dado que la partida finaliza (por liquidación general o cierre del Banco)
  Cuando se calcula mi puntuación
  Entonces esta es igual a mi patrimonio neto: digital + papel declarado −
    deuda, en el momento de la liquidación (assumption #3)
```

- **Dependencias:** US-055, US-056.
- **Notas técnicas:** Ninguna.

---

### US-101: Leaderboard global de mejores partidas
- **Épica:** EP-06 | **Prioridad:** Must | **Estimación:** M | **RF:** RF-33

Como usuario registrado, quiero ver un ranking global con las mejores
puntuaciones de partidas jugadas, para comparar mi desempeño con el de otros
jugadores.

```gherkin
Escenario: Entrada en el leaderboard
  Dado que finalicé una partida con una puntuación de 5000
  Cuando consulto el ranking global
  Entonces mi resultado aparece listado junto con los de otras partidas,
    ordenado de mayor a menor puntuación

Escenario: Sin acumulación entre partidas
  Dado que he jugado dos partidas con puntuaciones 5000 y 3000
  Cuando consulto el ranking global
  Entonces aparecen como dos entradas independientes (5000 y 3000), no como
    un acumulado de 8000 ni un promedio de 4000
```

- **Dependencias:** US-100.
- **Notas técnicas:** **Decisión del PO (2026-09-18): el ranking global es un
  leaderboard de mejores partidas individuales, sin acumular ni promediar el
  desempeño histórico de un usuario.** Ver US-205 en "Fuera del MVP" para la
  alternativa descartada.

---

### US-102: Exclusión de invitados del ranking
- **Épica:** EP-06 | **Prioridad:** Must | **Estimación:** S | **RF:** RF-34

Como sistema, no debo incluir a jugadores invitados en el ranking global,
para que el ranking refleje solo a usuarios con cuenta registrada.

```gherkin
Escenario: Invitado finaliza partida con buena puntuación
  Dado un invitado que finaliza una partida con la puntuación más alta
  Cuando se actualiza el ranking global
  Entonces su resultado no aparece en el leaderboard
```

- **Dependencias:** US-004, US-100, US-101.
- **Notas técnicas:** Ninguna.

---

## EP-07: Administración (baja prioridad, no bloqueante para el MVP)

### US-110: Gestión de monedas soportadas
- **Épica:** EP-07 | **Prioridad:** Could | **Estimación:** M | **RF:** sin número (PRD §2, rol Admin general)

Como Admin general, quiero ver y editar la lista de monedas soportadas por
la aplicación, para mantenerla actualizada sin necesidad de un despliegue.

```gherkin
Escenario: Agregar moneda soportada
  Dado que soy Admin general
  Cuando agrego una nueva moneda con su símbolo y configuración por defecto
  Entonces queda disponible para elegir al crear una partida (US-032)
```

- **Dependencias:** US-001 (requiere rol Admin autenticado), US-032.
- **Notas técnicas:** Assumption #5: interfaz mínima, no prioritaria para el
  MVP; una partida puede jugarse completa sin este panel usando monedas
  precargadas por configuración/seed.

---

### US-111: Gestión de parámetros globales por defecto
- **Épica:** EP-07 | **Prioridad:** Could | **Estimación:** S | **RF:** sin número (PRD §2, rol Admin general)

Como Admin general, quiero definir valores por defecto (fee, interés,
duración de sesión) a nivel de aplicación, para no tener que repetirlos en
cada partida.

```gherkin
Escenario: Definir valores por defecto
  Dado que soy Admin general
  Cuando defino un fee de transferencia por defecto del 0%
  Entonces las nuevas partidas usan ese valor si el Banco no lo cambia
    explícitamente
```

- **Dependencias:** US-110.
- **Notas técnicas:** Baja prioridad; puede resolverse con configuración
  estática (variables de entorno / seed de base de datos) en el MVP en lugar
  de un panel de UI, sin bloquear otras épicas.

---

## Fuera del MVP

Historias identificadas pero explícitamente excluidas de esta iteración,
para dejar constancia de que no se olvidaron sino que se decidió posponerlas.

| ID | Descripción | Motivo |
|----|-------------|--------|
| US-200 | Reglas específicas de Tío Rico: compra/venta de acciones, dividendos por coincidencia de dado, cambios de cotización y cartas de "La Trombetta" en los dobles, y condición de victoria por patrimonio ($50,000/$100,000) | **Actualizado 2026-09-18:** ya no es una pregunta abierta — las reglas están documentadas en `11-reglamento-juegos.md` (sección 1) y su correspondencia propuesta con CashCastor en la sección 4 (borrador). No se implementa en el MVP: requiere entidades nuevas (`Stock`, `StockQuotation`) y una condición de fin de partida por meta de patrimonio que no existen hoy. |
| US-201 | Lógica funcional de inflación (afectar precios de propiedades, tasa de cambio o saldos digitales) | **Decisión del PO (2026-09-18):** requiere un POC jugando una ronda real antes de definir el mecanismo; no se implementa en el MVP (ver US-036). |
| US-202 | Modo offline con cola de operaciones pendientes de sincronizar | Exclusión explícita del brief (Scope MVP, fase 2). El Service Worker del MVP solo cachea assets estáticos (ADR-006). |
| US-203 | Pagos con dinero real / integración con pasarelas de pago | Exclusión explícita del brief (Scope MVP, fase 2). |
| US-204 | Notificaciones push nativas | Exclusión explícita del brief (Scope MVP, fase 2). |
| US-205 | Ranking acumulado o promediado por usuario entre partidas | Alternativa evaluada y descartada por el PO (2026-09-18) a favor del leaderboard por partida (US-101). |
| US-206 | Declaración de papel moneda por parte del jugador (autoservicio) | Alternativa evaluada y descartada por el PO (2026-09-18) a favor de que solo el Banco declare el arqueo (US-056). |
| US-207 | Historial de transacciones visible para espectadores | Alternativa evaluada y descartada por el PO (2026-09-18) a favor de que el espectador vea solo saldos en vivo (US-084). |
| US-208 | Emisión de dinero de emergencia por el Banco durante la partida | Pregunta abierta #2 del documento de supuestos: actualmente "no"; contradice ADR-003 (masa monetaria fija). No se implementa salvo cambio de decisión. |
| US-209 | Integración con el tablero físico (NFC, QR en casillas) | Pregunta abierta #6 del documento de supuestos, sin definir; no forma parte del MVP. |
