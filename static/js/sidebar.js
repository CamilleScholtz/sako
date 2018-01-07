ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "sidebar") {
		return;
	}

	// Update card info.
	document.getElementById("balance").innerHTML = m.Balance.Balance;
	document.getElementById("unbalance").innerHTML = m.Balance.UnBalance;
	document.getElementById("address").innerHTML = m.Address;
	document.getElementById("qr").src = "/static/images/qr/" + m.Address +
		".png";

	// Display "Syncing..." in case the current block height doesn't match the
	// maximum block height.
	if (m.CurHeight != m.MaxHeight) {
		document.getElementById("sync").animate("fadeIn", showSync);
	} else {
		document.getElementById("sync").animate("fadeOut", hideSync);
	}
});

function showSync() {
	document.getElementById("sync").innerHTML = "<i class=\"fa fa-circle-o-notch fa-spin fa-2x fa-fw\"></i> Syncing...";
}

function hideSync() {
	document.getElementById("sync").innerHTML = "";
}

// selectAll selects all the text in an element, used to select the public key
// on click.
function selectAll(e) {
	var r = document.createRange();
	r.selectNode(document.getElementById(e));
	window.getSelection().removeAllRanges();
	window.getSelection().addRange(r);
}
