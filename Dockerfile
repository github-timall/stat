FROM golang:alpine

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    mkdir -p /go/src/gitlab.com/drcash/stat

COPY . /go/src/gitlab.com/drcash/stat

WORKDIR /go/src/gitlab.com/drcash/stat

RUN go-wrapper download && go-wrapper install

CMD ["go-wrapper", "run"]

EXPOSE 8080