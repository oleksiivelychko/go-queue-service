deploy-ingress:
	helm repo update
	helm upgrade \
		--install ingress-nginx ingress-nginx \
		--repo https://kubernetes.github.io/ingress-nginx \
		--namespace ingress-nginx \
		--create-namespace \
		--set controller.hostPort.enabled=true \
		--cleanup-on-fail \
		--verify \
		--wait

metrics_server := https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
metrics_server_yaml := .ops/k8s-apps/metrics-server.yaml
download-metrics:
	curl -Lo $(metrics_server_yaml) $(metrics_server)
apply-metrics:
	kubectl apply -f $(metrics_server_yaml)
delete-metrics:
	kubectl delete -f $(metrics_server_yaml)

install-skaffold:
	sudo -S rm -f /usr/local/bin/skaffold
	curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-arm64 && \
    sudo -S install skaffold /usr/local/bin/
	rm skaffold
	skaffold version

kube-delete-ns:
	kubectl delete ns go-ns

kube-init-ns:
	kubectl create namespace go-ns

kube-delete-secret:
	kubectl delete secret go-queue-service-secret-tls

kube-init-secret:
	kubectl create secret tls go-queue-service-secret-tls --key .ops/certs/localhost.key --cert .ops/certs/localhost.crt --namespace=go-ns

kube-delete-ingress:
	kubectl delete ingress -n default go-queue-ingress

kube-get-ingress:
	kubectl get ingress -n go-ns

kube-purge-ingress:
	kubectl delete all --all -n ingress-nginx

skaffold-debug:
	skaffold debug \
		--auto-build \
		--auto-deploy \
		--auto-sync \
		--verbosity info \
		--no-prune=false \
		--cache-artifacts=false \
		-f .ops/skaffold.yaml

skaffold-watch:
	skaffold dev --no-prune=false --cache-artifacts=false --verbosity error -f .ops/skaffold.yaml

