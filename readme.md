# gui
docker run -d --rm --hostname rab --name rab -p 5672:5672 -p 5673:5673 -p 15672:15672 rabbitmq:3.8-management
localhost:15672
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

# publisher
cd publisher
docker build . -t keys4words/rabbit-pub:1.0

# cluster in docker-compose
docker network create rabbitmq-cluster
docker-compose up -d