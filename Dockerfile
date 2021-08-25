FROM golang:1.16.2-alpine
ENV APP_HOME /go/src/minh-chat-bot
WORKDIR $APP_HOME
#WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
##COPY *.go ./
##COPY /test/ test/
COPY . $APP_HOME
RUN go build -o /minh-chat-bot
EXPOSE 1323
CMD [ "/minh-chat-bot" ]