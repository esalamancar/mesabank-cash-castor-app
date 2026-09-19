# ADR-004: PostgreSQL como fuente de verdad, Redis como caché y pub/sub

**Estado:** Aceptado

## Contexto
Necesitamos persistencia entre reinicios de pods y velocidad para tiempo real.

## Decisión
PostgreSQL para datos transaccionales; Redis para sesiones, idempotencia y
pub/sub de WebSockets.

## Consecuencias
Dos sistemas que mantener; Redis no es persistente por defecto, pero se usa
para datos efímeros.
