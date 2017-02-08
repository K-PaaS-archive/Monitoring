/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["angular","lodash","app/core/core_module","app/core/app_events"],function(a){var b,c,d,e,f;return{setters:[function(a){b=a},function(a){c=a},function(a){d=a},function(a){e=a}],execute:function(){f=function(){function a(a,b,c,d){this.$timeout=a,this.$sce=b,this.$rootScope=c,this.$modal=d,this.list=[]}return a.$inject=["$timeout","$sce","$rootScope","$modal"],a.prototype.init=function(){var a=this;this.$rootScope.onAppEvent("alert-error",function(b,c){a.set(c[0],c[1],"error",0)},this.$rootScope),this.$rootScope.onAppEvent("alert-warning",function(b,c){a.set(c[0],c[1],"warning",5e3)},this.$rootScope),this.$rootScope.onAppEvent("alert-success",function(b,c){a.set(c[0],c[1],"success",3e3)},this.$rootScope),e["default"].on("confirm-modal",this.showConfirmModal.bind(this))},a.prototype.set=function(a,d,e,f){var g=this,h={title:a||"",text:d||"",severity:e||"info"},i=b["default"].toJson(h);return c["default"].remove(this.list,function(a){return b["default"].toJson(a)===i}),this.list.push(h),f>0&&this.$timeout(function(){g.list=c["default"].without(g.list,h)},f),this.$rootScope.$$phase||this.$rootScope.$digest(),h},a.prototype.clear=function(a){this.list=c["default"].without(this.list,a)},a.prototype.clearAll=function(){this.list=[]},a.prototype.showConfirmModal=function(a){var b=this.$rootScope.$new();b.title=a.title,b.text=a.text,b.text2=a.text2,b.onConfirm=a.onConfirm,b.onAltAction=a.onAltAction,b.altActionText=a.altActionText,b.icon=a.icon||"fa-check",b.yesText=a.yesText||"Yes",b.noText=a.noText||"Cancel";var c=this.$modal({template:"public/app/partials/confirm_modal.html",persist:!1,modalClass:"confirm-modal",show:!1,scope:b,keyboard:!1});c.then(function(a){a.modal("show")})},a}(),a("AlertSrv",f),d["default"].service("alertSrv",f)}}});