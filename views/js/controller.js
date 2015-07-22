// define hutu app
var app = angular.module('hutu', [])
		.controller('homeCtrl', homeCtrl)
		.controller('createNewCntr', createNewCntr);



function homeCtrl($scope) {
	$scope.name = 'test';
	$scope.text = 'dfsdsf';
}




function createNewCntr($scope) {
	$scope.createNew = false;
	$scope.fileds = [
		{name: '', desc: ''}];


	$scope.createHost = function () {
		$scope.createNew = true;
	};
}

