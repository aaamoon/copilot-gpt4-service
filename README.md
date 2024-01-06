<h1 align="center">copilot-gpt4-service</h1>

<p align="center">
⚔️ 将 Github Copilot 转换为 ChatGPT
</p>

<p align="center">
简体中文 | <a href="README_EN.md">English</a>
</p>

## 使用方法

1、访问 https://gpt4copilot.tech

2、在设置的接口地址填入本仓库项目部署出来的服务端 API 地址 `https://gpt4copilot.tech`（**强烈建议自行部署服务端，因为不清楚后续 GitHub 会不会检测到从该服务端 IP 发出太多不同 Token 的请求导致有风险存在**）

3、在 API Key 中填入 GitHub Copilot Plugin Token

**大佬们如果有开通 GitHub Copilot 的话，可以使用自己的 Token，通过 [copilot-token接口](https://cocopilot.org/copilot/token) 来获取，目前太多不同的IP请求了，我提供出去的 Token 半个钟就失效了，如果是内部几个人用的话，Token 有效期一般是好几个月**

![步骤1](/assets/step1.png)

4、自行切换模型，支持 GPT-4 模型 **（据测试：模型参数仅支持 GPT-4 和 GPT-3.5-turbo ，实测使用其他模型均会以默认的 3.5 处理（对比 OpenAI API 的返回结果，猜测应该是最早的版本 GPT-4-0314 和 GPT-3.5-turbo-0301 ））**

5、接下来我们就可以无限制使用 GPT-4 模型了~

## 异常 HTTP 响应状态码解析

- 401: 使用的 GitHub Copilot Plugin Token 过期了或者错误，请重新获取
- 403: 使用的账号没有开通 GitHub Copilot'

## Super Token

在 `docker-compose.yml` 中存在两个可自定义的环境变量字段 `SUPER_TOKEN` 和 `DEFAULT_COPILOT_TOKEN`。这两个字段设计目的是将 copilot-gpt4-service 服务以更安全的方式分享给朋友：当用户请求的 `API Key` 为 `SUPER_TOKEN` 时，copilot-gpt4-service 服务端将调用内置的 `DEFAULT_COPILOT_TOKEN` 处理该请求。而服务端维护者只需要将 `SUPER_TOKEN` 分享给对方，而不用分享 `GitHub Copilot Plugin Token`



当 `SUPER_TOKEN` 不存在，为空字串符或为默认值时，该功能不会启用

## 个人部署

### 客户端

客户端使用的是 [ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web)，里面有详细的部署教程

### 服务端

#### Docker 部署

##### 一键部署方式

```bash
docker run -d \
  --name copilot-gpt4-service \
  --restart always \
  --env SUPER_TOKEN=your_super_token \
  --env DEFAULT_COPILOT_TOKEN=your_default_copilot_token \
  -p 8080:8080 \
  aaamoon/copilot-gpt4-service:latest
```

##### 实时构建方式

```bash
git clone https://github.com/aaamoon/copilot-gpt4-service && cd copilot-gpt4-service
# 可以在`docker-compose.yml`中修改端口
# 如需启用 Super Token 可在`docker-compose.yml`中修改对应的变量值
docker compose up -d
```

如需更新容器，可在源代码文件夹重新拉取代码及构建镜像，命令如下：  

```bash
git pull
docker compose up -d --build
```

#### Cloudflare Worker 部署

不方便使用 Docker 部署的话，可以使用 [Cloudflare Worker](https://github.com/wpv-chan/cf-copilot-service) 版本部署

## 实现原理

<a href="principle.md">原理链接</a>

原理流程图：
![实现原理](/assets/principle.png)

## 如何判断是不是 GPT-4 模型

鲁迅为什么暴打周树人？

- GPT-3.5 会一本正经的胡说八道
- GPT-4 表示鲁迅和周树人是同一个人

我爸妈结婚时为什么没有邀请我？

- GPT-3.5 他们当时认为你还太小，所以没有邀请你。
- GPT-4 他们结婚时你还没出生。

## 鸣谢

### 贡献者

<a href="https://github.com/aaamoon/copilot-gpt4-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aaamoon/copilot-gpt4-service" />
</a>

## 开源协议

[MIT](https://opensource.org/license/mit/)
