sudo mkdir -p /storage/docker/mysql-data

docker run -d --name [container_name] -e "MYSQL_ROOT_PASSWORD=pass" -p 3306:3306
-v /root/docker/[container_name]/conf.d:/etc/mysql/conf.d 
-v /storage/docker/mysql-data:/var/lib/mysql 
mysql:latest

resource: https://hevodata.com/learn/docker-mysql/