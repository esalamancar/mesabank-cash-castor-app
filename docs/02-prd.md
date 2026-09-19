# PRD: CashCastor MVP

## 1. Objetivo del Producto
Digitalizar la gestión de la caja del banco en juegos de mesa económicos,
complementando el papel moneda con una billetera digital y una herramienta de
auditoría para el banco.

## 2. Usuarios y Roles

| Rol | Descripción | Permisos |
|-----|-------------|----------|
| Admin general | Administrador de la aplicación | Configuración global, gestión de usuarios, monedas base. |
| Banco | Jugador que crea la partida y actúa como árbitro | Configura partida, ve todos los saldos, otorga préstamos, inicia liquidación, expulsa jugadores en bancarrota, reinicia partida. |
| Jugador | Participante de la partida | Ve su saldo e historial, transfiere a otros, consigna papel moneda, solicita préstamos. |
| Espectador | Observador sin participación activa | Ve saldos y movimientos en vivo, pero no opera. |

## 3. Requisitos Funcionales

### 3.1 Autenticación y Sesión
- RF-01: Registro con usuario y PIN.
- RF-02: Login con usuario y PIN.
- RF-03: Recuperación de sesión tras desconexión (retoma donde quedó).
- RF-04: Acceso como invitado (sin ranking).

### 3.2 Creación y Gestión de Partidas
- RF-05: Crear partida con configuración de masa monetaria.
- RF-06: Generar código de sala, QR y link de invitación.
- RF-07: Unirse a partida por código, QR o link.
- RF-08: Asignación automática de rol Banco al creador.
- RF-09: Soportar múltiples partidas simultáneas.
- RF-10: Reinicio de partida con confirmaciones múltiples.
- RF-11: Expulsión de jugador en bancarrota (solo banco).

### 3.3 Configuración de Partida (Banco)
- RF-12: Definir masa monetaria total (digital y papel).
- RF-13: Definir asignación inicial por jugador.
- RF-14: Definir moneda (USD, COP, EUR, personalizada) y tasa de cambio.
- RF-15: Definir inflación (configurable).
- RF-16: Definir fee de transferencia (por defecto 0).
- RF-17: Definir interés por ronda o por tiempo (porcentaje, periodicidad).
- RF-18: Definir deuda máxima como múltiplo del valor inicial.

### 3.4 Operaciones Financieras
- RF-19: Transferencia P2P instantánea e idempotente.
- RF-20: Consignación de papel moneda a cuenta digital.
- RF-21: Préstamo del banco a jugador (monto, interés, plazo).
- RF-22: Pago de intereses automático según configuración.
- RF-23: Registro de deuda con el banco (saldo negativo contable).
- RF-24: Liquidación de bancarrota (banco paga a acreedores).
- RF-25: Arqueo de caja en tiempo real (banco).
- RF-26: Historial de ingresos y egresos por jugador.
- RF-27: Auditoría completa para el banco.

### 3.5 Tiempo Real y Turnos
- RF-28: Sincronización en tiempo real vía WebSockets.
- RF-29: Asignación de turnos desde la app.
- RF-30: Cálculo de intereses por ronda (cuando el turno vuelve al jugador).
- RF-31: Cálculo de intereses por tiempo (cada N minutos).

### 3.6 Ranking y Puntuación
- RF-32: Puntuación al final de la partida.
- RF-33: Ranking global para usuarios registrados.
- RF-34: Invitados no entran al ranking.

## 4. Requisitos No Funcionales
- RNF-01: Backend en Go con Gin.
- RNF-02: Frontend React con Vite, PWA.
- RNF-03: PostgreSQL para persistencia, Redis para caché y sesiones.
- RNF-04: Despliegue en Kubernetes siguiendo reglas del repo `single-iac`.
- RNF-05: API documentada con OpenAPI.
- RNF-06: HTTPS obligatorio.
- RNF-07: Tolerancia a conexiones lentas/inestables.
- RNF-08: Persistencia entre reinicios de pods.
- RNF-09: Responsive, principalmente móvil.

## 5. Modelo de Datos (Entidades Principales)

| Entidad | Campos clave |
|---------|--------------|
| User | id, username, pin_hash, is_guest, created_at |
| Game | id, code, qr_url, link, status, created_by, config_json |
| GameConfig | currency, exchange_rate, inflation, fee, interest_type, interest_rate, interest_interval, max_debt_multiplier, total_digital, total_paper, initial_per_player |
| Player | id, game_id, user_id, role, digital_balance, paper_balance, debt, is_bankrupt |
| Transaction | id, game_id, from_player_id, to_player_id, amount, type, idempotency_key, created_at |
| Loan | id, game_id, player_id, amount, interest_rate, due_round, status |
| AuditLog | id, game_id, player_id, action, details, created_at |

## 6. API Endpoints (Resumen)

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | /auth/register | Registro |
| POST | /auth/login | Login |
| POST | /games | Crear partida |
| GET | /games/{code} | Info de partida |
| POST | /games/{code}/join | Unirse |
| POST | /games/{code}/transfer | Transferencia P2P |
| POST | /games/{code}/deposit | Consignación papel a digital |
| POST | /games/{code}/loan | Préstamo (solo banco) |
| GET | /games/{code}/audit | Auditoría (solo banco) |
| GET | /games/{code}/balance | Saldo propio |
| WS | /games/{code}/ws | Tiempo real |

## 7. Criterios de Aceptación (Ejemplos)
- CA-01: Dado un banco con masa monetaria configurada, cuando un jugador se
  une, recibe la asignación inicial correcta en digital y papel.
- CA-02: Dado un jugador con saldo suficiente, cuando transfiere a otro, el
  saldo se actualiza en tiempo real en ambos dispositivos.
- CA-03: Dado un banco, cuando inicia liquidación de un jugador en bancarrota,
  los acreedores reciben pago inmediato y el jugador es expulsado.
- CA-04: Dado un reinicio de partida, se requieren al menos tres confirmaciones
  y el arqueo de papel moneda debe coincidir.
