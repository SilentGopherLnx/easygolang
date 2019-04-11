package easygolang

import (
	"os"
)

type IFolderWalker interface {
	WithFile(f os.FileInfo, regular bool, path_src string, path_dst string)
	WithFolderBefore(f os.FileInfo, is_mount bool, path_src string, path_dst string)
	WithFolderAfter(f os.FileInfo, is_mount bool, list_err bool, path_src string, path_dst string)
	WithLink(f os.FileInfo, is_folder bool, path_src string, path_dst string) bool
}

func FoldersRecursively_Walk(mount_list [][2]string, file_or_dir os.FileInfo, path_src_real string, path_dst_real string, method IFolderWalker, deep int) {
	name_src := file_or_dir.Name()
	path_dst := path_dst_real + name_src
	if !FileIsLink(file_or_dir) {
		if file_or_dir.Mode().IsDir() {
			path_src := FolderPathEndSlash(path_src_real)
			path_dst += GetOS_Slash()
			is_mount := false
			for _, pair := range mount_list[:] {
				if FolderPathEndSlash(pair[1]) == path_src {
					is_mount = true
				}
			}
			deep2 := MAXI(1, deep)
			if is_mount {
				deep2 += 1
			}
			method.WithFolderBefore(file_or_dir, is_mount, path_src, path_dst)
			if deep == 0 || !is_mount {
				sub_files, err := Folder_ListFiles(path_src)
				folder_err := false
				if err == nil {
					for j := 0; j < len(sub_files); j++ {
						FoldersRecursively_Walk(mount_list, sub_files[j], path_src+sub_files[j].Name(), path_dst, method, deep2)
					}
				} else {
					folder_err = true
				}
				method.WithFolderAfter(file_or_dir, is_mount, folder_err, path_src, path_dst)
			}
		} else {
			path_src := FilePathEndSlashRemove(path_src_real)
			regular := file_or_dir.Mode().IsRegular()
			method.WithFile(file_or_dir, regular, path_src, path_dst)
		}
	} else {
		link_folder := FileLinkIsDir(path_src_real)
		next := method.WithLink(file_or_dir, link_folder, "", "")
		if next { // TODO!!!??
			if link_folder {

			} else {

			}
		}
	}
}

// ================

type folderWalker_size struct {
	counter_size      *AInt64
	counter_files     *AInt64
	counter_folders   *AInt64
	counter_unread    *AInt64
	counter_irregular *AInt64
	counter_mount     *AInt64
	counter_symlinks  *AInt64
}

func (m *folderWalker_size) WithFile(f os.FileInfo, regular bool, path_src string, path_dst string) {
	if regular {
		if m.counter_files != nil {
			m.counter_files.Add(1)
		}
		if m.counter_size != nil {
			m.counter_size.Add(f.Size())
		}
	} else {
		if m.counter_irregular != nil {
			m.counter_irregular.Add(1)
		}
	}
}

func (m *folderWalker_size) WithFolderBefore(f os.FileInfo, is_mount bool, path_src string, path_dst string) {
	if m.counter_folders != nil {
		m.counter_folders.Add(1)
	}
	if is_mount {
		if m.counter_mount != nil {
			m.counter_mount.Add(1)
		}
	}
	//return !is_mount
}

func (m *folderWalker_size) WithFolderAfter(f os.FileInfo, is_mount bool, list_err bool, path_src string, path_dst string) {
	if list_err {
		Prln("FOLDER READ ERROR:" + path_src)
		if m.counter_unread != nil {
			m.counter_unread.Add(1)
		}
	}
}

func (m *folderWalker_size) WithLink(f os.FileInfo, is_folder bool, path_src string, path_dst string) bool {
	if m.counter_symlinks != nil {
		m.counter_symlinks.Add(1)
	}
	return false
}

func FoldersRecursively_Size(mount_list [][2]string, file_or_dir os.FileInfo, path_real string,
	counter_size *AInt64,
	counter_files *AInt64,
	counter_folders *AInt64,
	counter_unread *AInt64,
	counter_irregular *AInt64,
	counter_mount *AInt64,
	counter_symlinks *AInt64) {
	m := &folderWalker_size{
		counter_size:      counter_size,
		counter_files:     counter_files,
		counter_folders:   counter_folders,
		counter_unread:    counter_unread,
		counter_irregular: counter_irregular,
		counter_mount:     counter_mount,
		counter_symlinks:  counter_symlinks}
	FoldersRecursively_Walk(mount_list, file_or_dir, path_real, "", m, 0)
}

// ================

type folderWalker_delete struct {
	counter_objects *AInt64
	current_file    *AString
	counter_size    *AInt64
}

func (m *folderWalker_delete) WithFile(f os.FileInfo, regular bool, path_src string, path_dst string) {
	if regular {
		m.current_file.Set(path_src)
		Prln("deleting file: " + path_src)
		//ok := true
		ok := FileDelete(path_src)
		if ok {
			if m.counter_objects != nil {
				m.counter_objects.Add(1)
			}
			if m.counter_size != nil {
				m.counter_size.Add(f.Size())
			}
		}
	} else {

	}
}

func (m *folderWalker_delete) WithFolderBefore(f os.FileInfo, is_mount bool, path_src string, path_dst string) {

}

func (m *folderWalker_delete) WithFolderAfter(f os.FileInfo, is_mount bool, list_err bool, path_src string, path_dst string) {
	list, err := Folder_ListFiles(path_src)
	if err == nil && len(list) == 0 {
		Prln("deleting folder: " + path_src)
		ok := FileDelete(path_src)
		//ok := true
		if ok && m.counter_objects != nil {
			m.counter_objects.Add(1)
		}
	} else {
		//skip
	}
}

func (m *folderWalker_delete) WithLink(f os.FileInfo, is_folder bool, path_src string, path_dst string) bool {
	return false
}

func FoldersRecursively_Delete(mount_list [][2]string, file_or_dir os.FileInfo, path_real string, counter_size *AInt64, counter_objects *AInt64, current_file *AString) {
	m := &folderWalker_delete{counter_objects: counter_objects, current_file: current_file, counter_size: counter_size}
	FoldersRecursively_Walk(mount_list, file_or_dir, path_real, "", m, 0)
}

// ================

const FILE_INTERACTIVE_RETRY = "retry"
const FILE_INTERACTIVE_REPLACE = "replace"
const FILE_INTERACTIVE_SKIP = "skip"
const FILE_INTERACTIVE_NEWNAME = "newname"

const FILE_INTERACTIVE_ASK_EXIST = 0
const FILE_INTERACTIVE_ASK_ERROR = 1
const FILE_INTERACTIVE_ASK_PANIC = 2

type FileInteractiveResponse struct {
	Command    string
	SaveChoice bool
}

type FileInteractiveRequest struct {
	Attempt  int
	FileName string
	AskType  int
}

type folderWalker_copymove struct {
	counter_size           *AInt64
	counter_files_done     *AInt64
	counter_files_replaced *AInt64
	counter_files_skipped  *AInt64
	buffer                 int
	chan_cmd               chan FileInteractiveResponse
	chan_ask               chan FileInteractiveRequest
	cmd_saved              string
	current_file           *AString
	move                   bool
	disk_equal             bool
}

func (m *folderWalker_copymove) WithFile(f os.FileInfo, regular bool, path_src string, path_dst string) {
	if regular {
		Prln(path_src + " >>> " + path_dst)
		m.current_file.Set(path_src)

		cmd := ""
		exist := true
		if FilePathEndSlashRemove(path_src) == FilePathEndSlashRemove(path_dst) {
			if m.move {
				cmd = FILE_INTERACTIVE_SKIP
			} else {
				cmd = FILE_INTERACTIVE_NEWNAME
			}
		} else {
			exist = FileExists(path_dst)
		}
		ask := exist
		can_do := true
		path_dst2 := path_dst

		num_ask_this := 0
	ask_label:
		num_ask_this++
		if ask {
			if cmd == "" {
				cmd = m.cmd_saved
			}
			if num_ask_this > 1 {
				cmd = ""
			}
			if cmd == "" {
				if exist {
					m.chan_ask <- FileInteractiveRequest{Attempt: num_ask_this, FileName: path_dst, AskType: FILE_INTERACTIVE_ASK_EXIST}
				} else {
					m.chan_ask <- FileInteractiveRequest{Attempt: num_ask_this - 1, FileName: path_src, AskType: FILE_INTERACTIVE_ASK_ERROR}
				}
				tcmd := <-m.chan_cmd
				if len(tcmd.Command) > 0 {
					cmd = tcmd.Command
					if tcmd.SaveChoice {
						m.cmd_saved = cmd
					}
				}
			}
			if cmd == FILE_INTERACTIVE_RETRY || cmd == FILE_INTERACTIVE_REPLACE {
				can_do = true
			}
			if cmd == FILE_INTERACTIVE_SKIP {
				can_do = false
			}
			if cmd == FILE_INTERACTIVE_NEWNAME {
				COPY_LABEL := "copy"
				path_dst2 = FileNameForCopy(path_dst, COPY_LABEL)
			}
		}
		if can_do {
			if !m.move || !m.disk_equal {
				Prln("atom copy file: [" + path_src + " >> " + path_dst2 + "]")
				size_old := m.counter_size.Get()
				err := FileCopyAtom(path_src, path_dst2, m.counter_size, m.buffer)
				if err != nil {
					m.counter_size.Set(size_old)
					Prln("copy err: " + err.Error())
					ask = true
					cmd = ""
					goto ask_label
				} else {
					if m.counter_files_done != nil {
						m.counter_files_done.Add(1)
					}
					if m.move {
						Prln("delete old file after move:" + path_src)
						ok2 := FileDelete(path_src)
						if !ok2 {
							Prln("deleting old file after move PROBLEM: [" + path_src + "]")
							m.chan_ask <- FileInteractiveRequest{Attempt: 0, FileName: path_src, AskType: FILE_INTERACTIVE_ASK_PANIC}
							<-m.chan_cmd
						}
					}
				}
			} else {
				Prln("renaming file: [" + path_src + " >> " + path_dst2 + "]")
				path_dst_back := ""
				if exist && path_dst == path_dst2 {
					path_dst_back = FileNameForCopy(path_dst, "BACKUP")
					Prln("renaming copy file before rename: [" + path_dst + " >> " + path_dst_back + "]")
					if !FileRename(path_dst, path_dst_back) {
						ask = true
						goto ask_label
					}
				}
				ok := FileRename(path_src, path_dst2)
				if ok {
					if exist {
						Prln("delete old file after rename:" + path_dst_back)
						ok2 := FileDelete(path_dst_back)
						if !ok2 {
							Prln("deleting old file after rename PROBLEM: [" + path_dst_back + "]")
							m.chan_ask <- FileInteractiveRequest{Attempt: 0, FileName: path_dst_back, AskType: FILE_INTERACTIVE_ASK_PANIC}
							<-m.chan_cmd
						}
					}
					if m.counter_files_done != nil {
						m.counter_files_done.Add(1)
					}
					if m.counter_size != nil {
						m.counter_size.Add(f.Size())
					}
				} else {
					fail_skip := false
					if exist && path_dst == path_dst2 {
						Prln("restoring of copy file of rename: [" + path_dst_back + " >> " + path_dst + "]")
						ok2 := FileRename(path_dst_back, path_dst)
						if !ok2 {
							Prln("restoring of copy file of rename PROBLEM: [" + path_dst_back + " >> " + path_dst + "]")
							fail_skip = true
							m.chan_ask <- FileInteractiveRequest{Attempt: 0, FileName: path_dst_back, AskType: FILE_INTERACTIVE_ASK_PANIC}
							<-m.chan_cmd
						}
					}
					if !fail_skip {
						ask = true
						goto ask_label
					}
				}
			}
		}
	} else {
		Prln("skip irregular:" + path_src)
	}
}

func (m *folderWalker_copymove) WithFolderBefore(f os.FileInfo, is_mount bool, path_src string, path_dst string) {
	ok := FolderMake(path_dst)
	if ok {
		// if m.counter_objects != nil {
		// 	m.counter_objects.Add(1)
		// }
	}
}

func (m *folderWalker_copymove) WithFolderAfter(f os.FileInfo, is_mount bool, list_err bool, path_src string, path_dst string) {
	//if m.move {
	Prln("delete folder after move:" + path_src)
	list, err := Folder_ListFiles(path_src)
	if err == nil && len(list) == 0 {
		ok := FileDelete(path_src)
		if !ok {
			m.chan_ask <- FileInteractiveRequest{Attempt: 0, FileName: path_src, AskType: FILE_INTERACTIVE_ASK_PANIC}
			<-m.chan_cmd
		}
	}
	//}
}

func (m *folderWalker_copymove) WithLink(f os.FileInfo, is_folder bool, path_src string, path_dst string) bool {
	return false
}

func FoldersRecursively_Copy(mount_list [][2]string, file_or_dir os.FileInfo, path_src_real string, path_dst_real string, counter_size *AInt64, counter_files_done *AInt64, buffer int, chan_cmd chan FileInteractiveResponse, chan_ask chan FileInteractiveRequest, current_file *AString) {
	m := &folderWalker_copymove{counter_size: counter_size, counter_files_done: counter_files_done, buffer: buffer, chan_cmd: chan_cmd, chan_ask: chan_ask, current_file: current_file, move: false}
	FoldersRecursively_Walk(mount_list, file_or_dir, path_src_real, FolderPathEndSlash(path_dst_real), m, 0)
}

func FoldersRecursively_Move(mount_list [][2]string, file_or_dir os.FileInfo, path_src_real string, path_dst_real string, counter_size *AInt64, counter_files_done *AInt64, buffer int, chan_cmd chan FileInteractiveResponse, chan_ask chan FileInteractiveRequest, current_file *AString, disk_equal bool) {
	m := &folderWalker_copymove{counter_size: counter_size, counter_files_done: counter_files_done, buffer: buffer, chan_cmd: chan_cmd, chan_ask: chan_ask, current_file: current_file, move: true, disk_equal: disk_equal}
	FoldersRecursively_Walk(mount_list, file_or_dir, path_src_real, FolderPathEndSlash(path_dst_real), m, 0)
}
