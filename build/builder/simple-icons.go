package builder

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func TaskForSimpleIcons(res string, gofile string) {
	initSimpleIconsResourceTemplate(res, gofile)
}

func initSimpleIconsResourceTemplate(src string, dest string) {
	// https://www.npmjs.com/package/simple-icons
	file := src
	fileRaw, err := os.ReadFile(filepath.Clean(file))
	siJSON := ""
	if err != nil {
		fmt.Println("读取文件出错", file)
	} else {
		var re = regexp.MustCompile(`[{,](si.+?):{title:.+?},path:"(.+?)",`)

		icons := make(map[string]string)

		for _, match := range re.FindAllStringSubmatch(string(fileRaw), -1) {
			icons[strings.ToLower(match[1])] = match[2]
		}

		file, _ := json.MarshalIndent(icons, "", " ")
		siJSON = string(file)
	}

	goFile := "package simpleicons\nvar IconMap = map[string]string" + siJSON
	goFile = strings.Replace(goFile, "\"\n}", "\",\n}", 1)
	content, _ := format.Source([]byte(goFile))

	err = os.WriteFile(dest, content, os.ModePerm)
	if err != nil {
		fmt.Println("保存文件出错", err)
	} else {
		fmt.Println("保存 Simple Icons 资源文件完毕", dest)
	}
}
