package home

import (
	"html/template"

	FlareData "github.com/soulteary/flare/internal/data"
	FlareModel "github.com/soulteary/flare/internal/model"
)

func GenerateSidebarTemplate() template.HTML {
	bookmarksData := FlareData.LoadNormalBookmarks()
	categories := bookmarksData.Categories
	tpl := ""

	if len(categories) > 0 {
		tpl += renderSidebar(&categories)
	}

	return template.HTML(tpl)
}

func renderSidebar(categories *[]FlareModel.Category) string {
	tpl := ""
	isEmpty := true

	for _, category := range *categories {
		if category.Name != "" {
			tpl += `<li><a href="#` + category.ID + `">` + category.Name + `</a></li>`
			isEmpty = false
		}
	}

	if isEmpty {
		return ""
	}

	return `<div class="category-group-container"><div class="category-group-header"><i><svg viewBox="0 3 24 24" width="26"><path d="M17,3H7A2,2 0 0,0 5,5V21L12,18L19,21V5C19,3.89 18.1,3 17,3Z"></path></svg></i><span>书签分类</span></div><div class="catetory-group-content"><ul class="category-list">` + tpl + `</ul></div></div><div class="category-button-container"><div class="category-btn-bg category-btn-list"><span id="btn-category-list" alt="Category List"><span>Category</span><svg viewBox="0 0 24 24" width="24"><path d="M12,20A8,8 0 0,1 4,12A8,8 0 0,1 12,4A8,8 0 0,1 20,12A8,8 0 0,1 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M12,12.5A1.5,1.5 0 0,1 10.5,11A1.5,1.5 0 0,1 12,9.5A1.5,1.5 0 0,1 13.5,11A1.5,1.5 0 0,1 12,12.5M12,7.2C9.9,7.2 8.2,8.9 8.2,11C8.2,14 12,17.5 12,17.5C12,17.5 15.8,14 15.8,11C15.8,8.9 14.1,7.2 12,7.2Z"></path></svg></span></div></div>`
}
