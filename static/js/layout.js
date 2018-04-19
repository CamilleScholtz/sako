document.addEventListener("DOMContentLoaded", function() {
	// Flip the card on click.
	["click", "ontouchstart"].forEach(function(evt) {
		document.getElementById("card").addEventListener(evt, function() {
			if (this.classList.contains("flip")) {
				this.classList.remove("flip");
			} else {
				this.classList.add("flip");
			}
		});
	});

	// Select the address on click.
	["click", "ontouchstart"].forEach(function(evt) {
		document.getElementById("address").addEventListener(evt, function() {
			window.getSelection().selectAllChildren(this);
		});
	});

	// Toggle the nav on responsive layouts.
	["click", "ontouchstart"].forEach(function(evt) {
		document.getElementById("hamburger").addEventListener(evt, function() {
			var nav = document.getElementsByTagName("nav")[0];

			this.animate("jello");
			if (this.classList.contains("fa-bars")) {
				nav.classList.add("expanded");
				this.classList.replace("fa-bars", "fa-times");
			} else {
				nav.classList.remove("expanded");
				this.classList.replace("fa-times", "fa-bars");
			}
		});
	});
});

ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "layout") {
		return;
	}

	// Update card info.
	document.getElementById("balance").innerHTML = m.Balance;
	document.getElementById("unbalance").innerHTML = m.UnBalance;
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
