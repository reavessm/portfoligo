FROM docker.io/library/golang:alpine as builder

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux go build -buildvcs=false -a -ldflags '-extldflags "-static"' -o main .

FROM scratch

COPY --from=builder /build/main /app/
COPY --from=builder /build/template /app/template/
COPY --from=builder /build/config /app/config/
COPY --from=builder /build/public/static /app/public/static

WORKDIR /app

EXPOSE 8080

CMD ["./main"]
