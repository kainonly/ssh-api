FROM alpine:edge

COPY dist /app
WORKDIR /app

EXPOSE 3000

VOLUME [ "app/data" ]

CMD [ "./ssh-api" ]