/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["../../../../../test/lib/common","jquery","../graph_tooltip"],function(a){function b(a,b){var d={};d.ctrl=f.ctrl,d.ctrl.panel={tooltip:{shared:!0},legend:{},stack:!1},d.setup=function(a){d.setupFn=a},c.describe(a,function(){c.beforeEach(function(){d.setupFn();var a=new e["default"](g,h,f);d.results=a.getMultiSeriesPlotHoverInfo(d.data,d.pos)}),b(d)})}var c,d,e,f,g,h;return{setters:[function(a){c=a},function(a){d=a},function(a){e=a}],execute:function(){f={appEvent:c.sinon.spy(),onAppEvent:c.sinon.spy(),ctrl:{}},g=d["default"]("<div></div>"),h={},b("steppedLine false, stack false",function(a){a.setup(function(){a.data=[{data:[[10,15],[12,20]],lines:{}},{data:[[10,2],[12,3]],lines:{}}],a.pos={x:11}}),c.it("should return 2 series",function(){c.expect(a.results.length).to.be(2)}),c.it("should add time to results array",function(){c.expect(a.results.time).to.be(10)}),c.it("should set value and hoverIndex",function(){c.expect(a.results[0].value).to.be(15),c.expect(a.results[1].value).to.be(2),c.expect(a.results[0].hoverIndex).to.be(0)})}),b("one series is hidden",function(a){a.setup(function(){a.data=[{data:[[10,15],[12,20]]},{data:[]}],a.pos={x:11}})}),b("steppedLine false, stack true, individual false",function(a){a.setup(function(){a.data=[{data:[[10,15],[12,20]],lines:{},datapoints:{pointsize:2,points:[[10,15],[12,20]]},stack:!0},{data:[[10,2],[12,3]],lines:{},datapoints:{pointsize:2,points:[[10,2],[12,3]]},stack:!0}],a.ctrl.panel.stack=!0,a.pos={x:11}}),c.it("should show stacked value",function(){c.expect(a.results[1].value).to.be(17)})}),b("steppedLine false, stack true, individual false, series stack false",function(a){a.setup(function(){a.data=[{data:[[10,15],[12,20]],lines:{},datapoints:{pointsize:2,points:[[10,15],[12,20]]},stack:!0},{data:[[10,2],[12,3]],lines:{},datapoints:{pointsize:2,points:[[10,2],[12,3]]},stack:!1}],a.ctrl.panel.stack=!0,a.pos={x:11}}),c.it("should not show stacked value",function(){c.expect(a.results[1].value).to.be(2)})}),b("steppedLine false, stack true, individual true",function(a){a.setup(function(){a.data=[{data:[[10,15],[12,20]],lines:{},datapoints:{pointsize:2,points:[[10,15],[12,20]]},stack:!0},{data:[[10,2],[12,3]],lines:{},datapoints:{pointsize:2,points:[[10,2],[12,3]]},stack:!1}],a.ctrl.panel.stack=!0,a.ctrl.panel.tooltip.value_type="individual",a.pos={x:11}}),c.it("should not show stacked value",function(){c.expect(a.results[1].value).to.be(2)})})}}});