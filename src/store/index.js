import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'

import { logger } from '../middleware'
import rootReducer from '../reducers'

const configureStore = (initialState) => {
  const create = window.devToolsExtension
    ? window.devToolsExtension()(createStore) : createStore

  const createStoreWithMiddleware = applyMiddleware(
    thunk,
    logger
  )(create)

  const store = createStoreWithMiddleware(rootReducer, initialState)

  if (module.hot) {
    module.hot.accept('../reducers', () => {
      const nextReducer = require('../reducers')
      store.replaceReducer(nextReducer)
    })
  }

  return store
}

export default configureStore
