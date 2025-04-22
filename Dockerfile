FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/app /app

USER nonroot:nonroot

ENTRYPOINT ["/app"]

