# 开发中... 

> 请勿使用 thx~

#### 项目规划

2018年12月06日10:30:21

client 放主要实现

​	|-- charge 支付

​	|-- notify 回调

common  公共接口的定义，数据的定义，基础类公共实现  

​	分ali和wx 里面会放公共的实现 

constant 常量

errors 所有的错误信息放这里

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

还在修改中的支付主流程：
```
// 1. 初始化数据，设置各种支付要用到的属性
// 2. 检测一些配置的数据，是否有问题
// 3. 构建要请求三方的数据【不管是否要请求】
// 3.1 构建初始请求数据
// 3.2 构建签名
// 3.3 拼接完成
// 4. 发起请求，获取返回的数据
// 5. 利用返回的数据，构建返回给前端的数据
```

