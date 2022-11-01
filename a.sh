docker run --name git -it -v /Users/yisar/repo/:/git/repo -v /Users/yisar/id_rsa:/root/.ssh/id_rsa -v /Users/yisar/id_rsa.pub:/root/.ssh/id_rsa.pub wuliangxue/git:0.1
docker run -id -p 3306:3306 --name=mysql -v /Users/yisar/mysql/conf:/etc/mysql/conf.d -v /Users/yisar/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 mysql:latest

# 备份
docker exec mysql(容器名) sh -c 'exec mysqldump --all-databases -uroot -p123456 --all-databases' > /var/backup/music_`date +%F`.sql

# 还原
docker cp /var/backup/emp_2022-03-17.sql mysql:/var
sudo docker exec -it mysql bin/bash
mysql -uroot -p123456 < /var/backup/emp_2022-03-15.sql


