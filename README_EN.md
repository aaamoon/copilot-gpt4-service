<h1 align="center">copilot-gpt4-service</h1>

<p align="center">
⚔️ Convert Github Copilot to ChatGPT
</p>

<p align="center">
English | <a href="README.md">简体中文</a>
</p>

## How to use
1. Visit https://gpt4copilot.tech

2. Enter the server API address deployed in the repository project in the set interface address: `https://gpt4copilot.tech` (**It is strongly recommended to deploy the server yourself, because it is unclear whether GitHub will detect that there are too many requests for different tokens from this server IP and cause risk**)

3. Enter your GitHub Copilot Plugin Token in the API Key field

**If you already have a GitHub Copilot account, you can use your own token by obtaining it through the [copilot-token API](https://cocopilot.org/copilot/token)，Currently, due to the high number of different IP requests, the tokens I provide become invalid within half an hour. If it's for internal use within a few people, the token is generally valid for several months.**

![step](/assets/step1_EN.png)

4. Switch models on your own, support the GPT-4 model. **(Based on testing, the model parameters only support GPT-4 and GPT-3.5-turbo. Other models tested were processed with the default 3.5 version (compared to the returned results from the OpenAI API, it is speculated that they are likely the earliest versions, GPT-4-0314 and GPT-3.5-turbo-0301)).**

5. Now, we can make unlimited use of the GPT-4 model.

## Exception HTTP response status code parsing

- 401: The GitHub Copilot Plugin token used has expired or is incorrect, please obtain it again.
- 403: The account you are using does not have GitHub  Copilot activated.

## Super Token

In the `docker-compose.yml` file, there are two customizable environment variable fields: `SUPER_TOKEN` and `DEFAULT_COPILOT_TOKEN`. The purpose of these fields is to share the copilot-gpt4-service with friends in a more secure way: when the `API Key` in a user's request is `SUPER_TOKEN`, the copilot-gpt4-service server will handle the request using the built-in `DEFAULT_COPILOT_TOKEN`. In this way, the server maintainer only needs to share the `SUPER_TOKEN` with others, without having to share the `GitHub Copilot Plugin Token`.

If `SUPER_TOKEN` does not exist or is an empty string or is the default placeholder value, this feature will not be activated.

## Self-Deployment

### Client

The client uses [ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web), where detailed deployment instructions are available

### Server

#### Docker Deployment

##### One-click Deployment

```bash
docker run -d \
  --name copilot-gpt4-service \
  --restart always \
  --env SUPER_TOKEN=your_super_token \
  --env DEFAULT_COPILOT_TOKEN=your_default_copilot_token \
  -p 8080:8080 \
  aaamoon/copilot-gpt4-service:latest
```

##### Real-time Build

```bash
git clone https://github.com/aaamoon/copilot-gpt4-service && cd copilot-gpt4-service
# You can modify the port in `docker-compose.yml`  
# Modify environment variables in `docker-compose.yml` to enable Super Token
docker compose up -d
```

If you need to update the container, you can re-pull the code and build the image in the source code folder. The commands are as follows:

```bash
git pull
docker compose up -d --build
```

#### Cloudflare Worker Deployment

If Docker deployment is inconvenient, you can use the [Cloudflare Worker](https://github.com/wpv-chan/cf-copilot-service) version for deployment.

## Implementation Principle

<a href="principle.md">Principle Link</a>

Principle process image:
![Implementation Principle](/assets/principle.png)

## How to Determine if It's the GPT-4 Model

There are 9 birds in the tree, the hunter shoots one, how many birds are left in the tree？

- GPT-3.5 8 birds(Only able to answer eight.)
- GPT-4 None (other birds scared away, there may be no birds left in the trees.)

Why weren't I invited when my parents got married?

- GPT-3.5 They considered you too young at that time, so they didn't invite you.
- GPT-4 They got married before you were born.

## Special Thanks

### Contributor

<a href="https://github.com/aaamoon/copilot-gpt4-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aaamoon/copilot-gpt4-service" />
</a>


## LICENSE

[MIT](https://opensource.org/license/mit/)
