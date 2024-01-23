## Docker initial setup
overall look of the project components:
![Screenshot from 2024-01-23 20-55-51](https://github.com/Morning1139Angel/web-backend-project/assets/127003598/fcae876d-f313-45f5-b3e0-7c33d23b33d4)
further detials can be found in this document : [web_hw1_1402 _final.pdf](https://github.com/Morning1139Angel/web-backend-project/files/14027850/web_hw1_1402._final.pdf)
create the network and build the images required
```bash
docker network create --driver bridge project-network
docker build -t gateway-server ./gateway
docker build -t auth-server ./auth/
```
##Redis
run the folowing command for running redis and the redis monitor ... the redis monitor will be on port 8001 of local host
the redis password is "SuperSecretSecureStrongPass"
```bash
docker run -d --rm --name redis -v `pwd`/config:/etc/redis/ redis:6.0-alpine redis-server /etc/redis/redis.conf
docker network connect project-network redis --alias redis
docker run -d --rm --name gateway_redis -v `pwd`/config:/etc/redis/ redis:6.0-alpine redis-server /etc/redis/redis.conf
docker network connect project-network gateway_redis --alias gateway_redis
docker run -p 8001:8001 -d --rm --network=project-network redislabs/redisinsight:latest
```
u can also use the redis CLI to view the expiration time... run this command on a new terminal first:
```bash
docker exec -it redis redis-cli -a "SuperSecretSecureStrongPass"
```
and then u can use ```TTL <key-name>``` inside the cli to see the expiration time

##AUTH server
for the auth server run the following commands:
```bash
docker run --network=project-network --network-alias=auth-server  -p 9000:9000 -d --name auth-server auth-server
```

## Golang gateway server
to start the gateway server run the following docker commands:
```bash
docker run -d --name gateway-server gateway-server
docker network connect project-network gateway-server --alias gateway-server
```
## Nginx inital setup 
start the Nginx container and connect it to the network created
```bash
docker run --name nginx --network=project-network -p 80:80 -v `pwd`/config/default.conf:/etc/nginx/conf.d/default.conf -d nginx
```

##GRPC
the folowing commands were used to generate the grpc boilerplate code 
```bash
protoc --go_out=./auth --go-grpc_out=./auth proto/auth.proto
protoc --go_out=./gateway --go-grpc_out=./gateway proto/auth.proto
```
and then use ```go tidy``` in both ./gateway and ./auth to get the dependencies
and for locust u can use : 
```bash
python -m grpc_tools.protoc -I./proto --python_out=./locust/ --grpc_python_out=./locust
 auth.proto
```
##Locust

you can use the following command to run locust locally to test gateway:
```bash
locust -f ./locust/locustfile.py 
```

##Swagger

you can access here to see the swagger UI : 

```bash
http://localhost/docs/index.html
```
