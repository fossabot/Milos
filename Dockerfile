FROM golang:alpine as builder

RUN apk --no-cache add git \
    && cd /go/src \
    && git clone https://github.com/coolrc136/Milos.git \
    && cd Milos \
    && go get ./... \
    && CGO_ENABLED=0 GOOS=linux go build


FROM alpine


COPY --from=0 /go/src/Milos/ /Milos/

RUN apk upgrade && apk add --no-cache ca-certificates

CMD cd /Milos && /Milos/Milos
EXPOSE 8080
