FROM golang:1.16.2-alpine
WORKDIR $APP_HOME
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /MinhChatBotDoc
CMD [ "/MinhChatBotDoc" ]