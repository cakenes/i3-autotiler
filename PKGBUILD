
# PKGBUILD for i3-autotiler-git
# Maintainer: Cakenes <me@cakenes.com>

pkgname="i3-autotiler-git"
pkgver=1.0
pkgrel=1
pkgdesc="Automatically tile windows in i3wm"
arch=('x86_64')
url="https://github.com/cakenes/i3-autotiler"
license=("MIT")
depends=("i3-wm")
source=("https://github.com/cakenes/i3-autotiler/raw/main/src/i3-autotiler")
sha256sums=("SKIP")

package() {
    install -Dm755 "i3-autotiler" "$pkgdir/usr/bin/i3-autotiler"
}
