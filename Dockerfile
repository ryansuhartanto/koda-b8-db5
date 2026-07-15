FROM golang:alpine AS build

WORKDIR /var/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM postgres:alpine

ENV POSTGRES_USER=app
ENV POSTGRES_PASSWORD=app

COPY init.sql /docker-entrypoint-initdb.d
COPY --from=build /var/app/main /bin/app

WORKDIR /var/app
COPY --chmod=755 entrypoint.sh .

CMD ["/var/app/entrypoint.sh"]
