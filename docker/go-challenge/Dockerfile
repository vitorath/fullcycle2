FROM golang AS builder
ENV GOOS=linux
WORKDIR /usr/local/app
COPY . .
RUN go build main.go

FROM scratch
WORKDIR /usr/local/app
COPY --from=builder /usr/local/app/main .
CMD ["./main"]