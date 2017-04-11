package main

import (
	. "github.com/siongui/godom"
	regen "github.com/siongui/responsive-embed-generator"
)

func main() {
	i := Document.QuerySelector("#oricode")
	hc := Document.QuerySelector("#htmlcode")
	cc := Document.QuerySelector("#csscode")

	Document.QuerySelector("#submit").AddEventListener("click", func(e Event) {
		html, css, err := regen.GetResponsiveYouTubeCode(i.Value())
		if err != nil {
			hc.SetValue("Fail to generate code")
			return
		}
		hc.SetValue(html)
		cc.SetValue(css)
	})

	Document.QuerySelector("#copyhtml").AddEventListener("click", func(e Event) {
		hc.Call("select")
		isSuccessful := Document.Call("execCommand", "copy").Bool()
		if isSuccessful {
			hc.SetValue("Succeed to copy HTML code")
		} else {
			hc.SetValue("Fail to copy HTML code")
		}
	})

	Document.QuerySelector("#copycss").AddEventListener("click", func(e Event) {
		cc.Call("select")
		isSuccessful := Document.Call("execCommand", "copy").Bool()
		if isSuccessful {
			cc.SetValue("Succeed to copy CSS code")
		} else {
			cc.SetValue("Fail to copy CSS code")
		}
	})
}
