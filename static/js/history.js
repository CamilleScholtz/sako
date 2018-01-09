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
		var icon = "";
		switch (transfer.Status) {
		case "incoming":
			icon = "<i class=\"icon fa fa-send fa-flip-horizontal fa-2x fa-fw incoming\"></i>";
			break;
		case "outgoing":
			icon = "<i class=\"icon fa fa-send fa-2x fa-fw outgoing\"></i>";
			break;
		case "pending":
			icon = "<i class=\"icon fa fa-hourglass fa-2x fa-fw pending\"></i>";
			break;
		case "failed":
			icon = "<i class=\"icon fa fa-remove fa-2x fa-fw failed\"></i>";
			break;
		}

		document.getElementById("transfers").innerHTML += " \
			<li class=\"transfer\"> \
				" + icon + " \
				<h1 class=\"amount\">" + transfer.Amount + "</h1> \
				<span class=\"time\">" + moment.unix(transfer.Timestamp).fromNow() + "</span> \
				<a href=\"#\" class=\"details\">Details</a> \
			</li> \
		";
	});
});
