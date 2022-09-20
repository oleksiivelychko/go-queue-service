# go-queue-service

### Local development environment is running on Kubernetes and handles working service with queue.

📌 Generate certificate:
```
openssl req -x509 -days 365 -out .ops/certs/localhost.crt -keyout .ops/certs/localhost.key \
      -newkey rsa:2048 -nodes -sha256 \
      -subj '/CN=go-queue-service.local' -extensions EXT -config <( \
       printf "[dn]\nCN=go-queue-service.local\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:go-queue-service.local\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

echo 127.0.0.1 go-queue-service.local >> /etc/hosts
```

📌 To show all available logs in **k9s**: set `k9s.logger.sinceSeconds` to `-1` (use `k9s info` to find config location).


