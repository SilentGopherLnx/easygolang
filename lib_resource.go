package easygolang

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func ImageToCode(fname string, funcffix string) string {
	return FileToCodeGenerator(fname, "Image", funcffix)
}

func FileToCodeGenerator(fname string, ftype string, funcffix string) string {
	imgdata, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	stepRound := 5
	percentDone := 0
	var buffer bytes.Buffer
	//str := "func GetImage_" + funcffix + "() (*[]byte) {\n var imgdata = []byte{"
	buffer.WriteString("func Get" + ftype + "_" + funcffix + "() (*[]byte) {\n var imgdata = []byte{")
	flen := len(imgdata)
	for i, v := range imgdata {
		if i > 0 {
			//str += ", "+strconv.Itoa(int(v))
			buffer.WriteString("," + strconv.Itoa(int(v)))
		} else {
			//str += strconv.Itoa(int(v))
			buffer.WriteString(strconv.Itoa(int(v)))
		}
		percentDone2 := i * 100 / flen / stepRound * stepRound
		if percentDone2 > percentDone {
			percentDone = percentDone2
			Prln("file: [" + fname + "] done: " + I2S(percentDone) + "%")
		}
	}
	//str += "}\n return &imgdata\n}\n"
	buffer.WriteString("}\n return &imgdata\n}\n")
	return buffer.String()
}

//https://golangcode.com/unzip-files-in-go/
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}
		defer rc.Close()

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {

			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)

		} else {

			// Make File
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return filenames, err
			}

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return filenames, err
			}

			_, err = io.Copy(outFile, rc)

			// Close the file without defer to close before next iteration of loop
			outFile.Close()

			if err != nil {
				return filenames, err
			}

		}
	}
	return filenames, nil
}
