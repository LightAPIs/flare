package redir

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"

	FlareData "github.com/soulteary/flare/data"
	FlareState "github.com/soulteary/flare/state"
)

// TODO 错误提示统一处理
// TODO 针对已跳转过的地址添加内存缓存，减少被恶意利用的可能

func RegisterRouting(router *gin.Engine) {

	internalError := []byte(`<html><p>找不到匹配的跳转地址，请确认地址未被人为修改。</p><p>或前往 <a href="https://github.com/soulteary/docker-flare/issues/" target="_blank">https://github.com/soulteary/docker-flare/issues/</a> 反馈使用中的问题，谢谢！</html>`)

	router.GET(FlareState.MiscPages.RedirHome.Path, func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, FlareState.RegularPages.Home.Path)
		c.Abort()
	})

	router.GET(FlareState.MiscPages.RedirHelper.Path, func(c *gin.Context) {
		encoded := c.Param("url")
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			c.Data(http.StatusBadRequest, "text/html; charset=utf-8", internalError)
			c.Abort()
			return
		}
		decodeURL := string(decoded)

		appsData := FlareData.LoadFavoriteBookmarks()
		for _, bookmark := range appsData.Items {
			if bookmark.URL == decodeURL {
				c.Redirect(http.StatusTemporaryRedirect, string(decoded))
				c.Abort()
				return
			}
		}

		bookmarksData := FlareData.LoadNormalBookmarks()
		for _, bookmark := range bookmarksData.Items {
			if bookmark.URL == decodeURL {
				c.Redirect(http.StatusTemporaryRedirect, string(decoded))
				c.Abort()
				return
			}
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", internalError)
		c.Abort()
	})
}
