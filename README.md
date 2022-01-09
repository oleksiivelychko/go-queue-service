# go-queue-service

### Local development environment on Kubernetes handles working services with queue.

Install development toolkit:
```
brew install skaffold
brew install helm
brew install k9s
```

Deploy Ingress controller:
```
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

Create/delete secret:
```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout .ops/certs/openssl.key -out .ops/certs/openssl.crt -subj "/CN=go-queue-service.local/O=go-queue-service.local"
kubectl create secret tls go-queue-service-tls --key .ops/certs/openssl.key --cert .ops/certs/openssl.crt
kubectl delete secret go-queue-service-tls
```
