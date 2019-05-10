package easygtk

import (
	. "github.com/SilentGopherLnx/easygolang"

	//"github.com/gotk3/gotk3/cairo"
	//"github.com/gotk3/gotk3/gdk"
	//"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	//"github.com/gotk3/gotk3/pango"
)

//var selected_1 = [4]float64{0.4, 0.7, 0.8, 1.0}   // BLUE DARK
var selected_1 = [4]float64{0.0, 0.0, 0.8, 1.0}   // BLUE DARK
var selected_2 = [4]float64{0.85, 0.9, 0.95, 1.0} // BLUE LIGHT

func init() {

}

//https://github.com/surajmandalcell/Gtk-Theming-Guide/blob/master/creating_gtk_themes.md
func GTK_ColorsLoad(win *gtk.Window) {
	style, err := win.GetStyleContext()
	if err == nil {
		gdkcol, ok := style.LookupColor("selected_bg_color")
		if ok {
			fl := gdkcol.Floats()
			if len(fl) >= 3 {
				selected_1[0] = fl[0]
				selected_1[1] = fl[1]
				selected_1[2] = fl[2]
			}
		} else {
			Prln("GTK_ColorsLoad(LookupColor) fail")
		}
	} else {
		Prln("GTK_ColorsLoad: " + err.Error())
	}
}

func GTK_ColorOfSelected() [4]float64 {
	return selected_1
}
