import { createAction } from 'redux-actions'
import 'isomorphic-fetch'

import {
  AUTH_USER,
  UNAUTH_USER,
  REGISTER_SUCCESS,
  REGISTER_FAILURE,
  SIGNIN_SUCCESS,
  SIGNIN_FAILURE,
  SIGNOUT_SUCCESS,
  SIGNOUT_FAILURE
} from '../constants/ActionTypes'

import { searchParams } from '../utils/auxiliary'

const authUser = createAction(AUTH_USER)
const unauthUser = createAction(UNAUTH_USER)

export function register(username, email, password) {
  let params = {
    username: username,
    email: email,
    password: password
  }

  return dispatch => {
    fetch('/account/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: searchParams(params)
    })
    .then(response => {
      if (response.ok) {
        response.json().then(json => dispatch(authUser(json)))
      }
    })
    .catch(error => {
      console.log(error.message)
    })
  }
}

export function signin(account, password) {
  let params = {
    account: account,
    password: password
  }

  return dispatch => {
    fetch('/account/signin', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      body: searchParams(params)
    })
    .then(response => {
      if (response.ok) {
        response.json().then(json => {
          localStorage.setItem('token', json.token)
          dispatch(authUser(json))
        })
      }
    })
    .catch(error => {
      console.log(error.message)
    })
  }
}

export function signout() {
  return {
    type: UNAUTH_USER
  }
}

export function signinFromToken(token) {
  return dispatch => {
    return fetch('/api/me', {
      method: 'GET',
      headers: {
        'Accept': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    })
      .then(response => {
        if (response.ok) {
          response.json().then(json => dispatch(authUser(json)))
        }
      })
      .catch(error => {
        console.log(error.message)
      })
  }
}
