var ConfigViewModel = function(graph) {
	var self = this;
	self.graph = graph;

	self.edges = {};
	self.edges.arrows = {
		to: ko.observable(true) 
	};

	self.nodes = {};
	self.nodes.font = ko.observable("12 sans-serif");
	self.nodes.shape = ko.observable("ellipse");
	self.nodes.size = ko.observable(10);

	self.init = function() {
		self.setOptions();
		return self;
	}

	self.setOptions = function() {
		self.graph.setOptions(self.options());
		return self;
	}

	self.options = function() {
		return {
			edges: {
				arrows: {
					to: { 
						enabled: self.edges.arrows.to(), 
						scaleFactor: 0.0 
					},
				},
				smooth: false,
			},
			nodes: {
				font: self.nodes.font(),
				shape: self.nodes.shape(),
				size: self.nodes.size(),
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