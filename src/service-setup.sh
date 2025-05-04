SCMSET_BIN=/usr/local/bin/scmset
SCMUPDATE_BIN=usr/local/bin/scmupdate
SCMUPDATE_TARGET=${SYNOPKG_PKGDEST}/${SCMUPDATE_BIN}

service_postinst() {
    echo "SYNOPKG_PKG_STATUS: ${SYNOPKG_PKG_STATUS}"
    if [ "${SYNOPKG_PKG_STATUS}" == "INSTALL" -o "${SYNOPKG_PKG_STATUS}" == "UPGRADE" ]; then
        # 为 scm-update 设置 cap
        if [ -f ${SCMSET_BIN} ]; then
            echo "run command: ${SCMSET_BIN} scm ${SCMUPDATE_BIN} CAP_SETFCAP,CAP_DAC_OVERRIDE"
            ${SCMSET_BIN} scm ${SCMUPDATE_BIN} CAP_SETFCAP,CAP_DAC_OVERRIDE 2>&1
            
            # 如果执行失败，则退出
            if [ $? -eq 0 ]; then
                # 运行 scm-update
                echo "run command: ${SCMUPDATE_TARGET}"
                ${SCMUPDATE_TARGET} 2>&1
            fi
        fi
    fi
}
