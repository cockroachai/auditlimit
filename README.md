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
      LIMIT: 40  # 限制每个userToken允许的次数
      PER: "3h" # 限制周期 1s, 1m, 1h, 1d, 1w, 1y
      O1LIMIT: 10  # 限制每个userToken允许的O1模型次数
      O1PER: "1w" # 限制周期 1s, 1m, 1h, 1d, 1w, 1y

    

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