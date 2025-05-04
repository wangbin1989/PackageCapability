// 程序需要CAP: CAP_SETFCAP CAP_DAC_OVERRIDE CAP_CHOWN

package main

import (
	"log"
	"os"

	"github.com/wangbin1989/scm/src/common"
)

func main() {
	scmset_src := "/var/packages/scm/target/usr/local/bin/scmset"
	scmset_bin := "/usr/local/bin/scmset"

	var err error

	log.Printf("移除文件 %s.", scmset_bin)
	err = os.Remove(scmset_bin)
	common.EnsureNoError(err, "移除文件失败. %v")

	file, err := os.ReadFile(scmset_src)
	log.Printf("读取文件 %s.", scmset_bin)
	common.EnsureNoError(err, "读取文件失败. %v")

	err = os.WriteFile(scmset_bin, file, 0755)
	log.Printf("写入文件 %s.", scmset_bin)
	common.EnsureNoError(err, "写入文件失败. %v")

	// err = os.Chown(scmset_bin, 0, 0)
	// common.EnsureNoError(err, "修改文件所有者失败. %v")

	// err = os.Chmod(scmset_bin, 6755)
	// common.CheckError(err, "修改文件权限失败. %v")

	log.Printf("设置文件CAP %s.", scmset_bin)
	common.SetFileCapabilities(scmset_bin, []string{"CAP_SETFCAP", "CAP_DAC_OVERRIDE"})
}
