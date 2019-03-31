package easygtk

import (
	. "github.com/SilentGopherLnx/easygolang"

	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/gotk3/gotk3/pango"
	//"github.com/gotk3/gotk3/gdk"
)

type GTK_RemoveAble interface {
	GetChildren() *glib.List
	Remove(gtk.IWidget)
}

type GTK_DestroyAble interface {
	Destroy()
}

type GTK_RemoveDestroyAble interface {
	GetChildren() *glib.List
	Remove(gtk.IWidget)
	Destroy()
}

// for l := container.GetChildren(); l != nil; l = l.Next() {
// 	if l.Data().(*gtk.Widget).GObject == myWidget.GObject {
// 		fmt.Println("found")
// 		break
// 	}
// }

func GTK_Childs(w GTK_RemoveDestroyAble, remove_all bool, destroy_all bool) []gtk.IWidget {
	chl := w.GetChildren()
	arr := []gtk.IWidget{}
	chl.Foreach(func(item interface{}) {
		switch el := item.(type) {
		case gtk.IWidget:
			if remove_all {
				w.Remove(el)
			}
			if destroy_all {
				var elem1 GTK_RemoveDestroyAble
				var ok1 bool
				elem1, ok1 = item.(GTK_RemoveDestroyAble)
				if ok1 {
					GTK_Childs(elem1, true, true)
					//Prln("Destroy-ed 1")
				} else {
					//Prln("ERROR OF <Destroy> 1")
				}
				var elem2 GTK_DestroyAble
				var ok2 bool
				elem2, ok2 = item.(GTK_DestroyAble)
				if ok2 {
					elem2.Destroy()
					//Prln("Destroy-ed 2")
				} else {
					//Prln("ERROR OF <Destroy> 2")
				}
			} else {
				arr = append(arr, el)
			}
		default:
			Prln("GTK_Childs ERROR:" + TypeOf(item))
		}
	})
	return arr
}

func GTK_WidgetExist(w interface {
	IsVisible() bool
	GetParent() (*gtk.Widget, error)
}) bool {
	if w != nil {
		if w.IsVisible() {
			par, err := w.GetParent()
			if err == nil && par != nil {
				return true
			}
		}
	}
	return false
}

// ==============

func GTK_MenuSub(rightmenu interface{ Add(gtk.IWidget) }, title string) *gtk.Menu {
	item, _ := gtk.MenuItemNewWithLabel(title)
	submenu, _ := gtk.MenuNew()
	item.SetSubmenu(submenu)
	rightmenu.Add(item)
	return submenu
}

func GTK_MenuItem(rightmenu *gtk.Menu, title string, func_event func()) *gtk.MenuItem {
	/*item, _ := gtk.MenuItemNewWithLabel(title)

	rightmenu.Add(item)
	if func_event != nil {
		item.Connect("button-press-event", func_event)
	} else {
		item.SetSensitive(false)
	}
	return item*/
	return GTK_MenuItemIcon(rightmenu, title, "", func_event)
}

func GTK_MenuItemIcon(rightmenu *gtk.Menu, title string, iconname string, func_event func()) *gtk.MenuItem {
	item, _ := gtk.MenuItemNewWithLabel(title)
	/*item, _ := gtk.MenuItemNew()
	lbl, _ := gtk.LabelNew(title)
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	//box.SetMarginStart(-16)
	//box.SetHAlign(gtk.ALIGN_START)
	icon, _ := gtk.ImageNew()
	icon.SetSizeRequest(16, 16)
	//icon.SetFromPixbuf(pixbuf)
	icon.SetFromIconName("edit-delete", gtk.ICON_SIZE_MENU)
	box.Add(icon)
	box.Add(lbl)
	//btn, _ := gtk.ButtonNewWithLabel("lbl")
	item.Add(box)
	//item.Add(btn)
	//item.SetProperty("always-show-image", false)
	//box.SetHExpand(true)
	//item.SetHAlign(gtk.ALIGN_FILL)*/

	// icon, _ := gtk.ImageNew()
	// icon.SetFromIconName("edit-delete", gtk.ICON_SIZE_MENU)
	// pixbuf2 := icon.GetPixbuf()
	// pixbuf2.

	if len(iconname) > 0 {
		b, _ := FileBytesRead(FolderLocation_App() + "gui/button_" + iconname + ".png")
		img := ImageDecode(b)
		if !InterfaceNil(img) {
			w := img.Bounds().Max.X
			h := img.Bounds().Max.Y
			//Prln(I2S(w) + ":" + I2S(h))
			item.Connect("draw", func(g *gtk.MenuItem, ctx *cairo.Context) {
				dx := 5
				dy := (g.GetAllocatedHeight() - h) / 2
				for y := 0; y < h; y++ {
					for x := 0; x < w; x++ {
						r, g, b, a := img.At(x, y).RGBA()
						ctx.SetSourceRGBA(float64(r/MAX_A)/255.0, float64(g/MAX_A)/255.0, float64(b/MAX_A)/255.0, float64(a/MAX_A)/255.0)
						// if x == 0 || y == 0 || x == w-1 || y == h-1 {
						// 	ctx.SetSourceRGBA(0, 0, 0, 0.1) // for test
						// }
						ctx.Rectangle(float64(x+dx), float64(y+dy), float64(1), float64(1))
						ctx.Fill()
					}
				}
			})
		}
	}

	rightmenu.Add(item)
	if func_event != nil {
		item.Connect("button-press-event", func_event)
	} else {
		item.SetSensitive(false)
	}
	return item
}

func GTK_MenuSeparator(rightmenu *gtk.Menu) {
	rightmenu.Add(A0(gtk.SeparatorMenuItemNew()).(gtk.IWidget))
}

// ===========

func GTK_LabelPair(title string, value string) (*gtk.Box, *gtk.Label) {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	lbl_name, _ := gtk.LabelNew(title)
	lbl_name.SetVAlign(gtk.ALIGN_START)
	lbl_name.SetMarkup("<b>" + HtmlEscape(title) + "</b>")
	lbl_value, _ := gtk.LabelNew(value)
	box.Add(lbl_name)
	box.Add(lbl_value)
	return box, lbl_value
}

func GTK_LabelWrapMode(label *gtk.Label, lines int) {
	label.SetLineWrap(true)
	label.SetLineWrapMode(pango.WRAP_CHAR)
	label.SetEllipsize(pango.ELLIPSIZE_MIDDLE)
	label.SetLines(lines)
}

// ===========

func GTK_Dialog(w *gtk.Window, title string) (*gtk.Dialog, *gtk.Box) {
	dial, _ := gtk.DialogNew()
	dial.SetTransientFor(w)
	dial.SetTitle(title)
	box, _ := dial.GetContentArea()
	return dial, box
}
