FROM golang:1.11.5 as builder

WORKDIR /src/github.com/apstndb/go-httptoolbox
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v -o httptoolbox

FROM alpine

RUN apk add --no-cache curl

COPY --from=builder /src/github.com/apstndb/go-httptoolbox/httptoolbox /httptoolbox

ENTRYPOINT ["/httptoolbox"]
EXPOSE 8080