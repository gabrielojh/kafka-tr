FROM --platform=linux/$TARGETARCH golang:alpine3.18

ARG TARGETARCH
RUN echo $TARGETARCH

RUN apk update
RUN apk add \
    gcc \
    musl-dev \
    # explicitly install SASL package
    cyrus-sasl-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=1 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=$TARGETARCH \
    CGO_LDFLAGS="-lsasl2" \
    go build -mod=vendor -o /main -tags musl -ldflags "-w -s" .

CMD ["/main"]