# 检查18位身份证号是否有效
> 本程序只能身份证格式进行校对，以及是否通过计算校验；不校验身份证号是否真实存在。

## 使用方法
### 安装
```
go get github.com/yzha5/cci
```
### 使用
```go
package main
import "github.com/yzha5/cci"

func main() {
    // 下面所展示的身份证号码仅为测试，如有巧合碰上你的号码纯为巧合
    info, ok, err := cci.Check("110101199901011232") // 有效
    // info, ok, err := Check("11010119990101123x") // 无效
    if err != nil {
        panic(err)
    }
    fmt.Println("是否有效", ok) // 是否有效 true
    fmt.Println(info) // &{北京 东城  1999-01-01 00:00:00 +0000 UTC 男}
}

```