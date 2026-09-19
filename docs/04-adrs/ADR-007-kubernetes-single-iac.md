# ADR-007: Despliegue en Kubernetes siguiendo reglas de single-iac

**Estado:** Aceptado

## Contexto
La organización ya tiene un repositorio de infraestructura como código
(`single-iac`) con pipelines y convenciones.

## Decisión
Todos los manifests, pipelines y configuraciones de despliegue deben seguir las
reglas definidas en `single-iac`. El agente de CI/CD es el existente en ese repo.

## Consecuencias
Consistencia con el resto de proyectos; requiere conocer las convenciones de
`single-iac`; no se crean pipelines ad-hoc.
