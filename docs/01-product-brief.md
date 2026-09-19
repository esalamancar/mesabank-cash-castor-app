# Product Brief: MesaBank

## Executive Summary
MesaBank es una aplicación web que digitaliza la gestión de la caja del banco en
juegos de mesa económicos como Monopoly y Tío Rico. No reemplaza el papel moneda;
lo complementa. El banco de la partida (un jugador con rol de árbitro) administra
una masa monetaria fija, dividida entre dinero digital y papel moneda. Los jugadores
usan sus propios dispositivos móviles para consultar saldos, transferirse dinero
digital y hacer consignaciones de efectivo. La app introduce la analogía de las
billeteras digitales modernas (Nequi, Bizum, Alipay) en el contexto analógico del
tablero, dando al banco una herramienta de auditoría en tiempo real.

## The Problem
En juegos como Monopoly o Tío Rico, el banco es un rol tedioso y propenso a errores.
El jugador-banquero debe repartir billetes, cobrar rentas, gestionar hipotecas y
calcular intereses manualmente, lo que ralentiza el juego y genera disputas. Además,
no hay trazabilidad: al final de la partida nadie sabe con certeza cuánto dinero
circuló, quién pagó qué o si la caja cuadra. El papel moneda se desgasta, se pierde
o se mezcla entre jugadores.

## The Solution
Una PWA responsive (mobile-first) donde cada partida tiene un código de sala. El
creador asume el rol de **Banco** y configura la masa monetaria inicial: cuánto
dinero digital y cuánto papel moneda existe en total, y cuánto recibe cada jugador
al inicio. Los jugadores se unen con código, QR o link, y desde su teléfono pueden:

- Ver su saldo digital y su efectivo declarado.
- Transferir dinero digital a otros jugadores (instantáneo, idempotente).
- Hacer consignaciones al banco (entregan papel moneda y reciben saldo digital).
- Ver su historial de ingresos y egresos.

El Banco ve todos los saldos, el total de papel moneda en caja, el total de digital
en circulación y un arqueo en tiempo real. Puede otorgar préstamos, iniciar
liquidaciones y expulsar jugadores en bancarrota.

## What Makes This Different
- **No es un juego digital**: el tablero y las fichas siguen siendo físicos.
  La app es una capa de gestión financiera, no una simulación.
- **Dualidad digital/físico**: el banco controla cuánto papel moneda existe y
  cuánto digital, y debe cuadrar al final.
- **Rol de banco con superpoderes de auditoría**: a diferencia de una billetera
  normal, el banco ve todo el flujo de dinero.
- **Multi-partida tipo Kahoot**: cualquiera crea una sala, los demás se unen con
  un código.

## Who This Serves
- **Jugadores de juegos de mesa económicos**: personas que juegan Monopoly,
  Tío Rico u otros juegos de gestión de recursos y quieren una experiencia más
  fluida sin perder el componente físico.
- **El banco/árbitro**: un jugador que quiere dejar de ser un cajero manual y
  convertirse en un administrador con visibilidad total.
- **Espectadores**: personas que quieren seguir la partida sin jugar, viendo
  saldos y movimientos en vivo.

## Success Criteria
- Una partida completa de principio a fin sin necesidad de calcular saldos a mano.
- El arqueo final del banco cuadra: el papel moneda en caja más el declarado por
  los jugadores es igual al inicial.
- Los jugadores pueden transferirse dinero sin fricción y sin errores de
  concurrencia.
- El sistema soporta múltiples partidas simultáneas sin interferencia.

## Scope (MVP)
**Incluye:**
- Registro/login con usuario y PIN.
- Creación de partida con código, QR y link.
- Roles: Banco, Jugador, Espectador.
- Configuración de masa monetaria (digital y papel), moneda, tasa de cambio,
  inflación, fee de transferencia (por defecto 0), interés por ronda o por tiempo.
- Transferencias P2P instantáneas e idempotentes.
- Consignaciones (papel a digital).
- Préstamos del banco a jugadores.
- Historial por jugador y auditoría completa para el banco.
- Arqueo de caja en tiempo real.
- Ranking al final de la partida (solo usuarios registrados).
- Expulsión por bancarrota.
- Reinicio de partida con confirmaciones múltiples.

**Excluye (fase 2):**
- Reglas específicas de Tío Rico (hipotecas, rentas, etc.).
- Modo offline.
- Pagos con dinero real.
- Notificaciones push nativas.

## Vision
Convertirse en la capa financiera estándar para juegos de mesa económicos, con
soporte para reglas personalizadas por juego, torneos, ligas y un ranking global
de jugadores. A largo plazo, la app podría integrarse con tableros digitales o
asistentes de voz, pero sin perder el papel moneda como elemento tangible del juego.
