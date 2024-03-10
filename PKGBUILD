
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
source=("git+${url}")
sha256sums=("SKIP")

package() {
    install -Dm755 "$srcdir/i3-autotiler/src/i3-autotiler" "$pkgdir/usr/bin/i3-autotiler"
    install -Dm644 "$srcdir/i3-autotiler/src/i3-autotiler.service" "$pkgdir/usr/lib/systemd/user/i3-autotiler.service"
}

post_install() {
    systemctl enable --user --now i3-autotiler.service
}

post_remove() {
    systemctl disable -user --now i3-autotiler.service
}
