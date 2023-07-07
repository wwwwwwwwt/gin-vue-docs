/*
 * @Author: zzzzztw
 * @Date: 2023-07-07 22:54:03
 * @LastEditors: Do not edit
 * @LastEditTime: 2023-07-08 00:36:29
 * @FilePath: /gin-vue-docs/gvd_server/utils/valid/valid.go
 */
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
	InitTrans("zh")
}

func InitTrans(local string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//注册一个获取json tag的自定义方法
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

		trans, ok = uni.GetTranslator(local)
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

// func GetValidMsg(err error, obj any) string {

// 	//使用的时候需要传入obj的指针

// 	getobj := reflect.TypeOf(obj)

// 	//将err接口断言为具体类型

// 	errs, ok := err.(validator.ValidationErrors)

// 	if ok {
// 		//断言成功
// 		for _, e := range errs {
// 			//循环每一个错误消息
// 			//根据报错字段名，获取结构体的具体字段

// 			f, exits := getobj.Elem().FieldByName(e.Field())
// 			if exits {
// 				msg := f.Tag.Get("msg")
// 				if msg == "" {
// 					continue
// 				}
// 				return msg
// 			}
// 		}
// 	}
// 	return Error(err)
// }

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

func ValidError(err error, obj any) (ret string, data map[string]string) {

	data = make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)

	if !ok {
		return err.Error(), data
	}

	//通过反射拿到obj的类型
	getobj := reflect.TypeOf(obj)

	for _, e := range validationErrors {
		f, exits := getobj.Elem().FieldByName(e.Field())
		msg := e.Translate(trans)
		filedname := e.Field()
		// 需要将Name替换为alias
		//先取tag
		alias := filedname
		label, tagok := f.Tag.Lookup("label")
		jsonField, jsonOk := f.Tag.Lookup("json")
		if tagok {
			alias = label
		} else {
			if jsonOk {
				alias = jsonField
			}
		}
		ret += e.Translate(trans) + ";"
	}
	return ret, data
}
