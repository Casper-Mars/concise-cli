package template

func NewMakefile() string {
	return `# Should be edit by user
# local
db_username = root
db_password = root
db_url = localhost
db_port = 3306
db_name = demo
db_container_name = test-mysql
redis_container_name = redis
redis_port = 6379
rabbitmq_container_name = rabbitmq
rabbitmq_port = 5672
rabbitmq_username = zhisheng
rabbitmq_password = zhisheng

# test
test_db_username = root
test_db_password = !Zhisheng2020
test_db_url = pv.zhisheng.com
test_db_port = 30306
test_db_name = cloud_restaurant

# image info
# User should set this var
IMAGE_NAME=

define run-mysql
	DB_USERNAME=${db_username} \
	DB_PASSWORD=${db_password} \
	DB_PORT=${db_port} \
	DB_CONTAINER_NAME=${db_container_name} \
	DB_DBS=demo \
	/bin/sh hack/mysql.sh
endef

define run-redis
	CONTAINER_NAME=${redis_container_name} \
	PORT=${redis_port} \
	/bin/sh hack/redis.sh
endef

define run-rabbitmq
	CONTAINER_NAME=${rabbitmq_container_name} \
	PORT=${rabbitmq_port} \
	USER=${rabbitmq_username} \
	PASSWORD=${rabbitmq_password} \
	/bin/sh hack/rabbitmq.sh
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

define migrate-up-test-db
	migrate -path db/test/migration \
	-database 'mysql://$(test_db_username):$(test_db_password)@tcp($(test_db_url):$(test_db_port))/${test_db_name}' \
	-verbose \
	up $(1)
endef

define migrate-down-test-db
	migrate -path db/test/migration \
	-database 'mysql://$(test_db_username):$(test_db_password)@tcp($(test_db_url):$(test_db_port))/${test_db_name}' \
	-verbose \
	down $(1)
endef

define start-test-container
    	docker start ${db_container_name}
    	docker start ${redis_container_name}
    	docker start ${rabbitmq_container_name}
endef

define delete-mysql
	docker stop ${db_container_name} && docker container rm ${db_container_name}
endef

define delete-redis
	docker stop ${redis_container_name} && docker rm ${redis_container_name}
endef

define delete-rabbitmq
	docker stop ${rabbitmq_container_name} && docker rm ${rabbitmq_container_name}
endef


define test-container
	docker run -i --rm --privileged \
	 --name test-env \
	 -e MYSQL_SERVICE=on \
	 -e REDIS_SERVICE=on \
	 -e RABBITMQ_SERVICE=on \
	 -w /work \
	 --mount type=bind,source=${PWD},target=/work \
	 harbor.zhisheng.com:5000/public/it-env:latest \
	 /bin/sh -c $(1)
endef

# 帮助提示
.PHONY: help
help:
	echo "Nothing happen"

# 安装到本地的命令
.PHONY: install
install:
	mvn -Dmaven.test.skip=true install

# 打包命令
.PHONY: package
package:
	mvn clean package -Dmaven.test.skip=true

# 构建命令
.PHONY: build
build:
	docker build -t ${IMAGE_NAME} .
	docker push ${IMAGE_NAME}

# 部署命令
.PHONY: deploy
deploy:
	sed -i "s/VERSION_PLACEHOLDER/${CI_COMMIT_SHA}/g" k8s/dp.yaml
	wget -O ~/.kube/config ${K8S_CONFIG}
	kubectl apply -f k8s/dp.yaml



# ----------------------------本地开发常用脚本------------------------------
# 建立本地开发环境
.PHONY: setup-local-test-env
setup-local-test-env:
	$(call run-mysql)
	$(call run-redis)
	$(call run-rabbitmq)
	$(call migrate-up-db)


# 清除本地开发环境
.PHONY: delete-local-test-env
delete-local-test-env:
	$(call delete-mysql)
	$(call delete-redis)
	$(call delete-rabbitmq)

# 启动测试环境
.PHONY: start-test-env
start-test-env:
	$(call start-test-container)

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

# 初始化测试环境的数据库
.PHONY: migrate-test-up
migrate-test-up:
	$(call migrate-up-test-db,$(version))

# 清理测试环境的数据库
.PHONY: migrate-test-down
migrate-test-down:
	$(call migrate-down-test-db,$(version))

# ----------------------------本地开发常用脚本 end------------------------------

# ----------------------------测试脚本-------------------------------

# 使用本地环境进行测试
.PHONY: test-local
test-local:
	echo "配置本地测试环境"
	make setup-local-test-env
	mvn clean
	make test-ut
	make test-it

# 进行单元测试，默认本地已经有测试环境
.PHONY: test-ut
test-ut:
	echo "进行单元测试"
	mvn test -Ptest_ut

# 进行集成测试，默认本地已经有测试环境
.PHONY: test-it
test-it:
	echo "进行集成测试"
	mvn integration-test -Ptest_it

# 使用集成环境进行测试
.PHONY: test
test:
	$(call test-container,'make test-local')

# 使用集成环境进行测试并清理测试后的结果
.PHONY: test-with-clean
test-with-clean:
	$(call test-container,'make test-local && mvn clean')

# ----------------------------测试脚本 end-------------------------------
`
}
