# ADR-008: Migraciones de esquema vía GORM AutoMigrate, sin scripts SQL manuales

**Estado:** Aceptado

## Contexto
Necesitamos gestionar el esquema de PostgreSQL de forma reproducible en
desarrollo, CI y despliegues. La arquitectura (`03-arquitectura.md` §2.1)
había dejado abierta la elección entre `sqlc` (SQL-first) y GORM (ORM) para
la capa de repositorios.

## Decisión
Se usa **GORM** como ORM y capa de acceso a datos, con `AutoMigrate` para
gestionar el esquema a partir de los modelos Go — no se escriben migraciones
SQL manuales ni se usa una herramienta aparte tipo `golang-migrate`.

El binario del backend acepta un flag de arranque `--migrations-only`:
cuando se pasa, el proceso se conecta a la base de datos, corre
`AutoMigrate` sobre todos los modelos, y termina sin levantar el servidor
HTTP. Está pensado para correr como Job de Kubernetes (o paso de un
pipeline) antes de desplegar una nueva versión de la API — equivalente en
intención a `dotnet ef database update`, pero como parte del mismo binario
en vez de una herramienta de CLI separada.

## Consecuencias
- El esquema queda acoplado a la definición de los modelos Go (un único
  lugar de verdad, sin archivos `.sql` que puedan desincronizarse).
- Se pierde el control fino de una migración escrita a mano (backfills
  complejos, cambios de tipo no triviales, renombrados sin pérdida de
  datos): `AutoMigrate` solo agrega columnas/índices/tablas faltantes, no
  los elimina ni renombra. Esos casos se resuelven con lógica adicional en
  Go si aparecen, no con SQL suelto.
- Reemplaza el plan original de usar `golang-migrate` con scripts
  `.up.sql`/`.down.sql` (tarea T-05 del backlog de Sprint 0,
  `10-backlog-sprint-0.md`).
