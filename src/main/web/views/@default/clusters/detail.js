Tea.context(function () {
	this.deleteNode = function (nodeId) {
		if (!window.confirm("确定要删除此节点吗？")) {
			return;
		}
		this.$post("/clusters/node/delete")
			.params({
				"clusterId": this.cluster.id,
				"nodeId": nodeId
			})
			.refresh();
	};
});