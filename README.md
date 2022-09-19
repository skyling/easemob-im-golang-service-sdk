## 环信 IM golang SDK

### 介绍
此项目是对环信 IM [服务端 API](https://docs-im.easemob.com/ccim/rest/overview) 的封装，这样做是为了节省服务器端开发者对接环信 API 的时间，只需要配置自己的 appkey 相关信息即可使用。

此项目是参照官方 [PHP SDK](https://github.com/easemob/im-php-server-sdk) 的一个golang 版本

### 功能
此项目提供了用户、消息、群组、聊天室等资源的操作管理能力。

### 依赖
- Go 1.18+

### 安装
```
go get github.com/skyling/easemob-im-golang-service-sdk
```

### 目录结构

- agora 声网服务
- cache 缓存组件实例化
- easemob_im 环信IM 接口实现
- types 请求参数与返回数据 结构体

### 功能点

- [x] 用户体系集成
- [ ] 推送设置
- [ ] 消息管理
- [ ] 用户属性
- [ ] 用户关系管理
- [ ] 群组
- [ ] 聊天室
- [ ] 在线状态订阅
- [ ] 消息表情恢复
- [x] 即时推送
