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
	"time"
	//"strings"
	//"bufio"
)

const BytesInMb uint64 = 1024 * 1024

const DEFAULT_FILE_PERMISSION os.FileMode = 776
const DEFAULT_FILE_PERMISSION_RUN os.FileMode = 777

var appdir = ""

var END_SLASH = ""

func init() {
	END_SLASH = string(os.PathSeparator)
	appdir = folderLocation_App()
}

func FolderPathEndSlash(path string) string {
	if path != "" && StringEnd(path, 1) != END_SLASH {
		return path + END_SLASH
	} else {
		return path
	}
}

func FilePathEndSlashRemove(filepath string) string {
	if StringEnd(filepath, 1) == END_SLASH {
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

func folderLocation_App() string {
	if len(os.Args) == 0 {
		return ""
	}
	location := os.Args[0]
	strs := StringSplit(location, END_SLASH)
	if len(strs) == 0 {
		return ""
	}
	strs2 := strs[:len(strs)-1]
	location = StringJoin(strs2, END_SLASH)
	location = FolderPathEndSlash(location)
	//Prln(">" + location + "]")
	if StringPart(location, 1, 2) != "./" { //home
		return location
	} else {
		location = FolderLocation_WorkDir() + StringPart(location, 3, 0)
		return FolderPathEndSlash(location)
	}
}

func FolderLocation_App() string {
	return appdir
}

func Folder_ListFiles(dirname string, fixlinks_isdir bool) ([]FileReport, error) {
	arr1, err := ioutil.ReadDir(dirname)
	if err != nil {
		return []FileReport{}, err
	}
	arr2 := make([]FileReport, len(arr1))
	path := FolderPathEndSlash(dirname)
	for j := 0; j < len(arr1); j++ {
		arr2[j] = NewFileReport(arr1[j], path, fixlinks_isdir)
	}
	return arr2, nil
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

func FileRename(path_src string, path_dst string) (bool, string) {
	if FileExists(path_dst) {
		return false, "Not exist"
	}
	err := os.Rename(path_src, path_dst)
	if err != nil {
		//Prln(err.Error())
		return false, err.Error()
	}
	return true, ""
}

//https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-file-using-go
/*func FileCopyAtom(src string, dst string, copied *AInt64, buffer_size int) error { //, dest_reopen bool
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

		if err0 := dest.Sync(); err0 != nil {
			return ErrorWithText("[dst_write]" + err0.Error()) //err
		}
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
}*/

func FileCopyAtom(src string, dst string, copied *AInt64, buffer_size int) error {
	source, err1 := os.Open(src)
	if err1 != nil {
		return ErrorWithText("[src_init]" + err1.Error()) // err1
	}
	defer source.Close()

	source_size := int64(1024)
	st, err3 := source.Stat()
	if err3 != nil {
		return ErrorWithText("[src_size]" + err3.Error())
	} else {
		if !st.Mode().IsRegular() {
			return ErrorWithText("[src_not_regulat_file]")
		} else {
			source_size = st.Size()
		}
	}

	safe_buffer := MAXI64(0, MINI64(int64(buffer_size), source_size))
	buf := make([]byte, safe_buffer)
	off := int64(0)

	dest, err2 := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0777) //os.O_APPEND|os.O_TRUNC|os.O_SYNC
	//dest, err2 := os.Create(dst)
	if err2 != nil {
		return ErrorWithText("[dst_init]" + err2.Error()) //err2
	}
	defer dest.Close()
	dest.Sync()
	//w := bufio.NewWriter(dest)

	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return ErrorWithText("[src_read]" + err.Error()) //err
		}
		if n == 0 {
			break
		}

		buf2 := buf[:n]
		n2, err0 := dest.Write(buf2)
		//dest.Sync()
		//n2, err0 := w.Write(buf2)
		if err0 != nil {
			msg := "[dst_write] " + I2S(n2) + "/" + I2S(n) + " - " + err0.Error()
			//Prln(msg)
			return ErrorWithText(msg)
		}
		//w.Flush()

		n64 := int64(n)
		off += n64

		if copied != nil {
			copied.Add(n64)
		}
	}
	dest.Sync()
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

type FileReport struct { // os.FileInfo interface
	Path        string
	NameOnly    string
	FullName    string
	IsRegular   bool
	IsDirectory bool
	IsLink      bool
	SizeBytes   int64
	mode        os.FileMode
	modTime     time.Time
	sys         interface{}
}

func (fr FileReport) Name() string {
	return fr.NameOnly
}

func (fr FileReport) IsDir() bool {
	return fr.IsDirectory
}

func (fr FileReport) Size() int64 {
	return fr.SizeBytes
}

func (fr FileReport) Mode() os.FileMode {
	return fr.mode
}

func (fr FileReport) ModTime() time.Time {
	return fr.modTime
}

func (fr FileReport) Sys() interface{} {
	return fr.sys
}

func (fr *FileReport) SetModeTime(mode string, modTime Time) {
	fr.modTime = time.Time(modTime)
}

func NewFileReport(fi os.FileInfo, path string, fixlinks_isdir bool) FileReport {
	fr := FileReport{}
	fr.Path = path
	fr.NameOnly = fi.Name()
	if fr.NameOnly != END_SLASH { // tested Linux only!!
		fr.FullName = path + fr.NameOnly
	} else {
		fr.FullName = path
	}
	fr.mode = fi.Mode()
	fr.modTime = fi.ModTime()
	fr.SizeBytes = fi.Size()
	fr.IsRegular = fr.mode.IsRegular()
	fr.sys = fi.Sys()
	fr.IsDirectory = fi.IsDir()
	if fr.IsRegular {
		fr.IsLink = (fr.mode&os.ModeSymlink != 0)
		if fr.IsLink && fixlinks_isdir {
			fr.IsDirectory = FileLinkIsDir(fr.FullName)
		}
	}
	return fr
}

// 'os.Lstat()' reads the link itself.
// 'os.Stat()' would read the link's target.
func FileInfo(fullname string, fixlinks_isdir bool) (FileReport, error) {
	f, err := os.Lstat(fullname)
	if err != nil {
		return FileReport{}, err
	}
	name := f.Name()
	parent := StringPart(fullname, 1, StringLength(fullname)-StringLength(name))
	return NewFileReport(f, FolderPathEndSlash(parent), fixlinks_isdir), nil
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

func FileNameForCopy(path_dst string, copy_label string, isdir bool) string {
	path2, name2 := FileSplitPathAndName(path_dst)
	name3, ext := FileSeparateNumberedName(name2, copy_label, isdir)
	name4 := FileFindFreeName(path2, name3, copy_label, ext)
	return path2 + name4
}

func FileSeparateNumberedName(filename string, copylabel string, isdir bool) (string, string) {
	ext := ""
	if !isdir {
		ext = FileExtension(filename)
	}
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
	fname := FilePathEndSlashRemove(filename)
	r1, r2 := filepath.Split(fname)
	//Prln(fname + ">>" + r1 + ">" + r2)
	return r1, r2
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
