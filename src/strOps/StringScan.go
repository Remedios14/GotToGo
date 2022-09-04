package strOps

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var sal float32
	var isPass bool
	fmt.Println("请输入姓名")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)

	fmt.Println("分空格输入 薪水，是否通过")
	fmt.Scanf("%f %t", &sal, &isPass)

	fmt.Printf("名字是 %v \n 年龄是 %v \n 薪水是 %v \n 是否通过 %v ", name, age)
}
