FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

ENV CLIENT_ID p@ssw0rd!
ENV CONN_STRING mongodb+srv://user:pass@127.0.0.1/bdname
ENV DB_HOST localhost
ENV DB_NAME dbname
ENV DB_PASS password
ENV DB_PORT 5432
ENV DB_TYPE postgres
ENV DB_USER user
ENV DEBUG true
ENV JWT_SECRET Aech7eepaesi8goo8phu8laech8aet4yie1kahsa4phohLuiHu9aeph6oa9Eoth7
ENV PAGE_SIZE 25
ENV PORT 8080
ENV READ_TIMEOUT 60
ENV REDIS_URL 172.0.17.2
ENV RUN_MODE debug
ENV WRITE_TIMEOUT 60


RUN go mod download