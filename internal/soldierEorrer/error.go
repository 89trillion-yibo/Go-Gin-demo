package soldierEorrer

var (
	OK = response(200,"ok")
	Error = response(500,"服务器错误")

	Parameters = response(1001,"有参数为空")
	NoData = response(1002,"数据为空,请检查输入参数")
)

//异常结构
type SoldierErr struct {
	Code    int             //错误码
	Data    interface{}     //返回数据
	Message string          //错误信息
}

//不需要返回数据
func response(code int,message string) *SoldierErr {
	return &SoldierErr{
		Code: code,
		Message: message,
		Data: nil,
	}
}

//需要返回数据
func (sol *SoldierErr) AddData(data interface{}) SoldierErr {
	return SoldierErr{
		Code: sol.Code,
		Message: sol.Message,
		Data: data,
	}
}
