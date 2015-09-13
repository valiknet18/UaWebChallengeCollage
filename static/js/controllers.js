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

				// $location.path("/collage/" + $scope.collage.width + "/" + $scope.collage.height);
				$location.path("/collage/").search({width: $scope.collage.width, height: $scope.collage.height, name: $scope.collage.name});
			}
		}])
		.controller('CollageCtrl', ['$scope', '$routeParams', 'getCollage', function ($scope, $routeParams, getCollage) {
			console.log($routeParams)
			$scope.params = $routeParams;

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

			getCollage($routeParams, function (data) {
				$scope.friends = data.data.users;

				clearInterval(interval);

				$('#preloader-area').remove();
			})
		}])
	;
})()