package builder

import "fmt"

func TaskForFavicon() {
	_Copy("embed/assets/favicon.ico", "internal/assets/favicon.ico")
	fmt.Println("复制静态资源 ... [OK]")
}
