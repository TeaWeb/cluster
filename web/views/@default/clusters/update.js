Tea.context(function () {
	firstFocus("name");

	this.submitSuccess = function () {
		alert("保存成功");
		window.location = "/clusters/detail?clusterId=" + this.cluster.id;
	};
});