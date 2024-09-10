FROM node:20 as build

WORKDIR /build

COPY . .

RUN npm install

RUN npm run build

FROM docker.io/library/golang:1.23rc1-alpine AS lgo

WORKDIR /back
COPY . .
RUN rm -rf ./backend/public
COPY --from=build /build/public /back/public

RUN cd backend && go mod download
RUN cd backend && go build .

EXPOSE 3000
ENTRYPOINT ["./backend/cideclasse"]
CMD ["./backed/cideclasse"]
