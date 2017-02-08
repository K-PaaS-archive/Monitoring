/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["lodash","app/core/config","app/core/core_module"],function(a){var b,c,d,e;return{setters:[function(a){b=a},function(a){c=a},function(a){d=a}],execute:function(){e=function(){function a(a,b,c,d,e){this.$http=a,this.alertSrv=b,this.$rootScope=c,this.$q=d,this.$timeout=e,this.inFlightRequests={},this.HTTP_REQUEST_CANCELLED=-1}return a.$inject=["$http","alertSrv","$rootScope","$q","$timeout"],a.prototype.get=function(a,b){return this.request({method:"GET",url:a,params:b})},a.prototype["delete"]=function(a){return this.request({method:"DELETE",url:a})},a.prototype.post=function(a,b){return this.request({method:"POST",url:a,data:b})},a.prototype.patch=function(a,b){return this.request({method:"PATCH",url:a,data:b})},a.prototype.put=function(a,b){return this.request({method:"PUT",url:a,data:b})},a.prototype.requestErrorHandler=function(a){if(!a.isHandled){var c=a.data||{message:"Unexpected error"};if(b["default"].isString(c)&&(c={message:c}),422===a.status)throw this.alertSrv.set("Validation failed",c.message,"warning",4e3),c;throw c.severity="error",a.status<500&&(c.severity="warning"),c.message&&this.alertSrv.set("Problem!",c.message,c.severity,1e4),c}},a.prototype.request=function(a){var b=this;a.retry=a.retry||0;var d=0===a.url.indexOf("/"),e=0===a.retry;return d&&!a.hasSubUrl&&(a.url=c["default"].appSubUrl+a.url,a.hasSubUrl=!0),this.$http(a).then(function(c){return"GET"!==a.method&&c&&c.data.message&&b.alertSrv.set(c.data.message,"","success",3e3),c.data},function(c){if(401===c.status&&e)return b.loginPing().then(function(){return a.retry=1,b.request(a)});throw b.$timeout(b.requestErrorHandler.bind(b,c),50),c})},a.prototype.datasourceRequest=function(a){var c=this;a.retry=a.retry||0;var d;a.requestId&&(d=this.inFlightRequests[a.requestId],d&&d.resolve(),d=this.$q.defer(),a.timeout=d.promise,this.inFlightRequests[a.requestId]=d);var e=0===a.url.indexOf("/"),f=0===a.retry;return e&&a.headers&&a.headers.Authorization&&(a.headers["X-DS-Authorization"]=a.headers.Authorization,delete a.headers.Authorization),this.$http(a)["catch"](function(g){if(g.status===c.HTTP_REQUEST_CANCELLED)throw{err:g,cancelled:!0};if(e&&f&&401===g.status)return c.loginPing().then(function(){return a.retry=1,d&&d.resolve(),c.datasourceRequest(a)});throw b["default"].isString(g.data)&&500===g.status&&(g.data={error:g.statusText}),!g.data.message&&b["default"].isString(g.data.error)&&(g.data.message=g.data.error),g})["finally"](function(){a.requestId&&delete c.inFlightRequests[a.requestId]})},a.prototype.loginPing=function(){return this.request({url:"/api/login/ping",method:"GET",retry:1})},a.prototype.search=function(a){return this.get("/api/search",a)},a.prototype.getDashboard=function(a,b){return this.get("/api/dashboards/"+a+"/"+b)},a.prototype.saveDashboard=function(a,b){return b=b||{},this.post("/api/dashboards/db/",{dashboard:a,overwrite:b.overwrite===!0})},a}(),a("BackendSrv",e),d["default"].service("backendSrv",e)}}});