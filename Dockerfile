FROM docker.io/library/golang:1.23rc1-alpine AS lgo

WORKDIR /back

COPY ./backend .

RUN go mod download

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./cideclasse"]
CMD ["./cideclasse"]
