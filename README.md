# web-hw1
## Nginx inital setup 
run these commands on the root project directory
```bash
docker network create --driver bridge Nginx
docker build -t gateway-server ./gateway
docker run -p 8080:8080 -d --name gateway-server gateway-server
docker run --name docker-nginx -p 80:80 -v `pwd`/default.conf:/etc/nginx/conf.d/default.conf -d nginx
```

## Golang gateway server
the golang server will log the incoming requests.
to see the logs of the server u can use the docker command :
```bash
docker logs -f gateway-server
```