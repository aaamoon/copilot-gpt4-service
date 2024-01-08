<h1 align="center">copilot-gpt4-service</h1>

<p align="center">
⚔️ Convert Github Copilot into ChatGPT
</p>

<p align="center">
English | <a href="README.md">简体中文</a>
</p>

## How to use

1. Deploy the copilot-gpt4-service and configure the API address, such as `https://youcopilotgpt4service.com`.
2. Obtain your GitHub account's Github Copilot Plugin Token (see below for details).
3. Use a third-party client, such as [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web), and enter the API address of the copilot-gpt4-service and the Github Copilot Plugin Token in the settings to use the GPT-4 model for conversation.

## Clients

To use copilot-gpt4-service, you need to use it with a third-party client. The following clients have been tested and are supported:

- [ChatGPT-Next-Web](https://github.com/ChatGPTNextWeb/ChatGPT-Next-Web) (recommended)
- [Chatbox](https://github.com/Bin-Huang/chatbox): Supports Windows, Mac, and Linux platforms
- [OpenCat APP](https://opencat.app/): Supports iOS and Mac platforms
- [ChatX APP](https://apps.apple.com/us/app/chatx-ai-chat-client/id6446304087): Supports iOS and Mac platforms

## Server

The deployment methods for copilot-gpt4-service currently include Docker deployment, source code deployment, Kubernetes deployment, and Cloudflare Worker implementation. They are described below.

### Configuration

Use environment variables or environment variable configuration file `config.env` to configure the service (environment variables take precedence over `config.env`), the default configuration items are as follows:  

```env
HOST=localhost # Service listening address
PORT=8080 # Service listening port
CACHE=true # Whether to enable persistence
CACHE_PATH=db/cache.sqlite3 # Path to persistent cache (only used when CACHE=true)
```

### Docker Deployment

#### One-click Deployment

```bash
docker run -d \
  --name copilot-gpt4-service \
  --restart always \
  -p 8080:8080 \
  aaamoon/copilot-gpt4-service:latest
```

#### Code Build

```bash
git clone https://github.com/aaamoon/copilot-gpt4-service && cd copilot-gpt4-service
# Modify the port in docker-compose.yml if necessary
docker compose up -d
```

To update the container, pull the code again and rebuild the image in the source code folder using the following command:

```bash
git pull && docker compose up -d --build
```

### Kubernetes Deployment

Supports deployment through Kubernetes, the specific deployment method is as follows:

```shell
helm repo add aaamoon https://charts.kii.la && helm repo update # Source by github pages
helm install copilot-gpt4-service aaamoon/copilot-gpt4-service


## Installation with Chat GPT Next Web
helm install copilot-gpt4-service aaamoon/copilot-gpt4-service \
  --set chatgpt-next-web.enabled=true \
  --set chatgpt-next-web.config.OPENAI_API_KEY=[ your openai api key ] \   #Token obtained by copilot
  --set chatgpt-next-web.config.CODE=[ backend access code ] \    # Access password for next chatgpt web ui
  --set chatgpt-next-web.service.type=NodePort \
  --set chatgpt-next-web.service.nodePort=30080
```

### Cloudflare Worker

Supports deployment through Cloudflare Worker, see [cf-copilot-service](https://github.com/wpv-chan/cf-copilot-service) for specific usage.

## Obtaining Copilot Token

Your account needs to have Github Copilot service enabled.

There are currently two ways to obtain the Github Copilot Plugin Token:

1. Obtain it by installing [Github Copilot CLI](https://githubnext.com/projects/copilot-cli/) and authorizing (recommended).
2. Obtain it through the third-party interface [https://cocopilot.org](https://cocopilot.org/copilot/token). Please note that this interface is provided by a third-party developer and its security cannot be guaranteed, so please use it with caution.

### Obtaining through Github Copilot CLI

**For Linux/MacOS Platforms**

```bash
# The script below will automatically install Github Copilot CLI and obtain the Github Copilot Plugin Token through authorization
bash -c "$(curl -fsSL https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.sh)"
```

**For Windows Platform**

Download the batch script and double-click to run it: [get_copilot_token.bat](https://raw.githubusercontent.com/aaamoon/copilot-gpt4-service/master/shells/get_copilot_token.bat).

### Obtaining through Third-Party Interface

Obtain it through the third-party interface [https://cocopilot.org](https://cocopilot.org/copilot/token). Please note that this interface is provided by a third-party developer and its security cannot be guaranteed, so please use it with caution.

## Frequently Asked Questions

### Model support

According to the test, the model parameters support GPT-4 and GPT-3.5-turbo, and the actual test will be processed at the default 3.5 when using other models (compared with the return results of the OpenAI API, guess it should be the earliest versions of GPT-4-0314 and GPT-3.5-turbo-0301)

## How to Determine if It's the GPT-4 Model

There are 9 birds in the tree, the hunter shoots one, how many birds are left in the tree？

- GPT-3.5 8 birds(Only able to answer eight.)
- GPT-4 None (other birds scared away, there may be no birds left in the trees.)

Why weren't I invited when my parents got married?

- GPT-3.5 They considered you too young at that time, so they didn't invite you.
- GPT-4 They got married before you were born.

### Explanation of HTTP Response Status Codes

- 401: The Github Copilot Plugin Token used has expired or is incorrect. Please obtain it again.
- 403: The account used does not have Github Copilot enabled.

## Acknowledgements

### Contributors

<a href="https://github.com/aaamoon/copilot-gpt4-service/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=aaamoon/copilot-gpt4-service&anon=0" />
</a>

## LICENSE

[MIT](https://opensource.org/license/mit/)
