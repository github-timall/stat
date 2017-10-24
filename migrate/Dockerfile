FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

RUN mkdir /go/src/github.com/mattes/migrate/cli/migrations

COPY . /go/src/github.com/mattes/migrate/cli/migrations

CMD ["migrate", "-path", "/go/src/github.com/mattes/migrate/cli/migrations", "-database", "postgres://stat:stat@stat-database:5432/stat?sslmode=disable", "up"]