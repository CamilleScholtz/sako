source.addEventListener("settings", event => {
	const msg = JSON.parse(event.data);

	document.getElementsByName("daemon")[0].value = msg.Daemon;
	document.getElementsByName("rpc")[0].value = msg.RPC;
	document.getElementsByName("host")[0].value = msg.Host;
	document.getElementsByName("currency")[0].value = msg.Currency;
});
