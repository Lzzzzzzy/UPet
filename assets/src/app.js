import { createApp } from 'vue'
import { Input, Button, Popup, Calendar, Price, CalendarCard, Steps, Grid, GridItem, Badge } from '@nutui/nutui-taro'
import { IconFont } from '@nutui/icons-vue-taro'
import './app.scss'
import "@nutui/nutui-taro/dist/style.css";

const App = createApp({
  onShow (options) {},
  // 入口组件不需要实现 render 方法，即使实现了也会被 taro 所覆盖
})
App.use(Button).use(Input).use(Popup).use(Calendar).use(Price).use(IconFont).use(CalendarCard).use(Steps).use(Grid).use(GridItem).use(Badge);

export default App
