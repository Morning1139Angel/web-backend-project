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
docker run --name nginx --network=project-network -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
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