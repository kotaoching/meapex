import React, { Component } from 'react'
import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'

import { Container, Grid, Form, Button } from 'semantic-ui-react'

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
        <Container>
          <Grid>
            <Grid.Row centered>
              <Grid.Column mobile={16} tablet={8} computer={6}>
                <Form onSubmit={this.onSubmit}>
                  <Form.Field>
                    <label>Username</label>
                    <input type="text" name="username" ref="username" placeholder='Username' />
                  </Form.Field>
                  <Form.Field>
                    <label>Email</label>
                    <input type="email" name="email" ref="email" placeholder='Email' />
                  </Form.Field>
                  <Form.Field>
                    <label>Password</label>
                    <input type="password" name="account" ref="password" placeholder='Password' />
                  </Form.Field>
                  <Button primary type='submit'>注册</Button>
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
)(Register)
