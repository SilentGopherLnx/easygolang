package easygtk

import (
	//	. "github.com/SilentGopherLnx/easygolang"

	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/gdk"
)

func GTK_Dialog(w *gtk.Window, title string) (*gtk.Dialog, *gtk.Box) {
	dial, _ := gtk.DialogNew()
	if w != nil {
		dial.SetTransientFor(w)
	}
	dial.SetTitle(title)
	box, _ := dial.GetContentArea()
	return dial, box
}

func GTK_DialogMessage(w *gtk.Window, msg_type gtk.MessageType, btns_type gtk.ButtonsType, title string, msg string) (*gtk.MessageDialog, *gtk.Box) {
	dial := gtk.MessageDialogNew(w, gtk.DIALOG_MODAL, msg_type, btns_type, msg)
	dial.SetTitle(title)
	box, _ := dial.GetContentArea()
	return dial, box
}
