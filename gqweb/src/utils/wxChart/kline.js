/**
 * Created by ChenChao on 2017/1/4.
 */

var common = require('./common');
var axis = require('./axis-k')();

module.exports = function (canvasId, Ctx) {
    return {
        unit: 20,  //不同K线，X轴的单位，默认60
        canvasId: canvasId,
        averageColors: ['#6A6969', '#F69A43', '#EDB2EB'],
        ctx: Ctx,
        canvasWidth: 0,
        canvasHeight: 0,
        paddingTop: 0,
        paddingBottom: 0,
        paddingLeft: 20,
        paddingRight: 0,
        options: null,
        dataStore: null,
        index: 0,
        yMax: 0,
        yMin: 1000000,
        isNew: false,
        offsetX: 0,
        startTime: '',
        endTime: '',
        txtColor: 'black',
        axisObj: null,
        reserve_left: 50,
        init: function (options) {
            // this.ctx = wx.createCanvasContext(this.canvasId);
            this.initConfig(options);
            return this;
        },
        initConfig: function (options) {
            var that = this;
            var axis = options.axis;
            var w = options.width;
            var h = options.height;
            if(w === 'auto') {
                w = that.canvasWidth = document.body.clientWidth*0.7;
            }
            if(h === 'auto'){
                h = 225;
            }
            this.canvasWidth = w;
            this.canvasHeight = h;
            this.unit = options.unit || this.unit;
            this.paddingTop = axis.paddingTop;
            this.paddingBottom = axis.paddingBottom;
            this.paddingLeft = axis.paddingLeft;
            this.paddingRight = axis.paddingRight;
            this.dataStore = options;
        },
        metaData1: function (origin, options) {
            var dataStore = options;
            var yMax = this.yMax = 0;
            var yMin = this.yMin = 1000000;
            var xAxis = dataStore.xAxis;
            var yAxis = dataStore.yAxis;
            var historyStep = Math.max.apply(null) || 0;
            var originData = origin.slice(0);
            var odl = origin.length;

            //处理小于 unit 条数据的情况
            dataStore.isNew = this.isNew;
            dataStore.offsetX = this.offsetX;

            var data = originData;

            var totalData = data
            yAxis.push({  //创建蜡烛趋势
                type: 'candle',
                gap: 0,
                showLabel: true,
                data_h: [],  //最高
                data_l: [],  //最低
                data_s: [],  //开盘
                data_c: [],  //收盘（现价）
                yin_yang: [] //阴阳: true为阳，false为阴
            });

            //"2017-01-06,17.00,17.10,17.40,16.90,1462083,25.1亿,2.95%" [日期，开盘价，现价，最高价，最低价，成交量，成交额，振幅]
            totalData.forEach(function (item, index) {
                var d = item.split(',');
                var t = d[0]; //时间
                var s = d[1]/1; //开盘价
                var h = d[2]/1; //最高价
                var l = d[3]/1; //最低价
                var c = d[4]/1; //现价
                var dataIndex = index;
                var candleOpt = yAxis[0];
                yMin = Math.min(l == 0 ? yMin : l, yMin);
                xAxis.data[dataIndex] = t;
                candleOpt['data_h'].push(h);
                candleOpt['data_l'].push(l);
                candleOpt['data_s'].push(s);
                candleOpt['data_c'].push(c);
                candleOpt['yin_yang'].push(c >= s);
                yMax = Math.max(h, yMax);

            });

            dataStore.axis.yMax = this.yMax = yMax;
            dataStore.axis.yMin = this.yMin = yMin;
            dataStore.unit = this.unit;
            dataStore.canvasWidth = this.canvasWidth;
            dataStore.canvasHeight = this.canvasHeight;
            this.setOptions(dataStore);
        },
        setOptions: function (options) {
            this.options = options;
        },
        axis: function (ctx, options) {
            this.axisObj = axis.init(ctx, options);
        },
        bezierLine: function (option) {
            common.bezierLine.call(this, option);
        },
        line: function (option) {
            if(option.hide){
                return;
            }
            var that = this;
            var ctx = this.ctx;
            var canvasHeight = this.canvasHeight;
            var canvasWidth = this.canvasWidth;
            var unit = this.unit;
            var step = (canvasWidth - this.paddingLeft - this.paddingRight) /  this.unit;
            var areaH = canvasHeight - this.paddingBottom - this.paddingTop;
            var max = this.yMax;
            var min = this.yMin;
            if(option.isBottomBar){
                min = 0;
            }
            var data = [];
            option.xAxis.data.map(function (item, index) {
                var d = option.data[index];
                var value = areaH - areaH * (d - min) / (max - min) + that.paddingTop;
                data.push([index * step - that.paddingLeft + step / 2, value]);
            });
            var barW = (canvasWidth - this.paddingLeft - this.paddingRight) / this.unit;
            if(this.offsetX > 0){
                ctx.translate(-(this.unit - this.offsetX) * barW + 1, 0);
            }
            ctx.beginPath();
            data.map(function (item, index) {
                var x0 = item[0];
                var x1 = item[1];
                if(option.isNew){
                    var startIndex = unit - (option.odl - option.val) - 1;
                    if(index == startIndex){
                        ctx['moveTo'](x0, x1);
                    }
                    if(index > startIndex){
                        ctx['lineTo'](x0, x1);
                    }
                }else{
                    ctx[index === 0 ? 'moveTo' : 'lineTo'](x0, x1);
                }
            });
            ctx.lineWidth = 1;
            ctx.lineCap = 'square';
            ctx.strokeStyle = option.lineColor;
            ctx.stroke();
            if(this.offsetX > 0){
                ctx.translate((this.unit - this.offsetX) * barW + 1, 0);
            }
        },
        bar: function (option) {
            var startTime = +new Date();
            var data = option.data;
            var ctx = this.ctx;
            var canvasHeight = this.canvasHeight;
            var canvasWidth = this.canvasWidth;
            var pb = this.paddingBottom;
            var barW = (canvasWidth - this.paddingLeft - this.paddingRight- this.reserve_left) / this.unit;
            barW -= 1;
            var max = Math.max.apply(null, data);
            var step = (canvasHeight - this.paddingTop - pb) / max;
            if(this.offsetX > 0){
                ctx.translate(-(this.unit - this.offsetX) * (barW + 1), 0);
            }
            data.forEach(function (item, index) {
                var barH = item * step;
                var color = option.color[index];
                /*if(color === 'red'){
                    ctx.setLineWidth(1);
                    ctx.strokeStyle = color);
                    ctx.strokeRect(index * barW - 2, canvasHeight - pb - barH, barW, barH);
                }else{*/
                    ctx.beginPath();
                    ctx.lineWidth = barW;
                    ctx.moveTo(index * (barW + 1) + barW/2 + this.reserve_left , canvasHeight - pb);
                    ctx.lineTo(index * (barW + 1) + barW/2 + this.reserve_left , canvasHeight - pb - barH);
                    ctx.strokeStyle = color;
                    ctx.stroke();
                //}
            });
            if(this.offsetX > 0){
                ctx.translate((this.unit - this.offsetX) * (barW + 1), 0);
            }
            if(option.complete){
                option.complete(+new Date() - startTime);
            }
            // if(option.showMax){
            //     ctx.setFillStyle(this.txtColor);
            //     ctx.fillText(common.metaUnit(max), this.paddingLeft + 3, this.paddingTop + 30);
            // }
        },
        candle: function (option) {
            var that = this;
            var ctx = this.ctx;
            var canvasWidth = this.canvasWidth;
            var canvasHeight = this.canvasHeight;
            var dataX = option.xAxis.data;
            var data_h = option.data_h;
            var data_l = option.data_l;
            var data_s = option.data_s;
            var data_c = option.data_c;
            var yin_yang = option.yin_yang;

            var max = this.yMax;//Math.max.apply(null, data_h);
            var min = this.yMin;//Math.min.apply(null, data_l);
            var areaH = canvasHeight - this.paddingBottom - this.paddingTop;
            var areaUnit = areaH / (max - min);

            var barW = (canvasWidth - this.paddingLeft - this.paddingRight - this.reserve_left) / this.unit;
            var yMin = this.yMin;
            var gap = option.gap;
            if(this.offsetX > 0){
                ctx.translate(-(this.unit - this.offsetX) * barW, 0);
            }
            ctx.translate(0, that.paddingTop);
            data_h.forEach(function (time, index) {
                var h = data_h[index];
                var l = data_l[index];
                var s = data_s[index];
                var c = data_c[index];
                var yy = yin_yang[index];
                var cx  = that.paddingLeft + (gap + barW) * index + that.reserve_left;
                if (index == data_h.length-1){
                  common.candle(ctx, cx, barW, h, l, s, c, yy, max, min, areaH, true);
                }
                else{
                common.candle(ctx, cx, barW, h, l, s, c, yy, max, min, areaH, false);
              }
            });
            ctx.translate(0, -that.paddingTop);
            if(this.offsetX > 0){
                ctx.translate((this.unit - this.offsetX) * barW, 0);
            }
        },
        clear: function () {
            this.ctx.clearRect(0, 0, this.canvasWidth, this.canvasHeight);
        },
        draw: function (opt) {
            var that = this;
            var ctx = this.ctx;
            var options = this.options;
            if(!options){
                console.log('Warn: No setting options!');
                return;
            }

            var xAxis = options.xAxis;
            var startTime = +new Date();
            this.clear();
            this.axis(ctx, options);
            options.yAxis.map(function (option, index) {
                option.xAxis = xAxis;
                that[option.type](option);
            });
            this.axisObj.drawYUnit();
            //this.ctx.draw();
            options.callback && options.callback(+new Date() - startTime);
        }
    };
};
