import React, { Component } from 'react'

import style from './style.scss'

class Footer extends Component {
  constructor(props) {
    super(props)
  }

  render() {
    return (
      <div id="footer" className={style.footer}>
        <div className="ui container">
          <a href="/" className="">
            MeApex
          </a>
        </div>
      </div>
    )
  }
}

export default Footer
