# Microservicio de Lectura Individual de Libros

## Desarrollador
Ricardo Florez

## Descripción
Este microservicio es responsable de obtener la información detallada de un libro específico por su ID. Forma parte de una arquitectura de microservicios para la gestión de una biblioteca digital.

## Características
- Implementado en Go 1.21
- Arquitectura limpia (Clean Architecture)
- Endpoints RESTful
- Integración con MongoDB
- Despliegue en Kubernetes

## Estructura del Proyecto
```
.
├── config/         # Configuración de la base de datos
├── models/         # Modelos de datos
├── repositories/   # Capa de acceso a datos
├── services/      # Lógica de negocio
├── k8s/           # Configuración de Kubernetes
└── tests/         # Pruebas unitarias
```

## API Endpoint
- **GET** `/libros/{id}`
  - Puerto: 30083 (NodePort)
  - Retorna un libro específico por su ID

### Ejemplo de Respuesta
```json
{
    "id": "5f7b5e1b9d3e2a1b4c7d8e9f",
    "titulo": "El Quijote",
    "autor": "Miguel de Cervantes"
}
```

### Respuestas de Error
```json
{
    "error": "Libro no encontrado"
}
```

## Configuración Kubernetes
- Deployment con 3 réplicas
- Service tipo NodePort (30083)
- Conexión a MongoDB mediante Service Discovery

## Variables de Entorno
- MONGODB_URI: URI de conexión a MongoDB

## Despliegue
```bash
# Construir la imagen
docker build -t libro-read-one:latest .

# Desplegar en Kubernetes
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## Pruebas
```bash
# Ejecutar pruebas unitarias
go test ./...

# Probar el endpoint
curl http://localhost:30083/libros/5f7b5e1b9d3e2a1b4c7d8e9f
```

## Monitoreo
El servicio puede ser monitoreado mediante:
- Logs de Kubernetes
- Métricas de contenedor
- Estado del Service y Deployment

## Manejo de Errores
- Validación de ID MongoDB
- Manejo de casos cuando el libro no existe
- Respuestas HTTP apropiadas para cada caso de error 