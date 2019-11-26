package wechat

import (
	"fmt"
	"sort"

	"github.com/adolfheir/go-wechat/consts"
	"github.com/adolfheir/go-wechat/setting"
	"github.com/adolfheir/go-wechat/util"
	"github.com/gin-gonic/gin"
)

// @Summary 微信校验
// @produce text/xml
// @Param echostr query string true "echostr"
// @Param signature query string true "signature"
// @Param timestamp query string true "timestamp"
// @Param nonce query string true "nonce"
// @Success 200 {string} string "return the querry echostr"
// @failure 400 {string} string "return the errormsg"
// @Router /wx/check [get]
func WxCheck(c *gin.Context) {
	echostr := c.DefaultQuery("echostr", "")
	signature := c.DefaultQuery("signature", "")
	timestamp := c.DefaultQuery("timestamp", "")
	nonce := c.DefaultQuery("nonce", "")
	token := setting.AppSetting.WxToken

	thisSignature := getSignature(token, timestamp, nonce)
	fmt.Println("log", signature, thisSignature)

	if thisSignature != signature {
		c.String(consts.ERROR, consts.GetMsg(consts.ERROR))
	} else {
		c.String(consts.SUCCESS, echostr)
	}
}

func getSignature(token, timestamp, nonce string) string {
	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	signature := util.SortSha1(str)
	return signature
}
