var graphData = {
	datasets: [{
		label: "XMR",
		yAxisID: "XMR",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#F6EFDC",
		backgroundColor: "#F6EFDC",
		pointRadius: 0,
	}, {
		label: "BTC",
		yAxisID: "BTC",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#FF7A73",
		backgroundColor: "#FF7A73",
		pointRadius: 0,
	}, {
		label: "ETH",
		yAxisID: "ETH",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "#E8CEDE",
		backgroundColor: "#E8CEDE",
		pointRadius: 0,
	}],
	labels: Array.apply(null, {length: 48}).map(Number.call, Number),
};

ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "info") {
		return;
	}

	// Stop if the current XMR value didn't change.
	if (document.title.match(/.[0-9]+\.[0-9]+|\?/) == m.Price.Symbol +
		m.Price.Value.toFixed(2)) {
		return;
	}

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/,
		m.Price.Symbol + m.Price.Value.toFixed(2));

	// Update graph.
	graphData.labels = m.GraphXMR.Time;
	graphData.datasets[0].data = m.GraphXMR.Value;
	graphData.datasets[1].data = m.GraphBTC.Value;
	graphData.datasets[2].data = m.GraphETH.Value;
	graph.update();

	// Determine if the XMR value went up or down and act accordingly.
	var dir = "";
	if (m.Price.Value < m.GraphXMR.Value[0]) {
		dir = "-";
		document.getElementById("change").className = "text-red";
	} else {
		dir = "+";
		document.getElementById("change").className = "text-green";
	}

	// Calculate the change in value.
	var changePercent = Math.abs(((m.Price.Value/m.GraphXMR.Value[0])-1)*100).
		toFixed(2);
	var changePrice = Math.abs(m.Price.Value-m.GraphXMR.Value[0]).toFixed(2);

	// Nice bounce animation on change.
	document.getElementById("price").animate("bounce");
	document.getElementById("change").animate("bounce");

	// Sleep a bit to change the XMR value text during the middle of the
	// animation.
	sleep(450).then(() => {
		document.getElementById("price").innerHTML = m.Price.Symbol + m.Price.
			Value.toFixed(2);
		document.getElementById("change").innerHTML = dir + " "  +
			changePercent + "% (" + m.Price.Symbol + changePrice + ")";
	});
});

document.addEventListener("DOMContentLoaded", function() {
	var ctx = document.getElementById("graph").getContext("2d");

	window.graph = new Chart(ctx, {
		type: "line",
		data: graphData,
		options: {
			maintainAspectRatio: false,
			tooltips: false,
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
});

function sleep (time) {
	return new Promise((resolve) => setTimeout(resolve, time));
}
