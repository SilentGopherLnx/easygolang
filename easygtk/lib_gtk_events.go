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

func GTK_TranslateKeyLayoutEnglish(key uint, state uint) (uint, uint) {
	key2 := key
	state2 := state
	if state2 == 8196 { //RUSSIAN Ctrl
		state2 = 4 //English Ctrl
	}
	switch key {
	case gdk.KEY_Cyrillic_ef: //RUSSIAN 'ф'
		key2 = gdk.KEY_a
	case gdk.KEY_Cyrillic_che: //RUSSIAN 'ч'
		key2 = gdk.KEY_x
	case gdk.KEY_Cyrillic_es: //RUSSIAN 'с'
		key2 = gdk.KEY_c
	case gdk.KEY_Cyrillic_em: //RUSSIAN 'м'
		key2 = gdk.KEY_v
		//etc
	}
	//Prln("key LOCALE : " + I2S(int(key)) + ", state=" + I2S(int(state)))
	//Prln("key ENGLISH: " + I2S(int(key2)) + ", state=" + I2S(int(state2)))
	return key2, state2
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
