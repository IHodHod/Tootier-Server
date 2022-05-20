package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type Status struct {
	Status string
	Message string
	Data interface{}
	Code int
}

func CreateStatus() Status {
	return Status{
		Status: "success" ,
		Message: GetLang().MSG_OK ,
		Code: 200 ,
	}
}

func (s *Status) Set(status string , message string , code int , data interface{})  {
	s.Status = status
	s.Code = code
	s.Message = message
	s.Data = data
}

func (s *Status) ToGin() gin.H {
	return gin.H{
		"status" : s.Status ,
		"message" : s.Message ,
		"data": s.Data ,
	}
}



func CurrentTimeUnix() string {
	return fmt.Sprintf("%v" , time.Now().Unix())
}

func STR(b interface{}) string {
	return fmt.Sprintf("%v" , b)
}
