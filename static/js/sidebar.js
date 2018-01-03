// selectAll selects all the text in an element, used to select the public key
// on click.
function selectAll(e) {
	var r = document.createRange();
	r.selectNode(document.getElementById(e));
	window.getSelection().removeAllRanges();
	window.getSelection().addRange(r);
}
