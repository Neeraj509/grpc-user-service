FROM golang:1.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /user_service

EXPOSE 50051

CMD [ "/user_service" ]
