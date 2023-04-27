# go-queue-service

### Local Kubernetes development environment is running on Skaffold and handles queue service.

📌 Use [generate-certs.sh](https://raw.githubusercontent.com/oleksiivelychko/laravel-starter-kit/main/.docker/shell/generate-certs.sh) to generate certificate:
```
CERTS_DIR=.ops/certs ./generate-certs.sh go-queue-service.local
```

📌 In order to run:

- delete PVC `kubectl delete pvc --all -n gons`
- make `deploy-ingress`
- make `create-secret`
- make `skaffold-dev`
- to obtain password `$(kubectl get secret --namespace gons rabbitmq -o jsonpath="{.data.rabbitmq-password}" | base64 -d)`

📌 In order to debug:
- make `skaffold-debug`
- run `debug` configuration in GoLand
- add breakpoints
- visit `go-queue-service.local/queue`

💡 To show all available logs in **k9s**: set `k9s.logger.sinceSeconds` to `-1` (use `k9s info` to find config location).

![Debugging an application using Skaffold and Delve](social_preview.png)
