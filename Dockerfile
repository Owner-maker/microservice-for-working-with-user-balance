FROM golang:1.19-alpine
WORKDIR /code
COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go build -o /app
CMD ["/app"]