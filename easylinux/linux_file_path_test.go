package easylinux

import (
	"testing"

	. "github.com/SilentGopherLnx/easygolang"
)

func Test_LinuxPath(t *testing.T) {

	arr := [][2]string{
		{"file:///", "/"},
		{"file:///mnt/dm-1/", "/mnt/dm-1/"},
		{"file:///mnt/dm-1/%25D1%2589/", "/mnt/dm-1/%D1%89/"},
		{"file:///mnt/dm-1/demo%20folder/plus+/", "/mnt/dm-1/demo folder/plus+/"},
		{"file:///mnt/dm-1/%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9/", "/mnt/dm-1/русский/"},

		{"smb://127.0.0.1/sharedfolder/", "/run/user/1000/gvfs/smb-share:server=127.0.0.1,share=sharedfolder/"},
		{"davs://username@webdav.yandex.ru/", "/run/user/1000/gvfs/dav:host=webdav.yandex.ru,ssl=true,user=username/"},
		{"ftp://anonymous@127.0.0.1/", "/run/user/1000/gvfs/ftp:host=127.0.0.1,user=anonymous/"},
		{"ftp://smbnas.local:169/", "/run/user/1000/gvfs/ftp:host=smbnas.local,port=169/"},
		{"mtp://%5Busb%3A001,006%5D/", "/run/user/1000/gvfs/mtp:host=%5Busb%3A001%2C006%5D/"},
		{"gphoto2://%5Busb%3A001,007%5D/", "/run/user/1000/gvfs/gphoto2:host=%5Busb%3A001%2C007%5D/"},
	}

	for j := 0; j < len(arr); j++ {
		path_url := NewLinuxPath(true)
		path_real := NewLinuxPath(true)
		path_url.SetUrl(arr[j][0])
		path_real.SetReal(arr[j][1])
		if arr[j][0] != path_real.GetUrl() || path_real.GetParseProblems() {
			t.Error("TEST[" + I2S(j) + "]#1: requested version {" + arr[j][0] + "} != converted version {" + path_real.GetUrl() + "} " + B2S_YN(path_real.GetParseProblems()) + ". Original:[" + arr[j][1] + "]")
		}
		if arr[j][1] != path_url.GetReal() || path_url.GetParseProblems() {
			t.Error("TEST[" + I2S(j) + "]#2: requested version {" + arr[j][1] + "} != converted version {" + path_url.GetReal() + "} " + B2S_YN(path_url.GetParseProblems()) + ". Original:[" + arr[j][0] + "]")
		}
	}
}
