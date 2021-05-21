package creator

import "os"

var makefileTemplate = `
db_username = root
db_password = root
db_url = localhost
db_port = 3306
db_name = demo
db_container_name = test-mysql


define run-mysql
	DB_USERNAME=${db_username} \
	DB_PASSWORD=${db_password} \
	DB_PORT=${db_port} \
	DB_CONTAINER_NAME=${db_container_name} \
	DB_DBS=demo \
	/bin/sh hack/mysql.sh
endef

define migrate-up-db
	migrate -path db/test/migration \
	-database 'mysql://$(db_username):$(db_password)@tcp($(db_url):$(db_port))/${db_name}' \
	-verbose \
	up $(1)
endef

define migrate-down-db
	migrate -path db/test/migration \
	-database 'mysql://$(db_username):$(db_password)@tcp($(db_url):$(db_port))/${db_name}' \
	-verbose \
	down $(1)
endef

define delete-mysql
	docker stop ${db_container_name} && docker container rm ${db_container_name}
endef

define test-container
	docker run -i --rm --privileged \
	 --name test-env \
	 -e MYSQL_SERVICE=on \
	 -w /work \
	 -v ${PWD}:/work \
	 harbor.zhisheng.com:5000/public/it-env \
	 /bin/sh -c $(1)
endef


# 安装到本地的命令
.PHONY: install
install:
	mvn -Dmaven.test.skip=true install

# 部署命令
.PHONY: deploy
deploy:
	mvn -Dmaven.test.skip=true deploy

# ----------------------------本地开发常用脚本------------------------------

# 建立本地开发环境
.PHONY: setup-local-test-env
setup-local-test-env:
	$(call run-mysql)
	$(call migrate-up-db)

# 清除本地开发环境
.PHONY: delete-local-test-env
delete-local-test-env:
	$(call delete-mysql)

# 初始化数据库
.PHONY: migrate-up
migrate-up:
	$(call migrate-up-db,$(version))

# 清理数据库
.PHONY: migrate-down
migrate-down:
	$(call migrate-down-db,$(version))

# 重置数据库
.PHONY: resetDB
resetDB:
	make migrate-down
	make migrate-up

# ----------------------------本地开发常用脚本 end------------------------------

# ----------------------------测试脚本-------------------------------

# 使用本地环境进行测试
.PHONY: test-local
test-local:
	mvn clean
	echo "配置本地测试环境"
	make setup-local-test-env
	make test-ut
	make test-it

# 进行单元测试，默认本地已经有测试环境
.PHONY: test-ut
test-ut:
	echo "进行单元测试"
	mvn test -Dspring.profiles.active=test -Ptest_ut

# 进行集成测试，默认本地已经有测试环境
.PHONY: test-it
test-it:
	echo "进行集成测试"
	mvn integration-test -Dspring.profiles.active=test -Ptest_it

# 使用集成环境进行测试
.PHONY: test
test:
	$(call test-container,"make test-local")

# 使用集成环境进行测试并清理测试后的结果
.PHONY: test-with-clean
test-with-clean:
	$(call test-container,"make test-local && mvn clean")

# ----------------------------测试脚本 end-------------------------------

`

func CreateMakefile(basePath string) error {
	makefile, err := os.Create(basePath + "/Makefile")
	if err != nil {
		return err
	}
	_, err = makefile.WriteString(makefileTemplate)
	if err != nil {
		return err
	}
	return makefile.Close()
}
