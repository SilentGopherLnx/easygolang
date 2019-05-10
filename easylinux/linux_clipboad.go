package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

// xclip -selection clipboard -t TARGETS -o

// find -maxdepth 1 -name "*.png"

/*func LinuxClipboard_GetText() {
	cmd := "xclip -o"
	Prln(cmd)
}

func LinuxClipboard_SetText(txt string) {
	cmd := "xclip -i"
	//Ctrl+shift+D
	Prln(cmd)
}*/

func LinuxClipBoard_CopyFiles(files []*LinuxPath, cut_mode bool) {
	// echo -e "copy\nfile:///mnt/dm-1/golang/my_code/screenshot.png\0"| xclip -i -selection clipboard -t x-special/gnome-copied-files
	if len(files) > 0 {
		input := "copy"
		if cut_mode {
			input = "cut"
		}
		for j := 0; j < len(files); j++ {
			input = input + "\n" + files[j].GetUrl()
		}
		ExecCommandBytes([]byte(input+"\000"), 1000, "xclip", "-i", "-selection", "clipboard", "-t", "x-special/gnome-copied-files")
	} else {
		Prln("COPY/CUT LIST EMPTY")
	}
}

func LinuxClipBoard_PasteFiles() string {
	//xclip -o -selection clipboard -t "x-special/gnome-copied-files"
	res, _, _ := ExecCommand("xclip", "-o", "-selection", "clipboard", "-t", "x-special/gnome-copied-files")
	//Prln(res)
	res = StringReplace(res, "\r", "")
	res = StringReplace(res, "\000", "")
	return res
}
