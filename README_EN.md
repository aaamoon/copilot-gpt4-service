# COPILOT-GPT4-SERVICE

English / [简体中文](./README.md)

## How to use
1. Visit http://gpt4copilot.tech

2. Fill in the interface address provided by this repository's deployed API: http://gpt4copilot.tech

3. Enter your Github Copilot token in the API Key field

Three pre-registered tokens for Github Copilot accounts are available for direct use:
- ~~**ghu_kEDPRczuQhVAxBxQD4Rkjv5uBba6zE3i0mNH**~~

**If you already have a Github Copilot account, you can use your own token by obtaining it through the [copilot-token API](https://cocopilot.org/copilot/token)，Currently, due to the high number of different IP requests, the tokens I provide become invalid within half an hour. If it's for internal use within a few people, the token is generally valid for several months.**

![step](/assets/step1_EN.png)

4. Various OpenAI models are supported, and the default is the GPT-4 model.(According to the expert's test, the model parameters only support GPT-4 and GPT-3.5-turbo. Other models tested were processed with the default 3.5 version (compared to the returned results from the OpenAI API, it is speculated that they are likely the earliest versions, GPT-4-0314 and GPT-3.5-turbo-0301)).

5. Now, we can make unlimited use of the GPT-4 model.

## Self-Deployment

### Client
The client uses [ChatGPT-Next-Web](https://github.com/Yidadaa/ChatGPT-Next-Web), where detailed deployment instructions are available

### Server

#### Docker Deployment
```bash
git clone https://github.com/aaamoon/copilot-gpt4-service
cd copilot-gpt4-service
docker compose up -d
```
P.S You can modify the port in `docker-compose.yml`  

update:  
```bash
git pull
docker compose up -d --build
```

#### Cloudflare Worker Deployment
If Docker deployment is inconvenient, you can use the [Cloudflare Worker](https://github.com/wpv-chan/cf-copilot-service) version for deployment.

## Implementation Principle

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
