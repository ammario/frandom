# frandom

Fast random

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [frandom](#frandom)
  - [Algorithm](#algorithm)
  - [Install](#install)
  - [As a package](#as-a-package)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

`/dev/urandom` is *slow*


/dev/urandom:

```bash
$ time head -c 1G < /dev/urandom > /dev/null

real	1m37.754s
user	0m0.048s
sys	1m37.684s
```

frandom:

```bash
$ time frandom | head -c 1G > /dev/null

real	0m2.704s
user	0m2.192s
sys	0m2.332s
```

__That's 10.3 mb/s to 370 mb/s__

## Algorithm

1) Generate iv/key with `crypto/rand`

2) Spam AES-CTR

## Install

```bash
go get -u github.com/ammario/frandom/cmd/frandom
if [ -z "$GOPATH" ]; then
    sudo cp ~/go/bin/frandom /usr/bin/frandom
else
    sudo cp $GOPATH/bin/frandom /usr/bin/frandom
fi  
```

## As a package
[Godoc](https://godoc.org/github.com/ammario/frandom)

`frandom` appears to outperform `crypto/rand`


## Disclaimer

I'm still investigating trade-offs.

Entropy tests appear ok but I'm not a crypto-man so maybe use this with caution.


