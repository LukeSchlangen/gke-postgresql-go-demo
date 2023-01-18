# Dockerfile
FROM golang:1.19-buster
WORKDIR /app
COPY . ./
# Could remove download here and add later
RUN go mod download
RUN go build -v -o server .
CMD ["/app/server"]