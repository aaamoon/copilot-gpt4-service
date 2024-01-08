<h1 align="center">copilot-gpt4-service</h1>

<p align="center">
⚔️ 将 Github Copilot 转换为 ChatGPT
</p>

<p align="center">
简体中文 | <a href="README_EN.md">English</a>
</p>

## 如何使用

1. 部署 copilot-gpt4-service 服务，并配置 API 地址，如：`https://youcopilotgpt4service.com`;
2. 获取你的 GitHub 账号 Github Copilot Plugin Token（详见下文）；
3. 使用第三方客户端，如：[ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web)，在设置中填入 copilot-gpt4-service 服务的 API 地址和 Github Copilot Plugin Token，即可使用 GPT-4 模型进行对话。

## 客户端

使用 copilot-gpt4-service，需要配合第三方客户端，目前已测试支持以下客户端：

- [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) (推荐)
- [Chatbox](https://github.com/Bin-Huang/chatbox)：支持 Windows, Mac, Linux 平台
- [OpenCat APP](https://opencat.app/)：支持 iOS、Mac 平台
- [ChatX APP](https://apps.apple.com/us/app/chatx-ai-chat-client/id6446304087) ：支持 iOS、Mac 平台

## 服务端

copilot-gpt4-service 服务的部署方式目前包含 Docker 部署、源码部署、Kubernetes 部署、Cloudflare Worker 实现、Zeabur 部署，下面分别介绍。

### 配置方式

使用环境变量或环境变量配置文件 `config.env` 配置服务（环境变量优先级高于 `config.env`），默认配置项如下：  

```env
HOST=localhost # 服务监听地址
PORT=8080 # 服务监听端口
CACHE=true # 是否启用持久化
CACHE_PATH=db/cache.sqlite3 # 持久化缓存的路径（仅当 CACHE=true 时有效）
```

### Docker 部署

#### 一键部署方式

```bash
docker run -d \
  --name copilot-gpt4-service \
  --restart always \
  -p 8080:8080 \
  aaamoon/copilot-gpt4-service:latest
```

#### 代码构建方式

```bash
git clone https://github.com/aaamoon/copilot-gpt4-service && cd copilot-gpt4-service
# 可在 docker-compose.yml 中修改端口  
docker compose up -d
```

如需更新容器，可在源代码文件夹重新拉取代码及构建镜像，命令如下：  

```bash
git pull && docker compose up -d --build
```

### Kubernetes 部署

支持通过 Kubernetes 部署，具体部署方式如下：

```shell
git clone https://github.com/aaamoon/copilot-gpt4-service.git
# git clone git@github.com:aaamoon/copilot-gpt4-service.git
cd copilot-gpt4-service/.chart
helm upgrade copilot-gpt4-service . --namespace copilot-gpt4-service --create-namespace --install  
```

### Cloudflare Worker

支持通过 Cloudflare Worker 部署，具体使用方式见 [cf-copilot-service](https://github.com/wpv-chan/cf-copilot-service)。

### Zeabur 部署

点击下方按钮一键部署：

[![Deploy on Zeabur](https://zeabur.com/button.svg)](https://zeabur.com/templates/73WZFG)

## 获取 Copilot Token

首先，你的账号需要开通 Github Copilot 服务

获取 Github Copilot Plugin Token 的方式目前有两种方式：

1. 通过安装 [Github Copilot CLI](https://githubnext.com/projects/copilot-cli/) 授权获取（推荐）。
2. 通过 [https://cocopilot.org](https://cocopilot.org/copilot/token) 第三方接口授权获取。

### 通过 Github Copilot CLI 授权获取

**Linux/MacOS平台获取**

```bash
# 如下脚本会自动安装 Github Copilot CLI 并通过授权获取 Github Copilot Plugin Token 
bash -c "$(curl -fsSL https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.sh)"
```

**Windows 平台获取**

下载批处理脚本，双击运行即可：[get_copilot_token.bat](https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.bat)。

### 第三方接口授权获取

通过 [https://cocopilot.org](https://cocopilot.org/copilot/token) 第三方接口授权获取，需要注意的是，该接口是第三方开发者提供的，不保证安全性，请谨慎使用。

## 常见问题

### 模型支持情况

据测试：模型参数支持 GPT-4 和 GPT-3.5-turbo ，实测使用其他模型均会以默认的 3.5 处理（对比 OpenAI API 的返回结果，猜测应该是最早的版本 GPT-4-0314 和 GPT-3.5-turbo-0301 ）

### 如何判断是不是 GPT-4 模型

鲁迅为什么暴打周树人？

- GPT-3.5 会一本正经的胡说八道
- GPT-4 表示鲁迅和周树人是同一个人

我爸妈结婚时为什么没有邀请我？

- GPT-3.5 他们当时认为你还太小，所以没有邀请你。
- GPT-4 他们结婚时你还没出生。

### HTTP 响应状态码解析说明

- 401: 使用的 Github Copilot Plugin Token 过期了或者错误，请重新获取
- 403: 使用的账号没有开通 Github Copilot

## 鸣谢

### 贡献者

<a href="https://github.com/aaamoon/copilot-gpt4-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aaamoon/copilot-gpt4-service&anon=0" />
</a>

## 开源协议

[MIT](https://opensource.org/license/mit/)
