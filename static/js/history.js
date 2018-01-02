ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);

	// Stop if the XMR value didn't change.
	if (document.title.match(/.[0-9]+\.[0-9]+|\?/) == m.Price.Symbol +
		m.Price.Value.toFixed(2)) {
		return;
	}

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/,
		m.Price.Symbol + m.Price.Value.toFixed(2));

	// Fill list with transfers.
	var content = document.getElementById("content");
	content.innerHTML = "";
	for (var i = m.Transfers.length-1; i >= 0; i--) {
		var div = document.createElement("div");
		div.innerHTML = m.Transfers[i].Amount;
		div.setAttribute("class", "transfer");

		content.appendChild(div);
	}
});
