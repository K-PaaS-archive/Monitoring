/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["app/features/panel/panel_ctrl","app/features/panel/metrics_panel_ctrl","app/features/panel/query_ctrl","app/core/config"],function(a){function b(a){f["default"].bootData.user.lightTheme?System["import"](a.light+"!css"):System["import"](a.dark+"!css")}var c,d,e,f;return a("loadPluginCss",b),{setters:[function(a){c=a},function(a){d=a},function(a){e=a},function(a){f=a}],execute:function(){a("PanelCtrl",c.PanelCtrl),a("MetricsPanelCtrl",d.MetricsPanelCtrl),a("QueryCtrl",e.QueryCtrl)}}});