# go-queue-service

### Local development environment on Kubernetes handles working services with queue.

📌 Install development tools:
```
brew install skaffold
brew install helm
brew install k9s
```

📌 Generate certificate:
```
openssl req -x509 -out .ops/certs/localhost.crt -keyout .ops/certs/localhost.key \
      -newkey rsa:2048 -nodes -sha256 \
      -subj '/CN=go-queue-service.local' -extensions EXT -config <( \
       printf "[dn]\nCN=go-queue-service.local\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:go-queue-service.local\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

echo 127.0.0.1 go-queue-service.local >> /etc/hosts
```

📌 To start debug process in GoLand do next run commands:
1. Run `skaffold debug`
2. Debug `debug`
