package global

import (
	"github.com/gin-gonic/gin"
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


func (s *Status) ToGin() gin.H {
	return gin.H{
		"status" : s.Status ,
		"message" : s.Message ,
		"data": s.Data ,
	}
}




