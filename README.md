<table width="100%">
	<tr>
		<td align="left" width="87%">
			<img src="https://punpun.moe/qsQX.svg" height="80" align="left">
			<h2>sako</h2>
		</td>
		<td align="right">
			<a href="https://goreportcard.com/report/github.com/onodera-punpun/sako">
				<img src="https://img.shields.io/badge/go report-A+-green.svg">
			</a>
			<a href="https://punpun.moe/donate">
				<img src="https://img.shields.io/badge/dontate-XMR-orange.svg">
			</a>
		</td>
	</tr>
	<tr>
		<td colspan="2">
			A self-hosted Monero wallet web-interface. <i>(work in progress)<i>
		</td>
	</tr>
</table>


## SCROTS

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
