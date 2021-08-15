FROM golang:1.15-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build/kodoo

COPY go.mod ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build

FROM alpine
WORKDIR /app
COPY --from=builder /build/kodoo/kodoo .
RUN chmod a+x kodoo && mkdir -p /.kodoo && chown -R 1000:1000 /.kodoo
USER 1000
ENTRYPOINT [ "./kodoo" ] 