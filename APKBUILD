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
source="https://nightly.link/BrycensRanch/Rokon/actions/runs/11875401640/rokon-vendored-source.zip"


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
a8f99e1b8d2cebc6e2dc10c9e043432901d992bff89afc4ae96a4e8d35c45006948838b129c3cb45ddaab99c5bd5127f6ce50b9a248e98348fdd5ccebb901bfe rokon-vendored-source.zip
"
