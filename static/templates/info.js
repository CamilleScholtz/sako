{{define "javascript"}}

var ws = new WebSocket("ws://" + window.location.host + "/socket");

ws.onmessage = function(evt) {
	var msg = JSON.parse(evt.data);

	if (document.getElementById("price").innerHTML == msg.CryptoComparePrice) {
		return;
	}

	document.title = document.title.replace(/\[.*\]/, "[" +
		msg.CryptoComparePrice + "]");

	graphData.labels = msg.CryptoCompareGraphTime;
	graphData.datasets[0].data = msg.CryptoCompareGraphPrice;
	window.graph.update();

	colorizeAndAnimatePrice();
	sleep(400).then(() => {
		document.getElementById("price").innerHTML = msg.CryptoComparePrice;
		document.getElementById("change").innerHTML = msg.CryptoCompareChangePercent +
			" (" + msg.CryptoCompareChangePrice + ")";
	});
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
	colorizeAndAnimatePrice();

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

function colorizeAndAnimatePrice() {
	if (document.getElementById("change").textContent.charAt(0) == "+") {
		document.getElementById("change").className = "text-green";
	} else {
		document.getElementById("change").className = "text-red";
	}

	document.getElementById("price").animate("bounce");
	document.getElementById("change").animate("bounce");
}

function sleep (time) {
	return new Promise((resolve) => setTimeout(resolve, time));
}

{{end}}
