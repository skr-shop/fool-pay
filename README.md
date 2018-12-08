# fool-pay
A fool person can use payment.

#### 项目规划

2018年12月06日10:30:21

client 放主要实现

​	|-- charge 支付

​	|-- notify 回调

common  公共接口的定义，数据的定义，基础类公共实现  

​	分ali和wx 里面会放公共的实现 

constant 常量

errors 所有的错误信息放这里

main 我放里测试

notify 回调的定义

```
    // func Run() (retdata interface{}, iswrong errors.PayError) {
	// 异常捕获的两种写法
	// 第一种
	// iswrong = errors.PayError{} //主要作用是分配内存
	// defer errors.Catch(&iswrong) //正常捕获
	// 第二种
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		switch err.(type) {
	// 		case errors.PayError:
	// 			pe := err.(errors.PayError)
	// 			iswrong.ErrorCode = pe.ErrorCode
	// 			iswrong.Message = pe.Message
	// 		}
	// 	}
	// }()
    // }
```