FROM golang:1.14-alpine AS build-env

RUN apk --no-cache add git
ADD . /src
ENV GOPROXY=direct
RUN cd /src && \
    go build -o app

FROM alpine
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ["/app/app"]