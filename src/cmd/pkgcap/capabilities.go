package main

import (
	"log"

	"github.com/wangbin1989/PackageCapability/common"
)

var allowedCapabilities = map[string][]string{
	"PackageCapability": {"CAP_DAC_OVERRIDE", "CAP_SETFCAP"},
	"AdGuardHome":       {"CAP_NET_BIND_SERVICE", "CAP_NET_RAW"},
}

func checkCapabilities(packageName string, caps []string) {
	// 检查权限列表
	allowedCaps, exists := allowedCapabilities[packageName]
	if !exists {
		log.Fatalf("Package [%s] is not allowed.", packageName)
	}

	sourceMap := make(map[string]bool)
	for _, cap := range allowedCaps {
		sourceMap[cap] = true
	}

	for _, cap := range caps {
		v, e := sourceMap[cap]

		if !e || !v {
			log.Fatalf("Capability [%s] is not allowed.", cap)
		}
	}
}

func setFileCapabilities(path string, caps []string) {
	common.SetFileCapabilities(path, caps)
}
