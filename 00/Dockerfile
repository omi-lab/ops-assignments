FROM golang:1.22.3-alpine3.19 as builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY main.go main.go

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build -o /build/api main.go

FROM scratch

COPY --from=builder /build/api /build/api 

ENTRYPOINT [ "/build/api" ]

CMD [ "--help" ]