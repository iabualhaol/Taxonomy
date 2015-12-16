var SessionViewModel = function() {
	var self = this;

	self.user = ko.observable("");
	self.baseUrl = "";

	self.init = function(host) {
		console.log("init session");
		self.baseUrl = "http://" + host + ":8080/";
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