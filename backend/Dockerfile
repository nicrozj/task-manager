FROM golang:1.24.1-alpine AS dev
WORKDIR /app

COPY wait-for-postgres.sh /wait-for-postgres.sh
RUN chmod +x /wait-for-postgres.sh

RUN go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
ENTRYPOINT ["/wait-for-postgres.sh"]
CMD ["air"]