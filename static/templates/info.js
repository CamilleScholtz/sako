{{define "javascript"}}

var ws = new WebSocket("ws://" + window.location.host + "/socket");

ws.onmessage = function(evt) {
	var msg = JSON.parse(evt.data);

	document.title = document.title.replace(/\[.*\]/, "[" +msg.coincapCurrent +
		"]");

	document.getElementById("current").innerHTML = msg.coincapCurrent;
	document.getElementById("change").innerHTML = msg.coincapChangePercent +
		" (" + msg.coincapChangePrice + ")";
	colorizeCurrent();

	graphData.labels = msg.coincapDate;
	graphData.datasets[0].data = msg.coincapPrice;
	window.graph.update();
}

var graphData = {
	datasets: [{
		label: "Price",
		data: {{.Coincap.Price}},
		borderColor: "rgba(255, 206, 86, 1)",
		backgroundColor: "rgba(255, 255, 255, 1)",
		pointBackgroundColor: "#FFCE56",
		pointRadius: 0,
		pointBorderWidth: 0,
	}],
	labels: {{.Coincap.Date}},
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
					ticks: {
						display: true,
					},
					gridLines: {
						display: false,
					},
				}],
				yAxes: [{
					display: false,
					ticks: {
						display: false,
					},
					gridLines: {
						display: false,
					},
				}],
			},
		},
	});
}

function colorizeCurrent() {
	if (document.getElementById("change").textContent.charAt(0) == "+") {
		document.getElementById("change").className = "text-green";
	} else {
		document.getElementById("change").className = "text-red";
	}
}

{{end}}
