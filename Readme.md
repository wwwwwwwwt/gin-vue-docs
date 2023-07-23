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

1. 从依赖项少的开始 前四张表最重要


### 关键表



#### 角色表
```go
type RoleModel struct {
	Model
	Title    string     `gorm:"size:16;not null;comment:角色名称" json:"title"`                                    //角色名字
	Pwd      string     `gorm:"size:64;comment:角色的密码" json:"-"`                                                //角色密码
	IsSystem bool       `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               //
	DocList  []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinRederences:DocID" json:"-"` //角色拥有的文档列表
}
```
#### 文档表



```go
type DocModel struct {
	Model
	Title           string      `gorm:"comment:文档标题" json:"title"`
	Content         string      `gorm:"comment:文档内容" json:"-"`
	DiggCount       int         `gorm:"comment:点赞量" json:"diggcount"`
	LookCount       int         `gorm:"comment:浏览量" json:"lookCount"`
	Key             string      `gorm:"comment:key;not null;unique" json:"key"`
	ParentID        *uint       `gorm:"comment:父文档id;column:parentID" json:"parentID"`
	ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"` //父文档
	Child           []*DocModel `gorm:"foreignKey:ParentID" json:"-"` //他会有子孙文档
	FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`
	UserCollDocList []UserModel `gorm:"many2many:user_coll_doc_models;joinForeignKey:DocID;JoinRederences:UserID" json:"-"`
}


```

#### 角色文档表

```go

type RoleModel struct {
	Model
	Title    string     `gorm:"size:16;not null;comment:角色名称" json:"title"`                                    //角色名字
	Pwd      string     `gorm:"size:64;comment:角色的密码" json:"-"`                                                //角色密码
	IsSystem bool       `gorm:"column:isSystem;comment:是否是系统角色" json:"isSystem"`                               //
	DocList  []DocModel `gorm:"many2many:role_doc_models;joinForeignKey:RoleID;JoinRederences:DocID" json:"-"` //角色拥有的文档列表
}


```

#### 用户表

```go
type UserModel struct {
	Model
	UserName  string    `gorm:"column:userName;size:36;unique;not null" json:"userName"` //用户名
	Password  string    `gorm:"column:password;size:256" json:"password"`                //密码
	Avatar    string    `gorm:"column:avatar;size:256" json:"avatar"`                    //头像
	NickName  string    `gorm:"column:nickName;size:36" json:"nickName"`                 //昵称
	Email     string    `gorm:"column:email;size:128" json:"email"`                      //邮箱
	Token     string    `gorm:"column:token;size:36" json:"token"`                       //其他平台的唯一id
	IP        string    `gorm:"column:ip;size:16" json:"ip"`                             //ip
	Addr      string    `gorm:"column:addr;size:64" json:"addr"`                         //地址
	RoleID    uint      `gorm:"column:roleId" json:"roleId"`                             //用户对应角色
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"-"`
}

```

#### 用户收藏文档表

```go
type UserCollDocModel struct {
	Model
	DocID     uint      `gorm:"column:docID" json:"docID"`
	DocMolel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"column:usrID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}

```

#### 用户密码访问文档表

```go
type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:userID" json:"userID"`
	DocID  uint `gorm:"column:docID" json:"docID"`
}


```

#### 图像表

```go
type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户id" json:"userID"`
	RoleModel RoleModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileNmae;size:128" json:"filename"`
	Size      int64     `gorm:"column:size;comment:文件大小，单位字节" json:"size"`
	Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`
	Hash      string    `gorm:"column:hash;size:64;comment:文件的哈希" json:"hash"`
}
```

#### 登陆记录表

```go
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` //登录ip
	NickName  string    `gorm:"column:nickname;size:42" json:"nickName"`
	UA        string    `json:"ua"` //ua
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` //登陆设备
	Addr      string    `gorm:"size:64" json:"addr"`
}

```

#### 文档数据量表

```go
type DocDataModel struct {
	Model
	DocID     uint   `gorm:"column:docID" json:"docID"`
	DocTitle  string `gorm:"column:docTitle" json:"docTitle"`
	LookCount int    `gorm:"column:lookCount" json:"lookCount"`
	DiggCount int    `gorm:"column:diggCount" json:"diggCount"`
	CollCount int    `gorm:"column:collCount" json:"collCount"`
}

```

## jwt 全程是json web token， 主要用于用户登录鉴权

## 用户管理api(/api  /service)
1. api：
* enter.go
* user_create.go : 用于验证发来post报文消息体内，有没有填好我们需要的字段
2. service：
* common：enter.go: 定义了我们使用回复报文的方法

## 创建自定义翻译器， utils/valid/valid.go

```go
package valid

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	trans ut.Translator
)

func init() {

}

func InitTrans(local string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//注册一个获取json tag的自定义方法，这一段作用是优先拿到前端传来的label字段，
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			if name == "" {
				//没有label
				name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			}
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() // 创建中文翻译器
		enT := en.New() // 创建英文翻译器

		// 第一个参数是备用fallback的语言环境
		// 后面的参数是支持的语言环境
		// uni := ut.New(zhT, zhT)也是可以的
		uni := ut.New(enT, zhT, enT)

		// local通常取决于额http请求头的"Aceept-language"
		var ok bool

		//也可以使用uni.FindTranslator(...)来传入多个local进行查找

		trans, ok := uni.FindTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTanslator(%s) faild", local)
		}

		//注册翻译器

		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

//GetValidMsg 返回结构体中的msg参数

func GetValidMsg(err error, obj any) string {

	//使用的时候需要传入obj的指针

	getobj := reflect.TypeOf(obj)

	//将err接口断言为具体类型

	errs, ok := err.(validator.ValidationErrors)

	if ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误消息
			//根据报错字段名，获取结构体的具体字段

			f, exits := getobj.Elem().FieldByName(e.Field())
			if exits {
				msg := f.Tag.Get("msg")
				if msg == "" {
					continue
				}
				return msg
			}
		}
	}
	return Error(err)
}

func Error(err error) (ret string) {
	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return err.Error()
	}

	for _, e := range validationErrors {
		ret += e.Translate(trans) + ";"
	}
	return ret
}



```

## 密码哈希与校验

```go
package pwd

import (
	"gvd_server/global"

	"golang.org/x/crypto/bcrypt"
)

func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err)
	}
	return string(hash)
}

func CheckPwd(hashpwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

```