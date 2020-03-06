FROM golang:1.14-alpine AS build-env

RUN apk --no-cache add git
ADD . /src
ENV GOPROXY=direct
RUN cd /src && \
    mkdir -p /src/bin && \
    go build -o ./bin/app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin/app /app/bin/
ENTRYPOINT ["/app/bin/app"]