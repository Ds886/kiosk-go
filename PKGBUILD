pkgname=kiosk-go
pkgver=1.1
pkgrel=1
pkgdesc="An implemntation in go for running a single app kiosk"
arch=(any)
url="https://github.com/Ds886/kiosk-go"
license=('APL')
depends=()
makedepends=("go" "go-md2man")
source=("https://github.com/Ds886/kiosk-go/archive/$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('ab1c6b1b4dadb09c0a45051386552f264954a8944ddbf2b1baf99385cc0bf24a')

prepare() {
	cd "$srcdir/$pkgname-$pkgver"
}

build() {
	cd "$srcdir/$pkgname-$pkgver"
	make DESTDIR="$pkgdir"
}

check() {
	cd "$srcdir/$pkgname-$pkgver"
}

package() {
	cd "$srcdir/$pkgname-$pkgver"
	make DESTDIR="$pkgdir/" PREFIX="usr"  install
}
