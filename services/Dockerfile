FROM golang:1.13-alpine as builder

WORKDIR /build

ARG project

COPY $project/go.mod $project/go.sum ./
RUN go mod download

COPY $project .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./cmd/main.go

FROM busybox:latest

WORKDIR /dist

COPY --from=builder /build/server .

CMD [ "./server" ]