package main

import (
	"CodeView"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("CodeView", func(w *unison.Window) {
		Codeview.New().Layout(w.Content())
	})
}
