# GCP チートシート
## GDP(DeploymentManagent)
### image-familyを確認するためのコマンド
~~~sh
gloud compute images list --project ${PROJECT名}
~~~
### deployment-managementの作成
~~~sh
gloud deployment-manager deployments create ${PROJECT名} --config ./deployment/gcp/deployment-management.yaml --preview
~~~
### deployment-managementの更新
~~~sh
gloud deployment-manager deployments update ${PROJECT名} --config ./deployment/gcp/deployment-management.yaml --preview
~~~
## GAP(ArtifactResitry)
### registryへのプッシュ方法
1. imageをビルド
~~~sh
docker build -f ./build/Dockerfile -t asia-northeast1-docker.pkg.dev/${PROJECT名}/map-talk/map-talk-image:latest .
~~~
2. プロジェクトへの認証
~~~sh
 gcloud auth configure-docker asia-northeast1-docker.pkg.dev
~~~
3. GARへプッシュ
~~~sh
docker push asia-northeast1-docker.pkg.dev/${PROJECT名}/map-talk/map-talk-image:latest   
~~~
## GCR(CouldRun)
### cloudRunへのデプロイ
~~~sh
gcloud run deploy map-talk-image --image asia-northeast1-docker.pkg.dev/${PROJECT名}/map-talk/map-talk-image:latest --platform managed --region asia-northeast1
~~~