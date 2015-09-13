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
				$scope.collage.height = $scope.collage.width * 1.3

				// $location.path("/collage/" + $scope.collage.width + "/" + $scope.collage.height);
				$location.path("/collage/").search({width: $scope.collage.width, height: $scope.collage.height, name: $scope.collage.name});
			}
		}])
		.controller('CollageCtrl', ['$scope', '$routeParams', '$location', 'getCollage', 'getAuthStatus', function ($scope, $routeParams, $location, getCollage, getAuthStatus) {
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

				console.log($scope.authStatus)

				if ($scope.authStatus == 'auth') {
					$scope.params = $routeParams;

					getCollage($routeParams, function (data) {
						$scope.lines = data.data;
						// $scope.imageParams = {width: data.data.width, height: data.data.height}

						clearInterval(interval);

						$('#preloader-area').remove();
					})
				} else {
					$location.path("/").search({})
				}
			});	

			$scope.collage = $routeParams;
			
			$scope.submit = function () {
				$scope.collage.width = ($scope.collage.width)?$scope.collage.width:300;
				$scope.collage.height = $scope.collage.width * 1.2

				// $location.path("/collage/" + $scope.collage.width + "/" + $scope.collage.height);
				$location.path("/collage/").search({width: $scope.collage.width, height: $scope.collage.height, name: $scope.collage.name});
			}			
		}])
	;
})()