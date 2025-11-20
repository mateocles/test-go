# API REST GO (Gin)

Ejemplo de API REST en Go usando Gin con CRUD de álbumes en memoria.

## Endpoints

| Método | Ruta        | Descripción             | Código éxito |
| ------ | ----------- | ----------------------- | ------------ |
| GET    | /albums     | Lista todos los álbumes | 200          |
| GET    | /albums/:id | Obtiene un álbum        | 200          |
| POST   | /albums     | Crea un álbum           | 201          |
| PUT    | /albums/:id | Actualiza un álbum      | 200          |
| DELETE | /albums/:id | Elimina un álbum        | 204          |

## Ejemplos curl

Listar:

```bash
curl -s http://localhost:8080/albums | jq
```

Crear:

```bash
curl -s -X POST http://localhost:8080/albums \
	-H 'Content-Type: application/json' \
	-d '{"title":"Kind of Blue","artist":"Miles Davis","year":1959}' | jq
```

Obtener:

```bash
curl -s http://localhost:8080/albums/4 | jq
```

Actualizar:

```bash
curl -s -X PUT http://localhost:8080/albums/4 \
	-H 'Content-Type: application/json' \
	-d '{"title":"Kind of Blue (Remastered)","artist":"Miles Davis","year":1959}' | jq
```

Eliminar:

```bash
curl -i -X DELETE http://localhost:8080/albums/4
```

## Ejecución local

```bash
go build -tags netgo -ldflags '-s -w' -o app
./app
```

La aplicación escucha en el puerto definido por `PORT` o `8080` por defecto.

## Notas

- Almacenamiento en memoria (no persistente).
- Mutex para concurrencia segura.
- Validación mínima: title, artist, year requeridos.

# test-go
