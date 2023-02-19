## Kubernetes 基础概念及入门操作

> Kubernetes 是一个可移植、可扩展的开源平台，用于管理容器化的工作负载和服务，可促进声明式配置和自动化。 Kubernetes 拥有一个庞大且快速增长的生态，其服务、支持和工具的使用范围相当广泛。

> 为什么需要Kubernetes？
>
> 容器是打包和运行应用程序的好方式。在生产环境中， 你需要管理运行着应用程序的容器，并确保服务不会下线。
>
>  Kubernetes 为你提供了一个可弹性运行分布式系统的框架。 Kubernetes 会满足你的扩展要求、故障转移你的应用、提供部署模式等。 例如，Kubernetes 可以轻松管理系统的 Canary (金丝雀) 部署。
>
> https://kuboard.cn/learning/k8s-bg/what-is-k8s.html
>
> https://kubernetes.io/zh-cn/docs/home/****

##### 目前持有的疑问：

###### 1.部署成功节点为什么不能直接访问（forward方式可用于调试，另一种方式通过k8s内部署nginx方式访问）

### 1. 重要概念

 #### 1.1 基础功能介绍

- **服务发现和负载均衡**

  Kubernetes 可以使用 DNS 名称或自己的 IP 地址来暴露容器。 如果进入容器的流量很大， Kubernetes 可以负载均衡并分配网络流量，从而使部署稳定。

- **存储编排**

  Kubernetes 允许你自动挂载你选择的存储系统，例如本地存储、公共云提供商等。

- **自动部署和回滚**

  你可以使用 Kubernetes 描述已部署容器的所需状态， 它可以以受控的速率将实际状态更改为期望状态。 例如，你可以自动化 Kubernetes 来为你的部署创建新容器， 删除现有容器并将它们的所有资源用于新容器。

- **自动完成装箱计算**

  你为 Kubernetes 提供许多节点组成的集群，在这个集群上运行容器化的任务。 你告诉 Kubernetes 每个容器需要多少 CPU 和内存 (RAM)。 Kubernetes 可以将这些容器按实际情况调度到你的节点上，以最佳方式利用你的资源。

- **自我修复**

  Kubernetes 将重新启动失败的容器、替换容器、杀死不响应用户定义的运行状况检查的容器， 并且在准备好服务之前不将其通告给客户端。

- **密钥与配置管理**

  Kubernetes 允许你存储和管理敏感信息，例如密码、OAuth 令牌和 ssh 密钥。 你可以在不重建容器镜像的情况下部署和更新密钥和应用程序配置，也无需在堆栈配置中暴露密钥

#### 1.2 重要组件

![Components of Kubernetes](https://d33wubrfki0l68.cloudfront.net/2475489eaf20163ec0f54ddc1d92aa8d4c87c96b/e7c81/images/docs/components-of-kubernetes.svg)

##### Master组件是集群的控制平台（control plane）

- master 组件负责集群中的全局决策（例如，调度）
- master 组件探测并响应集群事件（例如，当 Deployment 的实际 Pod 副本数未达到 `replicas` 字段的规定时，启动一个新的 Pod）

Master组件可以运行于集群中的任何机器上。但是，为了简洁性，通常在同一台机器上运行所有的 master 组件，且不在此机器上运行用户的容器

**kube-apiserver**：

此 master 组件提供 Kubernetes API。这是Kubernetes控制平台的前端（front-end），可以水平扩展（通过部署更多的实例以达到性能要求）。kubectl / kubernetes dashboard / kuboard 等Kubernetes管理工具就是通过 kubernetes API 实现对 Kubernetes 集群的管理

**etcd**：

支持一致性和高可用的名值对存储组件，Kubernetes集群的所有配置信息都存储在 etcd 中。请确保您 [备份 (opens new window)](https://kubernetes.io/docs/tasks/administer-cluster/configure-upgrade-etcd/#backing-up-an-etcd-cluster)了 etcd 的数据。关于 etcd 的更多信息，可参考 [etcd 官方文档(opens new window)](https://etcd.io/docs/)

**kube-scheduler**：

此 master 组件监控所有新创建尚未分配到节点上的 Pod，并且自动选择为 Pod 选择一个合适的节点去运行。

影响调度的因素有：

- 单个或多个 Pod 的资源需求
- 硬件、软件、策略的限制
- 亲和与反亲和（affinity and anti-affinity）的约定
- 数据本地化要求
- 工作负载间的相互作用

**kube-controller-manager**：

此 master 组件运行了所有的控制器

逻辑上来说，每一个控制器是一个独立的进程，但是为了降低复杂度，这些控制器都被合并运行在一个进程里。

kube-controller-manager 中包含的控制器有：

- 节点控制器： 负责监听节点停机的事件并作出对应响应
- 副本控制器： 负责为集群中每一个 副本控制器对象（Replication Controller Object）维护期望的 Pod 副本数
- 端点（Endpoints）控制器：负责为端点对象（Endpoints Object，连接 Service 和 Pod）赋值
- Service Account & Token控制器： 负责为新的名称空间创建 default Service Account 以及 API Access Token

**cloud-controller-manager**：

cloud-controller-manager 中运行了与具体云基础设施供应商互动的控制器。这是 Kubernetes 1.6 版本中引入的特性，尚处在 alpha 阶段。

cloud-controller-manager 只运行特定于云基础设施供应商的控制器。如果您参考 www.kuboard.cn 上提供的文档安装 Kubernetes 集群，默认不安装 cloud-controller-manager。

cloud-controller-manager 使得云供应商的代码和 Kubernetes 的代码可以各自独立的演化。在此之前的版本中，Kubernetes的核心代码是依赖于云供应商的代码的。在后续的版本中，特定于云供应商的代码将由云供应商自行维护，并在运行Kubernetes时链接到 cloud-controller-manager。

以下控制器中包含与云供应商相关的依赖：

- 节点控制器：当某一个节点停止响应时，调用云供应商的接口，以检查该节点的虚拟机是否已经被云供应商删除

  > 译者注：私有化部署Kubernetes时，我们不知道节点的操作系统是否删除，所以在移除节点后，要自行通过 `kubectl delete node` 将节点对象从 Kubernetes 中删除

- 路由控制器：在云供应商的基础设施中设定网络路由

  > 译者注：私有化部署Kubernetes时，需要自行规划Kubernetes的拓扑结构，并做好路由配置，例如 [离线安装高可用的Kubernetes集群](https://kuboard.cn/install/install-k8s.html) 中所作的

- 服务（Service）控制器：创建、更新、删除云供应商提供的负载均衡器

  > 译者注：私有化部署Kubernetes时，不支持 LoadBalancer 类型的 Service，如需要此特性，需要创建 NodePort 类型的 Service，并自行配置负载均衡器

- 数据卷（Volume）控制器：创建、绑定、挂载数据卷，并协调云供应商编排数据卷

  > 译者注：私有化部署Kubernetes时，需要自行创建和管理存储资源，并通过Kubernetes的[存储类](https://kuboard.cn/learning/k8s-intermediate/persistent/storage-class.html)、[存储卷](https://kuboard.cn/learning/k8s-intermediate/persistent/pv.html)、[数据卷](https://kuboard.cn/learning/k8s-intermediate/persistent/volume.html)等与之关联

> 译者注：通过 cloud-controller-manager，Kubernetes可以更好地与云供应商结合，例如，在阿里云的 Kubernetes 服务里，您可以在云控制台界面上轻松点击鼠标，即可完成 Kubernetes 集群的创建和管理。在私有化部署环境时，您必须自行处理更多的内容。幸运的是，通过合适的教程指引，这些任务的达成并不困难。

##### Node 组件

Node 组件运行在每一个节点上（包括 master 节点和 worker 节点），负责维护运行中的 Pod 并提供 Kubernetes 运行时环境。

**kubelet**：

此组件是运行在每一个集群节点上的代理程序。它确保 Pod 中的容器处于运行状态。Kubelet 通过多种途径获得 PodSpec 定义，并确保 PodSpec 定义中所描述的容器处于运行和健康的状态。Kubelet不管理不是通过 Kubernetes 创建的容器。

**kube-proxy**：

[kube-proxy](https://kuboard.cn/learning/k8s-intermediate/service/service-details.html#虚拟-ip-和服务代理) 是一个网络代理程序，运行在集群中的每一个节点上，是实现 Kubernetes Service 概念的重要部分。

kube-proxy 在节点上维护网络规则。这些网络规则使得您可以在集群内、集群外正确地与 Pod 进行网络通信。如果操作系统中存在 packet filtering layer，kube-proxy 将使用这一特性（[iptables代理模式](https://kuboard.cn/learning/k8s-intermediate/service/service-details.html#iptables-代理模式)），否则，kube-proxy将自行转发网络请求（[User space代理模式](https://kuboard.cn/learning/k8s-intermediate/service/service-details.html#user-space-代理模式)）

**容器引擎**：

容器引擎负责运行容器。Kubernetes支持多种容器引擎：[Docker (opens new window)](http://www.docker.com/)、[containerd (opens new window)](https://containerd.io/)、[cri-o (opens new window)](https://cri-o.io/)、[rktlet (opens new window)](https://github.com/kubernetes-incubator/rktlet)以及任何实现了 [Kubernetes容器引擎接口 (opens new window)](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-node/container-runtime-interface.md)的容器引擎

#### 

### 2. 安装k8s

> 注意：搭建本地环境，与开发环境有一定区别

1. minikube 基础操作

   > `minikube`是一个工具， 能让你在本地运行 Kubernetes。 `minikube` 在你的个人计算机（包括 Windows、macOS 和 Linux PC）上运行一个一体化（all-in-one） 或多节点的本地 Kubernetes 集群，以便你来尝试 Kubernetes 或者开展每天的开发工作

1.1 安装minikube

```shell
brew install minikube # 适合mac 在此方式中安装
```

1.2 minikube 操作指令

```shell
minikube start # 启动集群
minikube dashboard # 支持浏览器页面可视化面板操作
minikube pause # 不影響部署的應用程序的情況下暫停 Kubernetes
minikube unpause # 取消暂停
minikube stop # 关闭集群
minikube config set memory 9001 # 更新默认内存限制 需要重启集群
minikube delete --all # 删除全部集群
minikube addons list # 游览可安装的集群插件列表
minikube start -p aged --kubernetes-version=v1.16.1 # 创建一个指定版本的集群
```

2. kubectl 命令行工具

   > **kubectl**`minikube start`被配置為在執行命令時訪問 minikube 內的 kubernetes 集群控制平面
   >
   > https://jimmysong.io/kubernetes-handbook/guide/kubectl-cheatsheet.html 操作指令

   2.1 安装kubectl命令行工具

   ```shell
   brew install kubectl # mac上建议此安装方式 比较简单
   ```

   2.2 kubectl 操作指令

   ```shell
   minikube kubectl -- get pods # 获取所有节点信息
   minikube kubectl -- create deployment hello-minikube --image=kicbase/echo-server:1.0 # 创建一个部署
   minikube kubectl -- expose deployment hello-minikube --type=NodePort --port=8080 # 使用 NodePort 服務公開部署
   kubectl port-forward <pod_name> <forward_port> --namespace <namespace> --address <IP默认：127.0.0.1> # 需要调试部署的pod、svc等资源是否提供正常访问时使用
   kubectl apply -f project.yaml # 项目中给k8s写的yaml配置文件
   kubectl scale -n default deployment hello-gin --replicas=1 # 伸缩k8s 节点配置使用
   kubectl rollout restart -n default deployment gin-deployment # 重启default分组下 部署的节点
   kubectl delete -n default deployment gin-deployment # 删除节点
   ```

   2.3 kubectl 部署服务yaml配置文件

   ```yaml
   apiVersion: apps/v1 # api 版本
   kind: Deployment # 资源对象类型
   metadata: # Deployment 元数据
     name: hello-gin # 对象名称
   spec: # 对象规约
     selector: # 选择器，作用：选择带有下列标签的Pod
       matchLabels: # 标签匹配
         app: hello-gin # 标签KeyValue
     template: # Pod 模版
       metadata: # Pod元数据
         labels: # Pod 标签
           app: hello-gin # Pod 标签，与上述的 Deployment.selector中的标签对应
       spec: # Pod 对象规约
         containers: # 容器
           - name: hello-gin # 容器名称
             image: xjsun/gogin:demo # 镜像名称:镜像版本
             resources: # 资源限制
               limits: # 简单理解为max资源值
                 memory: "128Mi"
                 cpu: "500m"
               requests: # 简单理解为min资源值
                 memory: "128Mi"
                 cpu: "500m"
             ports: # 端口
               - containerPort: 8088 # 端口号
   ---
   apiVersion: v1 # api 版本
   kind: Service # 对象类型
   metadata: # 元数据
     name: hello-gin-svc # 对象名称
   spec: # 规约
     selector: # 选择器
       app: hello-gin # 标签选择器，与 Pod 的标签对应
     ports:
       - port: 8088 # Service 端口号
         targetPort: 8088 # Pod 暴露的端口号
   ```

   

   

