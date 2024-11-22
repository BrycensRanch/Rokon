# Contributor: Brycen G <brycengranville@outlook.com>
# Maintainer: Brycen G <brycengranville@outlook.com>
pkgname=rokon
pkgver=1.0.0
pkgrel=0
pkgdesc="A roku remote for your desktop"
url="https://github.com/BrycensRanch/Rokon"
arch="all"
license="AGPL-3.0-or-later"
makedepends="
    gtk4.0-dev
    gobject-introspection-dev
    go
    build-base
    unzip
	"
subpackages="
	$pkgname-doc
	"
source="https://nightly.link/BrycensRanch/Rokon/actions/runs/11943137586/rokon-vendored-source.zip"


build() {
    make TARGET=$pkgname PACKAGED=true PACKAGEFORMAT=apk EXTRAGOFLAGS="-buildmode=pie -trimpath -mod=vendor -modcacherw" EXTRALDFLAGS="-compressdwarf=false -linkmode=external" build
}

check() {
    ./rokon --version
}

package() {
    cd src
    make DESTDIR="$pkgdir" PREFIX="/usr" install
}

sha512sums="
6e1f3fffc4504ab9fd63d3820476d7f6d14f39323738cabe0423deb99146011a8e99bbf59d7549bbffb56f297b45ade0bc6f8595bc8e8712e0e219bcaf5b3d09 rokon-vendored-source.zip
"
