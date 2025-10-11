FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./main.go

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=build /out/server /app/server

EXPOSE 80

USER nonroot:nonroot

ENTRYPOINT ["/app/server"]