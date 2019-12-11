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