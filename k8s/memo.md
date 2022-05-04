# Kubernetes

## Ingress
Ingress Resource

### Ingress Controllers
- Nginx
- trafik
- Kong

## [minikube](https://minikube.sigs.k8s.io/docs/start/)
Minikubeは、Kubernetesをローカルで実行するツールです。MinikubeはシングルノードのKubernetesクラスターをパーソナルコンピューター上(Windows、macOS、Linux PCを含む)で実行することで、Kubernetesを試したり、日常的な開発作業のために利用できます。

``` sh
$ brew install minikube

minikube start --driver=docker
minikube addons enable ingress
kubectl get pods -n kube-system
kubectl create deployment web --image=gcr.io/google-samples/hello-app:1.0
kubectl expose deployment web --type=NodePort --port=8880
kubectl get service web
minikube service web --url

minikube delete --all --purge
```

### the terminal needs to be open to run it.
エラー内容

```
minikube service web --url


🏃  Starting tunnel for service web.
❗  Because you are using a Docker driver on darwin, the terminal needs to be open to run it.
```

kubectl  version

Client Version: version.Info{Major:"1", Minor:"22", GitVersion:"v1.22.5", GitCommit:"5c99e2ac2ff9a3c549d9ca665e7bc05a3e18f07e", GitTreeState:"clean", BuildDate:"2021-12-16T08:38:33Z", GoVersion:"go1.16.12", Compiler:"gc", Platform:"darwin/arm64"}
Server Version: version.Info{Major:"1", Minor:"23", GitVersion:"v1.23.3", GitCommit:"816c97ab8cff8a1c72eccca1026f7820e93e0d25", GitTreeState:"clean", BuildDate:"2022-01-25T21:19:12Z", GoVersion:"go1.17.6", Compiler:"gc", Platform:"linux/arm64"}


## Links
- [Ingress Controller](https://www.youtube.com/watch?v=MT7rnahKf90&ab_channel=Techtter)


curl \
  -H "Accept: application/vnd.github.v3+json" \
  http(s)://{hostname}/api/v3/repos/octocat/hello-world/stats/code_frequency

curl \
  -H "Authorization:token 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" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/api/v3/repos/p238049y/paractice/stats/code_frequency

curl \
  -H "Authorization:token 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" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/api/v3/repos/Kokoichi/til/stats/code_frequency

curl \
  -H "Authorization:token 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" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/users/kokoichi206/events

curl \
  -H "Authorization:token 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" \
  -H "Accept: application/vnd.github.v3+json" \
  https://api.github.com/users/p238049y/events


https://api.github.com/users/{username}/events

curl \
  -H "Accept: application/vnd.github.v3+json" \
  https://hostname/api/v3/repos/p238049y/paractice/stats/code_frequency




