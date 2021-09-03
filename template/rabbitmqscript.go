package template

func NewRabbitmqSh() string {
	return `#!/bin/sh

echo "启动rabbitmq:3.8.17服务"

docker run -d \
  --name ${CONTAINER_NAME} \
  --platform linux/amd64 \
  -p ${PORT}:5672 \
  -p 15672:15672 \
  -e RABBITMQ_DEFAULT_USER=${USER} \
  -e RABBITMQ_DEFAULT_PASS=${PASSWORD} \
  harbor.zhisheng.com:5000/public/rabbitmq-management:3.8.17
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

echo "rabbitmq服务启动完成"
`
}
