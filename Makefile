LOCAL_CONFIG_MK = $(CURDIR)local.mk
ifneq ($(wildcard $(LOCAL_CONFIG_MK)),)
include $(LOCAL_CONFIG_MK)
endif

PKG_NAME = PackageCapability
CROSS_DIR = $(SPKSRC_DIR)/cross/$(PKG_NAME)
SPK_DIR = $(SPKSRC_DIR)/spk/$(PKG_NAME)

default:

setup: local.mk

local.mk:
	@echo "Creating local configuration \"local.mk\"..."
	@echo "SPKSRC_DIR = /toolkit/spksrc" > $@

clean:
	@echo "=====> run go clean <====="
	cd src && make clean

build:
	@echo "=====> run go build <====="
	cd src && make build

spksrc-clean: cross-clean spk-clean

cross-clean:
	@echo "remove $(CROSS_DIR)"
	@if [ -e $(CROSS_DIR) ]; then rm -rf $(CROSS_DIR); fi

cross-copy: cross-clean
	cp -av cross $(CROSS_DIR)

cross: cross-copy
	cd $(CROSS_DIR) && make clean && make arch-x64-7.1

digests: cross-copy
	@echo "generate digest"
	@cd $(CROSS_DIR) && make digests

spk-clean:
	@echo "remove $(SPK_DIR)"
	@if [ -e $(SPK_DIR) ]; then rm -rf $(SPK_DIR); fi

pkg-copy: spk-clean
	cp -aV spk $(SPK_DIR)

pkg: cross-copy pkg-copy
	cd $(SPK_DIR) && make clean && make arch-x64-7.1