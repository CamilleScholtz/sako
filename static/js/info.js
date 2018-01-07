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
	graphData.labels = m.Graph.Time;
	graphData.datasets[0].data = m.Graph.Value;
	graph.update();

	// Determine if the XRM value went up or down and act accordingly.
	var dir = "";
	if (m.Price.Value < m.Graph.Value[0]) {
		dir = "-";
		document.getElementById("change").className = "text-red";
	} else {
		dir = "+";
		document.getElementById("change").className = "text-green";
	}

	// Calculate the change in value.
	var changePercent = Math.abs(((m.Price.Value/m.Graph.Value[0])-1)*100).
		toFixed(2);
	var changePrice = Math.abs(m.Price.Value-m.Graph.Value[0]).toFixed(2);

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

	// Fill feed with info.
	//m.Feed.forEach(function(item) {
	//	document.getElementById("feed").innerHTML += "<a href=\"" + item.Link +
	//		"\" target=\"_blank\">" + item.Title + "</br>";
	//});
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

function sleep (time) {
	return new Promise((resolve) => setTimeout(resolve, time));
}
