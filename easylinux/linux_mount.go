package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

//name,path
func LinuxGetMountList() [][2]string {
	arr := [][2]string{}

	out, _, _ := ExecCommand("mount")
	out_arr := StringSplitLines(out)
	for _, str := range out_arr[:] {
		if len(str) > 0 {
			ind_on := StringFind(str, " on ")
			ind_type := StringFindEnd(str, " type ")
			if ind_on > -1 && ind_type > -1 {
				diskname := StringPart(str, 1, ind_on-1)
				mountpath := FolderPathEndSlash(StringPart(str, ind_on+4, ind_type-1))
				//Prln("[" + diskname + "||" + mountpath + "]")
				arr = append(arr, [2]string{diskname, mountpath})
			}
		}
	}

	uid, _, _ := GetPC_UserUidLoginName()
	gvfs := "/run/user/" + uid + "/gvfs/"
	files, err := Folder_ListFiles(gvfs, false)
	if err == nil {
		for _, f := range files[:] {
			diskname := f.Name()
			mountpath := gvfs + f.Name() + "/"
			//Prln("[" + diskname + "||" + mountpath + "]")
			arr = append(arr, [2]string{diskname, mountpath})
		}
	}

	return arr
}

func LinuxFolderIsMountPoint(list [][2]string, path_real string) bool {
	path2 := FolderPathEndSlash(path_real)
	for _, pair := range list[:] {
		if FolderPathEndSlash(pair[1]) == path2 {
			return true
		}
	}
	return false
}

func LinuxFilePartition(list [][2]string, path string) (string, string) {
	if path == "" {
		return "", ""
	}
	list2 := make([][2]string, len(list))
	copy(list2, list)
	SortArray(list2, func(i, j int) bool {
		arr_i := StringSplit(list2[i][1], "/")
		arr_j := StringSplit(list2[j][1], "/")
		return len(arr_j) < len(arr_i)
	})
	// Prln("===" + path)
	// for j := 0; j < len(list2); j++ {
	// 	Prln(list2[j][1])
	// }
	// Prln("===")
	for j := 0; j < len(list2); j++ {
		//Prln(I2S(StringFind(path, list2[j][1])))
		if StringFind(path, list2[j][1]) == 1 {
			return list2[j][0], list2[j][1]
		}
	}
	return "", ""
}
