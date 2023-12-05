# AuditLimit

用于内容审核的限流器(施工中)

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
```

## 通用提示

状态码: 400

```json
{
  "detail": "别闹了"
}
```

## 正常返回

状态码: 200
