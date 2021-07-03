FROM golang:1.14

RUN go get github.com/cespare/reflex

WORKDIR /go/src/github.com/erry-azh/mqtt-on-go
RUN echo "-r '(\.go$|go\.mod)' -s go run ./webhook/" > /reflex.conf
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]