SPK_NAME = PackageCapability
SPK_VERS = 1.0.0
SPK_REV = 1
# SPK_ICON = src/scm.png

DEPENDS = 

MAINTAINER = wangbin1989
MAINTAINER_URL = https://github.com/wangbin1989

DISPLAY_NAME = PackageCapability
CHANGELOG = "Initial package release."
STARTABLE = no

LICENSE  = GPLv3

REQUIRED_MIN_DSM = 7.0

SPK_DEPENDS =

SERVICE_USER = auto
SERVICE_SETUP = src/service-setup.sh

# CONF_DIR = src/conf/
# DSM_UI_DIR = app
# DSM_UI_CONFIG = src/app/config

COPY_TARGET = nop
POST_STRIP_TARGET = PackageCapability_custom_post_strip

include ../../mk/spksrc.spk.mk

.PHONY: PackageCapability_custom_post_strip
PackageCapability_custom_post_strip:
	@mkdir -p $(STAGING_DIR)/usr/local/bin
	@cp -av bin/* $(STAGING_DIR)/usr/local/bin
	@cp -av src/first-install.sh $(STAGING_DIR)/first-install.sh
