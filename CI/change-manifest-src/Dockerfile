FROM golang:1.22.0-bookworm as build 

COPY ./change-manifest/go.mod ./change-manifest/go.sum ./change-manifest/change-manifest.go src/change-manifest/
COPY ./git  src/git/
COPY ./manifest  src/manifest/
WORKDIR src/change-manifest 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/src/change-manifest/main .
ENTRYPOINT ["./main"]
CMD ["-token=","-repoUrl=","-manifestfile=","-imageRef="]







