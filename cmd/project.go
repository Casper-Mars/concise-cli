/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/Casper-Mars/concise-cli/pkg/config"
	"github.com/Casper-Mars/concise-cli/pkg/dir"
	"github.com/Casper-Mars/concise-cli/pkg/file"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
)

var projectConfig = config.NewProjectConfig()

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "生成业务系统项目",
	Long: `使用时，必须指定父工程的版本、工程项目名称。
还可以指定项目需要用到的外部依赖的组件，例如：mysql、redis。目前支持的外部依赖组件由：mysql、redis、rabbitmq和minio
`,
	Run: createProject,
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.Flags().StringVarP(&projectConfig.ParentVersion, "parent", "p", "", "指定父项目的版本号")
	projectCmd.Flags().StringVarP(&projectConfig.Name, "name", "n", "", "指定项目名称")
	projectCmd.Flags().StringArrayVarP(&projectConfig.Dependence, "dependence", "d", []string{}, "指定需要使用的外部依赖")
}

func createProject(cmd *cobra.Command, args []string) {
	/*校验参数，工程项目名称和父工程版本号不能为空*/
	err := projectConfig.Check()
	if err != nil {
		log.Println(err.Error())
		return
	}
	group, _ := errgroup.WithContext(context.Background())
	/*生成目录*/
	dirTree := getProjectDirTree(projectConfig.Name)
	err = dir.Build([]byte(dirTree), ".")
	if err != nil {
		log.Fatalln(err)
	}
	rootPath := "./" + projectConfig.Name
	/*初始化关键文件*/
	group.Go(func() error {
		/*创建pom文件*/
		pom := file.NewPom("com.zhisheng.framework.concise", kitProjectConfig.Name, "0.1.0")
		pom.InitParent(fmt.Sprintf(`
    <parent>
        <groupId>com.zhisheng.framework.concise</groupId>
        <artifactId>parent</artifactId>
        <version>%s</version>
    </parent>
`, kitProjectConfig.ParentVersion))
		return pom.BuildFile(rootPath)
	})
	/*创建makefile*/
	group.Go(func() error {
		return buildProjectMakefile(rootPath)
	})
	/*创建gitlab-ci.yml*/
	group.Go(func() error {
		return buildProjectGitlabCi(rootPath)
	})
	/*创建.gitignore*/
	group.Go(func() error {
		return file.BuildGitIgnore(rootPath)
	})
	/*创建依赖脚本*/
	group.Go(func() error {
		hackFile := file.NewHackFile(rootPath, projectConfig.Dependence)
		return hackFile.BuildFile()
	})
	err = group.Wait()
	if err != nil {
		log.Fatalln(err)
	}

}

func getProjectDirTree(root string) string {
	return fmt.Sprintf(`
name: %s
child:
- name: db
  child:
  - name: test
    child:
    - name: migration
- name: hack
- name: k8s
- name: src
  child:
  - name: main
    child:
    - name: java
    - name: resources
  - name: test
    child:
    - name: java
      child:
      - name: unit
      - name: integration
    - name: resources
`, root)

}

func buildProjectMakefile(path string) error {
	target, err := os.OpenFile(path+"/.gitlab-ci.yml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
# 只有推送release分支才触发CICD管道
workflow:
  rules:
    - if: '$CI_COMMIT_BRANCH == "release"'

# 声明管道流程
stages:
  - package
  - build
  - deploy

# maven打包阶段
maven-package:
  stage: package
  tags:
    - maven-host
  script:
    - make package
  artifacts:
    paths:
      - target/*.jar

# 镜像构建阶段
build-master:
  stage: build
  tags:
    - docker
  script:
    - make build

#部署到k8s阶段
deploy:
  image: harbor.zhisheng.com:5000/public/kubectl:v1.2
  stage: deploy
  tags:
    - kubectl
  script:
    - make deploy`)
	return err
}

func buildProjectGitlabCi(path string) error {
	target, err := os.OpenFile(path+"/Makefile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
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

# image repo
image_repo_url = harbor.zhisheng.com:5000/concise/concise-demo-backend

# ---------------------------------按照需要添加的命令---------------------------------

define run-mysql
	DB_USERNAME=${db_username} \
	DB_PASSWORD=${db_password} \
	DB_PORT=${db_port} \
	DB_CONTAINER_NAME=${db_container_name} \
	DB_DBS=demo \
	/bin/sh hack/mysql.sh
endef

define delete-mysql
	docker stop ${db_container_name} && docker container rm ${db_container_name}
endef

define run-redis
	CONTAINER_NAME=${redis_container_name} \
	PORT=${redis_port} \
	/bin/sh hack/redis.sh
endef

define delete-redis
	docker stop ${redis_container_name} && docker rm ${redis_container_name}
endef

define run-rabbitmq
	CONTAINER_NAME=${rabbitmq_container_name} \
	PORT=${rabbitmq_port} \
	USER=${rabbitmq_username} \
	PASSWORD=${rabbitmq_password} \
	/bin/sh hack/rabbitmq.sh
endef

define delete-rabbitmq
	docker stop ${rabbitmq_container_name} && docker rm ${rabbitmq_container_name}
endef

define run-minio
	CONTAINER_NAME=${minio_container_name} \
	/bin/sh hack/minio.sh
endef

define delete-minio
	docker stop ${minio_container_name} && docker rm ${minio_container_name}
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

# ---------------------------------按照需要添加的命令 end---------------------------------

# ---------------------------------拉起完整的测试容器进行测试的命令----------------------

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

# ---------------------------------拉起完整的测试容器进行测试的命令 end----------------------

# 帮助提示
.PHONY: help
help:
	echo "Nothing happen"

# 安装到本地的命令
.PHONY: package
package:
	mvn clean package -Dmaven.test.skip=true

# 构建命令
.PHONY: build
build:
	echo "构建镜像"
	docker build -t ${image_repo_url} .
	echo "推送镜像"
	docker push ${image_repo_url}

# 部署命令
.PHONY: deploy
deploy:
	echo "替换版本号"
	sed -i "s/VERSION_PLACEHOLDER/${CI_COMMIT_SHA}/g" k8s/dp.yaml
	echo "替换镜像"
	sed -i "s/IMAGE_PLACEHOLDER/${image_repo_url}/g" k8s/dp.yaml
	echo "执行部署"
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

# ----------------------------测试脚本 end-------------------------------`)
	return err
}
