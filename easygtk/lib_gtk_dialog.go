package easygtk

import (
	//	. "github.com/SilentGopherLnx/easygolang"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gdk"
)

func GTK_Dialog(w *gtk.Window, title string) (*gtk.Dialog, *gtk.Box) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(w)
	dial.SetTitle(title)
	box, _ := dial.GetContentArea()
	return dial, box
}
