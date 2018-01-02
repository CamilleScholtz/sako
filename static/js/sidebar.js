ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);

	// Stop if the wallet balance didn't change.
	if (document.getElementById("unbalance").innerHTML ==
		m.Sidebar.Balance.Balance) {
		return;
	}

	document.getElementById("balance").innerHTML = m.Sidebar.Balance.Balance;
	document.getElementById("unbalance").innerHTML = m.Sidebar.Balance.UnBalance;
	document.getElementById("address").innerHTML = m.Sidebar.Address;
	document.getElementById("qr").src = "/static/images/qr/" +
		m.Sidebar.Address + ".png";
});

// selectAll selects all the text in an element, used to select the public key
// on click.
function selectAll(e) {
	var r = document.createRange();
	r.selectNode(document.getElementById(e));
	window.getSelection().removeAllRanges();
	window.getSelection().addRange(r);
}
