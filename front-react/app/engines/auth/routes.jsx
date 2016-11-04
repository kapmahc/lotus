import React from 'react'
import {Route} from 'react-router'

import {SignIn, SignUp, ForgotPassword, ChangePassword, Confirm, Unlock} from './users'

export default [
  <Route key="users.sign-in" path="users/sign-in" component={SignIn}/>,
  <Route key="users.sign-up" path="users/sign-up" component={SignUp}/>,
  <Route key="users.forgot-password" path="users/forgot-password" component={ForgotPassword}/>,
  <Route key="users.change-password" path="users/change-password" component={ChangePassword}/>,
  <Route key="users.confirm" path="users/confirm" component={Confirm}/>,
  <Route key="users.unlock" path="users/unlock" component={Unlock}/>
]
