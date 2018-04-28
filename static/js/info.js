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

document.addEventListener("DOMContentLoaded", function() {
	var ctx = document.getElementById("graph").getContext("2d");

	window.graph = new Chart(ctx, {
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
});

ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "info") {
		return;
	}

	// Update news feed.
	var newsHTML = "";
	m.News.forEach(function(d) {
		newsHTML += "<div class=\"story\">";
		newsHTML += "<span class=\"source\">" + d.Source + "</span> &nbsp;";
		newsHTML += "<small>" + moment.unix(d.Time).fromNow() + "</small>";
		newsHTML += "<br>";
		newsHTML += "<a href=\"" + d.URL + "\" target=\"_blank\">" + d.Title +
			"</a>";
		newsHTML += "</div>";
	});
	document.getElementById("news").innerHTML = newsHTML;

	// Update development feed.
	var developmentHTML = "";
	m.Funding.forEach(function(d) {
		developmentHTML += "<div class=\"project\">";
		developmentHTML += "<a href=\"" + d.URL + "\" target=\"_blank\">" +
			d.Title + "</a>";
		developmentHTML += "<br>";
		developmentHTML += "<meter low=\"" + d.Total + "\" max=\"" +
			d.Total + "\" value=\"" + d.Current + "\">" + d.Current +
			"</meter>";
		developmentHTML += "</div>";
	});
	document.getElementById("development").innerHTML = developmentHTML;

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
	var price = document.getElementById("price");
	var change = document.getElementById("change");
	if (m.Price.Value < m.GraphXMR.Value[0]) {
		dir = "-";
		change.className = "text-red";
	} else {
		dir = "+";
		change.className = "text-green";
	}

	// Calculate the change in value.
	var changePercent = Math.abs(((m.Price.Value/m.GraphXMR.Value[0])-1)*100).
		toFixed(2);
	var changePrice = Math.abs(m.Price.Value-m.GraphXMR.Value[0]).toFixed(2);

	// Nice bounce animation on change.
	price.animate("bounce");
	change.animate("bounce");

	// Sleep a bit to change the XMR value text during the middle of the
	// animation.
	sleep(450).then(() => {
		price.innerHTML = m.Price.Symbol + m.Price.Value.toFixed(2);
		change.innerHTML = dir + " "  + changePercent + "% (" + m.Price.Symbol +
			changePrice + ")";
	});
});

function sleep (time) {
	return new Promise((resolve) => setTimeout(resolve, time));
}
