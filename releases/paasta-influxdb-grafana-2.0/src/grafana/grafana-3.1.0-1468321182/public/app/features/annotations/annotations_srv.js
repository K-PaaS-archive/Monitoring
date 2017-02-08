/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

define(["angular","lodash","./editor_ctrl"],function(a,b){"use strict";var c=a.module("grafana.services");c.service("annotationsSrv",["$rootScope","$q","datasourceSrv","alertSrv","timeSrv",function(c,d,e,f,g){function h(a){console.log("Annotation error: ",a);var b=a.message||"Annotation query failed";f.set("Annotations error",b,"error")}var i,j=[],k=this;this.init=function(){c.onAppEvent("refresh",this.clearCache,c),c.onAppEvent("dashboard-initialized",this.clearCache,c)},this.clearCache=function(){i=null,j=[]},this.getAnnotations=function(f){if(0===f.annotations.list.length)return d.when(null);if(i)return i;k.dashboard=f;var l=b.where(f.annotations.list,{enable:!0}),m=g.timeRange(),n=g.timeRange(!1),o=b.map(l,function(b){return b.snapshotData?void k.receiveAnnotationResults(b.snapshotData):e.get(b.datasource).then(function(c){var d={range:m,rangeRaw:n,annotation:b};return c.annotationQuery(d).then(k.receiveAnnotationResults).then(function(c){f.snapshot&&(b.snapshotData=a.copy(c))}).then(null,h)},this)});return i=d.all(o).then(function(){return j})["catch"](function(a){c.appEvent("alert-error",["Annotations failed",a.message||a])})},this.receiveAnnotationResults=function(a){for(var b=0;b<a.length;b++)k.addAnnotation(a[b]);return a},this.addAnnotation=function(a){j.push({annotation:a.annotation,min:a.time,max:a.time,eventType:a.annotation.name,title:a.title,tags:a.tags,text:a.text,score:1})},this.init()}])});