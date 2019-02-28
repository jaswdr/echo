FROM golang as builder

WORKDIR /app
COPY ./main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

FROM scratch
COPY --from=builder /app/app /
ENTRYPOINT ["/app"]
