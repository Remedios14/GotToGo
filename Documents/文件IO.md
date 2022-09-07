# yosoro~

使用 os.File 结构体来抽象文件对象，完成文件相关操作

- os
- bufio

## 基本功能

- `func Open(name string) (file *File, err error)` ：打开文件，得到一个对象指针
- `func (f *File) Close() error` ：文件的关闭方法，推荐每次成功打开文件后都 defer 一个 Close 调用

### 文件读取

1. 使用 bufio.Reader 来进行读取，可以设置缓冲区实现逐内容读取，有具体的读取方法
2. 使用 io.ioutil.ReadFile 函数读取完整文件，适用于**不大的文件**，读入后为一个 `[]byte`，文件的开闭已经封装好

### 文件写入

位于 os 包内，使用 OpenFile 函数包含写入模式打开文件（其实有很多其他控制模式），然后借助 bufio.Writer 来写入

- `func OpenFile(name string, flag int, perm FileMode) (file *File, err error)`
    - flag ：同一个包内的一些常量，指示读写模式、是否创建文件等（TODO：后续包介绍内补充）
    - perm ：表示控制权限，r->4 / w->2 / x->1

### 文件存在与否

使用 os.Stat() 根据返回的 error 值来判断

1. 为 nil ：文件或目录存在
2. 使用 os.IsNotExist 判断 ：不存在
3. 其他错误 ：其他原因造成无法访问，无从判断

```go
_, err := os.Stat(path)
if err == nil {
    return true, nil
}
if os.IsNotExist(err) {
    return false, nil
}
```

### 其他综合

- 文件内容拷贝 from "io" ：`func Copy(dst Writer, src Reader) (written int64, err error)`

#### 命令行参数

os.Args 是一个 string 的切片，是传统的处理命令行参数的方法，但不利于配置，基本不用

- "flag" 包解析命令行 ：先声明变量，然后使用 `flag.XXXVar(dst interface{}, "parse var", "default val", "description")` 来赋值

```go
var user string
var port int
flag.StringVar(&user, "u", "", "用户名，默认为空")
flag.IntVar(&port, "port", 3306, "端口号，默认为 3306")
// 使用时
kk.exe -u mikyi -port 8808
```

#### json 序列化

使用 "encoding/json" 包

- 对象转 json ：`json.Marshal(v interface{})` 进行序列化，可以对 **结构体、map、slice** 进行，其中结构体的序列化可以在定义时用 tag 给出别名
- json 转对象 ：`json.Unmarshal(serial []byte, dst interface{})` 反序列化，注意总是先声明一致的结构再传入参数
    - 通常总是有一个 string 变量，记得调用时使用 `[]byte(str)` 强转