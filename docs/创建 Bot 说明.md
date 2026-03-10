创建使用 Bot 说明
===============

# 核心请求

# Bot 相关
## Bot 的配置（openclaw 配置）定义与 Json 表示
OpenClawConfig 定义最终渲染成 openclaw 的 openclaw.json 配置文件，请根据实际需要配置 创建 Bot 请求 config 部分（也就是如下定义配置）:

```go
type OpenClawConfig struct {
	Meta     *MetaConfig     `json:"meta,omitempty"`
	Models   *ModelsConfig   `json:"models,omitempty"`
	Agents   *AgentsConfig   `json:"agents,omitempty"`
	Channels ChannelsConfig  `json:"channels,omitempty"`
	Gateway  *GatewayConfig  `json:"gateway,omitempty"`
	Auth     *AuthConfig     `json:"auth,omitempty"`
	Plugins  *PluginsConfig  `json:"plugins,omitempty"`
	Messages *MessagesConfig `json:"messages,omitempty"`
	Commands *CommandsConfig `json:"commands,omitempty"`
	Hooks    *HooksConfig    `json:"hooks,omitempty"`
}
```

几乎完整版 Json 示例:
```json
{
  "meta": {
    "lastTouchedVersion": "2026.2.9",
    "lastTouchedAt": "2026-02-10T04:50:51.302Z"
  },
  "models": {
    "mode": "merge",
    "providers": {
      "tokenpony": {
        "baseUrl": "https://api.tokenpony.cn/v1",
        "apiKey": "Your TokenPony API KEY",
        "auth": "api-key", 
        "authHeader": false,
        "api": "openai-completions",
        "models": [
          {
            "id": "deepseek-v3.1-terminus",
            "name": "deepseek-v3.1-terminus",
            "reasoning": false,
            "input": [
              "text"
            ],
            "cost": {
              "input": 0,
              "output": 0,
              "cacheRead": 0,
              "cacheWrite": 0
            },
            "contextWindow": 128000,
            "maxTokens": 8192
          },
          {
            "id": "qwen3-vl-235b-a22b-instruct",
            "name": "qwen3-vl-235b-a22b-instruct",
            "reasoning": false,
            "input": [
              "text",
              "image"
            ],
            "cost": {
              "input": 0,
              "output": 0,
              "cacheRead": 0,
              "cacheWrite": 0
            },
            "contextWindow": 128000,
            "maxTokens": 8192
          }
        ]
      }
    }
  },
  "agents": {
    "defaults": {
      "model": {
        "primary": "tokenpony/deepseek-v3.1-terminus"
      },
      "models": {
        "tokenpony/deepseek-v3.1-terminus": {
          "alias": "qwen"
        },
        "tokenpony/qwen3-vl-235b-a22b-instruct": {}
      },
      "workspace": "/home/node/.openclaw/workspace",
      "compaction": {
        "mode": "safeguard"
      },
      "maxConcurrent": 4,
      "subagents": {
        "maxConcurrent": 8
      }
    }
  },
  "messages": {
    "ackReactionScope": "group-mentions"
  },
  "commands": {
    "native": "auto",
    "nativeSkills": "auto"
  },
  "gateway": {
    "port": 18889,
    "mode": "local",
    "bind": "loopback",
     "controlUi": {
       "enabled": true,
       "allowInsecureAuth": true
     },
    "auth": {
      "mode": "token",
      "token": "6a44596d342bebf38363591872e12c3b1427785356d631ab"
    },
    "tailscale": {
      "mode": "off",
      "resetOnExit": false
    },
    "nodes": {
      "denyCommands": [
        "camera.snap",
        "camera.clip",
        "screen.record",
        "calendar.add",
        "contacts.add",
        "reminders.add"
      ]
    }
  },
  "skills": {
    "install": {
      "nodeManager": "pnpm"
    }
  }
}
```

# 创建 Bot 步骤说明

对于每个用户，按如下步骤创建并使用 Bot

## 【Step 1】创建一个 App [可选]
每个 APP 对应会获取一个 API_TOKEN, 这个 API_TOKEN 用来管理 Bot 的生命周期。
如有人提供有效的 API_TOKEN，这一步骤可以忽略。

请求格式定义如下:
```shell
curl -s -X POST http://localhost:18080/bot/api/v1/admin/apps \
  -H "Authorization: Bearer my-admin-token" \
  -H "Content-Type: application/json" \
  -d '{"name": "my-app"}'
```

示例：
```shell
curl -s -X POST http://10.2.14.244:31080/bot/api/v1/admin/apps \
  -H "Authorization: Bearer shinemo@2026" \
  -H "Content-Type: application/json"  \
  -d '{"name": "fastclaw-app"}'
  
{"code":0,"message":"success","data":{"id":"1d18e470-7465-49b0-a84b-1014b40a7856","name":"fastclaw-app","api_token":"b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861","status":"active","created_at":"2026-03-06T02:24:58.260994355Z","updated_at":"2026-03-06T02:24:58.260994355Z"}}
```
其中:
* `http://10.2.14.244:31080` 是本 server 的访问地址（根据集群内部/外部访问 方式不同，IP 和端口可能不同），具体由管理员提供
  `"Authorization: Bearer shinemo@2026"` 的 `shinemo@2026` 是部署本 server 时指定的 adminToken，具体由管理员提供


保存 返回的  data.api_token 字段:
```shell
export API_TOKEN="b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"
```


## 【Step 2】创建一个 Bot
请求格式定义如下(其中 请求 body 的 config 部分最终渲染成 openclaw 的 openclaw.json 配置文件，请根据实际需要配置 config 部分):
```shell
curl -X POST http://localhost:18080/bot/api/v1/bots \
  -H "Authorization: Bearer $API_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-001",
    "name": "my-bot-001",
    "slug": "my-bot-001",
    "config": {
        "models": {
            "mode": "merge",
            "providers": {
                "tokenpony": {
                    "baseUrl": "https://api.tokenpony.cn/v1",
                    "apiKey": "sk-e29970*******************121f52",
                    "auth": "api-key",
                    "authHeader": false,
                    "api": "openai-completions",
                    "models": [
                        {
                            "id": "deepseek-v3-0324",
                            "name": "deepseek-v3-0324"
                        }
                    ]
                }
            }
        }
    }
}
```

示例：
```shell
curl -s -X POST http://10.2.14.244:31080/bot/api/v1/bots \
  -H "Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861" \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user-001",
    "name": "my-bot-001",
    "slug": "my-bot-001",
    "config": {
        "models": {
            "mode": "merge",
            "providers": {
                "tokenpony": {
                    "baseUrl": "https://api.tokenpony.cn/v1",
                    "apiKey": "sk-e29970*******************121f52",
                    "auth": "api-key",
                    "authHeader": false,
                    "api": "openai-completions",
                    "models": [
                        {
                            "id": "deepseek-v3-0324",
                            "name": "deepseek-v3-0324"
                        }
                    ]
                }
            }
        }
    }
}'

{"code":0,"message":"success","data":{"id":"79d95cd5-fb50-4f97-a9f7-ca9622ec9d45","app_id":"1d18e470-7465-49b0-a84b-1014b40a7856","user_id":"user-001","name":"my-bot-001","slug":"my-bot-001","access_token":"4ad6b29143f83ea431ee4f5bea827b33","status":"created","config":{"models":{"mode":"merge","providers":{"tokenpony":{"api":"openai-completions","apiKey":"sk-e29970**********121f52","auth":"api-key","authHeader":false,"baseUrl":"https://api.tokenpony.cn/v1","models":[{"id":"deepseek-v3-0324","name":"deepseek-v3-0324"}]}}}},"endpoint":"","created_at":"2026-03-06T02:35:45.153044068Z","updated_at":"2026-03-06T02:35:45.153044068Z","access_url":"https://my-bot-001.fastclaw.loc?token=4ad6b29143f83ea431ee4f5bea827b33"}}
```
其中:
* `http://10.2.14.244:31080` 是本 server 的访问地址（根据集群内部/外部访问 方式不同，IP 和端口可能不同），具体由管理员提供
* `"Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"` 的 `b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861` 是上一步创建 App 时获取的 API_TOKEN，具体由创建 APP 的用户提供
* 请求 body 的 config 部分最终渲染成 openclaw 的 openclaw.json 配置文件，请根据实际需要配置 config 部分

* 从创建 Bot 的接口返回中记录如下信息:
* Bot 的 ID: data.id
* Bot 的访问 Token: data.access_token
```shell
export BOT_ID="79d95cd5-fb50-4f97-a9f7-ca9622ec9d45"
export BOT_ACCESS_TOKEN="4ad6b29143f83ea431ee4f5bea827b33"
```


## 【Step 3】启动创建的 Bot
请求格式定义如下
```shell
curl -X POST http://localhost:18080/bot/api/v1/bots/$BOT_ID/start \
  -H "Authorization: Bearer $API_TOKEN"
```

示例:
```shell
curl -X POST http://10.2.14.244:31080/bot/api/v1/bots/79d95cd5-fb50-4f97-a9f7-ca9622ec9d45/start \
  -H "Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"

{"code":0,"message":"success","data":{"id":"79d95cd5-fb50-4f97-a9f7-ca9622ec9d45","app_id":"1d18e470-7465-49b0-a84b-1014b40a7856","user_id":"user-001","name":"my-bot-001","slug":"my-bot-001","access_token":"4ad6b29143f83ea431ee4f5bea827b33","status":"running","config":{"models":{"mode":"merge","providers":{"tokenpony":{"api":"openai-completions","auth":"api-key","apiKey":"sk-e29970****************121f52","models":[{"id":"deepseek-v3-0324","name":"deepseek-v3-0324"}],"baseUrl":"https://api.tokenpony.cn/v1","authHeader":false}}}},"endpoint":"oc-79d95cd5-svc.fastclaw.svc.cluster.local:18789","created_at":"2026-03-06T02:35:45.153044Z","updated_at":"2026-03-06T02:35:45.153044Z"}}
```
其中:
* `http://10.2.14.244:31080` 是本 server 的访问地址（根据集群内部/外部访问 方式不同，IP 和端口可能不同），具体由管理员提供
* `"Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"` 的 `b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861` 是上一步创建 App 时获取的 API_TOKEN，具体由创建 APP 的用户提供
* 请求 url 路径中的 `79d95cd5-fb50-4f97-a9f7-ca9622ec9d45` 是上一步创建生成的 BOT_ID


## 【Step 4】访问已启动运行的 Bot
启动 Bot 约 3 分钟之后，可以从浏览器中输入如下地址访问创建的 Bot。
其中:
* `http://10.2.14.244:31080` 是本 server 的访问地址（根据集群内部/外部访问 方式不同，IP 和端口可能不同），具体由管理员提供
* 请求 url 路径中的 `79d95cd5-fb50-4f97-a9f7-ca9622ec9d45` 是上一步创建生成的 id （BOT_ID）
* 请求 url query 参数 token 的值是创建 Bot 的调用中返回的 access_token （BOT_ACCESS_TOKEN）
```
http://10.2.14.244:31080/proxy/79d95cd5-fb50-4f97-a9f7-ca9622ec9d45/?token=4ad6b29143f83ea431ee4f5bea827b33
```


## 【Step 5】/重启/停止/删除 Bot
请求格式定义如下
```shell
# 重启 （Bot 必须是运行状态才支持重启）
curl -X POST http://localhost:18080/bot/api/v1/bots/$BOT_ID/restart -H "Authorization: Bearer $API_TOKEN"

# 停止
curl -X POST http://localhost:18080/bot/api/v1/bots/$BOT_ID/stop -H "Authorization: Bearer $API_TOKEN"

# 删除（删除 Bot 并不要求 Bot 停止状态）
curl -X DELETE http://localhost:18080/bot/api/v1/bots/$BOT_ID -H "Authorization: Bearer $API_TOKEN"
```

示例:
```shell
curl -X POST http://10.2.14.244:31080/bot/api/v1/bots/79d95cd5-fb50-4f97-a9f7-ca9622ec9d45/restart \
  -H "Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"

# 停止
curl -X POST http://10.2.14.244:31080/bot/api/v1/bots/79d95cd5-fb50-4f97-a9f7-ca9622ec9d45/stop \
  -H "Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"

# 删除（删除 Bot 并不要求 Bot 停止状态）
curl -X DELETE http://10.2.14.244:31080/bot/api/v1/bots/79d95cd5-fb50-4f97-a9f7-ca9622ec9d45 \
  -H "Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"
```
其中:
* `http://10.2.14.244:31080` 是本 server 的访问地址（根据集群内部/外部访问 方式不同，IP 和端口可能不同），具体由管理员提供
* 请求 url 路径中的 `79d95cd5-fb50-4f97-a9f7-ca9622ec9d45` 是上一步创建生成的 id （BOT_ID）
* `"Authorization: Bearer b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861"` 的 `b7985effef1b39ed5aadb1d24998d0ec5208a4960b572905b2ad457a1fd2b861` 是创建 App 时获取的 API_TOKEN，具体由创建 APP 的用户提供