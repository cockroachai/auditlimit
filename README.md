# AuditLimit

用于内容审核的限流器,本代码更多的是为了演示如果使用限流器,大家可以根据自己的需求进行修改。

如果不想部署,可以使用公共限流器: `https://auditlimit.closeai.biz/audit_limit`

## 部署方法

创建`docker-compose.yml`文件

```yml
version: '3'
services:
  auditlimit:
    image: xyhelper/auditlimit
    restart: always
    ports:
      - 9611:8080
    environment:
      OAIKEY: "" # OpenAI API key 用于内容审核
      AUTO: "200/3h"
      TEXT-DAVINCI-002-RENDER-SHA: "200/3h"
      GPT-4O-MINI: "200/3h"
      GPT-4O: "60/3h"
      GPT-4: "20/3h"
      GPT-4O-CANMORE: "30/3h"
      O1-PREVIEW: "7/24h"
      O1-MINI: "50/24h" # 模型名称: "次数/时间" 时间单位: h(小时) m(分钟) s(秒)  模型名称要改成大写

```

然后执行

```bash
docker-compose up -d
```

限流器接口地址为: `http://ip:9611/audit_limit`

## 超速返回格式

状态码: 429

```json
{
  "detail": {
    "clears_in": 252,
    "code": "model_cap_exceeded",
    "message": "You have sent too many messages to the model. Please try again later."
  }
}
````

## 通用提示

状态码: 400

```json
{
  "detail": "别闹了"
}
```

## 正常返回

状态码: 200


## 内容审核

配置环境变量

OAIKEY: "sk-xxxxxx"  # api.openai.com可用的key或sess

将启用moderations接口内容审核