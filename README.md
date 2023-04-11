# grpc_demo
grpc demo deploy to aws

## 安装golang环境

## 安装protoc

```shell
brew install automake
brew install libtool
brew install protobuf
```

## Protoc-gen-go的安装

```shell
go install github.com/golang/protobuf/protoc-gen-go@latest
```

## 部署到aws

### 1. github上安装流水线

### 2. ecr 和 ecs配置
- ECR: Amazon EC2 Container Registry (ECR) 是完全托管的 docker容器注册表
- ECS: Amazon EC2 Container Service（ECS）是一个高度可扩展的软件容器管理服务
