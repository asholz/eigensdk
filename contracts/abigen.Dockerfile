# use the below command to build this docker image locally
# docker build -t abigen -f abigen.Dockerfile .
FROM golang:1.22 AS build

RUN git clone https://github.com/ethereum/go-ethereum.git && \
    cd go-ethereum/cmd/abigen && \
    go build .

FROM debian:latest
COPY --from=build /go/go-ethereum/cmd/abigen/abigen /usr/local/bin/abigen
ENTRYPOINT [ "abigen"]