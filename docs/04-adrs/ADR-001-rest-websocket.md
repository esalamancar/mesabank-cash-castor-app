# ADR-001: Uso de REST + WebSocket en lugar de gRPC

**Estado:** Aceptado

## Contexto
Necesitamos operaciones CRUD y comunicación en tiempo real. El frontend es una
PWA que corre en navegadores móviles.

## Decisión
Usar REST para operaciones de negocio y WebSocket para eventos en tiempo real.

## Consecuencias
Mayor compatibilidad con navegadores; REST es más simple de documentar con
OpenAPI; WebSocket requiere gestión de conexiones y reconexión.
