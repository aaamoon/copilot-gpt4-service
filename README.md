# COPILOT-GPT4-SERVICE

简体中文 / [English](./README_EN.md)

## 使用方法
1、访问 http://gpt4copilot.tech

2、在设置的接口地址填入本仓库项目部署出来的api地址 http://gpt4copilot.tech

3、在API Key中填入 Github Copilot Token

提供三个已经开通了 Github Copilot 账号的 Token，可以直接使用：
- ~~**ghu_kEDPRczuQhVAxBxQD4Rkjv5uBba6zE3i0mNH**~~

**大佬们如果有开通 Github Copilot 的话，可以使用自己的 Token，通过 [copilot-token接口](https://cocopilot.org/copilot/token) 来获取，目前太多不同的IP请求了，我提供出去的 Token 半个钟就失效了，如果是内部几个人用的话，Token 有效期一般是好几个月**

![步骤1](/assets/step1.png)

4、支持各种 OpenAI 模型选择，目前默认使用的是 GPT-4 模型（据大佬测试：模型参数仅支持 GPT-4 和 GPT-3.5-turbo ，实测使用其他模型均会以默认的 3.5 处理（对比 OpenAI API 的返回结果，猜测应该是最早的版本 GPT-4-0314 和 GPT-3.5-turbo-0301 ））

5、接下来我们就可以无限制使用 GPT-4 模型了~

## 个人部署

### 客户端
客户端使用的是 [ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web)，里面有详细的部署教程

### 服务端

#### 使用Docker部署
```bash
git clone https://github.com/aaamoon/copilot-gpt4-service
cd copilot-gpt4-service
docker compose up -d
```
P.S 可以在`docker-compose.yml`中修改端口  

更新：  
```bash
git pull
docker compose up -d --build
```

#### Cloudflare Worker 部署
不方便使用Docker部署的话，可以使用 [Cloudflare Worker](https://github.com/wpv-chan/cf-copilot-service) 版本部署

## 实现原理

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
