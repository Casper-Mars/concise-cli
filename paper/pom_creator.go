package paper

import (
	"os"
	"strings"
)

var pomFileTemplate = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<project xmlns=\"http://maven.apache.org/POM/4.0.0\"\n         xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n         xsi:schemaLocation=\"http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd\">\n    <parent>\n        <groupId>com.zhisheng.framework.concise</groupId>\n        <artifactId>parent</artifactId>\n        <version>${parent_version}</version>\n    </parent>\n    <modelVersion>4.0.0</modelVersion>\n\n    <artifactId>${module_name}</artifactId>\n    <version>0.1.0-SNAPSHOT</version>\n\n    <properties>\n    </properties>\n\n    <dependencies>\n        <!--单元测试-->\n        <dependency>\n            <groupId>org.springframework.boot</groupId>\n            <artifactId>spring-boot-test-autoconfigure</artifactId>\n            <scope>test</scope>\n        </dependency>\n        <dependency>\n            <groupId>org.springframework.boot</groupId>\n            <artifactId>spring-boot-starter-test</artifactId>\n            <scope>test</scope>\n            <exclusions>\n                <exclusion>\n                    <groupId>org.junit.vintage</groupId>\n                    <artifactId>junit-vintage-engine</artifactId>\n                </exclusion>\n            </exclusions>\n        </dependency>\n        <dependency>\n            <groupId>org.jmockit</groupId>\n            <artifactId>jmockit</artifactId>\n            <scope>test</scope>\n        </dependency>\n        <dependency>\n            <groupId>org.testng</groupId>\n            <artifactId>testng</artifactId>\n            <scope>test</scope>\n        </dependency>\n        <!--单元测试end-->\n    </dependencies>\n\n    <repositories>\n        <repository>\n            <id>zhisheng-group</id>\n            <name>zhisheng Mirror</name>\n            <url>http://nexus.zhisheng.com:8081/repository/zhisheng/</url>\n            <releases>\n                <enabled>true</enabled>\n            </releases>\n            <snapshots>\n                <enabled>true</enabled>\n            </snapshots>\n        </repository>\n    </repositories>\n\n\n</project>"

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
