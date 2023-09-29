FROM golang:alpine AS builder

WORKDIR /go/src/app

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# # Declare arguments for environment variables
# ARG DB_USER
# ARG DB_PASSWORD
# ARG DB_HOST
# ARG DB_PORT
# ARG DB_NAME
# ARG JWT_SECRET

# # Bake the environment variables into the image with this:
# # docker build \
# #   --build-arg DB_USER=root \
# #   --build-arg DB_PASSWORD=123qwe \
# #   --build-arg DB_HOST=127.0.0.1 \
# #   --build-arg DB_PORT=3306 \
# #   --build-arg DB_NAME=session1_xx \
# #   --build-arg JWT_SECRET=wteqpiwqetcnyipyzsdbskndfmqwf \
# #   -t image-name .


ENV DB_USER=root
ENV DB_PASSWORD=123qwe
ENV DB_HOST=127.0.0.1
ENV DB_PORT=3306
ENV DB_NAME=session1_xx
ENV JWT_SECRET=wteqpiwqetcnyipyzsdbskndfmqwf

RUN go build -o ./run .

FROM alpine:latest
WORKDIR /root/

# Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8086
CMD ["./run"]
