Tea.context(function () {
	this.deleteCluster = function (clusterId) {
		if (!window.confirm("确定要删除此集群吗？")) {
			return;
		}
		this.$post(".delete")
			.params({
				"clusterId": clusterId
			})
			.refresh();
	};
});