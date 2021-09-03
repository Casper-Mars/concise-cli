package template

func NewDockerfile() string {
	return `# Should be edit by user
FROM raokii/minijdk:8u202
WORKDIR /
COPY /target/service.jar /
ENTRYPOINT ["java","-jar","-Xmx1G","-XX:+UseG1GC","-Dspring.profiles.active=test","/service.jar"]
`
}
