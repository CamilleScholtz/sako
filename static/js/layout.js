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
			const nav = document.getElementsByTagName("nav")[0];

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

source.addEventListener("sidebar", function(event) {
	const msg = JSON.parse(event.data);

	// Update card info.
	document.getElementById("balance").innerHTML = msg.Balance;
	document.getElementById("unbalance").innerHTML = msg.UnBalance;
	document.getElementById("address").innerHTML = msg.Address;
	document.getElementById("qr").src = "/static/images/qr/" + msg.Address +
		".png";

	// Display "Syncing..." in case the current block height doesn't match the
	// maximum block height.
	if (msg.CurHeight != msg.MaxHeight) {
		document.getElementById("sync").animate("fadeIn", showSync);
	} else {
		document.getElementById("sync").animate("fadeOut", hideSync);
	}
});

source.addEventListener("price", function(event) {
	const msg = JSON.parse(event.data);

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/, msg.Symbol +
		msg.Value.toFixed(2));
});

function showSync() {
	document.getElementById("sync").innerHTML = "<i class=\"fa fa-circle-o-notch fa-spin fa-2x fa-fw\"></i> Syncing...";
}

function hideSync() {
	document.getElementById("sync").innerHTML = "";
}
