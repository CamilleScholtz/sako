var list;

ws.addEventListener('message', function(evt) {
	var m = JSON.parse(evt.data);
	if (m.Type != "history") {
		return;
	}

	// Update title to display the current XMR value.
	document.title = document.title.replace(/.[0-9]+\.[0-9]+|\?/,
		m.Price.Symbol + m.Price.Value.toFixed(2));

	for (k in m.Transfers) {
		if (!Array.isArray(m.Transfers[k])) {
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

		m.Transfers[k].forEach(function(transfer) {
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

document.addEventListener("DOMContentLoaded", function() {
	// Create list for transfers.
	list = new List("history", {
		valueNames: ["txid", "timestamp", "date", "since", "amount", "fee", "height", "icon"],
		pagination: true,
		page:       20,
		item:       "<li><div class=\"info\"><span class=\"icon\"></span><span class=\"date\"></span><span class=\"since\"></span></div><h2 class=\"amount\"></h2></li>",
	});

	// Set default sorting.
	list.sort("timestamp", {order: "desc"});

	document.getElementById("filter").addEventListener("change", function() {
		var self = this;

		list.filter(function(item) {
			if (self.value == "all") {
				return true;
			} else {
				return item.values().icon.includes(self.value);
			}
		});
	});
});
