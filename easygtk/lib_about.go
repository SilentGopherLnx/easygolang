package easygtk

import (
	. "github.com/SilentGopherLnx/easygolang"

	"image"
	"image/color"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/cairo"
	//"github.com/gotk3/gotk3/glib"
	//"github.com/gotk3/gotk3/pango"
)

func Dialog_About(w *gtk.Window, version, author, mail, repository string, flag *gdk.Pixbuf) {
	dial, box := GTK_Dialog(w, "About")

	lbl_version, _ := gtk.LabelNew("version: " + version)
	lbl_author, _ := gtk.LabelNew("by: " + author)
	lbl_email, _ := gtk.LabelNew(mail)
	lbl_github, _ := gtk.LabelNew(repository)

	iconflag, _ := gtk.ImageNew()
	iconflag.SetFromPixbuf(flag)

	box.SetOrientation(gtk.ORIENTATION_VERTICAL)
	box.Add(iconflag)
	box.Add(lbl_version)
	box.Add(lbl_author)
	box.Add(lbl_email)
	box.Add(lbl_github)

	dial.SetResizable(false)
	dial.ShowAll()
	dial.Run()
	dial.Close()
}

//best country, of course
func GetFlag_Russian() *gdk.Pixbuf {
	img1 := image.NewRGBA(image.Rect(0, 0, 1, 3))
	img1.Set(0, 0, color.RGBA{255, 255, 255, 255})
	img1.Set(0, 1, color.RGBA{0, 57, 166, 255})
	img1.Set(0, 2, color.RGBA{213, 43, 30, 255})
	img2 := ImageResizeNearest(img1, 96, 48)
	return GTK_PixBuf_From_RGBA(img2)
}
