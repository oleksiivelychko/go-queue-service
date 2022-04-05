helm-update:
	helm repo update

deploy-ingress: helm-update
	helm upgrade --install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace

delete-ingress:
	kubectl delete ingress -n default go-queue-ingress

get-ingress:
	kubectl get ingress --all-namespaces

build-and-push-docker:
	docker buildx build --platform linux/amd64 --tag alexvelychko/goqueueservice -f .ops/docker/Dockerfile . && docker push alexvelychko/goqueueservice

re-create-ns:
	kubectl delete ns go-ns && kubectl create namespace go-ns

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

generate-cert:
	openssl req -x509 -out .ops/certs/localhost.crt -keyout .ops/certs/localhost.key \
      -newkey rsa:2048 -nodes -sha256 \
      -subj '/CN=go-queue-service.local' -extensions EXT -config <( \
       printf "[dn]\nCN=go-queue-service.local\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:go-queue-service.local\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

	$(warning echo 127.0.0.1 go-queue-service.local >> /etc/hosts)
