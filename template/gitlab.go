package template

func NewGitlabCI() string {
	return `# 只有推送release分支才触发CICD管道
# 只有推送release分支才触发CICD管道
workflow:
  rules:
    - if: '$CI_COMMIT_BRANCH == "release"'

# 声明管道流程
stages:
  - test
  - package
  - build
  - deploy

# maven测试阶段
maven-test:
  stage: test
  tags:
    - maven-host
  script:
    - make test-with-clean

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
  image: harbor.zhisheng.com:5000/public/kubectl:v1.3
  stage: deploy
  tags:
    - kubectl
  script:
    - make deploy
`
}
