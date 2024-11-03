#!/usr/bin/make -f
SHELL := $(shell command -v which >/dev/null 2>&1 && which sh || echo /bin/sh)
# Define the default install directory
PREFIX ?= /usr/local
VERSION ?= $(shell cat VERSION)
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git branch --show-current)
DATE := $(shell date -u +%Y-%m-%d)
PACKAGED ?= false
PACKAGEFORMAT ?=
NODOCUMENTATION ?= 0
NOTB ?= 0
EXTRALDFLAGS :=
EXTRAGOFLAGS :=
BUILDTAGS :=
UNAME_S := $(shell uname -s)
# Determine the system architecture using various methods
ARCH := $(shell \
    if command -v rpm > /dev/null 2>&1; then \
        rpm -E '%{_arch}'; \
    elif command -v dpkg > /dev/null 2>&1; then \
        dpkg --print-architecture; \
    elif command -v apk > /dev/null 2>&1; then \
        apk --print-arch; \
    elif command -v uname > /dev/null 2>&1; then \
        uname -m; \
    elif command -v arch > /dev/null 2>&1; then \
        arch; \
    elif command -v lscpu > /dev/null 2>&1; then \
        lscpu | grep Architecture | awk '{print $$2}'; \
    else \
        echo "unknown"; \
    fi)



ifneq ($(CFLAGS),)
    export CGO_CFLAGS := $(CFLAGS)
    $(info Using provided CFLAGS: $(CFLAGS))
endif

ifneq ($(CPPFLAGS),)
    export CGO_CPPFLAGS := $(CPPFLAGS)
    $(info Using provided CPPFLAGS: $(CPPFLAGS))
endif

ifneq ($(CXXFLAGS),)
    export CGO_CXXFLAGS := $(CXXFLAGS)
    $(info Using provided CXXFLAGS: $(CXXFLAGS))
endif

ifneq ($(LDFLAGS),)
    export CGO_LDFLAGS := $(LDFLAGS)
    $(info Using provided LDFLAGS: $(LDFLAGS))
endif

# Define target binary
TARGET = rokon

# Define install directories
BINDIR = $(DESTDIR)$(PREFIX)/bin
DATADIR = $(DESTDIR)$(PREFIX)/share/rokon
DOCDIR = $(DESTDIR)$(PREFIX)/share/doc/rokon
LICENSEDIR = $(DESTDIR)$(PREFIX)/share/licenses/rokon
APPLICATIONSDIR = $(DESTDIR)$(PREFIX)/share/applications
ICONDIR = $(DESTDIR)$(PREFIX)/share/icons/hicolor
METAINFODIR = $(DESTDIR)$(PREFIX)/share/metainfo
TARBALLDIR ?= ./tarball
SANITYCHECK ?= 1
RUNDIR ?= ./run
RUNLIBS ?= $(RUNDIR)/libs
ABS_RUNDIR := $(shell realpath $(RUNDIR))
# Check if selfextract exists in the PATH
SELFEXTRACT := $(shell command -v selfextract 2> /dev/null)

# Determine the selfextract path based on environment variables
ifneq ($(GOBIN),)
    SELFEXTRACT_PATH := $(GOBIN)/selfextract
else ifneq ($(GOPATH),)
    SELFEXTRACT_PATH := $(GOPATH)/bin/selfextract
else
    # Fallback to the safe default
    SELFEXTRACT_PATH := $(HOME)/go/bin/selfextract  # Safe default path
endif

# If selfextract command is found in the specified path
ifeq ($(wildcard $(SELFEXTRACT_PATH)),)
	SELFEXTRACT := $(SELFEXTRACT_PATH)
else
    SELFEXTRACT := $(SELFEXTRACT_PATH)
endif



TBPKGFMT ?= portable
ABS_TARBALLDIR := $(shell realpath $(TARBALLDIR))

TBLIBSDIR ?= $(TARBALLDIR)/libs
TAR_NAME ?= Rokon-$(UNAME_S)-$(VERSION)-$(ARCH).tar.gz
# Unix* users know .run is for them. DO NOT include it in the filename!
RUNFILE_NAME ?= Rokon-$(VERSION)-$(ARCH).run


make_wrapper_script = \
	echo '\#!/bin/sh' > $1/$(TARGET); \
	echo 'dir="$$(cd -P -- "$$(dirname -- "$$0")" && pwd -P)"' >> $1/$(TARGET); \
	echo 'cd "$$dir"' >> $1/$(TARGET); \
	echo 'export LD_LIBRARY_PATH="./libs:$$LD_LIBRARY_PATH"' >> $1/$(TARGET); \
	echo 'export LD_PRELOAD="./libs/libc.so.6"' >> $1/$(TARGET); \
	echo 'export XKB_DEFAULT_INCLUDE_PATH="./share/X11/xkb"' >> $1/$(TARGET); \
	echo 'export XKB_CONFIG_ROOT="./share/X11/xkb"' >> $1/$(TARGET); \
	echo 'if [ -e ./libs/ld*.so* ]; then exec ./libs/ld*.so* "./bin/$(TARGET)" "$$@"; else exec ./bin/$(TARGET) "$$@"; fi' >> $1/$(TARGET); \
	chmod +x $1/$(TARGET); \
	sed -i 's/rokon/\.\/$(TARGET)/g' $1/io.github.brycensranch.Rokon.desktop





copy_deps = \
	cp -L --no-preserve=mode -v $$(ldd $1 | grep -E '(^|[^a-zA-Z0-9])ld' | awk '{print $$1}') $2; \
	ldd $1 | awk '{print $$3}' | grep -v 'not found' | while read -r dep; do \
		if [ -n "$$dep" ]; then \
			cp -L --no-preserve=mode "$$dep" $2 || { echo "Failed to copy $$dep"; exit 1; }; \
		fi; \
		ldd "$$dep" | awk '{print $$3}' | grep -v 'not found' | while read -r subdep; do \
			if [ -n "$$subdep" ]; then \
				cp -L --no-preserve=mode "$$subdep" $2 || { echo "Failed to copy $$subdep"; exit 1; }; \
			fi; \
		done; \
	done; \
	chmod +x $2/*.so* || echo "Failed to set libraries executable! Oh well"; \
	strip --strip-all $2/*.so* || echo "Stripping libraries failed! Tarball *may* be larger than expected."


# Target to resolve dependencies
resolve:
	$(call copy_deps, $(TARGET))


.DEFAULT_GOAL := build
.PHONY: all
all: ## build pipeline
all: mod inst gen build tarball fatimage spell lint test

.PHONY: check
check: ## runs basic checks
check: spell lint test

.PHONY: tidy
tidy: ## go mod tidy
tidy: mod

.PHONY: precommit
precommit: ## validate the branch before commit
precommit: all vuln

.PHONY: ci
ci: ## CI checks pipeline
ci: precommit diff

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## remove files created during build pipeline
	$(call print-target)
	rm -rf dist .flatpak io.github.brycensranch.Rokon.desktop tarball io.github.brycensranch.Rokon.metainfo.xml macos/rokon .flatpak-builder flathub/builddir flathub/.flatpak-builder flathub/repo *.log *.zip modules.txt flathub/export macos/share flathub/*.flatpak AppDir src squashfs-root *.AppImage makeself* *.run run/ *.rpm *.pdf *.rtf windows/*.rtf *.deb *.msi *.exe pkg/ *.pkg.tar.zst .ptmp* *.tar* *.snap *.zsync rokon Rokon debian/tmp debian/rokon* *.changes *.buildinfo debian/.debhelper coverage.* '"$(shell go env GOCACHE)/../golangci-lint"'
	# go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: nuke
nuke: ## completely clean the repository of artifacts and clear cache
	$(call print-target)
	$(MAKE) clean
	go clean -i -cache -testcache -modcache -fuzzcache -x

.PHONY: version
version: ## software version e.g 1.0.0
	@echo "Version: $(VERSION)"

.PHONY: appimage
appimage: ## build AppImage using appimage-builder
	$(call print-target)
	@echo "Building AppImage version: $(VERSION)"
	rm -rf AppDir
	$(MAKE) PACKAGED=true PACKAGEFORMAT=AppImage EXTRAGOFLAGS="$(EXTRAGOFLAGS) -trimpath" EXTRALDFLAGS="$(EXTRALDFLAGS) -s -w" build
	$(MAKE) PREFIX=AppDir/usr BINDIR=AppDir install
	VERSION=$(VERSION) APPIMAGELAUNCHER_DISABLE=1 appimage-builder


# This was only added to not add duplicate version detection logic.
.PHONY: obsimage
obsimage: ## Turns AppDir into AppImage built on OpenSUSE Build Service
	$(call print-target)
	@echo "Building AppImage version: $(VERSION)"
	rm -rf AppDir
	mv AppDir/usr/share/metainfo/io.github.brycensranch.Rokon.metainfo.xml AppDir/usr/share/metainfo/io.github.brycensranch.Rokon.appdata.xml
	cp ./AppDir/usr/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png ./AppDir
	APPIMAGELAUNCHER_DISABLE=1 NO_STRIP=true linuxdeploy --appdir=AppDir --output appimage

.PHONY: fatimage
fatimage: ## build self contained AppImage that can run on older Linux systems while CI is on development branch
	$(call print-target)
	@echo "Building AppImage version: $(VERSION) (FAT)"
	rm -rf AppDir
	$(MAKE) PACKAGED=true PACKAGEFORMAT=AppImage EXTRAGOFLAGS="$(EXTRAGOFLAGS) -trimpath" EXTRALDFLAGS="$(EXTRALDFLAGS) -s -w" build
	$(MAKE) PREFIX=AppDir/usr install
	VERSION=$(VERSION) APPIMAGELAUNCHER_DISABLE=1 appimagetool -s deploy ./AppDir/usr/share/applications/io.github.brycensranch.Rokon.desktop
	@if command -v upx > /dev/null; then \
		echo "UPX found. Compressing binaries..."; \
		upx --best --lzma -v AppDir/usr/bin/$(TARGET) || echo "Failed to compress $(TARGET) binary."; \
	else \
		echo "UPX not found. Skipping compression."; \
	fi
	# My application follows the https://docs.fedoraproject.org/en-US/packaging-guidelines/AppData/ but this tool doesn't care lol
	mv AppDir/usr/share/metainfo/io.github.brycensranch.Rokon.metainfo.xml AppDir/usr/share/metainfo/io.github.brycensranch.Rokon.appdata.xml
	cp ./AppDir/usr/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png ./AppDir
	VERSION=$(VERSION) APPIMAGELAUNCHER_DISABLE=1 mkappimage --comp zstd --ll -u "gh-releases-zsync|BrycensRanch|Rokon|latest|Rokon-*$(ARCH).AppImage.zsync" ./AppDir

.PHONY: basedimage
basedimage: ## create AppImage from existing tarball directory
	$(call print-target)
	cp $(TARBALLDIR)/$(TARGET) $(TARBALLDIR)/AppRun
	mkdir -p $(TARBALLDIR)/usr/share/metainfo
	mkdir -p $(TARBALLDIR)/usr/share/applications
	cp $(TARBALLDIR)/io.github.brycensranch.Rokon.desktop $(TARBALLDIR)/usr/share/applications
	cp $(TARBALLDIR)/share/metainfo/io.github.brycensranch.Rokon.metainfo.xml $(TARBALLDIR)/usr/share/metainfo/io.github.brycensranch.Rokon.appdata.xml
	cp usr/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png $(TARBALLDIR)
	VERSION=$(VERSION) APPIMAGELAUNCHER_DISABLE=1 mkappimage --comp zstd --ll -u "gh-releases-zsync|BrycensRanch|Rokon|latest|Rokon-*$(ARCH).AppImage.zsync" $(TARBALLDIR)

.ONESHELL:
.PHONY: tarball
tarball: ## build self contained Tarball that auto updates
	$(call print-target)
	@echo "Building Rokon Tarball version: $(VERSION)"
	@echo "This process requires the following command line utils to work properly: awk, ldd, tar"
	@echo "SHELL: $(SHELL) ARCH: $(ARCH)"
	rm -rf $(TARBALLDIR) || sudo rm -v -rf $(TARBALLDIR)
	mkdir -p $(TARBALLDIR)
	mkdir -p $(TBLIBSDIR)
	$(MAKE) PACKAGED=true PACKAGEFORMAT=$(TBPKGFMT) EXTRAGOFLAGS="$(EXTRAGOFLAGS) -trimpath" EXTRALDFLAGS="$(EXTRALDFLAGS) -s -w -linkmode=external" build
	$(MAKE) PREFIX=$(TARBALLDIR) APPLICATIONSDIR=$(TARBALLDIR) install
	cp ./windows/portable.txt $(TARBALLDIR)
	$(call copy_deps,$(TARBALLDIR)/bin/$(TARGET),$(TBLIBSDIR))
	$(call make_wrapper_script,$(TARBALLDIR))
	@if command -v upx > /dev/null; then \
		echo "UPX found. Compressing binaries..."; \
		upx --best --lzma -v $(TARBALLDIR)/bin/$(TARGET) || echo "Failed to compress $(TARGET) binary."; \
	else \
		echo "UPX not found. Skipping compression."; \
	fi
	cd /usr && cp -r --parents -L --no-preserve=mode -r share/glib-2.0/schemas/gschemas.compiled share/X11 share/gtk-4.0 share/icons/Adwaita $(ABS_TARBALLDIR)
	cd -
	rm -rf $(TARBALLDIR)/share/gtk-4.0/emoji || true
	@if [ "$(SANITYCHECK)" == "1" ]; then \
		LD_DEBUG=libs $(TARBALLDIR)/$(TARGET) --version; \
		status=$$?; \
		if [ $$status -ne 0 ]; then \
			echo "Sanity check failed. See output above for details."; \
			exit $$status; \
		else \
			echo "Sanity check succeeded."; \
		fi; \
	else \
		echo "Sanity check skipped."; \
	fi

ifeq ($(NOTB),1)
		@echo "Finished making tarball directory. You have specified that a tarball shouldn't be created with NOTB=1"
else
		tar -czf $(TAR_NAME) $(TARBALLDIR)
		@if command -v zsyncmake >/dev/null 2>&1; then \
			zsyncmake $(TAR_NAME) -u "gh-releases-zsync|BrycensRanch|Rokon|latest|$(TAR_NAME).zsync"; \
		else \
			echo "zsyncmake not found. Please install it to generate the zsync file."; \
		fi
		rm $(TARGET)
		@echo "Tarball created: $(TAR_NAME)"
endif

.ONESHELL:
.PHONY: check_selfextract
check_selfextract:
	@if [ ! -x "$(SELFEXTRACT_PATH)" ]; then \
		echo "selfextract command not found in $(SELFEXTRACT_PATH), installing..."; \
		cd tools && go get github.com/synthesio/selfextract && go install -v github.com/synthesio/selfextract; \
		cd -; \
		echo "selfextract installed..."; \
	else \
		echo "Using selfextract located at: $(SELFEXTRACT_PATH)"; \
	fi

.ONESHELL:
.PHONY: run
run: ## create run "package"
	$(call print-target)
	$(MAKE) check_selfextract
	rm $(RUNFILE_NAME) || true
	$(MAKE) PACKAGED=true PACKAGEFORMAT="run" TBPKGFMT="run" TARBALLDIR=$(RUNDIR) NOTB=1 tarball
	cp $(RUNDIR)/$(TARGET) $(RUNDIR)/selfextract_startup
	$(SELFEXTRACT) -f $(RUNFILE_NAME) -C $(RUNDIR) .
	@if [ "$(SANITYCHECK)" == "1" ]; then \
		./$(RUNFILE_NAME) --version; \
		status=$$?; \
		if [ $$status -ne 0 ]; then \
			echo "Secondary Sanity check failed. See output above for details."; \
			exit $$status; \
		else \
			echo "Secondary Sanity check succeeded."; \
		fi; \
	else \
		echo "Secondary Sanity check skipped."; \
	fi
	@if command -v zsyncmake >/dev/null 2>&1; then \
		zsyncmake $(RUNFILE_NAME) -u "gh-releases-zsync|BrycensRanch|Rokon|latest|$(RUNFILE_NAME).zsync"; \
	else \
		echo "zsyncmake not found. Please install it to generate the zsync file."; \
	fi
	@echo "Cheers, the run file was successfully created. It is the file ./$(RUNFILE_NAME) 🚀"

.PHONY: dev
dev: ## go run -v .
	$(call print-target)
	@echo "Starting development server for Rokon: $(VERSION)"
	go run -v .

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	go mod tidy
	cd tools && go mod tidy

.PHONY: inst
inst: ## go install tools
	$(call print-target)
	cd tools && go get $(shell cd tools && go list -e -f '{{ join .Imports " " }}' -tags=tools)

.ONESHELL:
.PHONY: install
install: ## installs Rokon into $PATH and places desktop files
	$(call print-target)
	@echo "Version: $(VERSION)"
	@echo "Creating necessary directories..."
	mkdir -p $(BINDIR)
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/rokon
	mkdir -p $(APPLICATIONSDIR)
	mkdir -p $(LICENSEDIR)
	mkdir -p $(ICONDIR)/48x48/apps
	mkdir -p $(ICONDIR)/128x128/apps
	mkdir -p $(ICONDIR)/256x256/apps
	mkdir -p $(ICONDIR)/scalable/apps
	mkdir -p $(DESTDIR)$(PREFIX)/share/dbus-1/services
	mkdir -p $(METAINFODIR)
	@echo "Detected OS: $(UNAME_S)"

ifeq ($(UNAME_S),Darwin)
	install -m 0755 $(TARGET) $(BINDIR)
else
	install -Dpm 0755 $(TARGET) $(BINDIR)
endif

ifeq ($(UNAME_S),Darwin)
		install -m 0644 ./usr/share/applications/io.github.brycensranch.Rokon.desktop $(APPLICATIONSDIR)/io.github.brycensranch.Rokon.desktop
		install -m 0644 ./usr/share/dbus-1/services/io.github.brycensranch.Rokon.service $(DESTDIR)$(PREFIX)/share/dbus-1/services/io.github.brycensranch.Rokon.service
		install -m 0644 ./usr/share/icons/hicolor/48x48/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/48x48/apps/io.github.brycensranch.Rokon.png
		install -m 0644 ./usr/share/icons/hicolor/128x128/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/128x128/apps/io.github.brycensranch.Rokon.png
		install -m 0644 ./usr/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/256x256/apps/io.github.brycensranch.Rokon.png
		install -m 0644 ./usr/share/icons/hicolor/scalable/apps/io.github.brycensranch.Rokon.svg $(ICONDIR)/scalable/apps/io.github.brycensranch.Rokon.svg
		install -m 0644 ./usr/share/metainfo/io.github.brycensranch.Rokon.metainfo.xml $(METAINFODIR)/io.github.brycensranch.Rokon.metainfo.xml
		install -m 0644 ./LICENSE.md $(LICENSEDIR)/LICENSE.md;
else
		install -Dpm 0644 ./usr/share/applications/io.github.brycensranch.Rokon.desktop $(APPLICATIONSDIR)/io.github.brycensranch.Rokon.desktop
		install -Dpm 0644 ./usr/share/dbus-1/services/io.github.brycensranch.Rokon.service $(DESTDIR)$(PREFIX)/share/dbus-1/services/io.github.brycensranch.Rokon.service
		install -Dpm 0644 ./usr/share/icons/hicolor/48x48/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/48x48/apps/io.github.brycensranch.Rokon.png
		install -Dpm 0644 ./usr/share/icons/hicolor/128x128/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/128x128/apps/io.github.brycensranch.Rokon.png;
		install -Dpm 0644 ./usr/share/icons/hicolor/256x256/apps/io.github.brycensranch.Rokon.png $(ICONDIR)/256x256/apps/io.github.brycensranch.Rokon.png
		install -Dpm 0644 ./usr/share/icons/hicolor/scalable/apps/io.github.brycensranch.Rokon.svg $(ICONDIR)/scalable/apps/io.github.brycensranch.Rokon.svg
		install -Dpm 0644 ./usr/share/metainfo/io.github.brycensranch.Rokon.metainfo.xml $(METAINFODIR)/io.github.brycensranch.Rokon.metainfo.xml
		install -Dpm 0644 ./LICENSE.md $(LICENSEDIR)/LICENSE.md
endif

	@if [ "$(NODOCUMENTATION)" != "1" ]; then \
		echo "Installing documentation (PRIVACY.md, README.md) to $(DESTDIR)$(PREFIX)/share/doc/rokon"; \
		if [ "$(UNAME_S)" = "Darwin" ]; then \
			install -m 0644 ./PRIVACY.md ./README.md $(DESTDIR)$(PREFIX)/share/doc/rokon; \
		else \
			install -Dpm 0644 ./PRIVACY.md ./README.md $(DESTDIR)$(PREFIX)/share/doc/rokon; \
		fi; \
	else \
		echo "Skipping documentation installation. Please make sure you include PRIVACY notice."; \
	fi


.PHONY: uninstall
uninstall:
	$(call print-target)
	@echo "Uninstalling version $(VERSION)"
	rm -f $(BINDIR)/$(TARGET)
	rm -f $(APPLICATIONSDIR)/io.github.brycensranch.Rokon.desktop
	rm -f $(ICONDIR)/48x48/apps/io.github.brycensranch.Rokon.png
	rm -f $(ICONDIR)/128x128/apps/io.github.brycensranch.Rokon.png
	rm -f $(ICONDIR)/256x256/apps/io.github.brycensranch.Rokon.png
	rm -f $(ICONDIR)/scalable/apps/io.github.brycensranch.Rokon.svg
	rm -f $(METAINFODIR)/io.github.brycensranch.Rokon.metainfo.xml
	rm -rf $(DOCDIR)/rokon
	rm -f $(DESTDIR)$(PREFIX)/share/dbus-1/services/io.github.brycensranch.Rokon.service
	rm -rf $(DESTDIR)$(PREFIX)/share/licenses/rokon

.PHONY: gen
gen: ## go generate
	$(call print-target)
	go generate ./...

.PHONY: build
build: ## go build -v -o rokon
	$(call print-target)
	@echo "Building version $(VERSION) commit $(COMMIT) on branch $(BRANCH)"
	go build -v -ldflags="-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.packaged=$(PACKAGED) -X main.packageFormat=$(PACKAGEFORMAT) -X main.branch=$(BRANCH) $(EXTRALDFLAGS)" $(EXTRAGOFLAGS) -o $(TARGET) -tags "$(BUILDTAGS)" .

.PHONY: spell
spell: ## misspell
	$(call print-target)
	misspell -error -locale=US -w **.md

.PHONY: lint
lint: ## golangci-lint
	$(call print-target)
	golangci-lint run --fix

.PHONY: vuln
vuln: ## govulncheck
	$(call print-target)
	govulncheck ./...

.PHONY: test
test: ## go test
	$(call print-target)
	go test -race -covermode=atomic -coverprofile=coverage.out -coverpkg=./... ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: diff
diff: ## git diff
	$(call print-target)
	git diff --exit-code
	RES=$$(git status --porcelain) ; if [ -n "$$RES" ]; then echo $$RES && exit 1 ; fi

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
