export default defineAppConfig({
  pages: ['pages/index/index', 'pages/my/index', 'pages/search/index'],
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
      pages: ['package-pet/index', 'package-todo/index', 'package-pet-management/index', 'package-pet-ownership/index', 'package-contact/index', 'package-register/index']
    }
  ],
  tabBar: {
    custom: true,
    color: '#000000',
    selectedColor: '#FF0000',
    list: [
      {
        pagePath: 'pages/index/index',
        text: ''
      },
      {
        pagePath: 'pages/search/index',
        text: ''
      },
      {
        pagePath: 'pages/my/index',
        text: ''
      }
    ]
  }
});
