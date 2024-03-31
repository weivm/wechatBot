## 采用基于openwechat微信扫描登陆和gpt自动回复功能

### 如何启动：
1. git clone https://github.com/weivm/wechatBot.git
2. cd cmd
3. 在config.yaml填入自己的key和gpt域名
4. go run main.go
5. 自动拉起二维码扫码授权



### v1版本目前实现：

1. 添加用户回复功能

2. 用户私聊回复功能

3. 用户进群回复功能

4. 群聊@回复功能

#### Tip:
目前结构比较简单后续会持续优化新功能
项目中third_party有两个gpt文件一个是自定义中转key请求 另外一个是直接请求

