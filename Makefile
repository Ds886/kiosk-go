DOCS = man/kiosk-go.1 man/kiosk.toml.5
TARGETS_OBJ = bin/kiosk-go ${DOCS}

DESTDIR ?= /
PREFIX ?= usr/local
BINDIR = ${DESTDIR}${PREFIX}/bin
MAN1DIR = ${DESTDIR}${PREFIX}/share/man1
MAN5DIR = ${DESTDIR}${PREFIX}/share/man5
ETCDIR ?= ${DESTDIR}${PREFIX}/etc

GOARCH ?= amd64
GOOS ?= linux
GOOPTS = 

GO ?= $(shell which go)
INSTALL ?= $(shell which install)

.PHONY: clean uninstall install release

all: build

build: ${TARGETS_OBJ}

doc: ${DOCS}

clean:
	${RM} ${TARGETS_OBJ}
	${RM} -r ${PWD}/deb/usr
	${RM} -r ${PWD}/deb/etc

run:
	./bin/kiosk-go

install: ${TARGET_OBJ}
	$(INSTALL) -D bin/kiosk-go ${BINDIR}/kiosk-go
	$(INSTALL) -Dm 644 man/kiosk-go.1 ${MAN1DIR}/kiosk-go.1
	$(INSTALL) -Dm 644 man/kiosk.toml.5 ${MAN5DIR}/kiosk.toml.5
	${INSTALL} -Dm 644 res/kiosk.toml ${ETCDIR}/kiosk.toml

uninstall:
	${RM} ${BINDIR}/kiosk-go
	${RM} ${MAN1DIR}/kiosk-go.1
	${RM} ${MAN5DIR}/kiosk.toml.5
	${RM} ${ETCDIR}/kiosk.toml

man/%: man/%.md
	go-md2man -in $< -out $@

bin/kiosk-go:
	GOARCH=${GOARCH} GOOS=${GOOS} ${GO} build  ${GOOPTS} -o $@

kiosk-go.deb: prefix=${PWD}/deb/usr
kiosk-go.deb: DESTDIR=${PWD}/deb
kiosk-go.deb: clean release install
	${RM} $@
	${RM} -r deb/share
	mkdir -p ${DESTDIR}/usr/share/doc/kiosk-go
	echo "Apache-2" > ${PWD}/deb/usr/share/doc/kiosk-go/copyright
	touch ${prefix}/share/doc/kiosk-go/changelog
	echo "Not implemented" ${prefix}/share/doc/kiosk-go/changelog
	gzip -n -9 ${prefix}/share/doc/kiosk-go/changelog
	chmod 755 deb/usr/share/doc
	chmod 755 deb/usr/share/doc/kiosk-go
	chmod 644 deb/usr/share/doc/kiosk-go/*
	dpkg-deb --build --root-owner-group deb $@

deb: kiosk-go.deb

release: GOOPTS += -ldflags="-s -w"
release: build
