# link-art-api

## 项目结构

```
.
├── LICENSE
├── README.md
├── application.ini
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
├── main.go
└── route
    ├── api
    │   └── account.go
    ├── middleware
    │   └── requestid
    │       └── requestid.go
    ├── param_bind
    │   └── account.go
    └── route.go
```