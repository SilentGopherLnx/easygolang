package easygtk

import (
	//. "github.com/SilentGopherLnx/easygolang"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func GTK_MouseKeyOfEvent(event *gdk.Event) (int, int, int) {
	if event != nil {
		eventObject := &gdk.EventButton{event}
		key := 0
		btn := eventObject.ButtonVal()
		// switch btn {
		// case gdk.KEY_leftpointer:
		// 	key = 1
		// case gdk.KEY_rightpointer:
		// 	key = 3
		// default:
		key = int(btn)
		//}
		return key, int(eventObject.X()), int(eventObject.Y())
	}
	return 0, 0, 0
}

func GTK_KeyboardKeyOfEvent(event *gdk.Event) (uint, uint) {
	if event != nil {
		eventObject := &gdk.EventKey{event}
		key := eventObject.KeyVal()
		state := eventObject.State()
		//Prln("key:" + I2S(int(key)))
		return key, state
	}
	return 0, 0
}

func GTK_ScrollGetValues(scroll *gtk.ScrolledWindow) (int, int) {
	if scroll != nil {
		dx := int(scroll.GetHAdjustment().GetValue())
		dy := int(scroll.GetVAdjustment().GetValue())
		return dx, dy
	}
	return 0, 0
}

func GTK_ScrollReset(scroll *gtk.ScrolledWindow) {
	if scroll != nil {
		scroll.GetHAdjustment().SetValue(0)
		scroll.GetVAdjustment().SetValue(0)
	}
}
