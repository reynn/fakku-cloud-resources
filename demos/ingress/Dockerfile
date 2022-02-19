FROM golang:1.17 as BUILDER

WORKDIR /app
COPY go.mod main.go ./
RUN go build -o app main.go

FROM debian

COPY --from=BUILDER /app/app /usr/bin/app

EXPOSE 8000

ENTRYPOINT [ "/usr/bin/app" ]
