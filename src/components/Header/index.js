import React, { Component } from 'react'
import { Link } from 'react-router'

import style from './style.scss'

class Header extends Component {
  constructor(props) {
    super(props)
  }

  renderLoggedIn() {
    return (
      <ul className="nav navbar-nav navbar-right">
        <li className="dropdown">
          <a href="#" className="dropdown-toggle" data-toggle="dropdown" role="button" aria-haspopup="true" aria-expanded="false">{this.props.account.user.username}<span className="caret"></span></a>
          <ul className="dropdown-menu">
            <li><a href="#">我的主页</a></li>
            <li><a href="#">设置</a></li>
            <li><a href="#">退出</a></li>
          </ul>
        </li>
      </ul>
    )
  }

  renderLoggedOut() {
    return (
      <ul className="nav navbar-nav navbar-right">
        <li><Link to="/signup">注册</Link></li>
        <li><Link to="/signin">登录</Link></li>
      </ul>
    )
  }

  render() {
    return (
      <div id="header" className={style.header}>
        <div className={`navbar navbar-default ${style.nav}`}>
          <div className="container">
            <div className="navbar-header">
              <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#main-nav-menu" aria-expanded="false">
                <span className="sr-only">Toggle navigation</span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
                <span className="icon-bar"></span>
              </button>
              <a className="navbar-brand" href="/">MeApex</a>
            </div>

            <div className="collapse navbar-collapse" id="main-nav-menu">
              <ul className="nav navbar-nav">
                <li className="active">
                  <Link to="/">首页</Link>
                </li>
              </ul>
              {this.props.account.loggedIn ? this.renderLoggedIn() : this.renderLoggedOut()}
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default Header
