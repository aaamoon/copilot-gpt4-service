# README

![copilot gpt4 service](assets/copilot%20gpt4%20service.svg)

该`HLEM Chart`用于部署一个提供公开的或内部的API服务，作为Chat GPT Next的调用代理，将请求转换为GitHub Copilot API。

## 快速安装
使用`helm`命令安装`HLEM Chart`，命令如下：
```bash
git clone https://github.com/aaamoon/copilot-gpt4-service.git
# git clone git@github.com:aaamoon/copilot-gpt4-service.git
cd copilot-gpt4-service/.chart
helm upgrade copilot-gpt4-service . --namespace copilot-gpt4-service --create-namespace --install  
```

## Values 字段说明
下面是`HLEM Chart`中`Values`字段的详细解释：
下面是对`hlem chart`中Values字段的解释：
请根据实际需求调整`Values`字段的值，以满足您的部署需求。
以下是使用Markdown的表格格式输出对默认值进行解释的文档：
好的，下面是更详细的层级描述和默认值：

| 字段 | 默认值 | 描述                              |
| --- | --- |---------------------------------|
| `replicaCount` | 1 | 部署的副本数量                         |
| `image.repository` | aaamoon/copilot-gpt4-service | 容器镜像的仓库名                        |
| `image.pullPolicy` | Always | 容器镜像的拉取策略                       |
| `image.tag` | latest | 容器镜像的标签                         |
| `config.HOST` | 0.0.0.0 | 应用的主机配置                         |
| `config.PORT` | 8080 | 应用的端口配置                         |
| `persistent.cache.enabled` | false | 是否启用缓存                          |
| `persistent.cache.type` | pvc | 缓存的类型，可以是 pvc 或 hostPath        |
| `persistent.cache.name` | cache | 缓存的名称                           |
| `persistent.cache.mountPath` | /var/copilot-gpt4-service/cache.sqlite3 | 缓存的挂载路径                         |
| `persistent.cache.pvc.storageClassName` | "" | PVC 的存储类，如果为空则使用默认的存储类          |
| `persistent.cache.pvc.claimName` | copilot-gpt4-service-cache | PVC 的声明名称                       |
| `persistent.cache.pvc.accessModes` | [ReadWriteOnce] | PVC 的访问模式                       |
| `persistent.cache.pvc.size` | 1Gi | PVC 的大小                         |
| `persistent.cache.hostPath.path` | /var/copilot-gpt4-service/cache.sqlite3 | hostPath 的路径                    |
| `persistent.cache.hostPath.type` | DirectoryOrCreate | hostPath 的类型                    |
| `imagePullSecrets` | [ ] | 拉取私有镜像所需的密钥                     |
| `nameOverride` | "" | 用于覆盖默认的 Helm 发布名称               |
| `fullnameOverride` | "copilot-gpt4-service" | 用于覆盖默认的 Helm 完整发布名称             |
| `serviceAccount.create` | true | 是否创建服务账户                        |
| `serviceAccount.automount` | true | 是否自动挂载服务账户的 API 凭证              |
| `serviceAccount.annotations` | { } | 服务账户的注解                         |
| `serviceAccount.name` | "" | 服务账户的名称                         |
| `podAnnotations` | { } | Pod 的注解                         |
| `podLabels` | { } | Pod 的标签                         |
| `podSecurityContext` | { } | Pod 的安全上下文                      |
| `securityContext` | { } | 容器的安全上下文                        |
| `service.type` | ClusterIP | 服务的类型                           |
| `service.port` | 8080 | 服务的端口                           |
| `ingress.enabled` | false | 是否启用 Ingress                    |
| `ingress.className` | "nginx" | Ingress 的类名                     |
| `ingress.annotations` | { } | Ingress 的注解                     |
| `ingress.hosts` | [{host: example.com, paths: [{path: /, pathType: ImplementationSpecific}]}] | Ingress 的主机和路径配置                |
| `ingress.tls` | [ ] | Ingress 的 TLS 配置                |
| `resources` | { } | 资源限制和请求的配置                      |
| `autoscaling.enabled` | false | 是否启用自动扩缩                        |
| `autoscaling.minReplicas` | 1 | 自动扩缩的最小副本数                      |
| `autoscaling.maxReplicas` | 100 | 自动扩缩的最大副本数                      |
| `autoscaling.targetCPUUtilizationPercentage` | 80 | 自动扩缩的 CPU 利用率目标                 |
| `volumeMounts` | [ ] | 额外的卷挂载配置                        |
| `nodeSelector` | { } | 节点选择器的配置                        |
| `tolerations` | [ ] | 容忍度的配置                          |
| `affinity` | { } | 亲和性的配置                          |
| `chatgpt-next-web.enabled` | false | 是否启用下一代 ChatGPT web             |
| `chatgpt-next-web.config.BASE_URL` | http://copilot-gpt4-service:8080 | Next ChatGPT web 的基础 URL         |
| `chatgpt-next-web.config.OPENAI_API_KEY` | [ your openai api key ] | Next ChatGPT web 的 OpenAI API 密钥 |
| `chatgpt-next-web.config.CODE` | [ backend access code ] | Next ChatGPT web 的后端访问码         |
