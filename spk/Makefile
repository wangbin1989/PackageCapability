SPK_NAME = PackageCapability
SPK_VERS = 1.1.0
SPK_REV = 2
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

COPY_TARGET = nop
POST_STRIP_TARGET = PackageCapability_custom_post_strip

include ../../mk/spksrc.spk.mk

.PHONY: PackageCapability_custom_post_strip
PackageCapability_custom_post_strip:
	@mkdir -p $(STAGING_DIR)/bin
	@cp -av bin/* $(STAGING_DIR)/bin
	@cp -av src/first-install.sh $(STAGING_DIR)/first-install.sh
