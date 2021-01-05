FROM alpine:3.8

WORKDIR /opt/pilosa/

COPY ./assets/ .

EXPOSE 8000

CMD ["./pilosa-console"]