var EvidenceViewModel = function(graph) {
	var self = this;

	console.log("EvidenceViewModel.graph:", graph);
	self.graph = graph;

	self.items = ko.observableArray([]);

    self.init = function() {
    	self.graph.on("click", self.showEvidence);
    	self.graph.on("dragStart", self.showEvidenceWhileSelected);
    	return self;
    }

    self.showEvidence = function(event) {
	    if (event.nodes.length > 0) {
		    console.log("clicked on a node: " + event.nodes[0]);
		    self.graph.selectedNode(event.nodes[0]);
		    $.getJSON("http://localhost:8080/nodes/" + self.graph.selectedNode() + "/evidence", 
		        function(data) {
		        	self.items(data);
		    	});
	    } else if (event.edges.length > 0) {
	      console.log("clicked on an edge: " + params.edges[0]);
	      self.graph.selectedNode("");
	      self.items([]);
	    } else {
	      self.selectedNode("");
	      view.items([]);
	    }
	}

	self.showEvidenceWhileSelected = function(event) {
		if (event.nodes.length > 0 || event.edges.length > 0) {
      		self.showEvidence(event);
    	} 
	}
}