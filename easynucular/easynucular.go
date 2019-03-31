package easynukular

import (
	. "github.com/SilentGopherLnx/easygolang"

	ui "github.com/aarzilli/nucular"
)

func NUCULAR_MakeTextEditor(multiline bool, maxlen int, text string) ui.TextEditor {
	var te ui.TextEditor
	if multiline {
		te.Flags = ui.EditBox
	} else {
		te.Flags = ui.EditSimple
	}
	te.Maxlen = maxlen
	te.Paste(text)
	return te
}

func NUCULAR_GetText(te ui.TextEditor, trim bool) string {
	if trim {
		return StringTrim(string(te.Buffer))
	} else {
		return string(te.Buffer)
	}
}
