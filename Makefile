LOCAL_CONFIG_MK = $(CURDIR)/local.mk
ifneq ($(wildcard $(LOCAL_CONFIG_MK)),)
include $(LOCAL_CONFIG_MK)
endif

PKG_NAME = PackageCapability
CROSS_DIR = $(SPKSRC_DIR)/cross/$(PKG_NAME)
SPK_DIR = $(SPKSRC_DIR)/spk/$(PKG_NAME)

default: spk

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

spk-clean:
	@echo "remove $(SPK_DIR)"
	@if [ -e $(SPK_DIR) ]; then rm -rf $(SPK_DIR); fi

spk-copy: spk-clean build
	cp -a spk $(SPK_DIR)
	cp -a src/bin $(SPK_DIR)/bin

spk: spk-copy
	cd $(SPK_DIR) && make clean && make arch-x64-7.1
