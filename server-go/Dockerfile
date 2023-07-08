FROM golang:1.20 as builder
ARG SERVICE_NAME
ENV APP_NAME=$SERVICE_NAME

WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./src ./src
RUN CGO_ENABLED=0 go build -o ./bin/$APP_NAME ./src/cmd/$APP_NAME

FROM alpine:3.15
ARG SERVICE_NAME
ENV APP_NAME=$SERVICE_NAME

WORKDIR /app
COPY --from=builder /build/bin/$APP_NAME ./$APP_NAME

ENTRYPOINT ["/app/server"]