# cloud-element
To facilitate use for cloud-sealer

# cloud-sealer
sealer for cloud

# 如何在没有K8S的环境下调试sealer?
设置环境变量 FILTAB_DEBUG=true 则不会调度label，可以打散调试所有 TASK

# 如何在Linux 系统上编译filtab-sealer的docker镜像?

1. ubuntu 18.04安装lotus环境：
```
sudo apt update
sudo apt install mesa-opencl-icd ocl-icd-opencl-dev gcc git bzr jq pkg-config curl
```
2. ubuntu 18.04安装rust并且修改.cargo/config为国内源

```
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```
3. 下载filtab-sealer源代码并且下载好submodul: extern/filecoin-ffi

```
git submodule update --init --recursive
```

4. 进入extern/filecoin-ffi, 编译
```
make
```
5. docker build 镜像

注意: 以下 $filtab-sealer-image-name 变量需要自定义

```
$ cd cloud-sealer
$ docker build -t $filtab-sealer-image-name .

```
====================================================
#### k8s 1.18的包安装

```
$ go clean -i k8s.io/client-go...
$ go get k8s.io/client-go@kubernetes-1.18.0
```
===================================================

### sealer_test.go 测试 Test_Labeling_Job 所需要的环境变量

#### 1. 环境变量FILTAB_K8S_CONFIG_IN_CLUSTER=true会告诉miner(pod)从集群k8s service account获取k8s config, 否则使用本地的.kube/config

```
FILTAB_K8S_CONFIG_IN_CLUSTER=true
```

#### 2. K8s 的 cloud-sealer job需要在default的k8s ns
#### K8s 的 使用这个default空间的service_account的名字，[运维需要提前创建k8s sa](https://github.com/cloud/k8s-cloud-sealer/blob/master/cloud-cluster-config.yaml)

```
FILTAB_SERVICE_ACCOUNT = cloud-job-service-account
```

#### 3. 挂载到filtab-sealer job中的物理机目录(例如吧/tmp/demo挂载到filtab-sealer job的Pod中)

```
SECTOR_DATA_HOST_PATH = /tmp/demo
```

#### 4. 物理机的镜像名字，创建的 cloud-sealer job会使用物理机的这个docker 镜像
```
FILTAB_SEALER_IMAGE = registry.cn-shanghai.aliyuncs.com/cloud/filecoin-ubuntu:18.04
```

### 调试ReadPiece的labeling的环境变量

```
TMP_PATH=/Users/terrill;SECTOR_DIR=.lotusminer;SECTOR_MINER_ID=1000;SECTOR_NUMBER=2;TASK_SECTOR_TYPE=2KiB;TASK_TYPE=READ_PIECE;EVENTING=true;NATS_SERVER=http://localhost:4222;PARAMS=eyJPZmZzZXQiOiIwIiwiU2l6ZSI6IjI1NCIsIlJhbmRvbW5lc3MiOiJNdUZGNUtVb0JqQlltT0J4d28zZk0vK3doMW9mL1pWS2N4ZGx3d2E1NlZnPSIsIkNvbW1kIjoiZXlJdklqb2lZbUZuWVRabFlUUnpaV0Z4Ykc5amEyMHlkbkUxWm5KMmFYQTNlSEpzZDI5MGMzaHpOM1JwZW5wNmMzVnVaR1J3TW5BemEyZHhkM051WW5sNVp6UnBhU0o5IiwiTWluZXJJcCI6IjEyNy4wLjAuMSJ9;JOB_NODE_NAME=docker-desktop;RESERVE_GIB_FOR_SYSTEM_AND_LAST_UNSEALED_SECTOR=81;
```

#### sealer本地调试的环境变量
MINER_IP
