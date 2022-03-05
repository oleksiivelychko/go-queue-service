helm-update:
	helm repo update

deploy-ingress: helm-update
	helm upgrade --install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace

purge-ingress:
	kubectl delete all --all -n ingress-nginx

delete-ingress:
	kubectl delete ingress -n default go-queue-ingress

deploy-metrics:
	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

get-ingress:
	kubectl get ingress --all-namespaces

build-and-push-docker:
	docker buildx build --platform linux/amd64 --tag alexvelychko/goqueueservice -f .ops/docker/Dockerfile .
	docker push alexvelychko/goqueueservice

init-ns:
	kubectl delete ns go-ns
	kubectl create namespace go-ns

kube-contexts:
	kubectl config get-contexts

kube-use-default:
	kubectl config use-context docker-desktop

kube-services:
	kubectl get services

kube-deployments:
	kubectl get deployment

kube-pods:
	kubectl get pods -o wide --all-namespaces

create-secret:
	kubectl create secret tls go-queue-service-secret-tls --key .ops/certs/localhost.key --cert .ops/certs/localhost.crt --namespace=go-ns

delete-secret:
	kubectl delete secret go-queue-service-secret-tls

kube-secrets:
	kubectl get secrets
