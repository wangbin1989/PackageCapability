#!/bin/bash

# 此命令需要以 root 权限运行

scmset_src=/var/packages/scm/target/usr/local/bin/scmset
scmset_bin=/usr/local/bin/scmset
rm -f ${scmset_bin}
cp ${scmset_src} ${scmset_bin}
# chown root:root ${scmset_bin}
chmod 755 ${scmset_bin}
setcap CAP_SETFCAP,CAP_DAC_OVERRIDE=ep ${scmset_bin}
