FROM golang:latest as builder
LABEL maintainer="Mikhail Tikhonov <wbe7@ya.ru>"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest
WORKDIR /application
RUN adduser --uid 1001 --disabled-password --no-create-home big-boss big-boss
RUN chown -R big-boss:big-boss /application
USER big-boss
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]