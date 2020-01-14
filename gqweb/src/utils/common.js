module.exports = {
  // 返回一个分时走势图绘画规则的对象
  getOptionTimeSharing1: function (type, width) {
    return {
      name: type || 'time-sharing',
      width: width || 'auto',
      height: 200,
      axis: {
        row: 4,
        col: 4,
        paddingTop: 0,
        paddingBottom: 0,
        paddingLeft: 0,
        paddingRight: 0,
        color: '#cdcdcd'
      },
      xAxis: {
        data: []
      },
      yAxis: [
        {
          type: 'line',
          lineColor: '#2F6098',
          background: 'rgba(53,125,222,0.1)',
          data: []
        },
        {
          type: 'line',
          lineColor: '#A96F3E',
          data: []
        }
      ],
      callback: function (time) {
        // var page = getCurrentPages()
        // page = page[page.length - 1]
        // page.setData({
        //   ts1RenderTime: time
        // })
        console.log('getOptionTimeSharing1回调函数')
      }
    }
  },

  // 返回一个分时图成交量绘画规则的对象
  getOptionTimeSharing2: function (type, width) {
    return {
      name: type || 'time-sharing-b',
      width: width || 'auto',
      height: 80,
      axis: {
        row: 2,
        col: 4,
        showEdg: true,
        showX: true,
        showY: true,
        paddingTop: 5,
        paddingBottom: 14,
        paddingLeft: 0,
        paddingRight: 0,
        color: '#cdcdcd'
      },
      xAxis: {
        times: ['09:30', '15:00'],
        data: []
      },
      yAxis: [
        {
          type: 'bar',
          color: [],
          data: [],
          showMax: true
        }
      ],
      callback: function (time) {
        // var page = getCurrentPages()
        // page = page[page.length - 1]
        // page.setData({
        //   ts2RenderTime: time
        // })
        console.log('getOptionTimeSharing2的回调函数')
      }
    }
  },

  // 获得绘制日K价格走势图规则对象
  getOptionKline1: function (type) {
    return {
      name: type || 'ddk',
      width: 'auto',
      height: 160,
      // average: [5, 10, 20],
      axis: {
        row: 4,
        col: 5,
        showX: false,
        showY: true,
        showEdg: true,
        paddingTop: 0,
        paddingBottom: 10,
        paddingLeft: 0,
        paddingRight: 0,
        color: '#ababab'
      },
      xAxis: {
        data: [],
        averageLabel: []
      },
      yAxis: [],
      callback: function (time) {

      }
    }
  },

  // 获得绘制日K买卖走势图规则对象
  getOptionKline2: function (type) {
    return {
      name: type || 'dk',
      width: 'auto',
      height: 80,
      average: [5, 10, 20],
      axis: {
        row: 1,
        col: 5,
        showX: false,
        showY: true,
        showEdg: true,
        paddingTop: 0,
        paddingBottom: 14,
        paddingLeft: 0,
        paddingRight: 0,
        color: '#cdcdcd'
      },
      xAxis: {
        times: [],
        data: [],
        averageLabel: []
      },
      yAxis: [],
      callback: function (time) {

      }
    }
  }
}
