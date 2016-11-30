import { handleActions } from 'redux-actions'

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

const initialState = {
  loggedIn: false,
  user: null,
  error: null
}

export default handleActions({
  AUTH_USER: (state, action) => {
    return {...state,
      loggedIn: true,
      user: action.payload.data
    }
  },

  UNAUTH_USER: (state, action) => {
    return {
      loggedIn: false,
      user: null,
      errorMessage: null
    }
  }
}, initialState)
