package template

func NewMysqlSh() string {
	return `#!/bin/sh

echo "启动mysql:5.7服务"
docker run -id \
 --name ${DB_CONTAINER_NAME} \
 --platform linux/amd64 \
 -p ${DB_PORT}:${DB_PORT} \
 -e MYSQL_ROOT_PASSWORD=${DB_PASSWORD} \
 mysql:5.7

i=0
while true; do
    docker exec -i ${DB_CONTAINER_NAME} mysql -uroot -proot -e 'show databases;' >> /dev/null 2>&1
    if [ $? -eq 0 ]; then
        break
    fi
    sleep 1
    i=$(($i + 1))
    echo $i
    if [ $i -gt 20 ]; then
        echo "mysql服务启动失败"
        exit 1
    fi
done

if [ ${DB_DBS} != '' ]; then
    for var in ${DB_DBS}
    do
        echo ${var}
        echo "mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'"
        docker exec -i ${DB_CONTAINER_NAME} /bin/sh -c "mysql -u${DB_USERNAME} -p${DB_PASSWORD} -e 'create database ${var}'"
    done
fi
`
}
