# frandom

Fast random

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

## How?

`frandom` leverages your CPU's lovely AES optimization to quickly generate ciphertext which has the characteristic of being cryptographically random.

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

I'm still investigating trade-offs.