package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

const SMB_SLASH2 = "smb://"

type SMB_Name struct {
	Name string
	IPv4 string
	IPv6 string
	Port int
}

// smbs := SMB_ScanNetwork()
// for j := 0; j < len(smbs); j++ {
// 	Prln("{" + smbs[j].Name + "/" + smbs[j].IPv4 + "/" + smbs[j].IPv6 + "/" + I2S(smbs[j].Port) + "}")
// }
func SMB_ScanNetwork() ([]SMB_Name, error) {
	//arp -n - all pc in net
	//avahi-browse -a
	//nmblookup -T WORKGROUP
	arg_f := func(str string) string {
		ind := StringFind(str, "=")
		v := StringPart(str, ind+1, 0)
		v = StringTrim(v)
		if StringFind(v, "[") == 1 {
			v = StringPart(v, 2, 0)
		}
		vl := StringLength(v)
		if StringFind(v, "]") == vl {
			v = StringPart(v, 1, vl-1)
		}
		return v
	}
	out, _, _ := ExecCommand("avahi-browse", "-r", "_smb._tcp", "-t")
	strs := StringSplitLines(out)
	arr := []SMB_Name{}
	var smb *SMB_Name = nil
	for j := 0; j < len(strs); j++ {
		line := strs[j]
		ch := StringPart(line, 1, 1)
		if ch == "+" || ch == "=" {
			if smb != nil && smb.Name != "" {
				arr = append(arr, *smb)
			}
			if ch == "=" {
				smb = &SMB_Name{}
			} else {
				smb = nil
			}
		} else {
			ab := StringSplit(line, "=")
			if len(ab) > 1 {
				arg := StringTrim(ab[0])
				if arg == "hostname" {
					smb.Name = arg_f(line)
				}
				if arg == "address" {
					ip := arg_f(line)
					if StringFind(ip, ".") > 0 {
						smb.IPv4 = ip
					} else {
						smb.IPv6 = ip
					}
				}
				if arg == "port" {
					smb.Port = S2I(arg_f(line))
				}
			}
		}
	}
	if smb != nil && smb.Name != "" {
		arr = append(arr, *smb)
	}
	arr2 := []SMB_Name{}
	for j := 0; j < len(arr); j++ {
		exist := false
		for k := 0; k < len(arr2); k++ {
			if arr[j].Name == arr2[k].Name {
				exist = true
				if arr2[k].IPv4 == "" {
					arr2[k].IPv4 = arr[j].IPv4
				}
				if arr2[k].IPv6 == "" {
					arr2[k].IPv6 = arr[j].IPv6
				}
			}
		}
		if !exist {
			arr2 = append(arr2, arr[j])
		}
	}
	return arr2, nil
}

// arr := SMB_GetPublicFolders("smbnas")
// for _, a := range arr[:] {
// 	Prln(a)
// }
func SMB_GetPublicFolders(name_or_ip string) ([]string, error) {
	out, _, _ := ExecCommand("smbclient", "-N", "-g", "-L", name_or_ip)
	strs := StringSplitLines(out)
	arr := []string{}
	for j := 0; j < len(strs); j++ {
		line := strs[j]
		if StringFind(line, "Disk|") == 1 {
			ab := StringSplit(line, "|")
			if len(ab) > 1 {
				arr = append(arr, ab[1])
			}
		}
	}
	if len(arr) == 0 {
		return arr, ErrorWithText("No Public Folders For This PC")
	}
	return arr, nil
}

func SMB_IsMount(p *LinuxPath, folder string, mountlist [][2]string) bool {
	p2 := NewLinuxPath(true)
	p2.SetVisual(p.GetVisual() + folder + "/")
	return LinuxFolderIsMountPoint(mountlist, p2.GetReal())
}

func SMB_CheckVirtualPath(url string) (bool, string, string) {
	len_smb := StringLength(SMB_SLASH2)
	if url == SMB_SLASH2 {
		return true, "", ""
	}
	if StringPart(url, 1, len_smb) == SMB_SLASH2 {
		pc_name := StringPart(url, len_smb+1, 0)
		pc_name = FilePathEndSlashRemove(pc_name)
		if StringFind(pc_name, GetOS_Slash()) == 0 {
			return false, pc_name, ""
		}
		arr := StringSplit(pc_name, "/")
		if len(arr) == 2 {
			return false, arr[0], arr[1]
		}
	}
	return false, "", ""
}

func SMB_UnMount(pc_name string, folder_name string) error {
	//cmd:="gio mount -u smb://"+server+"/"+StringDown(folder_name )+"/"
	//ExecCommand(cmd)
	Prln("UNmounting smb [" + pc_name + "][" + folder_name + "]")
	c1, c2, c3 := ExecCommand("gio", "mount", "-u", "smb://"+pc_name+"/"+StringDown(folder_name)+"/")
	res := c1 + "/" + c2 + "/" + c3
	if len(res) > 2 {
		return ErrorWithText(res)
	}
	return nil
}

func SMB_MountLoginAsk(pc_name string, folder_name string) (bool, error) {
	return sMB_MountLoginUse(pc_name, folder_name, "", "")
}

func SMB_MountLoginUse(pc_name string, folder_name string, login string, password string) error {
	_, err := sMB_MountLoginUse(pc_name, folder_name, login, password)
	return err
}

func sMB_MountLoginUse(pc_name string, folder_name string, login string, password string) (bool, error) {
	Prln("mounting smb [" + pc_name + "][" + folder_name + "] start")
	c1, c2, c3 := ExecCommandBytes([]byte{}, 2000, nil, "gio", "mount", "smb://"+pc_name+"/"+StringDown(folder_name)+"/")
	res := string(c1) + "/" + string(c2) + "/" + c3
	Prln("mounting smb [" + pc_name + "][" + folder_name + "] finish: " + res)
	if false {
		return true, nil
	}
	if len(res) > 2 {
		return false, ErrorWithText(res)
	}
	return false, nil
}
