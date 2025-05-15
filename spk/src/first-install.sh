#!/bin/bash

# 此命令需要以 root 权限运行

pkgcap_src=/var/packages/PackageCapability/target/usr/local/bin/pkgcap
pkgcap_bin=/usr/local/bin/pkgcap
rm -f ${pkgcap_bin}
cp ${pkgcap_src} ${pkgcap_bin}
# chown root:root ${pkgcap_bin}
chmod 755 ${pkgcap_bin}
setcap CAP_SETFCAP,CAP_DAC_OVERRIDE=ep ${pkgcap_bin}
