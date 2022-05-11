deploy-metrics:
	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

download-metrics:
	curl -Lo metrics-server.yaml https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

apply-metrics:
	kubectl apply -f metrics-server.yaml

delete-metrics:
	kubectl delete -f metrics-server.yaml
