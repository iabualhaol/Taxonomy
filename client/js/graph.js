var GraphViewModel = function(session) {
	var self = this;
	self.session = session;

   	self.selectedNode = ko.observable("");
   	self.newNodeLabel = ko.observable("");

   	self.isEmpty = ko.observable(true);

 	self.graph = {};
	self.nodes = new vis.DataSet([]);
    self.edges = new vis.DataSet([]);

 	self.init = function() {
		return self;
	}

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

    self.setOptions = function(options) {
    	console.log("graph.setOptions:", options);
    	self.graph.setOptions(options);
    }

    self.loadNodes = function() {
    	console.log('loadNodes');
    	$.getJSON("http://localhost:8080/nodes", function(data) {
		    self.nodes.update(data); 
		    if (self.nodes.length > 0) {
		    	self.isEmpty(false);
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

     self.selectedNodeLabel = ko.computed(function() {
    	if (self.selectedNode()) {
    		return self.nodes.get(self.selectedNode()).label;
    	} else {
    		return "";
    	}
    });

   self.addNewNode = function() {
      	console.log("Add node:", self.newNodeLabel());
      	$.post("http://localhost:8080/nodes",
      		{ label: self.newNodeLabel() }, function(data) {
      		self.nodes.add(data);
        	if (self.selectedNode()) {
          		console.log("Add edge from:", self.selectedNode(), "to:", data.id);
          		// console.log("Selected node:", self.graph.selectionHandler._getSelectedNode())
          		self.addNewEdge(self.selectedNode(), data.id);
        	}
        	self.graph.fit();
        	if (self.nodes.length > 0) {
          		self.isEmpty(false);
        	}
      	}, "json");
      	self.clearNewNode();
    }

    self.clearNewNode = function() {
        self.newNodeLabel("")
    }

    self.canAddNode = ko.computed(function() {
      	return self.selectedNode() || self.isEmpty();
    });

    self.clusterNode = function() {
    	console.log("Cluster node:", self.selectedNode());
    	self.graph.clusterOutliers(self.selectedNode(), {
    		clusterNodeProperties: { 
    			borderWidth: 3 
    		}
    	});
    }

    self.addNewEdge = function(from, to) {
      	console.log("New edge:", from, to);
      	$.post("http://localhost:8080/edges",
      		{ from: from, to: to }, function(data) {
        		self.edges.add(data);
      	}, "json");
    }

    self.on = function(eventType, callback) {
    	self.graph.on(eventType, callback);
    }
}