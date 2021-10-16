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




////////ngrok for mysqlDocker
./ngork tcp -ap 3306
then get an public ip and port 

////////////////////////////Mysql Setup step
1.pull docker images
docker pull mysql/mysql-server:latest
2. run image
   docker run -p 3306:3306 -d --name=mysql mysql/mysql-server:latest

2.1 to get container id
docker ps

3. get logs
   docker logs mysql

then save generate root password in logs

4.exececute command in docker
docker exec -it mysql bash

5.log in to db
mysql -uroot -password

5.1 add user to mysql
mysql> CREATE USER 'userName'@'%' IDENTIFIED BY 'PASSWORD';
mysql> GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
mysql> FLUSH PRIVILEGES;

5.2 create user to connect via all ip address
CREATE USER 'userName'@'%' IDENTIFIED BY 'PASSWORD';
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%';
FLUSH PRIVILEGES;

6 change user password
ALTER USER 'userName'@'localhost' IDENTIFIED BY 'newpassword';

7 To run images
docker exec -it {container id} mysql -u{username} -p{password}

