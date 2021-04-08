package paper

import "os"

var gitlabCiTemplate = "workflow:\n  rules:\n    - if: '$CI_COMMIT_BRANCH == \"release\"'\n\n# test:进行测试的阶段\n# deploy:测试阶段正常通过后，进入部署阶段，把构件部署到仓库中\n# notify:部署完成后，进入通知阶段，把新部署的构件的信息推送给各个订阅者\nstages:\n  - test\n  - deploy\n  - notify\n\ntest:\n  stage: test\n  tags:\n    - maven-host\n  script:\n    - make test-with-clean\n\ndeploy:\n  stage: deploy\n  tags:\n    - maven-host\n  script:\n    make deploy\n\nnotify:\n  stage: notify\n  tags:\n    - maven-host\n  script:\n    - VERSION=$(mvn help:evaluate -Dexpression=project.version -q -DforceStdout)\n    - NAME=$(mvn help:evaluate -Dexpression=project.artifactId -q -DforceStdout)\n    - curl --location --request POST 'http://192.168.123.210:9444/api/msg' --form \"name=${NAME}\" --form \"version=${VERSION}\"\n"

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
