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
