# Supuestos y Preguntas Abiertas

## Supuestos
1. Reglas de Tío Rico: **Actualizado 2026-09-18 — ya no se asume equivalencia
   funcional con Monopoly** (esta assumption reemplaza a la versión anterior,
   que resultó incorrecta). Según el reglamento oficial incorporado en
   `docs/11-reglamento-juegos.md`, Tío Rico es un juego de compra/venta de
   acciones (Stock Cards) con dividendos pagados cuando la tirada de dados
   coincide con el número de la acción poseída, cotizaciones que cambian con
   cada doble, y victoria por acumular $50,000 (o $100,000) en efectivo — no
   por ser el último jugador no quebrado, como en Monopoly. Su mecánica
   central (acciones, dividendos, cotizaciones) no mapea directamente a los
   conceptos actuales de CashCastor (Loan, liquidación de propiedades). La
   implementación específica de estas reglas se pospone a fase 2; hay una
   propuesta inicial de correspondencia (borrador, no implementada en el MVP)
   en `docs/11-reglamento-juegos.md`, sección 4.
2. Arqueo de papel moneda: El banco declara manualmente cuánto papel moneda
   tiene en caja y cuánto tienen los jugadores. La app no escanea billetes físicos.
3. Ranking: La puntuación final se basa en el patrimonio neto (digital + papel
   declarado menos deuda) al momento de la liquidación.
4. Invitados: No pueden crear partidas (solo usuarios registrados pueden ser
   Banco) y no entran al ranking.
5. Admin general: Se implementará una interfaz mínima de administración para
   gestionar monedas y parámetros globales, pero no es prioridad en el MVP.

## Preguntas Abiertas (para fase 2)
1. ~~¿Cuáles son las reglas exactas de Tío Rico para hipotecas, rentas, subastas y
   bancarrota?~~ **RESUELTA (2026-09-18):** Tío Rico no tiene hipotecas, rentas
   ni subastas; su mecánica es de acciones y dividendos. Ver el reglamento
   completo en `docs/11-reglamento-juegos.md`, sección 1, y la propuesta de
   correspondencia con CashCastor en la sección 4 del mismo documento (borrador
   de fase 2, aún sin implementar).
2. ¿El banco debe poder emitir dinero de emergencia? (Actualmente: no).
3. ~~¿Cómo se maneja la inflación exactamente? ¿Afecta a los precios de las
   propiedades o solo al valor del dinero?~~ **RESUELTA (2026-09-18):** queda
   fuera del MVP por completo; no se implementa ninguna lógica automática
   (ni sobre precios, ni tasa de cambio, ni saldos). Se requiere un POC
   jugando una ronda real antes de definir el mecanismo — ver
   `07-historias-usuario.md` (US-036, US-201) y la nota relacionada con las
   cotizaciones de Tío Rico en `11-reglamento-juegos.md`, sección 4.
4. ~~¿Los espectadores pueden ver el historial de transacciones o solo
   saldos?~~ **RESUELTA (2026-09-18):** solo ven saldos en vivo, sin acceso
   al historial detallado de transacciones. Ver `07-historias-usuario.md`
   (US-084, US-207).
5. ¿Se requiere un modo de solo papel moneda (sin digital)?
6. ¿La app debe integrarse con el tablero físico de alguna manera (NFC, QR en
   casillas)?
7. ¿Cuál es el TTL exacto de los JWT y la estrategia de refresh?
8. ¿El repo `single-iac` tiene convenciones específicas para nombres de
   namespaces, secrets o pipelines que deba seguir?
