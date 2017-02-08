/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["../../../../../test/lib/common","../../../../../test/specs/helpers","../module"],function(a){var b,c,d;return{setters:[function(a){b=a},function(a){c=a},function(a){d=a}],execute:function(){b.describe("SingleStatCtrl",function(){function a(a,c){b.describe(a,function(){e.setup=function(a){b.beforeEach(b.angularMocks.module("grafana.services")),b.beforeEach(b.angularMocks.module("grafana.controllers")),b.beforeEach(e.providePhase()),b.beforeEach(e.createPanelController(d.SingleStatCtrl)),b.beforeEach(function(){a();var b=[{target:"test.cpu1",datapoints:e.datapoints}];e.ctrl.onDataReceived(b),e.data=e.ctrl.data})},c(e)})}var e=new c["default"].ControllerTestContext;a("with defaults",function(a){a.setup(function(){a.datapoints=[[10,1],[20,2]]}),b.it("Should use series avg as default main value",function(){b.expect(a.data.value).to.be(15),b.expect(a.data.valueRounded).to.be(15)}),b.it("should set formated falue",function(){b.expect(a.data.valueFormated).to.be("15")})}),a("MainValue should use same number for decimals as displayed when checking thresholds",function(a){a.setup(function(){a.datapoints=[[99.999,1],[99.99999,2]]}),b.it("Should be rounded",function(){b.expect(a.data.value).to.be(99.999495),b.expect(a.data.valueRounded).to.be(100)}),b.it("should set formated falue",function(){b.expect(a.data.valueFormated).to.be("100")})}),a("When value to text mapping is specified",function(a){a.setup(function(){a.datapoints=[[9.9,1]],a.ctrl.panel.valueMaps=[{value:"10",text:"OK"}]}),b.it("value should remain",function(){b.expect(a.data.value).to.be(9.9)}),b.it("round should be rounded up",function(){b.expect(a.data.valueRounded).to.be(10)}),b.it("Should replace value with text",function(){b.expect(a.data.valueFormated).to.be("OK")})}),a("When range to text mapping is specifiedfor first range",function(a){a.setup(function(){a.datapoints=[[41,50]],a.ctrl.panel.mappingType=2,a.ctrl.panel.rangeMaps=[{from:"10",to:"50",text:"OK"},{from:"51",to:"100",text:"NOT OK"}]}),b.it("Should replace value with text OK",function(){b.expect(a.data.valueFormated).to.be("OK")})}),a("When range to text mapping is specified for other ranges",function(a){a.setup(function(){a.datapoints=[[65,75]],a.ctrl.panel.mappingType=2,a.ctrl.panel.rangeMaps=[{from:"10",to:"50",text:"OK"},{from:"51",to:"100",text:"NOT OK"}]}),b.it("Should replace value with text NOT OK",function(){b.expect(a.data.valueFormated).to.be("NOT OK")})})})}}});