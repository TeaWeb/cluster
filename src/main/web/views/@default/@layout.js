Tea.context(function () {
	this.urlPrefix = function () {
		for (var i = 0; i < arguments.length; i++) {
			var b = window.location.pathname.startsWith(arguments[i]);
			if (b) {
				return true;
			}
		}
		return false;
	};
});

function firstFocus(inputName) {
	Tea.delay(function () {
		Tea.element("input[name='" + inputName + "']").focus();
	});
};