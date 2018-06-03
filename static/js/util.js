var source = new EventSource("event");
window.addEventListener("beforeunload", event => {
	source.close();
});

function formatFiat(input) {
	const f = input.toFixed(2).split(".");
	return f[0] + "<span class=\"decimal\">" + f[1] + "</span>";
}

function formatMonero(input) {
	const f = input.split(".");
	return f[0] + "." + f[1].substring(0, 3) + "<span class=\"decimal\">" +
		f[1].substring(3) + "</span>";
}

function sleep(time) {
	return new Promise(resolve => setTimeout(resolve, time));
}
