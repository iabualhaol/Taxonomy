var GraphViewModel = function() {
	var self = this;

	self.nodes = new vis.DataSet([]);
    self.edges = new vis.DataSet([]);

    self.init = function() {
     	return self;
    }

    self.createNetwork = function(container, options) {
    	self.loadNodes();
    	self.loadEdges();
    }

    self.loadNodes = function() {
    	console.log('loadNodes');
    }

    self.loadEdges = function() {
    	console.log('loadEdges');
    }
}