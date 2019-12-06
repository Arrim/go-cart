ARG build_image=golang:alpine
ARG base_image=alpine:latest

FROM ${build_image} as build_stage

WORKDIR /app

COPY . /app

ENV GO111MODULE=on
RUN go mod vendor
RUN go build -v -o out/go-cart

FROM ${base_image}
COPY --from=build_stage /app/out/go-cart /usr/local/bin
ADD config /config

USER nobody

# Default if /config/test.cfg
ENTRYPOINT /usr/local/bin/go-cart -config=/config/${CONFIG_FILE:-test.toml}