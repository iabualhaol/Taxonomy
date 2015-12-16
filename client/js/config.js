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
		return self.defaultOptions();
	}

	self.customOptions = function() {
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
			},
		};
	}

	self.saveProfile = function() {
		localStorage["taxi.config.options"] = self.options() || "";
	}

	self.loadProfile = function() {
		self.user(localStorage["taxi.config.options"] || "");
	}

	self.defaultOptions = function() {
		return {
			layout: {
			    improvedLayout: false,
			    hierarchical: {
			      	enabled: true,
			      	levelSeparation: 150,
			      	direction: 'UD',
			      	sortMethod: 'directed'
			    }
			},
	 		nodes: {
			    fixed: false,
			    font: '12px sans-serif black',
			    shape: 'box',
	     		color: {
			      	border: '#black',
			      	background: '#13D59B',
			      	highlight: {
			        	border: '#black',
			        	background: '#D2E5FF'
			      	}
	            }
	    	}
	    }
  	}

  	self.defaultOptionsBackup = function() {
		return {
			interaction: {
	          	navigationButtons: true,
	          	keyboard: true
	        },
	 		physics:{
			    enabled: true,
			    hierarchicalRepulsion: {
			      	centralGravity: 0,
			      	springLength: 10,
			      	springConstant: 0.1,
			      	nodeDistance: 170,
			      	damping: 0.1
			    },
			    maxVelocity: 50,
			    minVelocity: 0.1
			},
			layout: {
			    improvedLayout: true,
			    hierarchical: {
			      	enabled: true,
			      	levelSeparation: 150,
			      	direction: 'UD',
			      	sortMethod: 'directed'
			    }
			},
	 		nodes: {
			    fixed: false,
			    font: '12px sans-serif black',
			    shape: 'box',
	     		color: {
			      	border: '#black',
			      	background: '#13D59B',
			      	highlight: {
			        	border: '#black',
			        	background: '#D2E5FF'
			      	}
	            }
	    	}
	    }
  	}
}