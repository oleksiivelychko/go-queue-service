deploy-ingress:
	helm repo update
	helm upgrade --install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace

metrics_server := https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
metrics_server_yaml := .ops/k8s-apps/metrics-server.yaml
download-metrics:
	curl -Lo $(metrics_server_yaml) $(metrics_server)
apply-metrics:
	kubectl apply -f $(metrics_server_yaml)
remove-metrics:
	kubectl delete -f $(metrics_server_yaml)
deploy-metrics:
	kubectl apply -f $(metrics_server)
delete-metrics:
	kubectl delete -f $(metrics_server)

install-skaffold:
	sudo -S rm -f /usr/local/bin/skaffold
	curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-arm64 && \
    sudo -S install skaffold /usr/local/bin/
	rm skaffold
	skaffold version

kube-contexts:
	kubectl config get-contexts

kube-create-ns:
	kubectl create namespace go-ns

kube-create-secret:
	kubectl create secret tls go-queue-service-secret-tls --key .ops/certs/localhost.key --cert .ops/certs/localhost.crt --namespace=go-ns

kube-delete-ingress:
	kubectl delete ingress -n default go-queue-ingress

kube-delete-ns:
	kubectl delete ns go-ns

kube-delete-secret:
	kubectl delete secret go-queue-service-secret-tls

kube-deployments:
	kubectl get deployment

kube-get-ingress:
	kubectl get ingress --all-namespaces

kube-pods:
	kubectl get pods -o wide --all-namespaces

kube-purge-ingress:
	kubectl delete all --all -n ingress-nginx

kube-secrets:
	kubectl get secrets

kube-services:
	kubectl get services

kube-use-default:
	kubectl config use-context docker-desktop

skaffold-debug:
	$(info 'https://github.com/GoogleContainerTools/skaffold/issues/7582')
	skaffold debug --auto-build --auto-deploy --auto-sync --no-prune=false --cache-artifacts=false --verbosity info -f .ops/skaffold.yaml

skaffold-watch:
	skaffold dev --no-prune=false --cache-artifacts=false --verbosity error -f .ops/skaffold.yaml
