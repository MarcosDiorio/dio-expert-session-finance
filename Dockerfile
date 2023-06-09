FROM golang:1.20-alpine


WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY  . .

RUN CGO_ENABLE=0 go build -v -a -installsuffix cgo -o finance ./cmd/server

EXPOSE 8080

CMD [ "./finance" ]