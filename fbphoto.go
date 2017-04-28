package regen

import (
	"strings"
)

func GetResponsiveFbPhotoCode(iframecode string) (html string, err error) {
	html = strings.Replace(strings.Replace(iframecode, "&width=500", "", 1), `width="500"`, `width="auto"`, 1)
	return
}
