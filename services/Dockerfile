FROM golang:1.13-alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG project
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $project

FROM busybox:latest

WORKDIR /dist

COPY --from=builder /build/server .

CMD [ "./server" ]