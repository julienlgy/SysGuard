FROM golang:alpine3.11

COPY . /go

WORKDIR /go/src/sysguard/

RUN go build -o sysguard .
RUN ./sysguard --origin=http://127.0.0.1:80 --listen=http://0.0.0.0:8080
RUN ./sysguard --origin=http://srv-web --listen=http://0.0.0.0:80

CMD [ "./sysguard" ]