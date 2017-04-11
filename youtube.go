package regen

import (
	"bytes"
	"errors"
	"regexp"
	"strconv"
	"text/template"
)

var ptn = `<iframe width="(?P<w>[0-9]+)" height="(?P<h>[0-9]+)" .*></iframe>`
var youtubeClassName = `youtube-responsive-{{.Width}}x{{.Height}}`
var youtubeHtmlTmpl = `<div class="{{.ClassName}}">{{.Code}}</div>`
var youtubeCssTmpl = `
@media (max-width: {{.Width}}px) {
  .{{.ClassName}} {
    position: relative;
    padding-bottom: {{.AspectRatio}}; /* aspect ratio (height/width) of iframe */
    padding-top: 30px;
    height: 0;
    overflow: hidden;
  }

  .{{.ClassName}} iframe {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
  }
}
`

type YouTubeIframeInfo struct {
	Width       string
	Height      string
	ClassName   string
	Code        string
	AspectRatio string
	ar100       float64
	w           float64
	h           float64
}

func TmplToString(tmpl string, data interface{}) (s string, err error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return
	}

	var b bytes.Buffer
	err = t.Execute(&b, &data)
	if err != nil {
		return
	}
	s = b.String()
	return
}

func GetResponsiveYouTubeCode(iframecode string) (html, css string, err error) {
	yii := YouTubeIframeInfo{Code: iframecode}

	re := regexp.MustCompile(ptn)
	matches := re.FindStringSubmatch(iframecode)
	names := re.SubexpNames()
	for i, match := range matches {
		if names[i] == "w" {
			yii.Width = match
		}
		if names[i] == "h" {
			yii.Height = match
		}
	}
	if yii.Width == "" {
		err = errors.New("Cannot get width of iframe")
		return
	}
	if yii.Height == "" {
		err = errors.New("Cannot get height of iframe")
		return
	}

	yii.w, err = strconv.ParseFloat(yii.Width, 64)
	if err != nil {
		return
	}
	yii.h, err = strconv.ParseFloat(yii.Height, 64)
	if err != nil {
		return
	}
	yii.ar100 = yii.h * 100 / yii.w
	yii.AspectRatio = strconv.FormatFloat(yii.ar100, 'f', -1, 64) + "%"

	yii.ClassName, err = TmplToString(youtubeClassName, yii)
	if err != nil {
		return
	}

	html, err = TmplToString(youtubeHtmlTmpl, yii)
	if err != nil {
		return
	}

	css, err = TmplToString(youtubeCssTmpl, yii)
	if err != nil {
		return
	}

	return
}
