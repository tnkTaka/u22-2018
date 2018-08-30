DBNAME:=wanted
ENV:=development

run:  ## bulid main.go 実行
	go run app/main.go

serve: ## GAE 実行
    goapp serve app

deps: ## Set Up 時に実行
	which dep || go get -u github.com/golang/dep/cmd/dep
	which sql-migrate || go get -u github.com/rubenv/sql-migrate/...
	dep ensure

migrate/init: ## ローカル Mysql DB に wanted作成
	mysql -u root -h localhost --protocol tcp -e "CREATE DATABASE \`$(DBNAME)\` DEFAULT CHARACTER SET utf8"

migrate/up:　## wanted DB　に　migration
	sql-migrate up -env=$(ENV)
