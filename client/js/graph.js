var GraphViewModel = function() {
	var self = this;

   	self.selectedNode = ko.observable("");

 	self.graph = {};
	self.nodes = new vis.DataSet([]);
    self.edges = new vis.DataSet([]);

    self.createGraph = function(container, options) {
    	var data = {
			nodes: self.nodes,
			edges: self.edges
	  	};
	  	self.graph = new vis.Network(container, data, options);
    	self.loadNodes();
    	self.loadEdges();
    	return self;
    }

    self.loadNodes = function() {
    	console.log('loadNodes');
    	$.getJSON("http://localhost:8080/nodes", function(data) {
		    self.nodes.update(data); 
		    if (self.nodes.length > 0) {
		    	// self.firstNode(false);
		    }
		    self.graph.fit(); 
		});
    }

    self.loadEdges = function() {
    	console.log('loadEdges');
    	$.getJSON("http://localhost:8080/edges", function(data) {
		    self.edges.update(data);  
		    self.graph.fit();
		});
    }

    self.on = function(eventType, callback) {
    	self.graph.on(eventType, callback);
    }
}