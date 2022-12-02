FROM golang:1.18.7-alpine

ARG SERVICE_NAME
ENV ENV_SERVICE_NAME=$SERVICE_NAME
ENV TZ=Asia/Jakarta

WORKDIR /go/src/github.com/erry-azh/mqtt-on-go

RUN apk add --update curl && \
    rm -rf /var/cache/apk/*

RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/v1.40.4/install.sh \
    && chmod +x install.sh && sh install.sh && mv ./bin/air /bin/air

RUN mkdir -p "/etc/air"
COPY .dockerfiles/air-conf.toml /etc/air/air-conf.toml
RUN sed -i "s/service_name/${SERVICE_NAME}/g" "/etc/air/air-conf.toml"
RUN cat "/etc/air/air-conf.toml"

ENTRYPOINT ["air", "-c", "/etc/air/air-conf.toml"]