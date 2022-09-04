# yosoro~

编写一个简单的

```go
package main

func main() {
    println("hello!")
}
```

在当前目录下 `go build` 编译所有 `.go` 文件并生成同名的二进制文件(或者 `go build <name.go>` 编译指定文件)，在终端直接执行；

或者 `go run <name.go>` 直接运行指定代码文件

主要特性：
- 自动垃圾回收 GC
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

## Go 语言结构

基础组成包括如下：
- 包声明（必须在第一行声明所属）；`package main` 表示一个可独立执行的程序，每个 Go 应用程序都至少要有一个
- 引入包，格式 `import "pname"`，其中 fmt 包进行 IO，`fmt.Println()`
- 函数（`main` 函数是可执行程序必须的，但若有 `init()` 函数会先执行它）
- 变量
- 语句&表达式
- 注释（同 Java）

标识符（包括常量、变量、类型、函数名、结构字段等）若以大写字母开头如 `Group1` 则可被外部的包使用（类似 public）；否则类似 protected，仅包内部可见

注意 `{` 不能单独构成一行

## Go 语言基础语法

通常一行完成一条语句无需 `;` 换行，但若多条写在同一行（虽然不推荐）则需要用 `;` 分隔

字符串直接用 `+` 拼接

### 关键字

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

以及预定义标识符

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

### 杂项

空格增加可读性

`fmt.Sprintf` 格式化输出字符串

## Go 语言数据类型

数据类型用于声明函数和变量，申请内存

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **布尔型** 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。默认值 **false** |
| 2    | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。默认值（包括 complex64/128） 为 **0** |
| 3    | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。默认值为 **""** （空字符串） |
| 4    | **派生类型:** 包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

派生类型中， 1) 指针；2) 数组；3) Map；4) Chan；5) Func；6) error (接口) 类型的默认值为 **nil**

**注意**：Go 语言变量类型只能 **显式转换**，方法为 `T(v)` 将变量 v 转换为 T 类型

### 数字类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **uint8** 无符号 8 位整型 (0 到 255)                         |
| 2    | **uint16** 无符号 16 位整型 (0 到 65535)                     |
| 3    | **uint32** 无符号 32 位整型 (0 到 4294967295)                |
| 4    | **uint64** 无符号 64 位整型 (0 到 18446744073709551615)      |
| 5    | **int8** 有符号 8 位整型 (-128 到 127)                       |
| 6    | **int16** 有符号 16 位整型 (-32768 到 32767)                 |
| 7    | **int32** 有符号 32 位整型 (-2147483648 到 2147483647)       |
| 8    | **int64** 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

#### 浮点型

| 序号 | 类型和描述                        |
| :--- | :-------------------------------- |
| 1    | **float32** IEEE-754 32位浮点型数 |
| 2    | **float64** IEEE-754 64位浮点型数 |
| 3    | **complex64** 32 位实数和虚数     |
| 4    | **complex128** 64 位实数和虚数    |

#### 其他数字类型

| 序号 | 类型和描述                               |
| :--- | :--------------------------------------- |
| 1    | **byte** 类似 uint8                      |
| 2    | **rune** 类似 int32                      |
| 3    | **uint** 32 或 64 位                     |
| 4    | **int** 与 uint 一样大小                 |
| 5    | **uintptr** 无符号整型，用于存放一个指针 |



## Go 语言变量

```go
var identifier1[, identifier2] type
```

一般使用 `var` 关键字，可同时声明多个变量，在最后指定类型

```go
package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a) // Runoob
    
    var b, c int = 1, 2
    fmt.Println(b, c) // 1 2
}
```

### 三种声明方法

1. 先初始化再赋值，需要指定类型
2. 直接赋值，可以自动判断类型无需显式写 `type`
3. `:=` 完成声明和赋值，只能用于初次声明变量（也只能用在函数体中）

```go
package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}
```

### 值类型和引用类型

#### 值类型

所有像 int、float、bool 和 string 这些基本类型都属于值类型，其变量直接指向内存中的值，在 `=` 赋值时在内存中进行了拷贝

`&i` 获取到变量 `i` 的内存地址，如 0xf840000040 （每次运行的地址都可能不一样）。值类型变量的值存储在栈中。

#### 引用类型

`r1` 存储的是 `r1` 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。

这个内存地址称之为指针，这个指针实际上也被存在另外的某一个值中。

当使用 `r2 = r1` 时，只有引用（地址）被复制

### 注意点

- 在代码块中不使用一个声明的局部变量会**编译错误**，仅赋值也不行
- 全局变量允许声明不使用
- 交换变量值 `a, b = b, a` 在同类型变量间可用
- 空白标识符 `_` 用于抛弃值，其是一个只写变量，不能获取值



## Go 语言常量 const

常量是一个简单值的标识符，在程序运行时不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。定义格式如下

```go
const identifier [type] = value
```

可以省略 `[type]` 交给编译器自行判断

常量不能用 `:=` 语法声明

常量还可以用作枚举，也可以用内置函数计算表达式的值

```go
package main

import "unsafe"
const (
    a = "abc"
    // 常量表达式用到的函数只能使用内置函数
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c) // abc 3 16
}
```

### iota

特殊常量，可以认为是一个可以被编译器修改的常量

在 `const` 语句块中构成行索引，初始化 0，每行 `+1`

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

```go
package main

import "fmt"
const ( // 进行了位运算
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i) // 1 —— 1 << 0
    fmt.Println("j=",j) // 6 —— 3 << 1
    fmt.Println("k=",k) // 12 —— 3 << 2
    fmt.Println("l=",l) // 24 —— 3 << 3
}
```

## Go 语言运算符

算数、关系、逻辑、位、赋值、其他

### 算术运算符

`+ - * / % ++ --` 完全一致，其中 `/` 对整数为整除

### 关系运算符

`== != > < >= <=`

### 逻辑运算符

`&& || !`

### 位运算符

`&` 与 `|` 或 `^` 异或 `<<` 左移 `>>` 右移

### 赋值运算符

略

### 其他运算符

| 运算符 | 描述             | 实例                       |
| :----- | :--------------- | :------------------------- |
| &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| *      | 解引用，用于指针  | *a; 将访问指针指向的值    |

### 运算符优先级

| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |

## Go 语言条件语句

`if ... else if ... else ... ` 类似 `Python` 无需括号，但 `{}` 必须

```go
if v := math.Pow(x, n); v < lim {
}
```

可以在条件语句中声明变量（仅在条件语句 `if ... else ...` 内部生效的局部变量）

### Switch 语句

自上到下逐一测试直到匹配（默认在 `case` 最后自带 `break` ，若不希望终止可以使用 `fallthrough` ）支持多值 match

```go
switch var1 {
    case val1:
    	...
    case val2, val3, val4:
    	...
	default:
    	...
}
```

其中 `var1` 可以是任何类型，而 `val1` 和 `val2` 只需要是同类型值（或表达式）即可

**注意**：
1. default 不是必须的

#### 当做 if 控制

可以省略 Switch 后的条件来试做 if 条件

```go
switch {
    case age == 10 :
        ...
    case age == 20 :
        ...
}
```

#### Type Switch

`switch` 的用法之一，用于判断某个 `interface` 变量中实际存储的变量类型

```go
switch x.(type){
    case type:
    	statement(s)
    case type:
    	statement(s)
    default: /* 可选 */
    	statement(s)
}
```

#### fallthrough

写在原来 `break` 的位置，当遇到时会无视下一条 `case` 的条件执行下一条

### Select 语句

类似于通信的 `switch` 语句，每个 `case` 必须是一个通信操作（发送或接收）

`select` 随机执行一个可运行的 `case` ，如果没有 `case` 可运行，则会阻塞直到有。

一个默认的子句应该总是可运行的

```go
package main

import "fmt"

func main() {
   var c1, c2, c3 chan int
   var i1, i2 int
   select {
      case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
      case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
      case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
         } else {
            fmt.Printf("c3 is closed\n")
         }
      default:
         fmt.Printf("no communication\n")
   }    
}
```

`select` 语句的语法如下：

- 每个 case 都必须是一个通信

- 所有 channel 表达式都会被求值

- 所有被发送的表达式都会被求值

- 如果人以某个通信可以进行，它就执行，**其他被忽略**

- 如果有多个 case 都可以运行，Select 会**随机**公平地选出一个执行。**其他不会执行**。

  否则：

  1. 如果有 default 子句，则执行该语句
  2. 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值



## Go 语言循环语句

`for` 循环 

- `for init; condition; post {}` 可以省略其中的 `init` 和 `post` （但要有分号）
- `for condition {}` 类同 `while` 

- 直接 `for {}` 类同 C 中的 `for (;;)` ，会直接无限循环

#### For-each range 循环

可以对字符串、数组、切片等进行迭代输出元素

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

### 循环控制语句

#### break 语句

- 用于循环语句中跳出循环，并开始执行**循环之后**的语句
- 在 switch 语句中执行一条 case 后跳出语句
- 在多重循环中用标号 label 标出想 break 的循环

```go
package main

import "fmt"

func main() {

    // 不使用标记
    fmt.Println("---- break ----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                        fmt.Printf("i2: %d\n", i2)
                        break
                }
        }

    // 使用标记
    fmt.Println("---- break label ----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                break re
            }
        }
}
```

#### continue 语句

可以使用 label ，其他无变化

```go
package main

import "fmt"

func main() {

    // 不使用标记
    fmt.Println("---- continue ---- ")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                continue
            }
    }

    // 使用标记
    fmt.Println("---- continue label ----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i: %d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                    fmt.Printf("i2: %d\n", i2)
                    continue re
                }
        }
}
```

#### goto 语句

不推荐使用

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++    
   }  
}
```

## Go 语言函数

最少有一个 main() 函数

函数声明告诉编译器 函数的名称，返回类型和参数

Go 函数不支持重载

### 函数定义

```go
func function_name( [parameter list] ) [return_types] {
    // 函数体
}
```

函数定义解析：

- func ：函数由 func 开始声明
- function_name ：函数名称，参数列表和返回值类型构成了函数签名
- parameter list ：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的时参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数
- return_types ：返回类型，函数返回一列值。return_types 时该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的
- 函数体：函数定义的代码集合

### 函数调用

可以前缀关键词实现不同形式的函数调用

### [defer](http://blog.go-zh.org/defer-panic-and-recover)

```go
package main
import "fmt"
func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
```

defer 会将函数推迟到外层函数返回之后执行

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用

推迟的函数调用会被压入一个栈中，最后按照后进先出顺序调用

- 函数本身也是一种数据类型，可以作为参数传入其他函数，在其他函数定义时需要声明传入函数的 **出入参**

### 函数返回多个值

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Google", "Runoob")
   fmt.Println(a, b)
}
```

### 函数参数

函数如果使用参数，该变量可称为函数的形参，形参就像定义在函数体内的局部变量

- 值传递

在调用函数时将实际参数复制一份传递到函数中，从而在函数中的修改不会影响到实际参数

**默认情况**下，Go 语言使用的是值传递，针对基本类型和数组

- 引用传递

在调用函数时将实际参数的地址传递到函数中，从而在函数中对参数所进行的修改将影响到实际参数

操作上是将指针参数传递到函数内

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前，a 的值 : %d\n", a )
   fmt.Printf("交换前，b 的值 : %d\n", b )

   /* 调用 swap() 函数
   * &a 指向 a 指针，a 变量的地址
   * &b 指向 b 指针，b 变量的地址
   */
   swap(&a, &b)

   fmt.Printf("交换后，a 的值 : %d\n", a )
   fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址上的值 */
   *x = *y      /* 将 y 值赋给 x */
   *y = temp    /* 将 temp 值赋给 y */
}
```

#### 可变参数

入参中定义 `args... type` 实现可变参数，在函数体内试做一个 slice，使用索引访问各个值

```go
func sum(args... int) res int {
    res := 0
    for _, v := range args {
        res += v
    }
    return
}
```

### 函数用法

- 传递给变量构成一个函数变量
- 闭包，借助匿名函数可直接使用函数内的变量而不必声明

```go
package main

import "fmt"

func getSequence() func() int {
   i:=0
   return func() int {
      i+=1
     return i  
   }
}

func main(){
   /* nextNumber 为一个函数，函数 i 为 0 */
   nextNumber := getSequence()  

   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   
   /* 创建新的函数 nextNumber1，并查看结果 */
   nextNumber1 := getSequence()  
   fmt.Println(nextNumber1())
   fmt.Println(nextNumber1())
}
```

- 方法，即一个包含了接受者的函数，接受者可以使命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集（类似面向对象）

```go
func (variable_name variable_data_type) function_name() [return_type] {
    /* 函数体 */
}
```

```go
package main

import (
   "fmt"  
)

/* 定义结构体 */
type Circle struct {
  radius float64
}

func main() {
  var c1 Circle
  c1.radius = 10.00
  fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
  //c.radius 即为 Circle 类型对象中的属性
  return 3.14 * c.radius * c.radius
}
```

## Go 语言变量作用域

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围

变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数（相当于函数体内的局部变量）

### 局部变量

在函数体内声明，包括参数和返回值变量

可以与全局变量**同名**，在函数内会优先使用局部变量

### 全局变量

在函数体外声明（`main` 函数外），可以在整个包甚至外部包（被导出后）使用

## Go 语言数组

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，类型可以是任意的原始类型（整型、字符串等）或自定义类型

一个数组变量表示整个数组，在赋值或传递时会复制整个数组，如果希望避免复制则可以传递数组的**指针**

### 声明数组

需要指定元素类型及元素个数

`var variable_name [SIZE] variable_type` —— `var balance [10] float32`

### 初始化数组

```go
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
balance1 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

var balance2 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
// 或（长度不确定时可以使用 ... 会在初始化时自行判断长度）
balance3 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

//  将索引为 1 和 3 的元素初始化
balance4 := [5]float32{1:2.0,3:7.0}
```

`{}` 中的元素个数不能大于 `[]`

### 多维数组

套几个 `[]` 来声明，`{}` 来初始化即可

### 访问数组元素

同其他语言

## Go 语言指针

取地址符 `&`

### 什么是指针

一个指针变量指向了一个值的内存地址，声明格式如下：

`var var_name *var-type` 例如：

```go
var ip *int /* 指向整型 */
var fp *float32 /* 指向浮点型 */
```

### 如何使用指针

流程：

1. 定义指针变量
2. 为指针变量赋值
3. 访问指针变量中指向地址的值（**在指针变量前加上 `*` 来获取指针所指向的值**）

```go
package main
import "fmt"
func main() {
    var a int = 20
    var ip *int
    
    ip = &a
    
    fmt.Printf("a 变量的地址是：%x\n", &a)
    
    fmt.Printf("ip 变量储存的指针地址：%x\n", ip)
    
    fmt.Printf("*ip 变量的值：%d\n", *ip)
}
```

### 特殊指针

#### 空指针

当一个指针被定义后没有分配到任何变量时，它的值为 `nil`，也被称为空指针

#### 指针数组

`var ptr [Max]*int` 声明一个存储指针的数组后逐个给其赋值地址

`ptr[i] = &a[i]`

#### 指向指针的指针

`var ptr **int` 即已有一个指针，在获取其地址赋值给一个指针

#### 指针作为函数参数

在函数定义时 `func swap(x *int, y *int) {}`

在函数调用时 `swap(&a, &b)` （回想**赋值格式**即可）

## Go 语言结构体 Struct

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

可同一般变量一样作为函数参数和指针（且仍用 `.` 访问成员）

### 定义结构体

```go
type struct_variable_type struct {
    member definition
    ...
}
variable_name := struct_variable_type {value1, value2, ...}
// 或
variable_name := struct_variable_type {key1: value1, key2: value2, ...}
```

### 访问结构体成员

`结构体.成员名` 可用于获取和赋值

## Go 语言[切片](https://blog.go-zh.org/go-slices-usage-and-internals) （Slice）

切片是对数组的抽象，是一种“动态数组”，可追加元素

切片的零值是 `nil` ； `nil` 切片的长度和容量为 0 且没有底层数组

对切片的修改会作用到实际数组上

`var identifier []type` 未指定大小的数组即为一个切片

```go
var slice1 []type = make([]type, len)
// 也可以简写为
slice1 := make([]type, len)
```

此处的 `make([]T, length, capacity)` 中 **capacity** 为可选参数 - 容量（最大长度）

具体切片操作同 `python` （貌似没有步长）

使用自带函数 `len()` 和 `cap()` 可获取长度和容量

切片在未初始化之前默认为 nil，长度为 0

`copy(toSlice, fromSlice)` 进行切片内容拷贝；

`Slice<T> = append(Slice<T>, Ts...)` 进行元素追加（不增加容量）

`Slice<T> = AppendByte(Slice<T>, Ts...)` 自动扩容的追加（原长加上追加长后 * 2）

```go
a := []string{"John", "Paul"}
b := []string{"George", "Ringo", "Pete"}
a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
// 两个切片 append 时要扩展为 ...
```



## Go 语言范围（Range）

range 关键字用于 for 循环中迭代数组（array）、切片、通道（channel）或集合（map）的元素，无需括号

在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key 对

```go
package main
import "fmt"
func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```

## Go 语言 Map （集合）

Map 是一种无序的键值对集合，通过 key 来快速检索数据

```go
/* 声明变量，默认 map 是 nil ，此时不能直接进行插入修改，故一般用 make 函数初始化 */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)

/* 通过双赋值检测某个键是否存在 */
elem, ok := m[key]
// 若存在则 ok 为 true
```

如果不初始化 map，则会创建一个 nil map，其不能用来存放键值对；**所以一定要 make 或者指向已存在的**

`delete(mapVar, mapKey)` 用于删除指定 Map 的指定 Key

## Go 语言类型转换

用于将一种数据类型的变量转换为另外一种类型的变量

`type_name(expression)`

## Go 语言接口

是一种数据类型，把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

```go
package main
import (
    "fmt"
)
type Phone interface {
    call()
}
type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}
type IPhone struct {
}
func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}
func main() {
    var phone Phone

    phone = new(NokiaPhone)
    phone.call()

    phone = new(IPhone)
    phone.call()
}
```

## Go 错误处理

通过 error 类型接口提供，其定义如下

```go
type error interface {
    Error() string
}
```

```go
func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    // 实现
}
```

用 `errors.New(str)` 返回一个带有错误描述的 error 对象

## Go 并发

只需要通过关键字来启动 goroutine 即可支持并发

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的

`go 函数名（参数列表）` 

以一个不同的、新创建的 goroutine 来执行一个函数

同一个程序中的所有 goroutine 共享同一个地址空间

### 通道（channel）

是用来传递数据的一个数据结构，可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯

操作符 `<-` 用于指定通道的方向，发送或接收。若未指定方向则为双向通道

```go
ch := make(chan int) // 声明
 
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v
```

```go
package main

import "fmt"

func sum(s []int, c chan int) {
        sum := 0
        for _, v := range s {
                sum += v
        }
        c <- sum // 把 sum 发送到通道 c
}

func main() {
        s := []int{7, 2, 8, -9, 4, 0}

        c := make(chan int)
        go sum(s[:len(s)/2], c)
        go sum(s[len(s)/2:], c)
        x, y := <-c, <-c // 从通道 c 中接收

        fmt.Println(x, y, x+y)
}
```

#### 通道缓冲区

默认没有缓冲区，可通过 `make` 的第二个参数指定缓冲区大小

`ch := make(chan int, 100)`

带缓冲区的通道允许发送端的数据发送和接收端的数据获取处于异步状态，就是说发送端发送的数据可以放在缓冲区里面，可以等待接收端去获取数据，而不是立刻需要接收端去获取数据。

不过由于缓冲区的大小是有限的，所以还是必须有接收端来接收数据的，否则缓冲区一满，数据发送端就无法再发送数据了。

**注意**：如果通道不带缓冲，发送方会阻塞直到接收方从通道中接收了值。如果通道带缓冲，发送方则会阻塞直到发送的值被拷贝到缓冲区内；如果缓冲区已满，则意味着需要等待直到某个接收方获取到一个值。接收方在有值可以接收之前会一直阻塞。

#### 通道遍历与关闭通道

```go
package main

import (
        "fmt"
)

func fibonacci(n int, c chan int) {
        x, y := 0, 1
        for i := 0; i < n; i++ {
                c <- x
                x, y = y, x+y
        }
        close(c)
}

func main() {
        c := make(chan int, 10)
        go fibonacci(cap(c), c)
        // range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
        // 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
        // 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
        // 会结束，从而在接收第 11 个数据的时候就阻塞了。
        for i := range c {
                fmt.Println(i)
        }
}
```
