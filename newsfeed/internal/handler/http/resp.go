package http

type RespCode int32

const (
	CodeSuccess        RespCode = 0
	CodeInvalidRequest RespCode = 1
	CodeGrpcCall       RespCode = 2
)

type CommonResponse struct {
	Code    RespCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
