package main //声明文件所在的包,每个go文件必须有归属的包
import "fmt" //import  "fmt" //引入程序中需要用的包，为了使用这个包下的函数   比如println
func main() { //main主函数，一个程序的入口
	fmt.Println("hello,golang!!!") //在控制台打印输出一句话，双引号中的内容会原样输出
	var age int = 18 + 29
	fmt.Println(age)
}
