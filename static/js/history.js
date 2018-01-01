ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);

	// Stop if the XMR value didn't change.
	if (document.title.match(/.[0-9]+\.[0-9]+|\?/) == m.Price.Symbol +
		m.Price.Value.toFixed(2)) {
		return;
	}

	console.log(m.Transfers);

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/,
		m.Price.Symbol + m.Price.Value.toFixed(2));

	// Update graph.
	graphData.labels = m.Graph.Time;
	graphData.datasets[0].data = m.Graph.Value;
	graph.update();
});

window.onload = function() {
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
					display: false,
				}],
			},
		},
	});
}

var graphData = {
	datasets: [{
		label: "Price",
		data: Array.apply(null, {length: 48}).map(Number.prototype.valueOf, 1),
		borderColor: "rgba(255, 206, 86, 1)",
		backgroundColor: "rgba(255, 255, 255, 1)",
		pointBackgroundColor: "#FFCE56",
		pointRadius: 0,
		pointBorderWidth: 0,
	}],
	labels: Array.apply(null, {length: 48}).map(Number.call, Number),
};
