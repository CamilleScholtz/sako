[![Go Report Card](https://goreportcard.com/badge/github.com/onodera-punpun/sako)](https://goreportcard.com/report/github.com/onodera-punpun/sako)

A self-hosted Monero web-interface. **WIP**

![](https://punpun.moe/9N93.png)


## SYNOPSIS

sako [arguments]


## INSTALLATION

* Install [Monero](https://getmonero.org/), and run `monero-wallet-rpc
--rpc-bind-port 18082 --disable-rpc-login`. If you don't want to download the
entire Monero blockchain you can use an external node using`monero-wallet-rpc
--daemon-host node.viaxmr.com --rpc-bind-port 18082 --disable-rpc-login`.

* Clone this repository and edit `runtime/config.toml` to fit your needs.

* Build `sako` using `go -d -u -v && go build`.

* Run `sako` using `./sako`.



## AUTHORS

Camille Scholtz


## DONATE

`43N66aiA9392qz7pTFAfSe1qJxrxDACDhMvcTVv5uPkWK37XSCMxaeqg2PTp8NeZMuaGcjatuQCaoCFrUdRxuQX71mBnwvr`
