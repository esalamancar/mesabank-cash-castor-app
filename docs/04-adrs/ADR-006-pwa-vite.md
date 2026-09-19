# ADR-006: PWA con Vite y Service Worker

**Estado:** Aceptado

## Contexto
La app debe funcionar en móviles sin necesidad de tiendas de aplicaciones.
Debe tolerar conexiones lentas.

## Decisión
React + Vite + PWA plugin. Service Worker cachea assets estáticos y colas de
operaciones offline (fase 2).

## Consecuencias
Instalable desde el navegador; no requiere App Store; el Service Worker añade
complejidad de caché.
