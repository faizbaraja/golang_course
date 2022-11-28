FROM golang:1.19.3-alpine

ARG PORT

WORKDIR /app

COPY . .

# COPY main.go /app/main.go
# COPY go.mod /app/go.mod
# COPY .env /app/.env

RUN go mod vendor
RUN go mod tidy

RUN go build -o /app/batch6 cmd/api/main.go

EXPOSE ${PORT}

ENTRYPOINT [ "/app/batch6" ]
# CMD ["go", "run", "/app/main.go"] 

# docker build --tag batch6-faiz:1.1 .
# docker container create --name batch6-faiz -p 8080:8080 batch6-faiz:1.1
# docker container start batch6-faiz
#docker container ls --all
#docker tag batch6-faiz:1.1 faizbaraja/batch6-faiz:1.1
#docker push faizbaraja/batch6-faiz:1.1