package error

//Error类型

//error类型是一个接口类型，这是它的定义：

type error interface {
	Error() string
}
//error是一个内置的接口类型

//自定义Error

//使用GDB调试
//另外建议纯go代码使用delve可以很好的进行Go代码调试
//https://github.com/derekparker/delve




//Go怎么写测试用例
//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md

//另外建议安装gotests插件自动生成测试代码:
//https://github.com/cweill/gotests