# CRUD DB

Go and Database exercise implementing Create, Read, Update, and Delete on a Postgres database.

## Running

You can either manually setup the program or just use our image by running:

```sh
docker run -it ghcr.io/ryansuhartanto/koda-b8-db5
```

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
