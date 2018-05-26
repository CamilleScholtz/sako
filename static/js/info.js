var graphData = {
	datasets: [{
		label: "XMR",
		yAxisID: "XMR",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#C95B55",
		backgroundColor: "#C95B55",
		pointRadius: 0,
	}, {
		label: "BTC",
		yAxisID: "BTC",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#AC4740",
		backgroundColor: "#AC4740",
		pointRadius: 0,
	}, {
		label: "ETH",
		yAxisID: "ETH",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#FFCA63",
		backgroundColor: "#FFCA63",
		pointRadius: 0,
	}],
	labels: Array.apply(null, {length: 48}).map(Number.call, Number),
};

source.addEventListener("graph", function(event) {
	const msg = JSON.parse(event.data);

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/, msg.Price.
		Symbol + msg.Price.Value.toFixed(2));

	// Update graph.
	graphData.labels = msg.XMR.Time;
	graphData.datasets[0].data = msg.XMR.Value;
	graphData.datasets[1].data = msg.BTC.Value;
	graphData.datasets[2].data = msg.ETH.Value;
	graph.update();

	// Determine if the XMR value went up or down and act accordingly.
	var dir = "";
	const price = document.getElementById("price");
	const change = document.getElementById("change");
	if (msg.Price.Value < msg.XMR.Value[0]) {
		dir = "<i class=\"fa fa-level-down text-red\"></i> ";
		change.className = "text-red";
	} else {
		dir = "<i class=\"fa fa-level-up text-green\"></i> ";
		change.className = "text-green";
	}

	// Nice bounce animation on change.
	price.animate("bounce");
	change.animate("bounce");

	// Sleep a bit to change the XMR value text during the middle of the
	// animation.
	sleep(450).then(() => {
		price.innerHTML = msg.Price.Symbol + formatFiat(msg.Price.Value);
		change.innerHTML = dir + Math.abs(((msg.Price.Value/msg.XMR.Value[0])-
			1)*100).toFixed(2) + "%";
	});
});

source.addEventListener("submissions", function(event) {
	const msg = JSON.parse(event.data);

	// Update submissions feed.
	var content = "";
	msg.forEach(function(d) {
		content += "<div class=\"story\">";
		content += "<span class=\"source\">" + d.Source + "</span> &nbsp;";
		content += "<small>" + moment.unix(d.Time).fromNow() + "</small>";
		content += "<br>";
		content += "<a href=\"" + d.URL + "\" target=\"_blank\">" + d. Title +
			"</a>";
		content += "</div>";
	});
	document.getElementById("submissions").innerHTML = content;
});

source.addEventListener("funding", function(event) {
	const msg = JSON.parse(event.data);

	// Update funding feed.
	var content = "";
	msg.forEach(function(d) {
		content += "<div class=\"project\">";
		content += "<a href=\"" + d.URL + "\" target=\"_blank\">" + d.Title +
			"</a>";
		content += "<br>";
		content += "<meter low=\"" + d.Total + "\" max=\"" + d.Total +
			"\" value=\"" + d.Current + "\">" + d.Current + "</meter>";
		content += "<small>" + d.Current.toFixed(2) + " / " + d.Total.
			toFixed(2) + " XMR - " + d.Contributions + " contributions</small>";
		content += "</div>";
	});
	document.getElementById("funding").innerHTML = content;
});

window.graph = new Chart(document.getElementById("graph").getContext("2d"), {
	type: "line",
	data: graphData,
	options: {
		maintainAspectRatio: false,
		events: [],
		responsive: true,
		legend: {
			display: false
		},
		scales: {
			xAxes: [{
				display: false,
				type: 'time',
			}],
			yAxes: [{
				id: "XMR",
				display: false,
			}, {
				id: "BTC",
				display: false,
			}, {
				id: "ETH",
				display: false,
			}],
		},
	},
});
