FROM golang:1.23.4-alpine3.19 AS builder

RUN apk --update add git make openssh

RUN addgroup -g 1001 -S builder && \
    adduser -u 1001 -S builder -G builder
USER builder:builder

ENV APP_HOME /home/builder
WORKDIR $APP_HOME

COPY --chown=builder:builder . $APP_HOME/

# build executable to ./bin/app
RUN make build-native

RUN ls -la bin/app

FROM public.ecr.aws/lambda/provided:al2023


ENV BUILD_DIR /home/builder
# create dummy .env file
ENV APP_HOME /home/runner/bin
# create dummy .env file
RUN mkdir -p $APP_HOME && touch $APP_HOME/.env

WORKDIR $APP_HOME

COPY --from=builder $BUILD_DIR/bin .
COPY --from=builder $BUILD_DIR/resources ./resources


ENTRYPOINT ["./app", "http"]
