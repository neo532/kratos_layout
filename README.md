# Kratos Project Template

## 项目初始化

### 让你的本地拉代码支持公司的代码库
```
export GOPROXY=https://goproxy.cn
#export GOPRIVATE=*.gitee.com
#export GOSUMDB="off"
#git config --global url."ssh://git@git.gitee.com/".insteadOf https://git.gitee.com/
```

### Install Kratos
```
go install github.com/go-kratos/kratos/v2
```

### Create a template project
```
kratos new git.gitee.com/group_name/{xxx}
cd {xxx}
```

### 一些工具初始化
```
make init
```

### 创建配置文件模板
！！！注意！！！，每个服务都要全局唯一的端口号，请在这里备注[链接](https://wiki.gitee.com/a.html)
```
make initConfig
```



## 运行
```
// 下面为开发环境的运行，正式部署可参考cmd目录下的Dockerfile文件
make runApi         //API服务
make runConsumer    //消费者服务
make runScript      //脚本，该命令会执行一个脚本例子，具体业务自行定义。
```



## 一些维护工具
### 修改配置文件结构时
```
cd {xxx}
make config // 或者make generate
```


### 修改依赖注入框架wire文件时
```
go generate
```
