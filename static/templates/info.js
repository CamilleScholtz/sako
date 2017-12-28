{{define "websocket"}}

if (document.getElementById("price").innerHTML == m.Symbol +
	m.Price.toFixed(2)) {
	return;
}

graphData.labels = m.GraphTime;
graphData.datasets[0].data = m.GraphPrice;

var dir = "";
if (m.Price < m.GraphPrice[0]) {
	dir = "-";
	document.getElementById("change").className = "text-red";
} else {
	dir = "+";
	document.getElementById("change").className = "text-green";
}

var changePercent = Math.abs(((m.Price/m.GraphPrice[0])-1)*100).toFixed(2);
var changePrice = Math.abs(m.Price-m.GraphPrice[0]).toFixed(2);

document.getElementById("price").animate("bounce");
document.getElementById("change").animate("bounce");
sleep(450).then(() => {
	document.getElementById("price").innerHTML = m.Symbol + m.Price.toFixed(2);
	document.getElementById("change").innerHTML = dir + " "  + changePercent +
		"% (" + changePrice + ")";
});

{{end}}


{{define "javascript"}}

var graphData = {
	datasets: [{
		label: "Price",
		data: [1, 1],
		borderColor: "rgba(255, 206, 86, 1)",
		backgroundColor: "rgba(255, 255, 255, 1)",
		pointBackgroundColor: "#FFCE56",
		pointRadius: 0,
		pointBorderWidth: 0,
	}],
	labels: [0, 1],
};

window.onload = function() {
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

function sleep (time) {
	return new Promise((resolve) => setTimeout(resolve, time));
}

{{end}}
