一、命名
统一的命名规则有利于提高代码的可读性，好的命名仅仅通过命名就可以获取到足够多的信息。

1. 包命名： package
保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义。
尽量和标准库不要冲突。
包名应该为小写单词，不要使用下划线或者混合大小写。
如：package demo   package main

2. 文件命名
尽量采取有意义的文件名，简短，有意义，应该为小写单词。
使用下划线分隔各个单词。
my_test.go

3. 结构体命名
	3.1. 采用驼峰命名法，首字母根据访问控制大小写。
	3.2. struct申明和初始化格式采用多行，如下：
//多行申明
type User strut{
		Username string
		Email string
}
//多行初始化
u := User{
	Username : "astaxie"
	Email: "astaxie@gmail.com"
}

4. 接口命名
	4.1 命名规则基本和结构体类似
	4.2 单个函数的结构名以“er”作为后缀，例如：Reader, Writer
type Reader interface{
	Read(p []byte)(n int, err error)
}

5. 变量命名
	5.1 和结构体类似，变量名遵循驼峰法。首字母大小写，但遇到特有名词时，需要遵循一下规则：
		如果变量为私有，且特有名词为首个单词，则使用小写，如apiClient
		其他情况都应当使用该名词原有的写法，如APIClicent、repoID、UserID
		错误示例：UrlArray，应该写成urlArray,或则URLArrray
	5.2 若变量类型为bool类型，则名称应以Has，Is，Can或Allow开头。
	var isExist bool
	var hasConflict bool
	var canManage bool
	var allowGitHook bool

6. 常量命名
常量均需使用全部大写字母组成，并使用下划线分词。
const APP_VER = "1.0"

如果是枚举类型的常量，需要先创建相应类型：
type Scheme string
const(
	HTTP Scheme = "http"
	HTTPS Scheme = "https"
)

7. 关键字
这些关键字不能用作常量或变量或任何其他标识符名称。
break default func interface select
case defer Go map Struct
chan else Goto package Switch
const fallthrough if range Type
continue for import return Var

二、注释
//当行注释
/*
多行注释
 */
1. 包注释
//包的基本简介（包名，简介）
//创建者		格式：创建人：rtx名
//创建时间	格式：创建时间：yyyyMMdd
如下：
//util包, 该包包含了项目共用的一些常量,封装了项目中一些共用函数.
//创建人: xxx
//创建时间: 20210601

2. 结构(接口)注释
每个自定义的结构体或者接口都应该有注释说明,该注释对结构进行简要介绍.
放在结构体定义的前一行,格式为:结构体名,结构体说明.
同时结构体内的每个成员变量都要有说明,该说明放在成员变量的后面(注意对齐),示例如下:
//User, 用户对象,定义了用户的基础信息.
type User struct{
	Username string //用户名
	Email string // 邮箱
}

3. 函数(方法)注释
每个函数,或者方法(结构体或者接口下的函数称为方法)都应该有注释说明,
函数的注释应该包括三个方面(严格按照此顺序撰写):
	3.1 简要说明,格式说明:以函数名开头,","分隔说明部分.
	3.2 参数列表:每行一个参数,参数名开头,","分隔说明部分
	3.3 返回值:每行一个返回值.
示例如下:
//NewAttrModel , 属性数据层操作类的工程方法
//参数:
//		ctx: 上下文信息
//返回值:
//		属性操作类指针
func NewAttrModel(ctx *common.Context) *AttModel{

}

4. 代码逻辑注释
对一些关键位置的代码逻辑,或局部较为复杂的逻辑,需要有相应的逻辑说明,方便阅读.
//详细写逻辑.

5. 注释风格
中英间空格, 中英标点也如此.
// 从 Redis 中批量读取属性, 对于没有读到的 id , 记录到一个数组里面,
	5.1 建议全部使用单行注释
	5.2 和代码的规范一样,单行注释不要过长,禁止超过120字符.

三 代码风格
四 常用工具
强烈建议使用goimport,该工具在gofmt的基础上增加了自动删除和引入包.

go get golang.org/x/tools/cmd/goimports

go vet vet工具可以帮我们静态分析我们的源码存在的各种问题,
例如多余的代码,提前return的逻辑,struct的tag是否符合标准等.

go get golang.org/x/tools/cmd/vet

使用如下:
	go vet .






























