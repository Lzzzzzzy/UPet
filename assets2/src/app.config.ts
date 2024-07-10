export default defineAppConfig({
  pages: ['pages/index/index', 'pages/my/index', 'pages/calendar/index'],
  window: {
    backgroundColor: '#fff',
    backgroundTextStyle: 'light',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'Taro3',
    navigationBarTextStyle: 'black',
    navigationStyle: 'custom'
  },
  subPackages: [
    {
      root: 'package',
      pages: ['package-add-pet/index', 'package-add-todo/index']
    }
  ],
  tabBar: {
    custom: true,
    color: '#000000',
    selectedColor: '#FF0000',
    list: [
      {
        pagePath: 'pages/index/index',
        text: '首页'
      },
      {
        pagePath: 'pages/calendar/index',
        text: '日历'
      },
      {
        pagePath: 'pages/my/index',
        text: '个人中心'
      }
    ]
  }
});
