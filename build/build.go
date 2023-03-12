package main

import (
	builder "github.com/soulteary/flare/build/builder"
)

func main() {
	builder.TaskForMdi(
		"embed/assets/vendor/mdi-cheat-sheets", "internal/icons/mdi-cheat-sheets",
		"embed/assets/vendor/mdi/mdi.js", "internal/resources/mdi/icons.go",
	)
	builder.TaskForSimpleIcons("embed/assets/vendor/simple-icons/index.js", "internal/resources/simpleicons/icons.go")
	builder.TaskForGuideAssets("embed/assets/vendor/guide-assets", "internal/guide/guide-assets")
	builder.TaskForEditorAssets("embed/assets/vendor/editor-assets", "internal/editor/editor-assets")
	builder.TaskForHomeAssets("embed/assets/vendor/home-assets", "internal/pages/home/home-assets")
	builder.TaskForStyles("internal/state/style.go")
	builder.TaskForFavicon("embed/assets/favicon.ico", "internal/resources/assets/favicon.ico")
	builder.TaskForTemplates("embed/templates", "internal/resources/templates/html")
}
