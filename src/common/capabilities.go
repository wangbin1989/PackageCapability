package common

import (
	"bytes"
	"encoding/binary"
	"log"
	"strings"

	"golang.org/x/sys/unix"
)

// CapabilityMap 是字符串到 unix.CAP_* 常量的映射
var CapabilityMap = map[string]uintptr{
	"CAP_CHOWN":            unix.CAP_CHOWN,
	"CAP_DAC_OVERRIDE":     unix.CAP_DAC_OVERRIDE,
	"CAP_DAC_READ_SEARCH":  unix.CAP_DAC_READ_SEARCH,
	"CAP_FOWNER":           unix.CAP_FOWNER,
	"CAP_FSETID":           unix.CAP_FSETID,
	"CAP_KILL":             unix.CAP_KILL,
	"CAP_SETGID":           unix.CAP_SETGID,
	"CAP_SETUID":           unix.CAP_SETUID,
	"CAP_SETPCAP":          unix.CAP_SETPCAP,
	"CAP_LINUX_IMMUTABLE":  unix.CAP_LINUX_IMMUTABLE,
	"CAP_NET_BIND_SERVICE": unix.CAP_NET_BIND_SERVICE,
	"CAP_NET_BROADCAST":    unix.CAP_NET_BROADCAST,
	"CAP_NET_ADMIN":        unix.CAP_NET_ADMIN,
	"CAP_NET_RAW":          unix.CAP_NET_RAW,
	"CAP_IPC_LOCK":         unix.CAP_IPC_LOCK,
	"CAP_IPC_OWNER":        unix.CAP_IPC_OWNER,
	"CAP_SYS_MODULE":       unix.CAP_SYS_MODULE,
	"CAP_SYS_RAWIO":        unix.CAP_SYS_RAWIO,
	"CAP_SYS_CHROOT":       unix.CAP_SYS_CHROOT,
	"CAP_SYS_PTRACE":       unix.CAP_SYS_PTRACE,
	"CAP_SYS_PACCT":        unix.CAP_SYS_PACCT,
	"CAP_SYS_ADMIN":        unix.CAP_SYS_ADMIN,
	"CAP_SYS_BOOT":         unix.CAP_SYS_BOOT,
	"CAP_SYS_NICE":         unix.CAP_SYS_NICE,
	"CAP_SYS_RESOURCE":     unix.CAP_SYS_RESOURCE,
	"CAP_SYS_TIME":         unix.CAP_SYS_TIME,
	"CAP_SYS_TTY_CONFIG":   unix.CAP_SYS_TTY_CONFIG,
	"CAP_MKNOD":            unix.CAP_MKNOD,
	"CAP_LEASE":            unix.CAP_LEASE,
	"CAP_AUDIT_WRITE":      unix.CAP_AUDIT_WRITE,
	"CAP_AUDIT_CONTROL":    unix.CAP_AUDIT_CONTROL,
	"CAP_SETFCAP":          unix.CAP_SETFCAP,
}

func SetFileCapabilities(path string, caps []string) {
	attr := "security.capability"
	version := uint32(0x02000001)

	effective := uint32(0)

	for _, v := range caps {
		key := strings.ToUpper(v)
		unixCap, exists := CapabilityMap[key]
		EnsureKeyExists(exists, v)
		effective |= 1 << unixCap
	}

	var buf bytes.Buffer
	addToBuf(&buf, version, effective, effective, 0, 0)
	data := buf.Bytes()
	log.Printf("Capabilities data: %v", data)

	err := unix.Setxattr(path, attr, data, 0)
	EnsureNoError(err, "Setxattr: %v")
}

func addToBuf(buf *bytes.Buffer, data ...uint32) {
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		EnsureNoError(err, "Add to bytes.Buffer error. %v")
	}
}
