``` sh
go run main.go

curl localhost:21829/hello
```

## k8s

Kubernetes (k8s) でPodを終了または削除する際、コンテナ内のアプリケーションにはシグナルが送信されます。このシグナルを適切に処理することで、graceful shutdownを実現することができます。

### KubernetesにおけるPodの終了プロセス:

1. **SIGTERM シグナル**: Kubernetesは、Podを終了する際にまずコンテナのメインプロセスにSIGTERMシグナルを送信します。これは、リソースの解放やクリーンアップなど、終了前の処理を行うためのシグナルです。

2. **Grace Period**: SIGTERMシグナルの送信後、Kubernetesはデフォルトで30秒のGrace Period（猶予期間）を持ちます。この期間中にコンテナが終了しない場合、次のステップに進みます。

3. **SIGKILL シグナル**: Grace Periodが経過した後、KubernetesはコンテナのメインプロセスにSIGKILLシグナルを送信します。これは、プロセスを強制的に終了するシグナルです。

### Graceful Shutdownのためのアプローチ:

1. **シグナルハンドラの実装**: アプリケーション内でSIGTERMシグナルをキャッチするハンドラを実装します。このハンドラ内で、オープンしているコネクションのクローズ、ファイルの保存、キャッシュのクリアなどのクリーンアップ処理を行います。

2. **Grace Periodの調整**: 必要に応じて、Kubernetesのマニフェストファイル内で`terminationGracePeriodSeconds`を設定して、Grace Periodを調整します。これにより、アプリケーションが必要なクリーンアップ処理を完了するのに十分な時間を確保することができます。

3. **ヘルスチェック**: ヘルスチェック（livenessProbeやreadinessProbe）を適切に設定することで、アプリケーションの健康状態をKubernetesに伝えることができます。これにより、アプリケーションが不健康な状態になった場合に、Kubernetesが適切なアクションを取ることができます。

これらのアプローチを組み合わせることで、EKS上のKubernetes環境でのgraceful shutdownを実現することができます。

### Links

- Graceful Node Shutdown
  - https://aws.amazon.com/jp/blogs/news/amazon-eks-1-21-released/
