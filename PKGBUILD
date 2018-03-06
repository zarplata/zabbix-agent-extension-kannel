pkgname=zabbix-agent-extension-kannel
pkgver=1
pkgrel=1
pkgdesc="Zabbix agent for Kannel stats."
arch=('any')
license=('GPL')
makedepends=('go')
depends=()
install='install.sh'
source=("git+http://a.kitsul@git.rn/scm/~a.kitsul/zabbix-agent-extension-kannel.git#branch=dev")
md5sums=('SKIP')

pkgver() {
    cd "$srcdir/$pkgname"

    make ver
}
    
build() {
    cd "$srcdir/$pkgname"

    make
}

package() {
    cd "$srcdir/$pkgname"

    install -Dm 0755 .out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 "${pkgname}.conf" "${pkgdir}/etc/zabbix/zabbix_agentd.conf.d/${pkgname}.conf"
}

