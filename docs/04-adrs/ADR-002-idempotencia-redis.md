# ADR-002: Idempotencia de transacciones en Redis

**Estado:** Aceptado

## Contexto
Las conexiones móviles pueden ser inestables. Un reintento de transferencia
podría duplicar el débito.

## Decisión
Cada transacción lleva una `idempotency_key` generada por el cliente. Redis
almacena el resultado por un TTL de 24h.

## Consecuencias
Evita duplicados; requiere que el cliente genere claves únicas; Redis debe ser
altamente disponible.
