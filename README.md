# 企业微信推送服务

由于server酱即将于2021年4月底停止服务，
故自己使用Go语言重新实现类似于server酱的推送服务以作代替。

## 使用方法
浏览器访问[http://localhost:8080/?text=test](http://localhost:8080/?text=test)
即可向企业微信推送`test`。
网页返回错误码`errcode`，具体参考[接口定义](https://open.work.weixin.qq.com/api/doc/90000/90135/90236#%E6%8E%A5%E5%8F%A3%E5%AE%9A%E4%B9%89)

目前仅实现URL的文本推送，之后可能会实现网页端入口以及其他推送方式。

### 配置文件
配置文件需要在项目目录的`config`中配置`config.json`和`token.json`。
其中`token.json`无需手动配置即可自动生成。

`config.json`配置如下
```json
{
  "id": "corpid",
  "secret": "corpsecret",
  "appid": "appid",
  "user": "username"
}
```
其中`id`与`secret`对应[文档](https://open.work.weixin.qq.com/api/doc/90000/90135/91039)
中`corpid`与`corpsecret`，`appid`为企业应用的id，`user`为要发送消息的用户名。