deploy-ingress:
	helm repo update
	helm upgrade \
		--install ingress-nginx ingress-nginx \
		--repo https://kubernetes.github.io/ingress-nginx \
		--namespace ingress-nginx \
		--create-namespace \
		--set controller.hostPort.enabled=true \
		--wait \
		--cleanup-on-fail

metrics_server := https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
metrics_server_yaml := .ops/k8s-apps/metrics-server.yaml
download-metrics:
	curl -Lo $(metrics_server_yaml) $(metrics_server)
apply-metrics:
	kubectl apply -f $(metrics_server_yaml)
delete-metrics:
	kubectl delete -f $(metrics_server_yaml)

skaffoldVer := 1.39.4
install-skaffold:
	@if [ ${skaffoldVer} = "latest" ]; then\
        echo "trying to install skaffold latest version...";\
		sudo -S rm -f /usr/local/bin/skaffold;\
    	curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-darwin-arm64 && \
        sudo -S install skaffold /usr/local/bin/;\
        rm skaffold;\
    else\
        echo "trying to install skaffold ${skaffoldVer} version...";\
        sudo -S rm -f /usr/local/bin/skaffold;\
    	curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/v${skaffoldVer}/skaffold-darwin-arm64 && \
    	chmod +x skaffold && \
    	sudo -S mv skaffold /usr/local/bin;\
    fi;\
    skaffold version

create-secret:
	kubectl create secret tls go-queue-service-secret-tls \
		--key .ops/certs/localhost.key \
		--cert .ops/certs/localhost.crt \
		--namespace=gons

delete-secret:
	kubectl delete secret go-queue-service-secret-tls -n gons

delete-ingress:
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

