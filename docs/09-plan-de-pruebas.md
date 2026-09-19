# Plan de Pruebas: CashCastor MVP

Estrategia de testing para el backend (Go + Gin), el frontend (React + Vite
PWA) y el flujo end-to-end del sistema completo, alineada con las historias
de `07-historias-usuario.md` y los criterios de aceptación CA-01 a CA-04 del
PRD.

## 1. Estrategia General

Pirámide de testing, de mayor a menor volumen de casos:

```
        /\
       /e2e\        <- pocos, flujos críticos completos (Playwright)
      /------\
     /integra-\     <- servicios + Postgres/Redis reales o en contenedor
    /   ción   \
   /------------\
  / unitarias    \  <- la mayoría: lógica de negocio pura, sin I/O
 /----------------\
```

| Nivel | Objetivo | % aprox. del esfuerzo |
|-------|----------|------------------------|
| Unitarias | Lógica de negocio aislada (cálculo de intereses, validación de saldo, idempotencia, arqueo) | 60% |
| Integración | Endpoints REST contra Postgres/Redis reales (en contenedor), WebSocket hub | 30% |
| E2E | Flujos completos multi-usuario simulando una partida real | 10% |

## 2. Herramientas

### 2.1 Backend (Go)

| Herramienta | Uso |
|-------------|-----|
| `testing` (stdlib) + `testify` (`assert`/`require`/`mock`) | Pruebas unitarias e integración. |
| `net/http/httptest` | Pruebas de handlers Gin sin levantar un servidor real. |
| `dockertest` (o `testcontainers-go`) | Levantar Postgres y Redis reales en contenedor para pruebas de integración reproducibles. |
| `sqlmock` (`go-sqlmock`) | Pruebas unitarias de repositorios sin base de datos real, cuando no se justifica un contenedor. |
| `miniredis` | Simular Redis en memoria para pruebas unitarias de idempotencia y pub/sub. |
| `gorilla/websocket` (cliente de prueba) o `nhooyr.io/websocket` | Pruebas de integración del hub de WebSocket. |
| `golangci-lint` | Calidad estática, parte del pipeline de CI. |
| `go test -race` | Detección de condiciones de carrera (crítico para transferencias concurrentes). |

### 2.2 Frontend (React + Vite)

| Herramienta | Uso |
|-------------|-----|
| Vitest | Pruebas unitarias de componentes, hooks y stores (Zustand/Redux Toolkit). |
| React Testing Library | Pruebas de componentes centradas en comportamiento observable por el usuario. |
| MSW (Mock Service Worker) | Mockear la API REST y simular eventos de WebSocket en pruebas de integración de UI. |
| Playwright | Pruebas e2e multi-navegador/multi-dispositivo (simula varios jugadores en paralelo). |
| `axe-core` / `@axe-core/playwright` | Verificación básica de accesibilidad en las vistas principales. |

### 2.3 Contrato de API

| Herramienta | Uso |
|-------------|-----|
| `@redocly/cli lint` | Validar sintácticamente `08-openapi.yaml` en cada cambio (usado ya en este repo para generar el spec). |
| Contract testing (ej. Dredd o pruebas Go generadas desde el spec) | Verificar que las respuestas reales de la API cumplen los schemas del OpenAPI, especialmente en endpoints financieros. |

## 3. Casos de Prueba Derivados de Criterios de Aceptación del PRD

### CA-01: Asignación inicial al unirse a una partida
| Caso | Tipo | Descripción |
|------|------|-------------|
| CA-01.1 | Integración | Un jugador que se une a una partida configurada recibe exactamente `initial_per_player` en digital y papel. |
| CA-01.2 | Unitaria | La suma de asignaciones iniciales no puede exceder `total_digital`/`total_paper`; el sistema rechaza la configuración si excede. |
| CA-01.3 | Integración | Unirse a un código de partida inexistente devuelve 404. |

### CA-02: Actualización de saldo en tiempo real tras una transferencia
| Caso | Tipo | Descripción |
|------|------|-------------|
| CA-02.1 | E2E | Dos dispositivos conectados a la misma partida; al transferir desde uno, el otro ve el saldo actualizado sin recargar, en menos de 2 segundos. |
| CA-02.2 | Integración | El evento `balance_updated` se publica en el canal de Redis pub/sub correspondiente a la partida tras una transferencia exitosa. |
| CA-02.3 | Integración | Un cliente conectado a la partida "ABCD" no recibe eventos de la partida "WXYZ" (aislamiento, US-013). |

### CA-03: Liquidación de bancarrota
| Caso | Tipo | Descripción |
|------|------|-------------|
| CA-03.1 | Integración | Al iniciar liquidación de un jugador con un acreedor, este recibe el pago disponible de inmediato. |
| CA-03.2 | Integración | Tras la liquidación, el jugador queda `is_bankrupt = true` y es expulsado (US-015) en la misma operación o en la siguiente confirmación del Banco. |
| CA-03.3 | Unitaria | Con múltiples acreedores y fondos insuficientes, el prorrateo de pago sigue la regla definida (proporcional a la deuda, ver nota técnica de US-055). |

### CA-04: Reinicio de partida con confirmaciones múltiples
| Caso | Tipo | Descripción |
|------|------|-------------|
| CA-04.1 | Integración | Una solicitud de reinicio con menos de 3 confirmaciones no se ejecuta. |
| CA-04.2 | Integración | Con 3 confirmaciones, la partida vuelve a su configuración inicial y el arqueo de papel coincide con la masa total original. |
| CA-04.3 | Unitaria | Una confirmación duplicada del mismo participante no cuenta dos veces. |

## 4. Casos de Prueba por Épica (resumen no exhaustivo)

Cada historia de `07-historias-usuario.md` aporta al menos un escenario
Gherkin que debe convertirse en un caso de prueba automatizado. Se destacan
aquí los de mayor riesgo:

| Historia | Caso crítico |
|----------|--------------|
| US-001/US-002 | PIN se almacena con hash; nunca se expone en logs ni respuestas. |
| US-004 | Un invitado no puede crear partidas ni aparecer en `/ranking`. |
| US-014 | Reinicio requiere el mínimo de confirmaciones configurado (ver CA-04). |
| US-030/US-031 | La configuración de masa monetaria es inmutable una vez la partida está activa (ADR-003). |
| US-035 | Un préstamo que excede `max_debt_multiplier` se rechaza con 422. |
| US-036 | El campo `inflation` se persiste pero no altera ningún cálculo (regresión: agregar un caso que falle si en el futuro se introduce lógica no autorizada sin pasar por el proceso de POC). |
| US-056 | El arqueo solo puede registrarse por el rol Banco; un jugador que intente `POST /games/{code}/cash-audit` recibe 403. |
| US-084 | Un espectador que intente `GET /games/{code}/transactions` recibe 403. |
| US-101 | El leaderboard no acumula puntuaciones de un mismo usuario entre partidas distintas. |

## 5. Concurrencia e Idempotencia en Transferencias

Sección crítica dado ADR-002 y RNF-07 (tolerancia a conexiones inestables).

| Caso | Tipo | Descripción |
|------|------|-------------|
| CONC-01 | Integración | Enviar la misma solicitud de transferencia (misma `idempotency_key`) 5 veces en paralelo produce exactamente un único movimiento de saldo. |
| CONC-02 | Integración | Dos transferencias simultáneas desde el mismo jugador, cuyo saldo solo alcanza para una, resultan en que solo una se ejecuta (sin saldo negativo por condición de carrera). Verificar con `go test -race` y bloqueo/transacción a nivel de fila en PostgreSQL (`SELECT ... FOR UPDATE` o equivalente). |
| CONC-03 | Integración | Transferencia y consignación simultáneas del mismo jugador no corrompen su saldo (orden serializado correctamente). |
| CONC-04 | Integración | Expira el TTL de 24h de una `idempotency_key` en Redis: una nueva solicitud con la misma clave después del TTL se trata como una operación nueva (comportamiento esperado documentado, no un bug). |
| CONC-05 | Integración | El cálculo de interés por ronda (US-082) y una transferencia concurrente sobre el mismo jugador no se pisan entre sí (ambos cambios de saldo quedan reflejados). |
| CONC-06 | Unitaria | Idempotency key mal formada (no UUID) es rechazada antes de tocar Redis o Postgres. |

## 6. Pruebas de Carga

Objetivo: validar RF-09 (múltiples partidas simultáneas) y RNF-07/RNF-09.

| Escenario | Herramienta | Descripción |
|-----------|-------------|-------------|
| Carga REST | k6 | Simular N partidas concurrentes (ej. 50, 200, 500) con M jugadores cada una (4-8) ejecutando transferencias y consignaciones a una tasa realista de juego de mesa (varias operaciones por minuto, no por segundo). Medir p95/p99 de latencia y tasa de error. |
| Carga WebSocket | k6 (extensión `xk6-websocket`) o Artillery | Mantener conexiones WS abiertas por partida y medir tiempo de broadcast de un evento a todos los conectados de esa partida bajo carga. |
| Estrés de Redis pub/sub | k6 / script dedicado | Verificar que el pub/sub no se degrada con cientos de canales (uno por partida) activos simultáneamente. |
| Recuperación tras reinicio de pod | Manual / chaos test | Matar el pod de la API durante una partida activa y verificar que, al reiniciar, las partidas persisten (RNF-08) y los clientes se reconectan (US-003, US-080). |

Criterios de salida sugeridos (a validar con el equipo antes de fase de
carga real): p95 de transferencia < 500ms con 200 partidas concurrentes
activas, sin errores 5xx bajo carga nominal.

## 7. Fuera de Alcance de Pruebas en el MVP

Consistente con `07-historias-usuario.md`, sección "Fuera del MVP": no se
diseñan casos de prueba para reglas específicas de Tío Rico, lógica
funcional de inflación, modo offline, pagos reales ni push notifications
nativas, dado que esas funcionalidades no se implementan en esta iteración.
