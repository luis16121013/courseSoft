#go parameters
GOCMD=go
PATHCMD=./cmd/main.go

run:
	@$(GOCMD) run $(PATHCMD)

#commands docker-compose
up:
	@docker-compose up -d
down:
	@docker-compose down
ps:
	@docker-compose ps

#commands docker
mysql:
	@docker exec -it mysqlService mysql -p
