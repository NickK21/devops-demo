FROM golang:1.22-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG GIT_SHA=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=${GIT_SHA}" -o /out/server ./main.go

FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=build /out/server /app/server

EXPOSE 80
ENTRYPOINT ["/app/server"]