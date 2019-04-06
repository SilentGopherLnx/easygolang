package easylinux

import (
	. "github.com/SilentGopherLnx/easygolang"
)

func NetworkGetpcs() {
	//arp -n - all pc in net
	//net help share
	// nmblookup -S WORKGROUP
	//avahi-browse -r _smb._tcp
	//avahi-browse -a
	Prln("nmblookup -T WORKGROUP")
	// NAS, 192.168.999.111 WORKGROUP<00>
	// dell, 192.168.999.222 WORKGROUP<00>

	//smbclient -L ?
}
