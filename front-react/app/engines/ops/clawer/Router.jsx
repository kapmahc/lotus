import React from 'react'
import {Route} from 'react-router'

import {SignIn as LineSignIn, Index as LineIndex} from './line'

const Widget = (
  <Route path="ops/clawer">
    <Route path="line" component={LineIndex}/>
    <Route path="line/sign-in" component={LineSignIn}/>
  </Route>
)

export default Widget
