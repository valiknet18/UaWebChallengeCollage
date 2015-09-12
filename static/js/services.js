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
		}]);	
})();