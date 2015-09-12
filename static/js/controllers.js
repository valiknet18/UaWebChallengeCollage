(function () {
	angular
		.module('service.controllers', ['service.services'])
		.controller('MainCtrl', ['$scope', '$location', 'getAuthStatus', function ($scope, $location, getAuthStatus) {
			getAuthStatus.getStatus(function (response) {
				switch (response.data.status) {
					case 200: {
						$scope.authStatus = 'auth'
					}
						break;

					default: {
						$scope.authStatus = 'unauth'
					}
				}
			});

			$scope.collage = {};

			$scope.submit = function () {
				$scope.collage.width = ($scope.collage.width)?$scope.collage.width:300;
				$scope.collage.height = ($scope.collage.height)?$scope.collage.height:400;

				$location.path("/collage/" + $scope.collage.width + "/" + $scope.collage.height);
			}
		}])
		.controller('CollageCtrl', ['$scope', function ($scope) {
			$scope.status = 0;

			var interval = setInterval(function() {
				if ($scope.status >= 100) {
					$scope.status = 0
				}  else {
					$scope.status += 5;
				}

				$('#loader').css({
					width: $scope.status + "%"
				})
			}, 100);
		}])
	;
})()