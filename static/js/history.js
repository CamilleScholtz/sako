// Create list for transfers.
var list = new List("history", {
	valueNames: ["txid", "timestamp", "date", "since", "amount", "fee", "height", "icon"],
	page:       20,
	item:       "<li><div class=\"info\"><span class=\"icon\"></span><span class=\"date\"></span><span class=\"since\"></span></div><h2 class=\"amount\"></h2></li>",
});

source.addEventListener("history", function(event) {
	const msg = JSON.parse(event.data);

	for (k in msg.Transfers) {
		if (!Array.isArray(msg.Transfers[k])) {
			continue;
		}

		var icon = "";
		switch (k) {
		case "in":
			icon = "<i class=\"icon fa fa-send fa-flip-horizontal fa-2x fa-fw in\"></i>";
			break;
		case "out":
			icon = "<i class=\"icon fa fa-send fa-2x fa-fw out\"></i>";
			break;
		case "pending":
			icon = "<i class=\"icon fa fa-hourglass fa-2x fa-fw pending\"></i>";
			break;
		case "failed":
			icon = "<i class=\"icon fa fa-remove fa-2x fa-fw failed\"></i>";
			break;
		}

		msg.Transfers[k].forEach(function(transfer) {
			// Skip TXID's already in the list.
			if (list.get("txid", transfer.txid).length) {
				return;
			}

			list.add({
				txid:      transfer.txid,
				timestamp: transfer.timestamp,
				date:      moment.unix(transfer.timestamp).format("[<span class=\"day\">]D[</span><span class=\"month\">]MMM[</span>]"),
				since:     moment.unix(transfer.timestamp).fromNow(),
				amount:    (transfer.amount / 1e12).toFixed(12),
				fee:       (transfer.fee / 1e12).toFixed(12),
				height:    transfer.height,
				icon:      icon,
			});
		});
	}

	// Re-sort list.
	// TODO: Can I simplify this?
	if (typeof document.getElementsByClassName("desc")[0] !== "undefined") {
		list.sort(document.getElementsByClassName("desc")[0].dataset.sort, {
			order: "desc"
		});
	} else {
		list.sort(document.getElementsByClassName("asc")[0].dataset.sort, {
			order: "asc"
		});
	}
});

// Hide scrollbar.
// TODO: This is kind of a bloated of doing this, fix?
OverlayScrollbars(document.querySelectorAll('main'), {
	className: null,
});

// Set default sorting.
list.sort("timestamp", {order: "desc"});

document.getElementById("filter").addEventListener("change", function() {
	const self = this;

	list.filter(function(item) {
		if (self.value == "all") {
			return true;
		} else {
			return item.values().icon.includes(self.value);
		}
	});
});
