FROM golang:1.14-alpine as build
# RUN apk add --no-cache git
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY publisher.go ./
RUN go build publisher.go

FROM alpine as runtime
COPY --from=build /src/publisher /app/publisher
EXPOSE 80
CMD [ "/app/publisher" ]