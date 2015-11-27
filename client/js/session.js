var SessionViewModel = function() {
	var self = this;

	self.user = ko.observable("");

	self.init = function() {
		console.log("init session");
		self.loadProfile();
		return self;
	}

	self.saveProfile = function() {
		localStorage["taxi.session.user"] = self.user() || "";
		if (localStorage["taxi.session.user"]) {
			console.log("User saved profile:", self.user());
		}
	}

	self.loadProfile = function() {
		self.user(localStorage["taxi.session.user"] || "");
		if (localStorage["taxi.session.user"]) {
			console.log("User profile loaded:", self.user());
		}
	}
}