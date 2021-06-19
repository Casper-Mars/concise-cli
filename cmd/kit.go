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

var kitProjectConfig = config.NewKitConfig()

// kitCmd represents the kit command
var kitCmd = &cobra.Command{
	Use:   "kit",
	Short: "生成基础库项目",
	Long:  ``,
	Run:   createKit,
}

func init() {
	rootCmd.AddCommand(kitCmd)
	kitCmd.Flags().StringVarP(&kitProjectConfig.ParentVersion, "parent", "p", "", "指定父项目的版本号")
	kitCmd.Flags().StringVarP(&kitProjectConfig.Name, "name", "n", "", "指定项目名称")
	kitCmd.Flags().StringArrayVarP(&kitProjectConfig.Dependence, "dependence", "d", []string{}, "指定项目需要的外部依赖")
}

func createKit(cmd *cobra.Command, args []string) {
	/*校验参数，工程项目名称和父工程版本号不能为空*/
	err := kitProjectConfig.Check()
	if err != nil {
		log.Println(err.Error())
		return
	}
	group, _ := errgroup.WithContext(context.Background())
	/*生成目录*/
	err = dir.Build([]byte(getDirTree(kitProjectConfig.Name)), ".")
	if err != nil {
		log.Fatalln(err)
	}
	rootPath := "./" + kitProjectConfig.Name
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
		return buildKitMakefile(rootPath)
	})
	/*创建gitlab-ci.yml*/
	group.Go(func() error {
		return buildKitGitlabCi(rootPath)
	})
	/*创建.gitignore*/
	group.Go(func() error {
		return file.BuildGitIgnore(rootPath)
	})
	/*创建依赖脚本*/
	group.Go(func() error {
		hackFile := file.NewHackFile(rootPath, dependence)
		return hackFile.BuildFile()
	})
	err = group.Wait()
	if err != nil {
		log.Fatalln(err)
	}
}

//getDirTree 获取目录树
func getDirTree(root string) string {
	return fmt.Sprintf(`
name: %s
child:
- name: db
  child:
  - name: test
    child:
    - name: migration
- name: hack
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

func buildKitGitlabCi(path string) error {
	target, err := os.OpenFile(path+"/.gitlab-ci.yml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
workflow:
  rules:
    - if: '$CI_COMMIT_BRANCH == "release"'

# test:进行测试的阶段
# deploy:测试阶段正常通过后，进入部署阶段，把构件部署到仓库中
# notify:部署完成后，进入通知阶段，把新部署的构件的信息推送给各个订阅者
stages:
  - test
  - deploy
  - notify

test:
  stage: test
  tags:
    - maven-host
  script:
    - make test-with-clean

deploy:
  stage: deploy
  tags:
    - maven-host
  script:
    make deploy

notify:
  stage: notify
  tags:
    - maven-host
  script:
    - VERSION=$(mvn help:evaluate -Dexpression=project.version -q -DforceStdout)
    - NAME=$(mvn help:evaluate -Dexpression=project.artifactId -q -DforceStdout)
    - curl --location --request POST 'http://192.168.123.210:9444/api/msg' --form "name=${NAME}" --form "version=${VERSION}"
`)
	return err
}

func buildKitMakefile(path string) error {
	target, err := os.OpenFile(path+"/Makefile", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`db_username = root
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
`)
	return err
}
