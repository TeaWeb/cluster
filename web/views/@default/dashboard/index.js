Tea.context(function () {
	this.start = function () {
		this.$post(".start").refresh();
	};

	this.stop = function () {
		this.$post(".stop").refresh();
	};

	this.restart = function () {
		this.$post(".restart").refresh();
	};
});