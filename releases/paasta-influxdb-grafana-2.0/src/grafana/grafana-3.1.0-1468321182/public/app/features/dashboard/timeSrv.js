/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

define(["angular","lodash","moment","app/core/config","app/core/utils/kbn","app/core/utils/datemath"],function(a,b,c,d,e,f){"use strict";var g=a.module("grafana.services");g.service("timeSrv",["$rootScope","$timeout","$routeParams","timer",function(a,d,g,h){var i=this;this.init=function(a){h.cancel_all(),this.dashboard=a,this.time=a.time,this._initTimeFromUrl(),this._parseTime(),this.dashboard.refresh&&this.setAutoRefresh(this.dashboard.refresh)},this._parseTime=function(){b.isString(this.time.from)&&this.time.from.indexOf("Z")>=0&&(this.time.from=c(this.time.from).utc()),b.isString(this.time.to)&&this.time.to.indexOf("Z")>=0&&(this.time.to=c(this.time.to).utc())},this._parseUrlParam=function(a){if(a.indexOf("now")!==-1)return a;if(8===a.length)return c.utc(a,"YYYYMMDD");if(15===a.length)return c.utc(a,"YYYYMMDDTHHmmss");if(!isNaN(a)){var b=parseInt(a);return c.utc(b)}return null},this._initTimeFromUrl=function(){g.from&&(this.time.from=this._parseUrlParam(g.from)||this.time.from),g.to&&(this.time.to=this._parseUrlParam(g.to)||this.time.to)},this.setAutoRefresh=function(a){if(this.dashboard.refresh=a,a){var b=e.interval_to_ms(a),c=b-Date.now()%b;d(function(){i.start_scheduled_refresh(b),i.refreshDashboard()},c)}else this.cancel_scheduled_refresh()},this.refreshDashboard=function(){a.$broadcast("refresh")},this.start_scheduled_refresh=function(a){i.cancel_scheduled_refresh(),i.refresh_timer=h.register(d(function(){i.start_scheduled_refresh(a),i.refreshDashboard()},a))},this.cancel_scheduled_refresh=function(){h.cancel(this.refresh_timer)},this.setTime=function(e,f){b.extend(this.time,e),!f&&c.isMoment(e.to)?(this.old_refresh=this.dashboard.refresh||this.old_refresh,this.setAutoRefresh(!1)):this.old_refresh&&this.old_refresh!==this.dashboard.refresh&&(this.setAutoRefresh(this.old_refresh),this.old_refresh=null),a.appEvent("time-range-changed",this.time),d(this.refreshDashboard,0)},this.timeRangeForUrl=function(){var a=this.timeRange(!1);return b.isString(a.to)&&a.to.indexOf("now")&&(a=this.timeRange()),c.isMoment(a.from)&&(a.from=a.from.valueOf()),c.isMoment(a.to)&&(a.to=a.to.valueOf()),a},this.timeRange=function(a){var b=c.isMoment(this.time.from)?c(this.time.from):this.time.from,d=c.isMoment(this.time.to)?c(this.time.to):this.time.to;return a!==!1&&(b=f.parse(b,!1),d=f.parse(d,!0)),{from:b,to:d}}}])});