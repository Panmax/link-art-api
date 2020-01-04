# link-art-api

## 项目结构

```
.
├── LICENSE
├── README.md
├── application
│   ├── api
│   │   ├── account.go
│   │   ├── common.go
│   │   └── product.go
│   ├── command
│   │   ├── account.go
│   │   └── product.go
│   ├── middleware
│   │   ├── auth.go
│   │   └── requestid.go
│   ├── representation
│   │   ├── account.go
│   │   └── product.go
│   └── route.go
├── application.ini
├── application.ini.example
├── domain
│   ├── model
│   │   ├── account.go
│   │   ├── model.go
│   │   └── product.go
│   ├── repository
│   │   ├── account.go
│   │   └── product.go
│   └── service
│       ├── account.go
│       └── product.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── config
│   │   └── config.go
│   └── util
│       ├── bind
│       │   └── bind.go
│       ├── response
│       │   └── response.go
│       └── uuid
│           └── uuid.go
└── main.go
```

### TODO

- [x] 作品审核
- [ ] 发送短信验证码
- [ ] 注册时短信验证
- [ ] 密码找回
- [x] 拍卖状态处理
- [x] 关注艺术家
- [x] 取消关注
- [x] 关注列表
- [x] 粉丝列表
- [x] 艺术家搜索
- [x] 商品搜索
- [x] 首页-发现
- [x] 首页-推荐艺术家
- [x] 首页-我关注的
- [x] 首页-分类展示
- [x] 艺术家商品列表
- [x] 消息列表
- [x] 获取地区列表
- [x] 新增收获地址
- [x] 删除收获地址
- [x] 编辑收获地址
- [x] 设置默认收货地址
- [x] 获取收货地址列表
- [ ] 购物车


### 省市县库

https://github.com/wecatch/china_regions