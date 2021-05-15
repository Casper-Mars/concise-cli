package creator

import "os"

var mysqlShellTemplate = "#!/bin/sh\n\necho \"启动mysql:5.7服务\"\ndocker run -id \\\n --name ${DB_CONTAINER_NAME} \\\n -p ${DB_PORT}:${DB_PORT} \\\n -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \\\n mysql:5.7\n\ni=0\nwhile true; do\n    docker exec -i ${DB_CONTAINER_NAME} mysql -uroot -proot -e 'show databases;'\n    if [ $? -eq 0 ]; then\n        break\n    fi\n    sleep 1\n    i=$(($i + 1))\n    echo $i\n    if [ $i -gt 20 ]; then\n        exit 1\n    fi\ndone\n\nif [ ${DB_DBS} != '' ]; then\n    array=${DB_DBS//,/ }\n    for var in ${array}\n    do\n        echo ${var}\n        echo \"mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'\"\n        docker exec -i ${DB_CONTAINER_NAME} /bin/sh -c \"mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'\"\n    done\nfi\n\n"

func CreateMysqlShell(basePath string) error {
	mysqlShell, err := os.Create(basePath + "/mysql.sh")
	if err != nil {
		return err
	}
	_, err = mysqlShell.WriteString(mysqlShellTemplate)
	if err != nil {
		return err
	}
	return mysqlShell.Close()
}
