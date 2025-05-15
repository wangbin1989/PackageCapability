package main

import (
	"log"
	"strings"

	"github.com/wangbin1989/PackageCapability/common"
)

func UpdateCapabilities(packageName, relPath, capabilities string) {
	privilegePath := getPrivilegePath(packageName)
	common.EnsureFileExists(privilegePath)

	cmdPath := "/var/packages/" + packageName + "/target/" + relPath
	common.EnsureFileExists(cmdPath)

	capArray := strings.Split(capabilities, ",")
	checkCapabilities(packageName, capArray)

	log.Printf("修改 privilege 文件. %s", privilegePath)
	updatePrivilegeFile(privilegePath, relPath, capabilities)

	log.Printf("设置文件CAP. %s", cmdPath)
	setFileCapabilities(cmdPath, capArray)
}
