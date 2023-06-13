mysqlinit:mysqlstop
	docker run --name task-schedular-mysql-db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 -d mysql
mymongoinit:mongostop
	docker run --name my-mongo-container -d -p 27017:27017 mongo
	sleep 15
createmysqldb:
	 docker exec -it task-schedular-mysql-db mysql -u root -p -e "CREATE DATABASE task_schedular;"
dropmysqldn:
	 docker exec -it task-schedular-mysql-db mysql -u root -p -e "DROP DATABASE task_schedular;"
mysqlstop:
	bash ./scripts/mysql_stop.sh
mongostop:
	bash ./scripts/mongo_stop.sh
db_mysql_prepare:mysqlinit
	docker cp task_schedular.sql task-schedular-mysql-db:task_schedular.sql
	echo "Executing databases...wait for 15 seconds"
	sleep 15
	docker exec -i task-schedular-mysql-db sh -c 'mysql -uroot < task_schedular.sql'
db_mongo_prepare: mymongoinit
	docker exec -it my-mongo-container mongosh --eval "use task_schedular" --eval "db.createCollection('tasks')"
test: db_mongo_prepare db_mysql_prepare
	bash ./scripts/test.sh

.PHONY: mysqlinit mymongoinit createmysqldb dropmysqldn mysqlstop mongostop test
