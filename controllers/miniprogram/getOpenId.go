package miniprogram

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOpenIdHandle(c *gin.Context) {
	code := c.Query("code")
	var params string = "https://api.weixin.qq.com/sns/jscode2session?appid=wxc4a7ea41d4eff1ae&secret=29e8c7a2a8e6b078b8c657004aa5cd86&js_code=" + code + "&grant_type=authorization_code"
	res, _ := http.Get(params)
	body, _ := ioutil.ReadAll(res.Body)
	openid := string(body)
	c.String(http.StatusOK, openid)
}
