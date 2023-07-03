# gin-vue-docs 静态网页小项目

## 项目目录结构

1. api:接口目录
2. config: 存放记录配置的结构体目录
3. core: 初始化操作
4. docs： swag生成的api目录
5. flag：命令行相关的初始化
6. global: 存放全局变量
7. middleware：中间件
8. models：表结构
9. routers：gin路由目录
10. service：项目与服务目录
11. tastdata：测试目录
12. utiles：常用工具
13. settings：配置文件

## 配置文件读取

```
	"gopkg.in/yaml.v3"

```

## 配置lgorus日志

```
go get github.com/sirupsen/logrus
```

### logrus常用方法

```go
logrus.Debugln("Debugln")
logrus.Infoln("Infoln")
logrus.Warnln("Warnln")
logrus.Errorln("Errorln")
logrus.Println("Println")

// 输出如下
time="2022-12-17T14:02:01+08:00" level=info msg=Infoln   
time="2022-12-17T14:02:01+08:00" level=warning msg=Warnln
time="2022-12-17T14:02:01+08:00" level=error msg=Errorln 
time="2022-12-17T14:02:01+08:00" level=info msg=Println

//debug的没有输出，是因为logrus默认的日志输出等级是 info
fmt.Println(logrus.GetLevel())  // info

```
1. 日志等级：  
```
PanicLevel  // 会抛一个异常  
FatalLevel  // 打印日志之后就会退出  
ErrorLevel  
WarnLevel  
InfoLevel  
DebugLevel  `
TraceLevel  // 低级别  
```
2. 更改日志级别  

```go
logrus.SetLevel(logrus.DebugLevel)
```

### 日志模板直接复制粘贴


## 配置链接mysql

1. 依赖
```go
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
```
2. 事先在mysql中创建db名为”gvd_db“

```
>mysql -uroot -p123456
create database gvd_db;
```
3. settings.yaml中配置mysql

```yaml
mysql:
  host: 127.0.0.1 # 如果不配置host就不会连接 mysql
  port: 3306
  config: charset=utf8mb4&parseTime=True&loc=Local
  db: gvd_db
  username: root
  password: 123456
  logLevel: error
```

4. 配置config_mysql init_mysql global三个文件


## 配置redis
1. 依赖
```
	"github.com/go-redis/redis"
```
2. 连接
```go
func InitRedis() *redis.Client {

	redisconfig := global.Config.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     redisconfig.Addr(),
		Password: redisconfig.Password,
		DB:       0,
		PoolSize: redisconfig.PoolSize,
	})

	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := client.Ping().Result()
	if err != nil {
		global.Log.Fatal("redis 连接失败, err: ", err.Error())
	}
	return client
}

```

## 命令行参数控制
1. flags文件，建立一个option结构体，里面包含需要输入的参数，当Parse函数检查当前值有无被修改，如果被修改就调用对应的函数来改。


## 后端表结构

怎么做表结构？原则

1. 从依赖项少的开始