source.addEventListener("sidebar", event => {
	const msg = JSON.parse(event.data);

	// Update card info.
	document.getElementById("balance").innerHTML = msg.Balance;
	document.getElementById("unbalance").innerHTML = msg.UnBalance;
	document.getElementById("address").innerHTML = msg.Address;
	document.getElementById("qr").src = "/static/images/qr/" + msg.Address +
		".png";

	// Display "Syncing..." in case the current block height doesn't match the
	// maximum block height.
	// TODO: This shows "Syncing..." way too often.
	if (msg.CurHeight != msg.MaxHeight) {
		document.getElementById("sync").animate("fadeIn", showSync);
	} else {
		document.getElementById("sync").animate("fadeOut", hideSync);
	}
});

source.addEventListener("price", event => {
	const msg = JSON.parse(event.data);

	// Update title to display the current XMR value.
	document.title = document.title.replace(/\[.*\]/, "[" + msg.Symbol + msg.
		Value.toFixed(2)) + "]";
});

// Flip the card on click.
["click", "ontouchstart"].forEach(trigger => {
	document.getElementById("card").addEventListener(trigger, event => {
		event.currentTarget.classList.toggle("flip");
	});
});

// Select the address on click.
["click", "ontouchstart"].forEach(trigger => {
	document.getElementById("address").addEventListener(trigger, event => {
		window.getSelection().selectAllChildren(event.target);
	});
});

// Toggle the nav on responsive layouts.
["click", "ontouchstart"].forEach(trigger => {
	document.getElementById("hamburger").addEventListener(trigger, event => {
		const nav = document.getElementsByTagName("nav")[0];

		event.target.animate("jello");
		nav.classList.toggle("expanded");
		// TODO: Simplify this.
		if (this.classList.contains("fa-bars")) {
			this.classList.replace("fa-bars", "fa-times");
		} else {
			this.classList.replace("fa-times", "fa-bars");
		}
	});
});

function showSync() {
	document.getElementById("sync").innerHTML = "<i class=\"fa fa-circle-o-notch fa-spin fa-2x fa-fw\"></i> Syncing...";
}

function hideSync() {
	document.getElementById("sync").innerHTML = "";
}

// Add particle background the card.
particlesJS(
	"particles-card",
	{
		"particles": {
			"number": {
				"value": 32,
				"density": {
					"enable": false,
				},
			},
			"color": {
				"value": "#DFDDA2",
			},
			"shape": {
				"type": "polygon",
				"stroke": {
					"width": 0,
					"color": "#000000",
				}, "polygon": {
					"nb_sides": 5,
				},
			},
			"opacity": {
				"value": 0.6,
				"random": true,
			},
			"size": {
				"value": 4.09946952864299,
				"random": true,
			},
			"line_linked": {
				"enable": true,
				"distance": 150,
				"color": "#DFDDA2",
				"opacity": 0.4,
				"width": 1,
			},
			"move": {
				"enable": true,
				"speed": 0.25,
				"direction": "none",
				"random": false,
				"straight": false,
				"out_mode": "out",
				"bounce": false,
			},
		},
		"interactivity": {
			"detect_on": "canvas",
			"events": {
				"onhover": {
					"enable": false,
				},
				"onclick": {
					"enable": false,
				},
				"resize": true,
			},
		},
		"retina_detect": true,
	},
);
