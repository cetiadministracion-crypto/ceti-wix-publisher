# CETI Wix Publisher retirado

Retirado el 2026-10-03 tras la migración del sitio y del blog a Next.js y
PostgreSQL. Wix permanece como proveedor DNS; este publicador no administra DNS.

Se eliminaron de main el código Go, dependencias, ejemplo e instrucciones de
publicación Wix. El directorio posts solo contenía un archivo vacío .gitkeep;
no había artículos de producción en el árbol revisado.

El blog se administra con [ceti-website](https://github.com/cetiadministracion-crypto/ceti-website).
El código retirado permanece en la historia Git. Este cambio no elimina contenido
ni credenciales en Wix y no modifica ceti-production.
