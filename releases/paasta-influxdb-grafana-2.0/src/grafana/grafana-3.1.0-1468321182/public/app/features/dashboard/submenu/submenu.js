/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["angular"],function(a){function b(){return{restrict:"E",templateUrl:"public/app/features/dashboard/submenu/submenu.html",controller:d,bindToController:!0,controllerAs:"ctrl",scope:{dashboard:"="}}}var c,d;return a("submenuDirective",b),{setters:[function(a){c=a}],execute:function(){d=function(){function a(a,b,c,d){this.$rootScope=a,this.templateValuesSrv=b,this.templateSrv=c,this.$location=d,this.annotations=this.dashboard.templating.list,this.variables=this.dashboard.templating.list}return a.$inject=["$rootScope","templateValuesSrv","templateSrv","$location"],a.prototype.disableAnnotation=function(a){a.enable=!a.enable,this.$rootScope.$broadcast("refresh")},a.prototype.getValuesForTag=function(a,b){return this.templateValuesSrv.getValuesForTag(a,b)},a.prototype.variableUpdated=function(a){var b=this;this.templateValuesSrv.variableUpdated(a).then(function(){b.$rootScope.$emit("template-variable-value-updated"),b.$rootScope.$broadcast("refresh")})},a}(),a("SubmenuCtrl",d),c["default"].module("grafana.directives").directive("dashboardSubmenu",b)}}});