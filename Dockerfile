FROM golang:1.7

MAINTAINER david.guo18@yahoo.com

COPY ./ /go/src/stress_test/

WORKDIR /go/src/stress_test/

RUN go get gopkg.in/yaml.v2 && \
	go get && \
    go build

EXPOSE 80 8089

CMD ["go run", "main.go"]