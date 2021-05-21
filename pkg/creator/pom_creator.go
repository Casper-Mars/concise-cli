package creator

import (
	"os"
	"strings"
)

var pomFileTemplate = `
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <groupId>com.zhisheng.framework.concise</groupId>
        <artifactId>parent</artifactId>
        <version>${parent_version}</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>${module_name}</artifactId>
    <version>0.1.0-SNAPSHOT</version>

    <properties>
    </properties>

    <dependencies>
        <!--单元测试-->
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-test-autoconfigure</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
            <exclusions>
                <exclusion>
                    <groupId>org.junit.vintage</groupId>
                    <artifactId>junit-vintage-engine</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <dependency>
            <groupId>org.jmockit</groupId>
            <artifactId>jmockit</artifactId>
            <scope>test</scope>
        </dependency>
        <dependency>
            <groupId>org.testng</groupId>
            <artifactId>testng</artifactId>
            <scope>test</scope>
        </dependency>
        <!--单元测试end-->
    </dependencies>

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
</project>
`

func CreatePom(basePath, moduleName, parentVersion string) error {
	pom := strings.ReplaceAll(pomFileTemplate, "${module_name}", moduleName)
	pom = strings.ReplaceAll(pom, "${parent_version}", parentVersion)
	newFile, err := os.Create(basePath + "/pom.xml")
	if err != nil {
		return err
	}
	_, err = newFile.WriteString(pom)
	if err != nil {
		return err
	}
	err = newFile.Close()
	return err
}
