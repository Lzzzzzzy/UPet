export default defineAppConfig({
  pages: ['pages/index/index', 'pages/my/index'],
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
        pagePath: 'pages/my/index',
        text: '我的'
      }
    ]
  }
});
