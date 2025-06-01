FROM golang:alpine as builder

WORKDIR /build

COPY . .

RUN go build -o nameDemo .

FROM scratch

COPY --from=builder /build/nameDemo /

ENTRYPOINT ["/nameDemo"]