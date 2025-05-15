LOCAL_CONFIG_MK = $(CURDIR)/local.mk
ifneq ($(wildcard $(LOCAL_CONFIG_MK)),)
include $(LOCAL_CONFIG_MK)
endif

PKG_NAME = PackageCapability
SPK_VERS = 1.1.0
SPK_REV = 2
CROSS_DIR = $(SPKSRC_DIR)/cross/$(PKG_NAME)
SPK_DIR = $(SPKSRC_DIR)/spk/$(PKG_NAME)
PKG_DIR = $(SPKSRC_DIR)/packages
PKG_DIST_NAME = $(PKG_NAME)_$(ARCH)-$(TCVERSION)_$(SPK_VERS)-$(SPK_REV).spk

DEFAULT_ARCH = x64
DEFAULT_TCVERSION = 7.0

default: package

setup: local.mk

local.mk:
	@echo "Creating local configuration \"local.mk\"..."
	@echo "SPKSRC_DIR = /toolkit/spksrc" > $@
	@echo "ARCH = $(DEFAULT_ARCH)" >> $@
	@echo "TCVERSION = $(DEFAULT_TCVERSION)" >> $@

update-version:
	@echo "=====> run update version <====="
	@sed -i 's/SPK_VERS = .*/SPK_VERS = $(SPK_VERS)/g' spk/Makefile
	@sed -i 's/SPK_REV = .*/SPK_REV = $(SPK_REV)/g' spk/Makefile

clean:
	@echo "=====> run go clean <====="
	cd src && make clean

build:
	@echo "=====> run go build <====="
	cd src && make build

spk-clean:
	@echo "remove $(SPK_DIR)"
	@if [ -e $(SPK_DIR) ]; then rm -rf $(SPK_DIR); fi

spk-copy: spk-clean build
	cp -a spk $(SPK_DIR)
	cp -a src/bin $(SPK_DIR)/bin

spk: update-version spk-copy
	cd $(SPK_DIR) && make clean && make arch-$(ARCH)-$(TCVERSION)

package: spk
	@mkdir -p packages
	@cp -a $(PKG_DIR)/$(PKG_DIST_NAME) packages/
