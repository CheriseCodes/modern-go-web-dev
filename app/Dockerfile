FROM golang:1.21.3-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o runners-app main.go
EXPOSE 8080
CMD [ "/app/runners-app" ]