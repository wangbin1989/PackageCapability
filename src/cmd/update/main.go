// 程序需要CAP: CAP_SETFCAP CAP_DAC_OVERRIDE CAP_CHOWN

package main

import (
	"log"
	"os"

	"github.com/wangbin1989/PackageCapability/common"
)

func main() {
	pkgcap_src := "/var/packages/PackageCapability/target/bin/pkgcap"
	pkgcap_bin := "/usr/local/bin/pkgcap"

	var err error

	log.Printf("移除文件 %s.", pkgcap_bin)
	err = os.Remove(pkgcap_bin)
	common.EnsureNoError(err, "移除文件失败. %v")

	log.Printf("读取文件 %s.", pkgcap_bin)
	file, err := os.ReadFile(pkgcap_src)
	common.EnsureNoError(err, "读取文件失败. %v")

	log.Printf("写入文件 %s.", pkgcap_bin)
	err = os.WriteFile(pkgcap_bin, file, 0755)
	common.EnsureNoError(err, "写入文件失败. %v")

	log.Printf("设置文件CAP %s.", pkgcap_bin)
	common.SetFileCapabilities(pkgcap_bin, []string{"CAP_SETFCAP", "CAP_DAC_OVERRIDE"})
	common.EnsureNoError(err, "设置文件CAP失败. %v")
}
