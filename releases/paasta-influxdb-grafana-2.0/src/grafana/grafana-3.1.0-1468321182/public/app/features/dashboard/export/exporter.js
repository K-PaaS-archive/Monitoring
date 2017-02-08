/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["app/core/config","lodash","../dynamic_dashboard_srv"],function(a){var b,c,d,e;return{setters:[function(a){b=a},function(a){c=a},function(a){d=a}],execute:function(){e=function(){function a(a){this.datasourceSrv=a}return a.prototype.makeExportable=function(a){var e=this,f=new d.DynamicDashboardSrv;f.process(a,{cleanUpOnly:!0}),a.id=null;for(var g=[],h={},i={},j=[],k=function(a){j.push(e.datasourceSrv.get(a.datasource).then(function(b){var c="DS_"+b.name.replace(" ","_").toUpperCase();i[c]={name:c,label:b.name,description:"",type:"datasource",pluginId:b.meta.id,pluginName:b.meta.name},a.datasource="${"+c+"}",h["datasource"+b.meta.id]={type:"datasource",id:b.meta.id,name:b.meta.name,version:b.meta.info.version||"1.0.0"}}))},l=0,m=a.rows;l<m.length;l++){var n=m[l];c["default"].each(n.panels,function(a){void 0!==a.datasource&&k(a);var c=b["default"].panels[a.type];c&&(h["panel"+c.id]={type:"panel",id:c.id,name:c.name,version:c.info.version})})}for(var o=0,p=a.templating.list;o<p.length;o++){var q=p[o];"query"===q.type&&(k(q),q.options=[],q.current={},q.refresh=1)}for(var r=0,s=a.annotations.list;r<s.length;r++){var t=s[r];k(t)}return h.grafana={type:"grafana",id:"grafana",name:"Grafana",version:b["default"].buildInfo.version},Promise.all(j).then(function(){c["default"].each(i,function(a,b){g.push(a)});for(var b=0,d=a.templating.list;b<d.length;b++){var e=d[b];if("constant"===e.type){var f="VAR_"+e.name.replace(" ","_").toUpperCase();g.push({name:f,type:"constant",label:e.label||e.name,value:e.current.value,description:""}),e.query="${"+f+"}",e.options[0]=e.current={value:e.query,text:e.query}}}h=c["default"].map(h,function(a){return a});var j={};return j.__inputs=g,j.__requires=h,c["default"].defaults(j,a),j})["catch"](function(a){return console.log("Export failed:",a),{error:a}})},a}(),a("DashboardExporter",e)}}});