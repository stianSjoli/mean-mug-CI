FROM golang:1.21.3-bullseye as base
COPY go.mod go.sum main.go src/

RUN apt-get update
RUN apt-get install -y curl
RUN curl https://get.docker.com/builds/Linux/x86_64/docker-latest.tgz | tar xvz -C /tmp/ && mv /tmp/docker/docker /usr/bin/docker
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR src
RUN go mod download
RUN go build main.go
#FROM gcr.io/distroless/static-debian11
#COPY --from=base go/src/main .
ENTRYPOINT ./main
