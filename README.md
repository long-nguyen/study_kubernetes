### Run locally with docker
> docker build --tag myapp .
> docker run --env-file ./.env -p 8080:8080 myapp 

## Build to GCP Kubernetes

### Deploy code
> docker build -t asia.gcr.io/testkubenetes-370014/myapp:latest . 
Nếu là mac M1 thì
> docker build -t asia.gcr.io/testkubenetes-370014/myapp:latest . --platform linux/amd64 
Sau đó push
> docker push asia.gcr.io/testkubenetes-370014/myapp

Sau đó chạy build
> kubectl apply -f deployment/dev/deployment.yml


### When updating code
Ta chỉ cần update image và nó tự reload lại
Làm lại quá trình build trên
> kubectl rollout restart deployment myapp

### setup load balancer
Chỉ cần chạy 1 lần thôi
kubectl apply -f deployment/dev/service.yml
