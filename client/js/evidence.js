var EvidenceViewModel = function(graph) {
	var self = this;
	self.graph = graph;

	self.items = ko.observableArray([]);

    self.init = function() {
    	self.graph.on("click", self.showEvidence);
    	self.graph.on("dragStart", self.showEvidenceWhileSelected);
    	return self;
    }

    self.showEvidence = function(event) {
	    if (event.nodes.length > 0) {
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

	self.averageEvidence = ko.computed(function() {
	    var points = 0;
	    var items = self.items();
	    for (i in items) {
	    	points += items[i].vote;
	    }
	    if (items && items.length > 0) {
	    	var average = points/items.length;
	    	return self.convertToStars(Math.floor(average)) + " " + Math.round(10 * average)/10;
	    } else {
	        return "(NA)";
	    }
    });

    self.convertToStars = function(vote) {
      	var stars = "";
      	for (i=0; i<vote; i++) {
        	stars = stars + "&#9733;"
      	}
      	return stars;
    }
}