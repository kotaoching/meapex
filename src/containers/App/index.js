import '../../assets/styles/global.scss'

import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import * as AccountActions from '../../actions/account'

import Header from '../../components/Header'
import Footer from '../../components/Footer'

class App extends Component {
  constructor(props) {
    super(props)
  }

  componentWillMount() {
    const token = localStorage.getItem('token')
    if (token) {
      this.props.actions.signinFromToken(token)
    }
  }

  render() {
    const { account, actions, children } = this.props

    return (
      <div>
        <Header account={account} />
        {children}
        <Footer />
      </div>
    )
  }
}

function mapStateToProps(state) {
  return {
    account: state.account
  }
}

function mapDispatchToProps(dispatch) {
  return {
    actions: bindActionCreators(AccountActions, dispatch)
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
