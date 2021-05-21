package creator

import "os"

var mysqlShellTemplate = `
#!/bin/sh

echo "启动mysql:5.7服务"
docker run -id \
 --name ${DB_CONTAINER_NAME} \
 -p ${DB_PORT}:${DB_PORT} \
 -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \
 mysql:5.7

i=0
while true; do
    docker exec -i ${DB_CONTAINER_NAME} mysql -uroot -proot -e 'show databases;'
    if [ $? -eq 0 ]; then
        break
    fi
    sleep 1
    i=$(($i + 1))
    echo $i
    if [ $i -gt 20 ]; then
        exit 1
    fi
done

if [ ${DB_DBS} != '' ]; then
    array=${DB_DBS//,/ }
    for var in ${array}
    do
        echo ${var}
        echo "mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'"
        docker exec -i ${DB_CONTAINER_NAME} /bin/sh -c "mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'"
    done
fi
`

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
