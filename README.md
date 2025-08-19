# 🖼️ Image Processing Service

Este proyecto es una API REST escrita en Go para procesamiento de imágenes, con arquitectura modular y soporte para operaciones como registro, login, subida, búsqueda, paginación y transformación de imágenes. 

Los detalles de los requerimientos fueron extraídos de [Roadmap.sh](https://github.com/gin-gonic/gin).

## ✨ Características principales
- **Registro y autenticación de usuarios** (bcrypt + JWT)
- **Subida de imágenes**
- **Búsqueda y paginación de imágenes**
- **Transformaciones**: resize, crop, rotate, format
- **Almacenamiento local y SQLite**
- **Event Bus** para eventos internos
- **Arquitectura limpia**: separación en application, domain, infrastructure

## 📁 Estructura de carpetas
```
├── find/
├── images/
├── init/
├── login/
├── paginate/
├── register/
├── shared/
├── tests/
├── transform/
├── upload/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
```

## ⚙️ Instalación
1. Clona el repositorio.
2. Construye y ejecuta la aplicación usando Docker Compose:
    ```bash
    docker-compose up -d
    ```
3. La API estará disponible en el puerto especificado en `docker-compose.yml` (por defecto, suele ser el 8080).

## 🧪 Tests

Para ejecutar los tests, primero accede al bash dentro del contenedor Docker:

```bash
docker-compose exec go-app bash
```

Luego, ejecuta el comando:

```bash
tests
```

## 🌐 Endpoints
- `POST /register`
- `POST /login`
- `POST /upload`
- `GET /find/:uuid`
- `GET /images?page=N&limit=N`
- `POST /transform`

## 🏗️ Arquitectura
Consulta la documentación de la arquitectura en la carpeta `docs/`.

## 📦 Dependencias principales
- [Gin](https://github.com/gin-gonic/gin)
- [bimg](https://github.com/h2non/bimg)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [SQLite](https://www.sqlite.org/index.html)
