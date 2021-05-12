package paper

import "os"

var makefileTemplate = "db_username = root\ndb_password = root\ndb_url = localhost\ndb_port = 3306\ndb_name = demo\ndb_container_name = test-mysql\n\n\ndefine run-mysql\n\tDB_USERNAME=${db_username} \\\n\tDB_PASSWORD=${db_password} \\\n\tDB_PORT=${db_port} \\\n\tDB_CONTAINER_NAME=${db_container_name} \\\n\tDB_DBS=demo \\\n\t./hack/mysql.sh\nendef\n\ndefine migrate-up-db\n\tmigrate -path db/test/migration \\\n\t-database 'mysql://$(db_username):$(db_password)@tcp($(db_url):$(db_port))/${db_name}' \\\n\t-verbose \\\n\tup $(1)\nendef\n\ndefine migrate-down-db\n\tmigrate -path db/test/migration \\\n\t-database 'mysql://$(db_username):$(db_password)@tcp($(db_url):$(db_port))/${db_name}' \\\n\t-verbose \\\n\tdown $(1)\nendef\n\ndefine delete-mysql\n\tdocker stop ${db_container_name} && docker container rm ${db_container_name}\nendef\n\ndefine test-container\n\tdocker run -i --rm --privileged \\\n\t --name test-env \\\n\t -e MYSQL_SERVICE=on \\\n\t -w /work \\\n\t -v `pwd`:/work \\\n\t harbor.zhisheng.com:5000/public/it-env \\\n\t /bin/sh -c $(1)\nendef\n\n\n# 安装到本地的命令\n.PHONY: install\ninstall:\n\tmvn -Dmaven.test.skip=true install\n\n# 部署命令\n.PHONY: deploy\ndeploy:\n\tmvn -Dmaven.test.skip=true deploy\n\n# ----------------------------本地开发常用脚本------------------------------\n\n# 建立本地开发环境\n.PHONY: setup-local-test-env\nsetup-local-test-env:\n\t$(call run-mysql)\n\t$(call migrate-up-db)\n\n# 清除本地开发环境\n.PHONY: delete-local-test-env\ndelete-local-test-env:\n\t$(call delete-mysql)\n\n# 初始化数据库\n.PHONY: migrate-up\nmigrate-up:\n\t$(call migrate-up-db,$(version))\n\n# 清理数据库\n.PHONY: migrate-down\nmigrate-down:\n\t$(call migrate-down-db,$(version))\n\n# 重置数据库\n.PHONY: resetDB\nresetDB:\n\tmake migrate-down\n\tmake migrate-up\n\n# ----------------------------本地开发常用脚本 end------------------------------\n\n# ----------------------------测试脚本-------------------------------\n\n# 使用本地环境进行测试\n.PHONY: test-local\ntest-local:\n\tmvn clean\n\techo \"配置本地测试环境\"\n\tmake setup-local-test-env\n\tmake test-ut\n\tmake test-it\n\n# 进行单元测试，默认本地已经有测试环境\n.PHONY: test-ut\ntest-ut:\n\techo \"进行单元测试\"\n\tmvn test -Dspring.profiles.active=test -Ptest_ut\n\n# 进行集成测试，默认本地已经有测试环境\n.PHONY: test-it\ntest-it:\n\techo \"进行集成测试\"\n\tmvn integration-test -Dspring.profiles.active=test -Ptest_it\n\n# 使用集成环境进行测试\n.PHONY: test\ntest:\n\t$(call test-container,\"make test-local\")\n\n# 使用集成环境进行测试并清理测试后的结果\n.PHONY: test-with-clean\ntest-with-clean:\n\t$(call test-container,\"make test-local && mvn clean\")\n\n# ----------------------------测试脚本 end-------------------------------\n\n"

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
