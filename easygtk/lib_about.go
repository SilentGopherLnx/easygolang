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

func Dialog_About(w *gtk.Window, version, author, mail, repository, more string, flag *gdk.Pixbuf) {
	dial, box := GTK_Dialog(w, "About is App")

	lbl_app_version, _ := gtk.LabelNew("App version: " + version)
	lbl_author, _ := gtk.LabelNew("by: " + author)
	lbl_email, _ := gtk.LabelNew("contact: " + mail)
	lbl_github, _ := gtk.LabelNew(repository)

	iconflag, _ := gtk.ImageNew()
	iconflag.SetFromPixbuf(flag)

	box.SetOrientation(gtk.ORIENTATION_VERTICAL)
	box.Add(iconflag)
	box.Add(lbl_app_version)
	box.Add(lbl_author)
	box.Add(lbl_email)
	box.Add(lbl_github)

	lbl_go_version, _ := gtk.LabelNew("Build by compiler: " + GetGolangVersion())
	lbl_gtk_version, _ := gtk.LabelNew("GTK version: " + GTK_GetVersion())
	lbl_gtk_version_wrapper, _ := gtk.LabelNew("gotk3 (GTK wrapper) version: " + GTK_GetVersionWrapper())

	box.Add(lbl_go_version)
	box.Add(lbl_gtk_version)
	box.Add(lbl_gtk_version_wrapper)

	if StringLength(more) > 0 {
		lbl_more, _ := gtk.LabelNew("\n" + more)
		box.Add(lbl_more)
	}

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
