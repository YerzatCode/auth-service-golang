FROM golang:1.21.5-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["/go.mod","/go.sum" , "./"]

RUN go mod download

ENV CONFIG_PATH=./config.yaml

COPY ./ ./
RUN go build -o ./bin/app cmd/main.go

FROM alpine  AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY configs/config.yaml /config.yaml

CMD ["/app"]