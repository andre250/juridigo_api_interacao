FROM golang:1.8 as builder

LABEL MAINTAINER="GuilhermeCaruso"
LABEL COMPANY="Juridigo"

WORKDIR /go/src/github.com/juridigo/juridigo_api_interacao

COPY . ./

RUN apt-get update
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go build

FROM golang:1.8
WORKDIR /go/src/github.com/juridigo/juridigo_api_interacao
COPY --from=builder /go/src/github.com/juridigo/juridigo_api_interacao/juridigo_api_interacao .

CMD /bin/bash -c "./juridigo_api_interacao"