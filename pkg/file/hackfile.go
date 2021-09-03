package file

import "os"

type hackFile struct {
	dependence []string
	hackPath   string
}

func NewHackFile(path string, dependence []string) *hackFile {
	return &hackFile{
		dependence: dependence,
		hackPath:   path + "/hack",
	}
}

func (receiver hackFile) BuildFile() error {
	if len(receiver.dependence) == 0 {
		return nil
	}
	var err error
	for _, value := range receiver.dependence {
		switch value {
		case "mysql":
			err = receiver.buildMysqlHack()
		case "redis":
			err = receiver.buildRedisHack()
		case "rabbitmq":
			err = receiver.buildRabbitmqHack()
		case "minio":
			err = receiver.buildMinioHack()
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//buildMinioHack 构建minio脚本
func (receiver hackFile) buildMinioHack() error {
	target, err := os.OpenFile(receiver.hackPath+"/minio.sh", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
#!/bin/sh

echo "启动minio服务"

docker run -d \
 -p 9000:9000 \
 --name ${CONTAINER_NAME} \
  minio/minio server /data

`)
	return err
}

//buildRabbitmqHack 构建rabbitmq脚本
func (receiver hackFile) buildRabbitmqHack() error {
	target, err := os.OpenFile(receiver.hackPath+"/rabbitmq.sh", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
#!/bin/sh

echo "启动rabbitmq:3.8.17服务"

docker run -d \
  --name ${CONTAINER_NAME} \
  -p ${PORT}:5672 \
  -e RABBITMQ_DEFAULT_USER=${USER} \
  -e RABBITMQ_DEFAULT_PASS=${PASSWORD} \
  rabbitmq:3.8.17-rc.1-alpine
i=0
while true; do
    docker exec -i ${CONTAINER_NAME} rabbitmqctl ping >> /dev/null
    if [ $? -eq 0 ]; then
        break
    fi
    sleep 1
    i=$(($i + 1))
    echo $i
    if [ $i -gt 20 ]; then
        echo "rabbitmq服务启动失败"
        exit 1
    fi
done

echo "rabbitmq服务启动完成"`)
	return err
}

//buildRedisHack 构建redis脚本
func (receiver hackFile) buildRedisHack() error {
	target, err := os.OpenFile(receiver.hackPath+"/redis.sh", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
#!/bin/sh

echo "启动redis:5.0.8服务"

docker run -d \
  --name ${CONTAINER_NAME} \
  -p ${PORT}:6379 \
  redis:5.0.8-alpine3.11 \
  --requirepass "123456"
i=0
while true; do
    docker exec -i ${CONTAINER_NAME} redis-cli help >> /dev/null
    if [ $? -eq 0 ]; then
        break
    fi
    sleep 1
    i=$(($i + 1))
    echo $i
    if [ $i -gt 20 ]; then
        echo "redis服务启动失败"
        exit 1
    fi
done

echo "redis服务启动完成"`)
	return err
}

//buildMysqlHack 构建mysql脚本
func (receiver hackFile) buildMysqlHack() error {
	target, err := os.OpenFile(receiver.hackPath+"/mysql.sh", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer target.Close()
	_, err = target.WriteString(`
#!/bin/sh

echo "启动mysql:5.7服务"
docker run -id \
 --name ${DB_CONTAINER_NAME} \
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

`)
	return err
}
