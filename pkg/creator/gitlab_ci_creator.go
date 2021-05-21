package creator

import "os"

var gitlabCiTemplate = `
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
`

func CreateGitlabCi(basePath string) error {
	gitlabCi, err := os.Create(basePath + "/.gitlab-ci.yml")
	if err != nil {
		return err
	}
	_, err = gitlabCi.WriteString(gitlabCiTemplate)
	if err != nil {
		return err
	}
	return gitlabCi.Close()
}
