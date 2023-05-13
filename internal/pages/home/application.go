package home

import (
	"html/template"
	"strings"

	FlareData "github.com/soulteary/flare/internal/data"
	FlareIcons "github.com/soulteary/flare/internal/icons"
	FlareState "github.com/soulteary/flare/internal/state"
)

func GenerateApplicationsTemplate() template.HTML {
	options := FlareData.GetAllSettingsOptions()
	appsData := FlareData.LoadFavoriteBookmarks()
	apps := appsData.Items
	tpl := ""

	for _, app := range apps {

		desc := ""
		if app.Desc == "" {
			desc = app.URL
		} else {
			desc = app.Desc
		}

		// 如果以 chrome-extension:// 协议开头
		// 则使用服务端 Location 方式打开链接
		templateURL := ""
		if strings.HasPrefix(app.URL, "chrome-extension://") {
			templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(app.URL)
		} else {
			if options.EnableEncryptedLink {
				templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(app.URL)
			} else {
				templateURL = app.URL
			}
		}

		templateIcon := ""
		if strings.HasPrefix(app.Icon, "http://") || strings.HasPrefix(app.Icon, "https://") {
			templateIcon = `<img src="` + app.Icon + `"/>`
		} else if app.Icon != "" {
			templateIcon = FlareIcons.GetIconByName(app.Icon)
		} else {
			if options.IconMode == "FILLING" {
				templateIcon = FlareState.GetYandexFavicon(app.URL, FlareIcons.GetIconByName(app.Icon))
			} else {
				templateIcon = FlareIcons.GetIconByName(app.Icon)
			}
		}

		if options.OpenAppNewTab {
			tpl = tpl + `
			<div class="app-container" data-id="` + app.Icon + `">
			<a target="_blank" rel="noopener" href="` + templateURL + `" class="app-item" title="` + app.Name + `">
			  <div class="app-icon">` + templateIcon + `</div>
			  <div class="app-text">
				<p class="app-title">` + app.Name + `</p>
				<p class="app-desc">` + desc + `</p>
			  </div>
			</a>
			</div>
			`
		} else {
			tpl = tpl + `
			<div class="app-container" data-id="` + app.Icon + `">
			<a rel="noopener" href="` + templateURL + `" class="app-item" title="` + app.Name + `">
			  <div class="app-icon">` + templateIcon + `</div>
			  <div class="app-text">
				<p class="app-title">` + app.Name + `</p>
				<p class="app-desc">` + desc + `</p>
			  </div>
			</a>
			</div>
			`
		}
	}
	return template.HTML(tpl)
}
