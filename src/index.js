// import 'babel-polyfill'
import React from 'react'
import ReactDom from 'react-dom'
import { Router, Route, IndexRoute, browserHistory } from 'react-router'
import { syncHistoryWithStore } from 'react-router-redux'
import { Provider } from 'react-redux'

import configureStore from './store'

import App from './containers/App'
import Home from './containers/Home'
import Register from './containers/Register'
import Signin from './containers/Signin'
import Signout from './containers/Signout'

const store = configureStore()
const history = syncHistoryWithStore(browserHistory, store)

ReactDom.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <IndexRoute component={Home}/>
        <Route path="/signup" component={Register}></Route>
        <Route path="/signin" component={Signin}></Route>
        <Route path="/signout" component={Signout}></Route>
        <Route path="/u/:username" component={Home}></Route>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('app')
)
