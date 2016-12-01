import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { Container, Grid, Form, Button } from 'semantic-ui-react'

import * as AccountActions from '../../actions/account'

class Signin extends Component {
  constructor(props) {
    super(props)
    this.onSubmit = this.onSubmit.bind(this)
  }

  onSubmit(event) {
    event.preventDefault()

    const account = this.refs.account.value
    const password = this.refs.password.value

    this.props.actions.signin(account, password)
  }

  render() {
    return (
      <div id="main">
        <Container>
          <Grid>
            <Grid.Row centered>
              <Grid.Column mobile={16} tablet={8} computer={6}>
                <Form onSubmit={this.onSubmit}>
                 <Form.Field>
                  <label>Account</label>
                  <input type="text" name="account" ref="account" placeholder='Email or Username' />
                 </Form.Field>
                 <Form.Field>
                   <label>Password</label>
                   <input type="password" name="account" ref="password" placeholder='Password' />
                 </Form.Field>
                 <Button primary type='submit'>登录</Button>
                </Form>
              </Grid.Column>
            </Grid.Row>
          </Grid>
        </Container>
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
)(Signin)
