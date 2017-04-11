package regen

import (
	"testing"
)

const youtubeiframecode = `<iframe width="560" height="315" src="https://www.youtube.com/embed/YpWFR-ioQlE" frameborder="0" allowfullscreen></iframe>`
const youtubehtml = `<div class="youtube-responsive-560x315">` + youtubeiframecode + `</div>`
const youtubecss = `
@media (max-width: 560px) {
  .youtube-responsive-560x315 {
    position: relative;
    padding-bottom: 56.25%; /* aspect ratio (height/width) of iframe */
    padding-top: 30px;
    height: 0;
    overflow: hidden;
  }

  .youtube-responsive-560x315 iframe {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
  }
}
`

func TestGetResponsiveYouTubeCode(t *testing.T) {
	html, css, err := GetResponsiveYouTubeCode(youtubeiframecode)
	if err != nil {
		t.Error(err)
		return
	}
	if html != youtubehtml {
		t.Error("incorrect youtube html:")
		t.Error(html)
		return
	}
	if css != youtubecss {
		t.Error("incorrect youtube css:")
		t.Error(css)
		return
	}
}
