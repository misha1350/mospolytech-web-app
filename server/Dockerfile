FROM golang:alpine AS builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Declare arguments for environment variables
ARG DB_USER
ARG DB_PASSWORD
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME
ARG JWT_SECRET

# Bake the environment variables into the image with this:
# docker build \
#   --build-arg DB_USER=root \
#   --build-arg DB_PASSWORD=123qwe \
#   --build-arg DB_HOST=127.0.0.1 \
#   --build-arg DB_PORT=3306 \
#   --build-arg DB_NAME=session1_xx \
#   --build-arg JWT_SECRET=wteqpiwqetcnyipyzsdbskndfmqwf \
#   -t image-name .


ENV DB_USER=$DB_USER
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_HOST=$DB_HOST
ENV DB_PORT=$DB_PORT
ENV DB_NAME=$DB_NAME
ENV JWT_SECRET=$JWT_SECRET

RUN go build -o ./run .

FROM alpine:latest

# Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8086
CMD ["./run"]
