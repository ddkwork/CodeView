package main

import (
	"CodeView"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("CodeView", func(w *unison.Window) {
		w.Content().AddChild(Codeview.New().Layout())
	})
}
