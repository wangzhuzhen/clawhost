### 带token认证有效配置.md
但是访问有问题，页面没有到 token (b4329039e2ba4988a57711127f0a9b1b), 页面需要 
http://10.2.14.244:31080/proxy/92b953ef-f350-4070-9718-c7c5c132cebc/?token=b4329039e2ba4988a57711127f0a9b1b  访问

```
 {
        "gateway": {
          "port": 18789,
          "mode": "local",
          "bind": "lan",
          "auth": {
            "mode": "token",
            "token": "b4329039e2ba4988a57711127f0a9b1b"
          },
          "tailscale": {
            "mode": "off",
            "resetOnExit": false
          },
          "trustedProxies": ["10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "127.0.0.0/8"],
          "controlUi": {
            "dangerouslyAllowHostHeaderOriginFallback": true,
            "enabled": true,
            "dangerouslyDisableDeviceAuth": true,
            "allowedOrigins": [
              "http://localhost:18789",
              "http://127.0.0.1:18789",
              "http://10.2.14.244:31080",
              "http://oc-92b953ef-svc.fastclaw.svc.cluster.local:18789"   ------>  bot自己的 serice 域名地址，可以不添加
            ],
            "allowInsecureAuth": true
          }
        },
        "agents": {
          "defaults": {
            "model": {
              "primary": "tokenpony/deepseek-v3-0324"
            }
          }
        },
        "models": {
          "mode": "merge",
          "providers": {
            "tokenpony": {
              "api": "openai-completions",
              "apiKey": "sk-e29970*****************121f52",
              "auth": "api-key",
              "authHeader": false,
              "baseUrl": "https://api.tokenpony.cn/v1",
              "models": [
                {
                  "contextWindow": 200000,
                  "id": "deepseek-v3-0324",
                  "input": [
                    "text"
                  ],
                  "maxTokens": 8192,
                  "name": "deepseek-v3-0324",
                  "reasoning": false
                }
              ]
            }
          }
        }
      }
``` 