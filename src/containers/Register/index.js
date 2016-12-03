import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import * as AccountActions from '../../actions/account'

class Register extends Component {
  constructor(props) {
    super(props)
    this.onSubmit = this.onSubmit.bind(this)
  }

  onSubmit(event) {
    event.preventDefault()

    const username = this.refs.username.value
    const email = this.refs.email.value
    const password = this.refs.password.value

    this.props.actions.register(username, email, password)
  }

  render() {
    return (
      <div id="main">
        <div className="container">
          <div className="row">
            <div className="col-md-6 col-md-offset-3">
              <form onSubmit={this.onSubmit}>
                <h2 className="form-register-heading">Register</h2>
                <div className="form-group">
                  <label htmlFor="username">Username</label>
                  <input type="text" className="form-control" name="username" ref="username" placeholder="Username" />
                </div>
                <div className="form-group">
                  <label htmlFor="email">Email address</label>
                  <input type="email" className="form-control" name="email" ref="email" placeholder="Email" />
                </div>
                <div className="form-group">
                  <label htmlFor="password">Password</label>
                  <input type="password" className="form-control" name="password" ref="password" placeholder="Password" />
                </div>
                <button type="submit" className="btn btn-default">注册</button>
              </form>
            </div>
          </div>
        </div>
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
)(Register)
