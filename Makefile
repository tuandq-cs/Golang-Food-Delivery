setupdb:
	docker run -d --name Golang_Edu -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE="food_delivery" -e MYSQL_USER=root -e MYSQL_PASSWORD=12345 -e MYSQL_COLLATE= -p 3307:3306  --privileged=true bitnami/mysql:5.7
migrateup:
	migrate -path db/migration -database "mysql://root:12345@tcp(127.0.0.1:3307)/food_delivery" -verbose up
migratedown:
	migrate -path db/migration -database "mysql://root:12345@tcp(127.0.0.1:3307)/food_delivery" -verbose down
.PHONY: setupdb migrateup migratedown