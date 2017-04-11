package main

import (
	. "github.com/siongui/godom"
	regen "github.com/siongui/responsive-embed-generator"
)

func main() {
	i := Document.QuerySelector("#oricode")

	Document.QuerySelector("#submit").AddEventListener("click", func(e Event) {
		html, css, err := regen.GetResponsiveYouTubeCode(i.Value())
		if err != nil {
			return
		}
		println(html)
		println(css)
	})
}
