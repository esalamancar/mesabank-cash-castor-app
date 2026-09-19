# Roadmap MVP

## Fase 0: Fundaciones (Semana 1-2)
- Configuración del repositorio y estructura de carpetas.
- Definición de OpenAPI spec.
- Setup de PostgreSQL y Redis en local/desarrollo.
- Setup de Kubernetes (namespace, secrets, configmaps) siguiendo `single-iac`.
- Setup de CI/CD (pipeline base).

## Fase 1: Core Bancario (Semana 3-5)
- Autenticación: registro, login, JWT, sesión persistente.
- Gestión de partidas: crear, unirse, roles, código/QR/link.
- Configuración de masa monetaria y parámetros de partida.
- Modelo de datos y migraciones.
- API de saldos y consulta de balance.

## Fase 2: Operaciones Financieras (Semana 6-8)
- Transferencias P2P con idempotencia.
- Consignaciones (papel a digital).
- Préstamos del banco.
- Historial por jugador.
- Auditoría para el banco.
- Cálculo de intereses (por ronda y por tiempo).

## Fase 3: Tiempo Real y Arqueo (Semana 9-10)
- WebSocket hub y broadcast de eventos.
- Arqueo de caja en tiempo real.
- Panel de banco (saldos, totales, deudas).
- Liquidación de bancarrota.
- Expulsión de jugadores.

## Fase 4: Frontend y UX (Semana 11-13)
- PWA con Vite y Service Worker.
- Vistas: login, lobby, sala de espera, billetera, transferencia, historial,
  panel banco.
- UI mobile-first y responsive.
- Integración con WebSocket.

## Fase 5: Cierre y Ranking (Semana 14)
- Puntuación final de partida.
- Ranking global para usuarios registrados.
- Reinicio de partida con confirmaciones.
- Pruebas de carga y ajustes.

## Fase 6: Entrega y Documentación (Semana 15)
- Documentación técnica final.
- Guía de despliegue en `single-iac`.
- Pruebas end-to-end.
- Demo y validación con PO.
