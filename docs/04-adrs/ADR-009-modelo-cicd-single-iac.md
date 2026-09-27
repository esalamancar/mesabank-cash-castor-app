# ADR-009: Modelo real de responsabilidad CI/CD con `single-iac`

**Estado:** Aceptado

## Contexto
ADR-007 estableció que el despliegue sigue las reglas de `single-iac` y que
"el agente de CI/CD es el existente en ese repo, no se crean pipelines
ad-hoc". Al coordinar con `single-iac` (T-08/T-09/T-10 del backlog de
Sprint 0) se aclaró que esa frase se puede interpretar de dos formas muy
distintas:

1. `single-iac` construye, publica y despliega la imagen de cada app
   (pipeline centralizado).
2. `single-iac` solo aplica los manifiestos de Kubernetes (namespace,
   Deployments, Secrets, etc.); cada app trae su propio pipeline de
   build + push + deploy en su propio repo.

`single-iac` ya tuvo un incidente real por construir/gestionar el tag de
imagen de otra app (violación de sus propios ADR-0009 y ADR-0012), así que
la opción (1) está descartada ahí de forma permanente.

## Decisión
Se adopta el modelo (2): **CashCastor es dueño de su propio pipeline de
build, push y despliegue**, como workflow de GitHub Actions en este
repositorio (no en `single-iac`). `single-iac` provee:

- El namespace, quotas, NetworkPolicy, DNS/TLS y demás recursos de
  infraestructura (Terraform, en `single-iac`).
- Los manifiestos base de Kubernetes (Deployments, Services, Secrets de
  ejemplo), en `k8s-manifests/cashcastor/` de `single-iac`.
- Un workflow **genérico** que aplica esos manifiestos contra el cluster
  (ya existe, cubre `cashcastor` sin cambios adicionales).

Esto sigue cumpliendo ADR-007 ("no pipelines ad-hoc", "seguir las reglas de
`single-iac`"): el "agente de CI/CD existente" al que se refería esa
decisión es ese workflow genérico de aplicación de manifiestos, no un
pipeline de build centralizado.

Nombres de recursos: se usa `api`/`web` (sin prefijo `cashcastor-`) para
los Deployments/Services, siguiendo el patrón real del resto de apps de
`single-iac` (el namespace ya da el scope) — reemplaza el boceto original
de `cashcastor-api`/`cashcastor-web` de `03-arquitectura.md`.

## Consecuencias
- Este repo necesita su propio workflow de GitHub Actions para build +
  push de las imágenes (`backend/Dockerfile`, `frontend/Dockerfile`) y el
  despliegue (`kubectl set image` o equivalente) — pendiente, ver
  `backlog/tareas.md` (T-10).
- Sin separación de infraestructura por entorno (dev/qa/prod): un único
  namespace `cashcastor` para todo, mapeado a la rama `develop`. Revisar
  cuando haya necesidad real (ver nota en `03-arquitectura.md` §2.4).
- Nada de lo anterior está aplicado contra el cluster real todavía: vive
  en la rama `add-cashcastor-namespace` de `single-iac`, pendiente de
  revisión, merge y `terraform-apply.yml` manual por el PO.
