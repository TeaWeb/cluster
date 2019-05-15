Tea.context(function () {
	this.sync = function () {
		this.$post("/clusters/node/sync")
			.params({
				"clusterId": this.cluster.id,
				"nodeId": this.node.id
			})
			.success(function () {
				alert("已通知同步");
				window.location.reload();
			});
	};
});