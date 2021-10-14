# LineChatBot
Golang LineChatbot

/////////Install
-Line golang sdk
-Echo

//Fisrt Run
$go mod init
$go mod tidy
$go mod vendor


//Run Using
$go run *.go
*Run ngrok that connect to the opening port to expose port to internet



/////////For docker
//pull image
$docker run -dp 1000:2000 minhminhh12/minhlinebot:minh-chat-bot

//get image ID
docker images

//run image
$docker run {image ID}


////////for mysql docker

//create user to connect via all ip address
CREATE USER 'admin'@'%' IDENTIFIED BY 'the_secure_password';
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%';
FLUSH PRIVILEGES;

////////ngrok for mysqlDocker
./ngork tcp -ap 3306
then get an public ip and port 