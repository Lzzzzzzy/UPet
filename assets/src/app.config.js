export default defineAppConfig({
  pages: [
    'pages/index/index',
    // 'pages/classSelection/index'
  ],
  window: {
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  }
  // tabBar: {
  //   "color": "#a9b7b7",
  //   "selectedColor": "#11cd6e",
  //   "borderStyle": "white",
  //   "list": [
  //     {
  //       "pagePath": "pages/index/index",
  //       "text": "首页",
  //       // "iconPath": "static/tabs/home.png",
  //       // "selectedIconPath": "static/tabs/home-active.png"
  //     },
  //     {
  //       "pagePath": "pages/classSelection/index",
  //       "text": "自定义",
  //       // "iconPath": "static/tabs/orders.png",
  //       // "selectedIconPath": "static/tabs/orders-active.png"
  //     }
  //   ]
  // }
})
