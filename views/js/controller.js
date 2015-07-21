// define hutu app
var app = angular.module('hutu', []).controller('homeCtrl', homeCtrl);



function homeCtrl($scope) {
	$scope.name = 'test';
	$scope.text = 'dfsdsf';
}


