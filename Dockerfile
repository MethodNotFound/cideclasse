FROM node:14 as build

WORKDIR /build

COPY ./frontend ./

RUN npm install

RUN npm run build

FROM docker.io/library/golang:1.23rc1-alpine AS lgo

WORKDIR /back

COPY --from=build /Bundle.js /back/frontend/Bundle.js

COPY . .

RUN go mod download

RUN go build .

EXPOSE 3000

ENTRYPOINT ["./main"]
CMD ["./main"]
