# web-hw1
## Docker initial setup
create the network and build the images required
```bash
docker network create --driver bridge project-network
docker build -t gateway-server ./gateway
```
## Nginx inital setup 
run these commands on the root project directory
```bash
docker run --name nginx -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
docker network connect project-network docker-nginx --alias nginx
```

## Golang gateway server
the golang server will log the incoming requests.
to see the logs of the server u can use the docker command :
```bash
docker run -d --name gateway-server gateway-server
docker network connect project-network gateway-server --alias gateway-server
```