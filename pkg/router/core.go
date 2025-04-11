package router

import (
	"encoding/json"
	"fmt"
	"github.com/qiniu/x/log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var resErrors = map[int]ResError{}

// Context ceres Context of HTTP
type Context struct {
	*gin.Context
}

// HandlerFunc ceres handler function
type HandlerFunc func(c *Context)

// Wrap the gin router function with the ceres core context
func Wrap(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &Context{c}
		h(ctx)
	}
}

// OK return with 200
func (ctx *Context) OK(data interface{}) {
	ctx.JSON(
		http.StatusOK,
		data,
	)
}

type ResError struct {
	Status  int         `json:"-"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r ResError) Error() string {
	if r.Code == 0 {
		return ""
	}
	b, _ := json.Marshal(r)
	return string(b)
}

func (r ResError) WithMsg(message string) ResError {
	r.Message = message
	return r
}

func (r ResError) WithMsgf(format, message string) ResError {
	r.Message = fmt.Sprintf(format, message)
	return r
}

func NewResError(status, code int, message string, datas ...interface{}) ResError {
	if _, ok := resErrors[code]; ok {
		panic(fmt.Sprintf("apierror code:%d, message:`%s` has been taken", code, message))
	}
	e := ResError{
		status, code, message, datas,
	}
	resErrors[code] = e
	return e
}

var (
	ErrBadRequest     = NewResError(http.StatusBadRequest, 400, "bad request params")
	ErrUnauthorized   = NewResError(http.StatusUnauthorized, 401, "need login")
	ErrForbidden      = NewResError(http.StatusForbidden, 403, "forbidden")
	ErrNotFound       = NewResError(http.StatusNotFound, 404, "not found")
	ErrInternalServer = NewResError(http.StatusInternalServerError, 500, "internal server error")
)

func (ctx *Context) HandleError(err error) {
	log.Infof("#### [HandleError] : %v\n", err)
	switch err.(type) {
	case ResError:
		ctx.Context.JSON(
			err.(ResError).Status,
			err,
		)
	case validator.ValidationErrors:
		ctx.Context.JSON(
			http.StatusBadRequest,
			&err,
		)
	default:
		ctx.Context.JSON(
			http.StatusInternalServerError,
			&ErrInternalServer,
		)
	}
}
