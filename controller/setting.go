package controller

import (
	"goflylivechat/models"

	"github.com/gin-gonic/gin"
)

func GetConfigs(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	configs := models.FindConfigsByUserId(kefuName)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": configs,
	})
}
func GetConfig(c *gin.Context) {
	key := c.Query("key")
	config := models.FindConfig(key)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": config,
	})
}
func PostConfig(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")
	kefuName, _ := c.Get("kefu_name")
	if key == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "配置项 key 不能为空",
		})
		return
	}
	models.UpdateConfig(kefuName, key, value)

	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}

// 活码跳转：/go/:name → 读 configs 表中 redirect_url，302 跳转
func GetRedirect(c *gin.Context) {
	name := c.Param("name")
	config := models.FindConfigByUserId(name, "redirect_url")
	if config.ID == 0 || config.ConfValue == "" {
		c.String(404, "未配置跳转地址")
		return
	}
	c.Redirect(302, config.ConfValue)
}

// 读取跳转 URL
func GetRedirectUrl(c *gin.Context) {
	kefuName, _ := c.Get("kefu_name")
	config := models.FindConfigByUserId(kefuName, "redirect_url")
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": config.ConfValue,
	})
}

// 保存跳转 URL
func PostRedirectUrl(c *gin.Context) {
	url := c.PostForm("url")
	kefuName, _ := c.Get("kefu_name")
	if url == "" {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "链接不能为空",
		})
		return
	}
	models.UpdateConfig(kefuName.(string), "redirect_url", url)
	c.JSON(200, gin.H{
		"code":   200,
		"msg":    "ok",
		"result": "",
	})
}
