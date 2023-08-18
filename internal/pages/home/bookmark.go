package home

import (
	"html/template"
	"strings"

	FlareData "github.com/soulteary/flare/internal/data"
	FlareIcons "github.com/soulteary/flare/internal/icons"
	FlareModel "github.com/soulteary/flare/internal/model"
	FlareState "github.com/soulteary/flare/internal/state"
)

func GenerateBookmarkTemplate() template.HTML {
	options := FlareData.GetAllSettingsOptions()
	bookmarksData := FlareData.LoadNormalBookmarks()
	tpl := ""

	var bookmarks []FlareModel.Bookmark

	for _, bookmark := range bookmarksData.Items {
		bookmark.URL = FlareState.ParseDynamicUrl(bookmark.URL)
		bookmarks = append(bookmarks, bookmark)
	}

	if len(bookmarksData.Categories) > 0 {
		defaultCategory := bookmarksData.Categories[0]
		for _, category := range bookmarksData.Categories {
			tpl += renderBookmarksWithCategories(&bookmarks, &category, &defaultCategory, options.OpenBookmarkNewTab, options.EnableEncryptedLink, options.IconMode)
		}
	} else {
		tpl += renderBookmarksWithoutCategories(&bookmarks, options.OpenBookmarkNewTab, options.EnableEncryptedLink, options.IconMode)
	}

	return template.HTML(tpl)
}

func renderBookmarksWithoutCategories(bookmarks *[]FlareModel.Bookmark, OpenBookmarkNewTab bool, EnableEncryptedLink bool, IconMode string) string {
	tpl := ""
	for _, bookmark := range *bookmarks {

		// 如果以 chrome-extension:// 协议开头
		// 则使用服务端 Location 方式打开链接
		templateURL := ""
		if strings.HasPrefix(bookmark.URL, "chrome-extension://") {
			templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
		} else {
			if EnableEncryptedLink {
				templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
			} else {
				templateURL = bookmark.URL
			}
		}

		templateIcon := ""
		if strings.HasPrefix(bookmark.Icon, "http://") || strings.HasPrefix(bookmark.Icon, "https://") {
			templateIcon = `<img src="` + bookmark.Icon + `"/>`
		} else if bookmark.Icon != "" {
			templateIcon = FlareIcons.GetIconByName(bookmark.Icon)
		} else {
			if IconMode == "FILLING" {
				templateIcon = FlareState.GetYandexFavicon(bookmark.URL, FlareIcons.GetIconByName(bookmark.Icon))
			} else {
				templateIcon = FlareIcons.GetIconByName(bookmark.Icon)
			}
		}

		if OpenBookmarkNewTab {
			tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
		} else {
			tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
		}
	}
	return `<div class="bookmark-group-container pull-left"><ul class="bookmark-list">` + tpl + `</ul></div>`
}

func renderBookmarksWithCategories(bookmarks *[]FlareModel.Bookmark, category *FlareModel.Category, defaultCategory *FlareModel.Category, OpenBookmarkNewTab bool, EnableEncryptedLink bool, IconMode string) string {
	tpl := ""
	isEmpty := true

	for _, bookmark := range *bookmarks {

		// 如果以 chrome-extension:// 协议开头
		// 则使用服务端 Location 方式打开链接
		templateURL := ""
		if strings.HasPrefix(bookmark.URL, "chrome-extension://") {
			templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
		} else {
			if EnableEncryptedLink {
				templateURL = "/redir/url?go=" + FlareData.Base64EncodeUrl(bookmark.URL)
			} else {
				templateURL = bookmark.URL
			}
		}

		templateIcon := ""
		if strings.HasPrefix(bookmark.Icon, "http://") || strings.HasPrefix(bookmark.Icon, "https://") {
			templateIcon = `<img src="` + bookmark.Icon + `"/>`
		} else if bookmark.Icon != "" {
			templateIcon = FlareIcons.GetIconByName(bookmark.Icon)
		} else {
			if IconMode == "FILLING" {
				templateIcon = FlareState.GetYandexFavicon(bookmark.URL, FlareIcons.GetIconByName(bookmark.Icon))
			} else {
				templateIcon = FlareIcons.GetIconByName(bookmark.Icon)
			}
		}

		if bookmark.Category != "" {
			if bookmark.Category == category.ID {
				if OpenBookmarkNewTab {
					tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				} else {
					tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				}
				isEmpty = false
			}
		} else {
			if category.ID == defaultCategory.ID {
				if OpenBookmarkNewTab {
					tpl += `<li><a target="_blank" rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				} else {
					tpl += `<li><a rel="noopener" href="` + templateURL + `" class="bookmark">` + templateIcon + `<span>` + bookmark.Name + `</span></a></li>`
				}
				isEmpty = false
			}
		}
	}

	if isEmpty {
		return ``
	}

	return `<div class="bookmark-group-container pull-left"><h3 class="bookmark-group-title" data-set-category="` + category.ID + `">` + category.Name + `</h3><ul class="bookmark-list">` + tpl + `</ul></div>`
}
