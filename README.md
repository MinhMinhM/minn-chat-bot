# LineChatBot
Golang LineChatbot

//Install
-Line golang sdk
-Echo

//Fisrt Run
$go mod init
$go mod tidy
$go mod vendor


//Run Using
$go run *.go
*Run ngrok that connect to the opening port



//For docker
//pull image
$docker run -dp 1000:2000 minhminhh12/minhlinebot:minh-chat-bot

//get image ID
docker images

//run image
$docker run {image ID}