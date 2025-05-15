PKGCAP_BIN=/usr/local/bin/pkgcap
UPDATE_BIN=usr/local/bin/update
UPDATE_TARGET=${SYNOPKG_PKGDEST}/${UPDATE_BIN}

service_postinst() {
    echo "SYNOPKG_PKG_STATUS: ${SYNOPKG_PKG_STATUS}"
    if [ "${SYNOPKG_PKG_STATUS}" == "INSTALL" -o "${SYNOPKG_PKG_STATUS}" == "UPGRADE" ]; then
        # 为 update 设置 cap
        if [ -f ${PKGCAP_BIN} ]; then
            echo "run command: ${PKGCAP_BIN} PackageCapability ${UPDATE_BIN} CAP_SETFCAP,CAP_DAC_OVERRIDE"
            ${PKGCAP_BIN} PackageCapability ${UPDATE_BIN} CAP_SETFCAP,CAP_DAC_OVERRIDE 2>&1

            # 如果执行失败，则退出
            if [ $? -eq 0 ]; then
                # 运行 update
                echo "run command: ${UPDATE_TARGET}"
                ${UPDATE_TARGET} 2>&1
            fi
        fi
    fi
}
