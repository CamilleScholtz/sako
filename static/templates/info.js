{{define "javascript"}}

var ws = new WebSocket("ws://" + window.location.host + "/socket");

ws.onmessage = function(evt) {
	var msg = JSON.parse(evt.data);

	document.title = document.title.replace(/\[.[0-9\.]*]/, "[" +
		msg.CryptoCompareSymbol + msg.CryptoComparePrice + "]");

	document.getElementById("price-js").innerHTML = msg.CryptoComparePrice;
	document.getElementById("change-price-js").innerHTML = msg.CryptoCompareChangePrice;
	document.getElementById("change-percent-js").innerHTML = msg.CryptoCompareChangePercent;
	colorizeCurrent();

	graphData.labels = msg.CryptoCompareGraphTime;
	graphData.datasets[0].data = msg.CryptoCompareGraphPrice;
	window.graph.update();
}

var graphData = {
	datasets: [{
		label: "Price",
		data: {{.CryptoCompare.GraphPrice}},
		borderColor: "rgba(255, 206, 86, 1)",
		backgroundColor: "rgba(255, 255, 255, 1)",
		pointBackgroundColor: "#FFCE56",
		pointRadius: 0,
		pointBorderWidth: 0,
	}],
	labels: {{.CryptoCompare.GraphTime}},
};

window.onload = function() {
	colorizeCurrent();

	var ctx = document.getElementById("graph").getContext("2d");
	ctx.globalCompositeOperation = "destination-over";

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

function colorizeCurrent() {
	if (document.getElementById("change-percent-span").textContent > 0) {
		document.getElementById("change").className = "text-green";
	} else {
		document.getElementById("change").className = "text-red";
	}
}

{{end}}
