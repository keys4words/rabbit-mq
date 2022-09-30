# run a standalone instance
docker network create rabbits
docker run -d --rm --net rabbits --hostname rab1 --name rab1 rabbitmq:3.11

# gui
localhost:8080
guest/guest

# inside container
rabbitmqctl
rabbitmq-plugins

# how to grab existing erlang cookie
docker exec -it rab1 cat /var/lib/rabbitmq/.erlang.cookie
TZSWZEJJEPKTVXFMXPYRroot@rab1

# plugins management
docker run -d --rm --net rabbits -p 8080:15672 --hostname rab1 --name rab1 rabbitmq:3.11
rabbitmq-plugins enable rabbitmq_management


docker run -d --rm --net rabbits --hostname rab2 --name rab2 rabbitmq:3.11
docker run -d --rm --net rabbits --hostname rab3 --name rab3 rabbitmq:3.11

# publisher
cd publisher
docker build . -t keys4words/rabbit-pub:1.0
docker run -it --rm --net rabbits -e RABBIT_HOST=rab1 -e RABBIT_PORT=5672 -e RABBIT_USERNAME=guest -e RABBIT_PASSWORD=guest -p 80:80 keys4words/rabbit-pub:1.0