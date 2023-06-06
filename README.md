# web-hw1
run these commands on the root project directory
## Docker initial setup
create the network and build the images required
```bash
docker network create --driver bridge project-network
docker build -t gateway-server ./gateway
```
## Golang gateway server
the golang server will log the incoming requests.
to see the logs of the server u can use the docker command :
```bash
docker run -d --name gateway-server gateway-server
docker network connect project-network gateway-server --alias gateway-server
```
## Nginx inital setup 
start the Nginx container and connect it to the network created
```bash
docker run --name nginx --network=project-network -p 80:80 -v `pwd`/config/default.conf:/etc/nginx/conf.d/default.conf -d nginx
```
##Testing
To see the logs of the server u can use the following command 
```bash
docker logs -f gateway-server
```
And you can use these example requests to test the app: 
```bash
#GET
curl --location 'http://localhost:80'

#POST
curl --location 'http://localhost:80?sss=deeffer&sdwefsev=aaa&name=amir' \
--header 'Content-Type: text/plain' \
--data 'wdwef hellllllllo!!!'
```

##GRPC
the folowing commands were used to generate the grpc boilerplate code 
```bash
protoc --go_out=./auth --go-grpc_out=./auth proto/auth.proto
protoc --go_out=./gateway --go-grpc_out=./gateway proto/auth.proto
```
and then use ```go tidy``` in both ./gateway and ./auth to get the dependencies

##Redis
run the folowing command for running redis and the redis monitor ... the redis monitor will be on port 8001 of local host
the redis password is "SuperSecretSecureStrongPass"
```bash
docker run -d --rm --name redis -v `pwd`/config:/etc/redis/ redis:6.0-alpine redis-server /etc/redis/redis.conf
docker network connect project-network redis --alias redis
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
docker build -t auth-server ./auth/
docker run --network=project-network -p 9000:9000  -d --name auth-server auth-server
```