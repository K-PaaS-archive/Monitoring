/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

define(["./helpers","app/plugins/panel/graph/series_overrides_ctrl"],function(a){"use strict";describe("SeriesOverridesCtrl",function(){var b=new a.ControllerTestContext,c={};beforeEach(module("grafana.services")),beforeEach(module("grafana.controllers")),beforeEach(b.providePhase({popoverSrv:c})),beforeEach(inject(function(a,c){b.scope=a.$new(),b.scope.ctrl={refresh:sinon.spy(),render:sinon.spy(),seriesList:[]},b.scope.render=function(){},b.controller=c("SeriesOverridesCtrl",{$scope:b.scope})})),describe("When setting an override",function(){beforeEach(function(){b.scope.setOverride({propertyName:"lines"},{value:!0})}),it("should set override property",function(){expect(b.scope.override.lines).to.be(!0)}),it("should update view model",function(){expect(b.scope.currentOverrides[0].name).to.be("Lines"),expect(b.scope.currentOverrides[0].value).to.be("true")})}),describe("When removing overide",function(){it("click should include option and value index",function(){b.scope.setOverride(1,0),b.scope.removeOverride({propertyName:"lines"}),expect(b.scope.currentOverrides.length).to.be(0)})})})});