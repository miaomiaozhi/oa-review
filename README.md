# oa-review

## 项目介绍
做一个简要的oa审核 demo

## 项目需求
1. 服务设计：可通过一个配置文件（json、yaml）实例化一个流程模型，支持流程运转或者回退，运转要求有多个审批人（审批人通过参数模拟），且同时支持多个流程，并开放初始化流程、运转、回退接口
2. 服务部署至虚拟机docker上 => 部署到k8s集群
3. 代码托管到公司gitlab
4. 接口通过Nginx暴露、运用mysql、redis其中一个中间件

## 项目地址
[oa-review](https://gitlab.moresec.cn/mozezhao/oa-review)

## 启动
```
go mod tidy && go run ./cmd/main.go --config=./conf/config.json
```