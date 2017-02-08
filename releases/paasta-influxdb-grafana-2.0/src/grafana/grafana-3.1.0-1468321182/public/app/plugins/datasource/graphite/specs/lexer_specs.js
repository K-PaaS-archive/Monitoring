/*! grafana - v3.1.0-1468321182 - 2016-07-12
 * Copyright (c) 2016 Torkel Ödegaard; Licensed Apache-2.0 */

System.register(["test/lib/common","../lexer"],function(a){var b,c;return{setters:[function(a){b=a},function(a){c=a}],execute:function(){b.describe("when lexing graphite expression",function(){b.it("should tokenize metric expression",function(){var a=new c.Lexer("metric.test.*.asd.count"),d=a.tokenize();b.expect(d[0].value).to.be("metric"),b.expect(d[1].value).to.be("."),b.expect(d[2].type).to.be("identifier"),b.expect(d[4].type).to.be("identifier"),b.expect(d[4].pos).to.be(13)}),b.it("should tokenize metric expression with dash",function(){var a=new c.Lexer("metric.test.se1-server-*.asd.count"),d=a.tokenize();b.expect(d[4].type).to.be("identifier"),b.expect(d[4].value).to.be("se1-server-*")}),b.it("should tokenize metric expression with dash2",function(){var a=new c.Lexer("net.192-168-1-1.192-168-1-9.ping_value.*"),d=a.tokenize();b.expect(d[0].value).to.be("net"),b.expect(d[2].value).to.be("192-168-1-1")}),b.it("should tokenize metric expression with equal sign",function(){var a=new c.Lexer("apps=test"),d=a.tokenize();b.expect(d[0].value).to.be("apps=test")}),b.it("simple function2",function(){var a=new c.Lexer("offset(test.metric, -100)"),d=a.tokenize();b.expect(d[2].type).to.be("identifier"),b.expect(d[4].type).to.be("identifier"),b.expect(d[6].type).to.be("number")}),b.it("should tokenize metric expression with curly braces",function(){var a=new c.Lexer("metric.se1-{first, second}.count"),d=a.tokenize();b.expect(d.length).to.be(10),b.expect(d[3].type).to.be("{"),b.expect(d[4].value).to.be("first"),b.expect(d[5].value).to.be(","),b.expect(d[6].value).to.be("second")}),b.it("should tokenize metric expression with number segments",function(){var a=new c.Lexer("metric.10.12_10.test"),d=a.tokenize();b.expect(d[0].type).to.be("identifier"),b.expect(d[2].type).to.be("identifier"),b.expect(d[2].value).to.be("10"),b.expect(d[4].value).to.be("12_10"),b.expect(d[4].type).to.be("identifier")}),b.it("should tokenize func call with numbered metric and number arg",function(){var a=new c.Lexer("scale(metric.10, 15)"),d=a.tokenize();b.expect(d[0].type).to.be("identifier"),b.expect(d[2].type).to.be("identifier"),b.expect(d[2].value).to.be("metric"),b.expect(d[4].value).to.be("10"),b.expect(d[4].type).to.be("number"),b.expect(d[6].type).to.be("number")}),b.it("should tokenize metric with template parameter",function(){var a=new c.Lexer("metric.[[server]].test"),d=a.tokenize();b.expect(d[2].type).to.be("identifier"),b.expect(d[2].value).to.be("[[server]]"),b.expect(d[4].type).to.be("identifier")}),b.it("should tokenize metric with question mark",function(){var a=new c.Lexer("metric.server_??.test"),d=a.tokenize();b.expect(d[2].type).to.be("identifier"),b.expect(d[2].value).to.be("server_??"),b.expect(d[4].type).to.be("identifier")}),b.it("should handle error with unterminated string",function(){var a=new c.Lexer("alias(metric, 'asd)"),d=a.tokenize();b.expect(d[0].value).to.be("alias"),b.expect(d[1].value).to.be("("),b.expect(d[2].value).to.be("metric"),b.expect(d[3].value).to.be(","),b.expect(d[4].type).to.be("string"),b.expect(d[4].isUnclosed).to.be(!0),b.expect(d[4].pos).to.be(20)}),b.it("should handle float parameters",function(){var a=new c.Lexer("alias(metric, 0.002)"),d=a.tokenize();b.expect(d[4].type).to.be("number"),b.expect(d[4].value).to.be("0.002")}),b.it("should handle bool parameters",function(){var a=new c.Lexer("alias(metric, true, false)"),d=a.tokenize();b.expect(d[4].type).to.be("bool"),b.expect(d[4].value).to.be("true"),b.expect(d[6].type).to.be("bool")})})}}});