package utils

type Option func(data *JsonData)

type JsonData struct {
	Code    int         `json:"code"`    //返回code
	Data    interface{} `json:"data"`    //返回的data
	Message string      `json:"message"` //返回的message
}

//设置参数
//func (json *JsonData) SetCode(code int) Option {
//	return func(data *JsonData) {
//		data.Code = code
//	}
//}
//
////设置data
//func (json *JsonData) SetData(data interface{}) Option {
//	return func(data *JsonData) {
//		data.Data = data
//	}
//}
//
////设置message
//func (json *JsonData) SetMessage(message string) Option {
//	return func(data *JsonData) {
//		data.Message = message
//	}
//}

//func (json *JsonData) New(option ...Option) *JsonData {
//	//先设置默认值
//	data := &JsonData{
//		Code:    200,
//		Data:    nil,
//		Message: "",
//	}
//	//再对设置值的参数进行赋值覆盖
//	for _, fun := range option {
//		fun(data) //把默认的值给调用的函数进行覆盖
//	}
//	return data
//}

func (json *JsonData) New() *JsonData {
	//先设置默认值
	var code = 200
	if json.Code != 0 {
		code = json.Code
	}
	message := "操作成功"
	if json.Message != "" {
		message = json.Message
	}
	data := &JsonData{
		Code:    code,
		Data:    json.Data,
		Message: message,
	}
	return data
}
