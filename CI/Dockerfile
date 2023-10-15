FROM golang:1.21.3-bullseye as base
COPY go.mod go.sum main.go src/
WORKDIR src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go mod download
RUN go build main.go
FROM gcr.io/distroless/static-debian11
COPY --from=base go/src/main .
ENTRYPOINT ./main
