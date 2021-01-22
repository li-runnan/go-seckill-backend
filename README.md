# go-seckill-backend
基于beego的抢票秒杀系统(后台管理端)

### 项目部署 ###

1.安装go环境 1.15.3

2.开启go module并配置 

```
export GOPROXY=https://goproxy.io
```

3.安装beego和bee(这里最好按照下方命令安装，否则可能出现bee和beego版本不匹配导致的错误)

```
go get -u github.com/astaxie/beego@develop

go get github.com/beego/bee@v1.12.3
```

4.用Goland打开项目，在命令行中通过bee工具启动项目，支持热编译

```
bee run
```

***

### 项目运行

1.后台管理端:http://localhost:8080/login

2.管理员用户名:admin

3.初始密码:admin