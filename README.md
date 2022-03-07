# go-queue-service

### Local development environment on Kubernetes handles working services with queue.

Deploy Ingress controller:
```
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

Create secret:
``` 
openssl req -x509 -out .ops/certs/localhost.crt -keyout .ops/certs/localhost.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=go-queue-service.local' -extensions EXT -config <( \
   printf "[dn]\nCN=go-queue-service.local\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:go-queue-service.local\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

echo 127.0.0.1 go-queue-service.local >> /etc/hosts

kubectl create secret tls go-queue-service-secret-tls --key .ops/certs/localhost.key --cert .ops/certs/localhost.crt --namespace=go-ns
```

To start debug process in GoLand do next run commands:
1. Run `skaffold debug`
2. Debug `debug`
