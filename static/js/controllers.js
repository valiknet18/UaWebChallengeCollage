(function () {
	angular
		.module('service.controllers', ['service.services'])
		.controller('MainCtrl', ['$scope', 'twitterService', function ($scope, twitterService) {
			twitterService.initialize();

			$scope.authenticate = function () {
				twitterService.connectTwitter().then(function () {
					if (twitterService.isReady()) {
						$scope.connectedTwitter = true;
					}
				});
			}

			$scope.signOut = function() {
		        twitterService.clearCache();
		        $scope.tweets.length = 0;
		        $('#getTimelineButton, #signOut').fadeOut(function() {
		            $('#connectButton').fadeIn();
		            $scope.$apply(function() {
		                $scope.connectedTwitter = false
		            })
		        });
		    }
		}]);
})()