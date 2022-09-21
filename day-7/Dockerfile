FROM golang:1.18-alpine AS build

WORKDIR /app

COPY ./go.* .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch AS prod

COPY --from=build /bin/app /bin/app

EXPOSE 3000

CMD [ "/bin/app"]