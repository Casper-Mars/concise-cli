package pom

import (
	"fmt"
	"strings"
)

const (
	projectInfoPlaceHolder = "${project-info}"
)

type pom struct {
	template   string
	dependence string
}

func NewPom(groupId, artifactId, version, parentVersion string) *pom {
	projectInfo := fmt.Sprintf(`
    <groupId>%s</groupId>
    <artifactId>%s</artifactId>
    <version>%s</version>
`, groupId, artifactId, version)

	return &pom{
		template: strings.Replace(getPomTemplate(), projectInfoPlaceHolder, projectInfo, 1),
	}
}

//AppendDependence 添加maven依赖
func (receiver *pom) AppendDependence() {

}

//Build 构造pom文件内容
func (receiver *pom) Build() string {
	return receiver.template
}

//getPomTemplate 获取pom文件的模板
func getPomTemplate() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

	${project-info}

    <groupId>${group-id}</groupId>
    <artifactId>${artifact-id}</artifactId>
    <version>${version}</version>

    <properties>
        <maven.compiler.source>8</maven.compiler.source>
        <maven.compiler.target>8</maven.compiler.target>
        <spring-boot.version>2.2.9.RELEASE</spring-boot.version>
    </properties>


    <dependencies>
        <!--spring boot-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter</artifactId>
            <version>${spring-boot.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
            <version>${spring-boot.version}</version>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-actuator</artifactId>
        </dependency>
        <!--spring boot end-->
       
        <!--Prometheus-->
        <dependency>
            <groupId>io.micrometer</groupId>
            <artifactId>micrometer-registry-prometheus</artifactId>
        </dependency>
        <!--Prometheus end-->
    </dependencies>
    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-dependencies</artifactId>
                <version>${spring-boot.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>
    <repositories>
        <repository>
            <id>zhisheng-group</id>
            <name>zhisheng Mirror</name>
            <url>http://nexus.zhisheng.com:8081/repository/zhisheng/</url>
            <releases>
                <enabled>true</enabled>
            </releases>
            <snapshots>
                <enabled>true</enabled>
            </snapshots>
        </repository>
    </repositories>

    <build>
        <finalName>service</finalName>
        <plugins>
            <plugin>
                <groupId>org.springframework.boot</groupId>
                <artifactId>spring-boot-maven-plugin</artifactId>
                <version>${spring-boot.version}</version>
                <executions>
                    <execution>
                        <goals>
                            <goal>repackage</goal>
                        </goals>
                    </execution>
                </executions>
            </plugin>
        </plugins>
    </build>
</project>`
}
