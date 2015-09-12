(function() {
	angular
		.module('service.services', [])
		.factory('twitterService', ['$q', function ($q) {
			var authorizationResult = false;

			return {
				initialize: function () {
					OAuth.initialize('vMqpqRdAb7Qd1Yu_Aqzk1zTi2xo', {
		                cache: true
		            });

		            authorizationResult = OAuth.create("twitter");
				},
				isReady: function () {
					return (authorizationResult);
				},
				connectTwitter: function () {
					var deferred = $q.defer();

					OAuth.popup('twitter', {
						cache: true
					}, function (error, result) {
						if (!error) {
							authorizationResult = result;
                    		deferred.resolve();
						}
					});

					return deferred.promise;
				},
				clearCache: function () {
					OAuth.clearCache('twitter');
					authorizationResult = false;
				}
			}
		}]);	
})();