```bash
docker compose build
docker compose up -d
docker build ./app -t k8s-api-prod:1.0.0

minikube start
minikube image load k8s-api-prod:1.0.0
kubectl apply -f api-deployment.yml
kubectl apply -f api-service.yml

minikube ip
minikube service api-service --url
minikube stop
```