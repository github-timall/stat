FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN go get -u -d github.com/mattes/migrate/cli github.com/lib/pq
RUN go build -tags 'postgres' -o /usr/local/bin/migrate github.com/mattes/migrate/cli

RUN mkdir /go/src/github.com/mattes/migrate/cli/migrations

COPY . /go/src/github.com/mattes/migrate/cli/migrations

CMD migrate -path /go/src/github.com/mattes/migrate/cli/migrations -database postgres://"$STAT_DB_USER":"$STAT_DB_PASSWORD"@"$STAT_DB_HOST":"$STAT_DB_PORT"/"$STAT_DB_NAME"?sslmode=disable up