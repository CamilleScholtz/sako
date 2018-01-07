ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "history") {
		return;
	}

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/,
		m.Price.Symbol + m.Price.Value.toFixed(2));

	// Fill feed with info.
	document.getElementById("transfers").innerHTML = "";
	m.Transfers.forEach(function(transfer) {
		document.getElementById("transfers").innerHTML += " \
			<div class=\"transfer\"> \
				<h1 class=\"amount\">" + transfer.Amount + "</h1> \
				<span class=\"size\">" + transfer.Size + " kB</span> \
				<span class=\"hash\">" + transfer.Hash + "</span> \
			</div> \
		";
	});
});
