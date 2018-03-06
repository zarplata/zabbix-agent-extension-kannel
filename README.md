# zabbix-agent-extension-kannel

zabbix-agent-extension-kannel - this extension for monitoring kannel statistic.

### Supported features

This extension obtains stats:

#### Kannel stats

- [x] status
- [ ] uptime

#### WDP stats

- [ ] received msg
- [ ] queued msg

#### SMS stats

- [x] sent msg
- [ ] sent queued msg
- [x] recv msg
- [ ] recv queued msg
- [x] store msg
- [ ] inbound rate
- [ ] outbound rate

#### DLR stats

- [ ] queued msg
- [ ] storage

#### Boxes stats

- [ ] type
- [ ] id
- [ ] IP
- [ ] queue
- [ ] status
- [ ] ssl

#### SMSC stats

- [ ] count

##### for each provider

- [x] ID
- [x] status
- [x] uptime
- [x] recv msg
- [x] sent msg
- [x] failed msg
- [x] queued msg


### Installation

#### Manual build

```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-kannel.git
cd zabbix-agent-extension-kannel
make

#Installing
make install

# By default, binary installs into /usr/bin/ and zabbix config in /etc/zabbix/zabbix_agentd.conf.d/ but,
# you may manually copy binary to your executable path and zabbix config to specific include directory
```

#### Arch Linux package
```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-kannel.git
git checkout pkgbuild

makepkg

#Installing
pacman -U *.tar.xz
```

### Dependencies

zabbix-agent-extension-kannel requires [zabbix-agent](http://www.zabbix.com/download) v2.4+ to run.

### Zabbix configuration
In order to start getting metrics, it is enough to import template and attach it to monitored node.

You should redefine macro `{$KANNEL_URL}` in global (template) or local (host) scope with url `http://127.0.0.1:13000/status.xml`

`WARNING:` You must define macro with name - `{$ZABBIX_SERVER_IP}` in global or local (template) scope with IP address of zabbix server.
