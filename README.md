# CRUD DB

Go and Database exercise implementing Create, Read, Update, and Delete on a Postgres database.

## Schema

```mermaid
erDiagram

contact {
    bigint id

    timestamp created_at
    timestamp updated_at

    string name
    date? dob
    string? address
    string? phone
    string? email
}
```

## License

[MIT](LICENSE)
