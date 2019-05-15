FROM golang:1.11.5 as builder

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/main

FROM alpine

RUN apk add --no-cache curl

COPY --from=builder /src/main /main

ENTRYPOINT ["/main"]
EXPOSE 8080