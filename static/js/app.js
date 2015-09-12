(function () {
	angular
		.module('twitter-service', ['service.controllers', 'service.services', 'ngRoute'])
		.config(['$routeProvider', function ($routeProvider) {
			$routeProvider
				.when('/', {
					controller: 'MainCtrl',
					templateUrl: 'static/templates/main.html'
				})
				.when('/collage/:width/:height', {
					controller: 'CollageCtrl',
					templateUrl: 'static/templates/collage.html'
				})
				.otherwise('/', {
					controller: 'MainCtrl',
					templateUrl: 'static/templates/main.html'
				})
			;
		}]);
})();