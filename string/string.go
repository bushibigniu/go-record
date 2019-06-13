package main

import "strings"

func main() {

	//\n：换行符
	//\r：回车符
	//\t：tab 键
	//\u 或 \U：Unicode 字符
	//\\：反斜杠自身

	//1.字符串的内容（纯字节）可以通过标准索引法来获取
	//字符串 str 的第 1 个字节：str[0]
	//第 i 个字节：str[i - 1]
	//最后 1 个字节：str[len(str)-1]
	//
	//注意事项 获取字符串中某个字节的地址的行为是非法的，例如：&str[i]。

	//2.字符串拼接符 +
	//在循环中使用加号 + 拼接字符串并不是最高效的做法，
	//更好的办法是使用函数 strings.Join()
	//
	//有没有更好地办法了？有！使用字节缓冲（bytes.Buffer）拼接更加给力

	//strings 和 strconv 包

	strings.hasPrefix(s, prefix string) bool //判断字符串 s 是否以 prefix 开头
	strings.hasSuffix(s, suffix string) bool //判断字符串 s 是否以 suffix 结尾
	strings.Contains(s, substr string) bool // 判断字符串 s 是否包含 substr
	strings.Index(s, str string) int  //Index 返回字符串 str 在字符串 s 中的索引
	// （str 的第一个字符的索引），-1 表示字符串 s 不包含字符串 str：
	s := "this is vary"
	n := strings.Index(s, "i")   //2
	n := strings.Index(s, "th")  //0
	n := strings.Index(s, "tha") //-1

	strings.LastIndex(s, str string) int //返回字符串 str 在字符串 s 中
	//最后出现位置的索引（str 的第一个字符的索引），
	// -1 表示字符串 s 不包含字符串str

	strings.IndexRune(s string ,r, rune) int //查询非 ASCII 编码的字符在父字符串中的位置

}