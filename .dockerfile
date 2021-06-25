FROM golang:1.14

ARG SERVICE_NAME
ENV ENV_SERVICE_NAME=$SERVICE_NAME

RUN go get github.com/cespare/reflex

WORKDIR /go/src/source-code
RUN echo "-r '(\.go$|go\.mod)' -s go run ./$ENV_SERVICE_NAME/" > /reflex.conf
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]