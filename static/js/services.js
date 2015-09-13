(function() {
	angular
		.module('service.services', [])
		.factory('authService', ['$http', function ($http) {
			
		}])
		.factory('getAuthStatus', ['$http', function ($http) {			
			var responseData = {};

			return {
				getStatus: function (callback) {
					$http
						.get('/api/verify_auth')
						.then(function (response) {
							responseData = response

							callback(response)
						});	
				}, 
			}
		}])
		.factory('getCollage', ['$http', function ($http) {
			return function (params, callback) {
				console.log(params)

				$http
					.get('/api/collage?name=' + params.name + "&width=" + params.width + "&height=" + params.height)
					.then(callback)
				;
			} 
		}]);	
})();