package easygtk

import (
	//	. "github.com/SilentGopherLnx/easygolang"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

//var table_keys map[uint]uint

func init() {
	/*table_keys = make(map[uint]uint)
	table_keys[gdk.KEY_Cyrillic_ya] = gdk.KEY_z  //RUSSIAN 'я'
	table_keys[gdk.KEY_Cyrillic_ef] = gdk.KEY_a  //RUSSIAN 'ф'
	table_keys[gdk.KEY_Cyrillic_che] = gdk.KEY_x //RUSSIAN 'ч'
	table_keys[gdk.KEY_Cyrillic_es] = gdk.KEY_c  //RUSSIAN 'с'
	table_keys[gdk.KEY_Cyrillic_em] = gdk.KEY_v  //RUSSIAN 'м'*/
}

func GTK_MouseKeyOfEvent(event *gdk.Event) (int, int, int, uint) {
	if event != nil {
		eventObject := &gdk.EventButton{event}
		//eventmotion:=gdk.EventMotionNewFromEvent(event)
		key := 0
		btn := eventObject.Button()
		state := eventObject.State()
		// switch btn {
		// case gdk.KEY_leftpointer:
		// 	key = 1
		// case gdk.KEY_rightpointer:
		// 	key = 3
		// default:
		key = int(btn)
		//}
		return key, int(eventObject.X()), int(eventObject.Y()), state
	}
	return 0, 0, 0, 0
}

func GTK_KeyboardKeyOfEvent(event *gdk.Event) (uint, uint, uint16) {
	if event != nil {
		eventObject := &gdk.EventKey{event}
		key := eventObject.KeyVal()
		en_key := eventObject.HardwareKeyCode()
		state := eventObject.State()
		//Prln("key:" + I2S(int(key)))
		return key, state, en_key
	}
	return 0, 0, 0
}

/*func GTK_KeyboardTranslateLayoutEnglish(key uint, state uint) (uint, uint) {
	key2 := key
	state2 := state
	if state2 > 8192 { //RUSSIAN Ctrl 8196 == English Ctrl 4
		state2 -= 8192
	}
	key3, ok := table_keys[key]
	if ok {
		key2 = key3
	}
	// Prln(">" + I2S(int(state&gdk.GDK_CONTROL_MASK)))
	// Prln(">" + I2S(int(gdk.GDK_CONTROL_MASK)))
	Prln("pressed: [" + string(gdk.KeyvalToUnicode(key)) + "]")
	Prln("key LOCALE : " + I2S(int(key)) + ", state=" + I2S(int(state)))
	Prln("key ENGLISH: " + I2S(int(key2)) + ", state=" + I2S(int(state2)))
	return key2, state2
}*/

func GTK_KeyboardCtrlState(state uint) bool {
	return state&gdk.GDK_CONTROL_MASK == gdk.GDK_CONTROL_MASK
}

// func GTK_KeyboardCtrl(event *gdk.Event) bool {
// 	//_, state := GTK_KeyboardKeyOfEvent(event)//replace to mouse!
// 	return GTK_KeyboardCtrlState(state)
// }

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
