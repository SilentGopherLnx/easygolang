package easylinux

import (
	"net/url"

	. "github.com/SilentGopherLnx/easygolang"
)

var linux_mount_gvfs string = ""
var linux_mount_gvfs_len int = 0

var hide_file_protocol = true

func init() {
	uid, _, _ := GetPC_UserUidLoginName()
	linux_mount_gvfs = "/run/user/" + uid + "/gvfs/"
	linux_mount_gvfs_len = len(linux_mount_gvfs)
}

type LinuxPath struct {
	add_end_slash bool
	path_url      string // smb://127.0.0.1/demo%20folder/%D1%80%D1%83%D1%81%D1%81%D0%BA%D0%B8%D0%B9.txt
	path_real     string // /run/user/1000/gvfs/smb-share:server=127.0.0.1,share=demo%20folder/русский.txt
	path_visual   string // smb://127.0.0.1/demo folder/русский.txt
	parse_error   bool
}

func NewLinuxPath(add_end_slash bool) *LinuxPath {
	return &LinuxPath{add_end_slash: add_end_slash, path_url: "file:///", path_real: "/", path_visual: "file:///"}
}

func (p *LinuxPath) SetUrl(path_url string) {
	if p.add_end_slash {
		p.path_url = FolderPathEndSlash(path_url)
	} else {
		p.path_url = path_url
	}
	p.path_visual = linuxFilePath_Unescape(p.path_url)
	p.path_real, p.parse_error = linuxFilePathProtocol_VisualToReal(p.path_visual)
	p.parse_error = !p.parse_error
}

func (p *LinuxPath) SetReal(path_real string) {
	if p.add_end_slash {
		p.path_real = FolderPathEndSlash(path_real)
	} else {
		p.path_real = path_real
	}
	p.path_visual = linuxFilePathProtocol_RealToVisual(p.path_real)
	//Prln(linuxFilePath_Escape(p.path_real))
	p.path_url = linuxFilePathProtocol_RealToVisual(linuxFilePath_Escape(p.path_real))
	p.parse_error = false
}

func (p *LinuxPath) SetVisual(path_visual string) {
	p.path_visual = path_visual
	if StringFind(p.path_visual, "/") == 1 {
		p.path_visual = "file://" + p.path_visual
	}
	if p.add_end_slash {
		p.path_visual = FolderPathEndSlash(p.path_visual)
	} else {
		p.path_visual = p.path_visual
	}
	p.path_url = linuxFilePath_Escape(p.path_visual)
	p.path_real, p.parse_error = linuxFilePathProtocol_VisualToReal(p.path_visual)
	p.parse_error = !p.parse_error
}

func (p *LinuxPath) GetUrl() string {
	return p.path_url
}

func (p *LinuxPath) GetReal() string {
	return p.path_real
}

func (p *LinuxPath) GetVisual() string {
	if hide_file_protocol {
		return removeFileProtocol(p.path_visual)
	}
	return p.path_visual
}

func (p *LinuxPath) GetParseProblems() bool {
	return p.parse_error
}

func (p *LinuxPath) GoUp() {
	p.SetReal(LinuxFileGetParent(p.path_real))
}

func (p *LinuxPath) GoDeep(subfolder string) {
	r := FolderPathEndSlash(p.path_real)
	p.SetReal(r + subfolder)
}

func (p *LinuxPath) GetLastNode() string {
	return LinuxFileNameFromPath(p.path_real)
}

func removeFileProtocol(path string) string {
	ind := StringFind(path, "file://")
	if ind != 1 {
		return path
	}
	return StringPart(path, 8, 0)
}

//what is correct convertation?
func linuxFilePath_Escape(path string) string {
	esc := UrlQueryEscape(path)
	esc = StringReplace(esc, "+", "%20")
	esc = StringReplace(esc, "%2B", "+")
	esc = StringReplace(esc, "%3A", ":")
	esc = StringReplace(esc, "%2F", "/")
	//??
	return linuxFilePath_EscapeArgs(esc)
}

func linuxFilePath_EscapeArgs(args string) string {
	esc := args
	esc = StringReplace(esc, "%2C", ",")
	esc = StringReplace(esc, "%3D", "=")
	//??
	return esc
}

//what is correct convertation?
func linuxFilePath_Unescape(path string) string {
	unesc := path
	//unesc := UrlQueryUnescape(unesc)
	// unesc = StringReplace(unesc, "/", "%2F")
	// unesc = StringReplace(unesc, ",", "%2C")
	// unesc = StringReplace(unesc, ":", "%3A")
	// unesc = StringReplace(unesc, "=", "%3D")
	unesc = StringReplace(unesc, "+", "%2B")
	unesc = UrlQueryUnescape(unesc)
	return unesc
}

func linuxFilePathProtocol_VisualToReal(path_visual string) (string, bool) {
	path2 := path_visual
	ind := StringFind(path2, "://")
	protocol := StringPart(path2, 1, ind-1)
	path2 = StringPart(path2, ind+3, 0)
	path_arr := StringSplit(path2, "/")
	//Prln(protocol)
	switch protocol {
	case "file":
		return path2, true
	case "smb":
		if len(path_arr) > 2 {
			return linux_mount_gvfs + "smb-share:server=" + UrlQueryEscape(path_arr[0]) + ",share=" + UrlQueryEscape(path_arr[1]) + "/" + StringJoin(path_arr[2:], "/"), true
		}
	case "mtp", "gphoto2":
		if len(path_arr) > 1 {
			return linux_mount_gvfs + protocol + ":host=" + UrlQueryEscape(path_arr[0]) + "/" + StringJoin(path_arr[1:], "/"), true
		}
	case "dav", "ftp", "davs", "ftps":
		pr := protocol
		ssl := ""
		if StringEnd(protocol, 1) == "s" {
			ssl = ",ssl=true"
			pr = StringPart(protocol, 1, StringLength(protocol)-1)
		}
		if len(path_arr) > 1 {
			user := ""
			host := ""
			port := ""
			ab := StringSplit(path_arr[0], "@")
			if len(ab) == 2 {
				host = ab[1]
				user = ab[0]
				ab = StringSplit(user, ":")
				if len(ab) == 2 {
					user = ab[0]
					port = ab[1]
				}
			} else {
				host = path_arr[0]
				ab = StringSplit(host, ":")
				if len(ab) == 2 {
					host = ab[0]
					port = ab[1]
				}
			}
			if len(user) > 0 {
				user = ",user=" + UrlQueryEscape(user)
			}
			if len(port) > 0 {
				port = ",port=" + UrlQueryEscape(port)
			}
			host = UrlQueryEscape(host)
			return linux_mount_gvfs + pr + ":host=" + host + ssl + user + port + "/" + StringJoin(path_arr[1:], "/"), true

		}
	}
	return "/", false
}

func linuxFilePathProtocol_RealToVisual(path_real string) string {
	path2 := path_real
	protocol := ""
	protocol_args := ""
	var q url.Values
	var qerr error = nil
	ind := StringFind(path_real, linux_mount_gvfs)
	if ind == 1 && len(path2) > linux_mount_gvfs_len {
		path2 = StringPart(path_real, linux_mount_gvfs_len+1, 0)
		ind2 := StringFind(path2, "/")
		protocol_args = StringPart(path2, 1, ind2-1)
		ind3 := StringFind(protocol_args, ":")
		protocol = StringPart(protocol_args, 1, ind3-1)
		protocol_args = StringPart(protocol_args, ind3+1, 0)
		path2 = StringPart(path2, ind2+1, 0)
		if path2 == "/" {
			path2 = ""
		}
		q, qerr = UrlQueryParse(StringReplace(StringReplace(protocol_args, ";", "#"), ",", ";"))
		//Prln(protocol)
		//Prln(protocol_args)
		if qerr == nil {
			switch protocol {
			case "smb-share":
				servername := StringTrim(q.Get("server"))
				foldershare := StringTrim(q.Get("share"))
				return "smb://" + servername + "/" + foldershare + "/" + path2
			case "mtp", "gphoto2":
				hostname := StringTrim(q.Get("host"))
				return protocol + "://" + linuxFilePath_EscapeArgs(hostname) + "/" + path2
			case "dav", "ftp":
				hostname := StringTrim(q.Get("host"))
				username := StringTrim(q.Get("user"))
				port := StringTrim(q.Get("port"))
				ssl := ""
				if StringFind(protocol_args, ",ssl=true,") > 0 {
					ssl = "s"
				}
				if len(username) > 0 {
					username += "@"
				}
				if len(port) > 0 {
					port = ":" + port
				}
				return protocol + ssl + "://" + username + hostname + port + "/" + path2
			}
		}
	} else {
		return "file://" + path2
	}
	return "file:///"
}
