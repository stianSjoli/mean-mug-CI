FROM golang:1.21.3-bullseye as base
COPY main.go src/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR src
RUN go build main.go
ENTRYPOINT ./main
