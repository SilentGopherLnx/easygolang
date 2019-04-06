package easygolang

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	//"strings"
)

const BytesInMb uint64 = 1024 * 1024

const DEFAULT_FILE_PERMISSION os.FileMode = 776
const DEFAULT_FILE_PERMISSION_RUN os.FileMode = 777

func FolderPathEndSlash(path string) string {
	slash := string(os.PathSeparator)
	if path != "" && StringEnd(path, 1) != slash {
		return path + slash
	} else {
		return path
	}
}

func FilePathEndSlashRemove(filepath string) string {
	separator := GetOS_Slash()
	if StringEnd(filepath, 1) == separator {
		return StringPart(filepath, 1, StringLength(filepath)-1)
	}
	return filepath
}

// =====================

func FolderLocation_WorkDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		cwd = ""
	}
	return FolderPathEndSlash(cwd)
}

func FolderLocation_UserHome() string {
	home := ""
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}
	return FolderPathEndSlash(home)
}

func FolderLocation_App() string {
	slash := string(os.PathSeparator)
	if len(os.Args) == 0 {
		return ""
	}
	location := os.Args[0]
	strs := StringSplit(location, slash)
	if len(strs) == 0 {
		return ""
	}
	strs2 := strs[:len(strs)-1]
	location = StringJoin(strs2, slash)
	return FolderPathEndSlash(location)
}

func Folder_ListFiles(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

func FolderMake(pathname string) bool {
	Prln(filepath.Dir(pathname))
	err := os.MkdirAll(pathname, os.ModePerm)
	if err == nil {
		return true
	} else {
		return false
	}
}

// =====================

func FileMake(pathname string) bool {
	Prln(filepath.Dir(pathname))
	_, err := os.Create(pathname)
	if err == nil {
		return true
	} else {
		return false
	}
}

func FileRename(path_src string, path_dst string) bool {
	if FileExists(path_dst) {
		return false
	}
	err := os.Rename(path_src, path_dst)
	if err != nil {
		Prln(err.Error())
		return false
	}
	return true
}

//https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go
func FileCopyAtom(src string, dst string, copied *AInt64, buffer_size int) error { //, dest_reopen bool
	source, err1 := os.Open(src)
	if err1 != nil {
		return ErrorWithText("[src]" + err1.Error()) // err1
	}
	//dest, err2 := os.Create(dst)
	dest, err2 := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0777)
	if err2 != nil {
		return ErrorWithText("[dst]" + err2.Error()) //err2
	}
	source_size := int64(1024)
	st, err3 := source.Stat()
	if err3 == nil {
		source_size = st.Size()
		if !st.Mode().IsRegular() {
			return ErrorWithText("Not a regular file")
		}
	}
	safe_buffer := MAXI64(0, MINI64(int64(buffer_size), source_size))
	buf := make([]byte, safe_buffer)
	off := int64(0)
	//closed := false
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return ErrorWithText("[src_read]" + err.Error()) //err
		}
		if n == 0 {
			break
		}
		// if closed {
		// 	dest, err = os.OpenFile(dst, os.O_WRONLY|os.O_APPEND, 0777)
		// 	if err != nil {
		// 		return ErrorWithText("[dst_reopen]" + err.Error()) //err
		// 	}
		// }

		if _, err := dest.WriteAt(buf[:n], off); err != nil {
			return ErrorWithText("[dst_write]" + err.Error()) //err
		}
		n64 := int64(n)
		off += n64
		if copied != nil {
			copied.Add(n64)
		}
		// if dest_reopen {
		// 	dest.Close()
		// 	closed = true
		// }
	}
	source.Close()
	//if !dest_reopen {
	dest.Close()
	//}
	return nil
}

/*func FileMove(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %s", err)
	}
	outputFile, err := os.Create(destPath)
	if err != nil {
		inputFile.Close()
		return fmt.Errorf("Couldn't open dest file: %s", err)
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, inputFile)
	inputFile.Close()
	if err != nil {
		return fmt.Errorf("Writing to output file failed: %s", err)
	}
	// The copy was successful, so now delete the original file
	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Failed removing original file: %s", err)
	}
	return nil
}*/

// =====================

func FileFindFreeName(folderpath string, filename string, copylabel string, ext string) string {
	folderpath2 := FolderPathEndSlash(folderpath)
	ext2 := ""
	if len(ext) > 0 {
		ext2 = "." + ext
	}
	// copylabel2 := ""
	// if len(copylabel) > 0 {
	// 	copylabel2 = " (" + copylabel + ")"
	// }
	tfile := filename + ext2
	if copylabel == "" && !FileExists(folderpath2+tfile) {
		return tfile
	} else {
		copylabel2 := ""
		if len(copylabel) > 0 {
			copylabel2 = copylabel + " "
		}
		start := 1
		if copylabel == "" {
			start = 2
		}
		for j := start; j < 100; j++ {
			tfile = filename + " (" + copylabel2 + StringEnd("0"+I2S(j), 2) + ")" + ext2
			if !FileExists(folderpath2 + tfile) {
				return tfile
			}
		}
	}
	return ""
}

func FileNameForCopy(path_dst string, copy_label string) string {
	path2, name2 := FileSplitPathAndName(path_dst)
	name3, ext := FileSeparateNumberedName(name2, copy_label)
	name4 := FileFindFreeName(path2, name3, copy_label, ext)
	return path2 + name4
}

func FileSeparateNumberedName(filename string, copylabel string) (string, string) {
	ext := FileExtension(filename)
	name := filename
	if len(ext) > 0 {
		name = StringPart(filename, 1, StringLength(filename)-StringLength(ext)-1)
	}
	copylabel2 := ""
	if len(copylabel) > 0 {
		copylabel2 = copylabel + " "
	}
	lbl := StringLength(copylabel2)
	sl := StringLength(name)
	if StringPart(name, sl, sl) == ")" {
		if StringPart(name, sl-lbl-4, sl-lbl-3) == " (" {
			if StringFind(name, copylabel2) == sl-lbl-2 {
				num := StringPart(name, sl-2, sl-1)
				if IsInt(num) {
					name = StringPart(name, 1, sl-lbl-5)
				}
			}
		}
	}
	return name, ext
}

func FileSplitPathAndName(filename string) (string, string) {
	return filepath.Split(filename)
	//path, name :=
	//Prln(path)
	//return path, name
}

func FileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if len(ext) > 1 {
		return StringDown(StringPart(ext, 2, 0))
	} else {
		return ""
	}
	// fn := StringSplit(filename, ".")
	// if len(fn) > 1 {
	// 	tfile := StringDown(fn[len(fn)-1])
	// 	if StringFind(tfile, "/") > 0 {
	// 		return ""
	// 	}
	// 	return tfile
	// }
	// return ""
}

// 'os.Lstat()' reads the link itself.
// 'os.Stat()' would read the link's target.
func FileInfo(name string) (os.FileInfo, bool) {
	f, err := os.Lstat(name)
	if err != nil {
		return nil, false
	}
	return f, true
}

func FilePermissionsString(name string) string {
	f, err := os.Stat(name)
	if err != nil {
		return ""
	}
	return f.Mode().String()
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func FileDelete(path string) bool {
	err := os.Remove(path)
	if err != nil {
		return false
	}
	return true
}

func FileIsLink(f os.FileInfo) bool {
	return (f.Mode()&os.ModeSymlink != 0)
}

func FileLinkIsDir(fullname string) bool {
	path2, err := filepath.EvalSymlinks(fullname)
	if err == nil {
		f, err := os.Stat(path2)
		if err == nil {
			return f.IsDir()
		}
	}
	return false
}

func FileEvalSymlinks(fullpath string) (string, error) {
	return filepath.EvalSymlinks(fullpath)
}

func FileSizeNiceString(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

/*[yes]:{root}drwx------    X
[yes]:{lost+found}drwx------  X
[no]:{swapfile}-rw-------  X
[no]:{ssh1c.cfg}-r----x---*/

// =====================

func FileTextRead(name string) (string, bool) {
	f, err1 := os.Open(name)
	if err1 != nil {
		return "", false
	}
	data, err2 := ioutil.ReadAll(f)
	if err2 != nil {
		return "", false
	} else {
		return string(data), true
	}
}

func FileTextWrite(fname string, data string) bool {
	bytedata := []byte(data)
	err := ioutil.WriteFile(fname, bytedata, DEFAULT_FILE_PERMISSION)
	if err != nil {
		return false
	}
	return true
}

func FileBytesRead(fname string) (*[]byte, bool) {
	fbytes, err := ioutil.ReadFile(fname)
	if err != nil {
		return &[]byte{}, false
	}
	bytedata := CloneBytesArray(fbytes)
	return &bytedata, true
}

func FileBytesWrite(fname string, bytedata *[]byte, executable bool) bool {
	rights := DEFAULT_FILE_PERMISSION
	if executable {
		rights = DEFAULT_FILE_PERMISSION_RUN
	}
	err := ioutil.WriteFile(fname, *bytedata, rights)
	if err != nil {
		return false
	}
	return true
}

//=================

func FileSortName(filename string) string {
	ext := filepath.Ext(filename)
	name := filename[:len(filename)-len(ext)]
	// split numeric suffix
	i := len(name) - 1
	for ; i >= 0; i-- {
		if '0' > name[i] || name[i] > '9' {
			break
		}
	}
	i++
	// string numeric suffix to uint64 bytes
	// empty string is zero, so integers are plus one
	b64 := make([]byte, 64/8)
	s64 := name[i:]
	if len(s64) > 0 {
		u64, err := strconv.ParseUint(s64, 10, 64)
		if err == nil {
			binary.BigEndian.PutUint64(b64, u64+1)
		}
	}
	// prefix + numeric-suffix + ext
	return name[:i] + string(b64) + ext
}
