# ADR-005: Autenticación con usuario y PIN (JWT)

**Estado:** Aceptado

## Contexto
El público objetivo son jugadores de mesa, no usuarios técnicos. No se manejan
datos sensibles reales.

## Decisión
Registro con usuario y PIN de 4-6 dígitos. Se emite un JWT de corta duración
con refresh token.

## Consecuencias
Simple y rápido; menos seguro que OAuth, pero adecuado para el contexto. Los
invitados no tienen PIN persistente.
