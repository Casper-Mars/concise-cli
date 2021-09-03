package template

func NewRedisSh() string {
	return `#!/bin/sh

echo "启动redis:5.0.8服务"

docker run -d \
  --name ${CONTAINER_NAME} \
  --platform linux/amd64 \
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

echo "redis服务启动完成"
`
}
