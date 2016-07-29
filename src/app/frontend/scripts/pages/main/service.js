define([
	'utils'
	], function(utils){
	'use strict';

	var getApis = function($http){
		var apis = {};

		apis.getNavlist = function(param, success, error){
			return utils.http($http, 'post', '/api/main/navlist', param, success, error);
		};

		return apis;
	};	

	var services = {
		module: 'yce-manage',
		name: 'mainService',	
		getApis: getApis
	};

	return services;
});