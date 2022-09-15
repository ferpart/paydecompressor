FROM golang:buster as build

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o bin/paydecompressor .

FROM scratch

COPY --from=build /app/bin .
EXPOSE 8082/tcp
CMD ["./paydecompressor"]

