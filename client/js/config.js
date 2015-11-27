var ConfigViewModel = function(graph) {
	var self = this;
	self.graph = graph;

	self.nodes = {};
	self.nodes.font = ko.observable("");

	self.init = function() {
		return self;
	}

	self.setOptions = function() {
		self.graph.setOptions(self.options());
		return self;
	}

	self.options = function() {
		return {
			nodes:  {
				font: self.nodes.font(),
			}		
		};
	}

	self.saveProfile = function() {
		localStorage["taxi.config.options"] = self.options() || "";
	}

	self.loadProfile = function() {
		self.user(localStorage["taxi.config.options"] || "");
	}
}