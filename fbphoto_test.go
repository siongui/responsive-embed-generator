package regen

import (
	"testing"
)

const fbphotoiframecode = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1492967057426485%3A0&width=500" width="500" height="518" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`
const fbphotohtml = `<iframe src="https://www.facebook.com/plugins/post.php?href=https%3A%2F%2Fwww.facebook.com%2FDDMCHAN%2Fposts%2F1492967057426485%3A0" width="auto" height="518" style="border:none;overflow:hidden" scrolling="no" frameborder="0" allowTransparency="true"></iframe>`

func TestGetResponsiveFbPhotoCode(t *testing.T) {
	html, err := GetResponsiveFbPhotoCode(fbphotoiframecode)
	if err != nil {
		t.Error(err)
		return
	}
	if html != fbphotohtml {
		t.Error("incorrect Facebook Photo html:")
		t.Error(html)
		return
	}
}
