package userHandlers

import (
	"log"
	
	"github.com/gin-gonic/gin"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/common/r"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/global"
	"github.com/rustuo-m/cicoyi/cicoyi_user/internal/models"
)

func RegisterUser(c *gin.Context) {
	var req models.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println("bind json error")
		r.Failed("bind json error", c)
		return
	}
	_, codeExist := req.ExistRegister()
	if codeExist == global.NotFind {
		// 密码加密
		code := req.Create()
		if code == global.SUCCESS {
			r.OKWithData(req.Name, "create user success", c)
			return
		} else {
			r.Failed("create user error", c)
		}
	} else {
		r.Failed("user exists", c)
	}
}
