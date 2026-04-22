# Series Tracker Backend

REST API for managing TV series. Returns data in JSON format and is consumed by a separate client.

## Technologies

- Go (net/http)
- PostgreSQL

## How to run the project

1. Clone the repository
2. Make sure PostgreSQL is running
3. Create the database table:

```sql
CREATE TABLE series (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    image TEXT
);
```

4. Run the server:

```bash
go run cmd/main.go
```
The server runs at
```bash
http://localhost:3000
```

## Endpoints
### GET /series

Return all series

**Optional query parameters:**

`q` -> search by name
`sort` -> field (id, name)
`order` -> ascending or descending

### GET /series/:id

Returns a series by id

### POST /series

Creates a new series

Request body:

```JSON
{
  "name": "Dark",
  "description": "Time travel",
  "image": "url"
}
```

### PUT /series/:id

Updates existing series

### DELETE /series/:id

Deletes a series

## CORS

CORS is a browser security policy that restricts requests between different origins.

The server is configured to allow all origins during development using:

`Access-Control-Allow-Origin: *`

## Status

Functional API with full CRUD, search and sorting are implemented.


