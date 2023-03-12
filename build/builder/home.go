package builder

import "log"

func TaskForHomeAssets(src string, dest string) {
	_PrepareDirectory(dest)
	if err := _CopyDirectoryWithoutSymlink(src, dest); err != nil {
		log.Fatal(err)
	}
}
