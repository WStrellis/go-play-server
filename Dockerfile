FROM golang:1.16 AS buildserver
WORKDIR /srv
COPY . ./
RUN GOOS=linux GOARCH=amd64 go build -o app_server main.go


FROM ubuntu:latest
WORKDIR /srv
COPY ./index.html ./
COPY --from=buildserver /srv/app_server ./

EXPOSE 80

CMD ["./app_server"] 