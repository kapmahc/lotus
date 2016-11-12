import SignIn from './SignIn'
import SignUp from './SignUp'
import Confirm from './Confirm'
import ForgotPassword from './ForgotPassword'
import ResetPassword from './ResetPassword'
import Unlock from './Unlock'
import Logs from './Logs'

export default [
  {
    name: 'users.sign-in',
    path: '/users/sign-in',
    component: SignIn
  },
  {
    name: 'users.sign-up',
    path: '/users/sign-up',
    component: SignUp
  },
  {
    name: 'users.forgot-password',
    path: '/users/forgot-password',
    component: ForgotPassword
  },
  {
    name: 'users.reset-password',
    path: '/users/reset-password',
    component: ResetPassword
  },
  {
    name: 'users.confirm',
    path: '/users/confirm',
    component: Confirm
  },
  {
    name: 'users.unlock',
    path: '/users/unlock',
    component: Unlock
  },

  {
    name: 'users.logs',
    path: '/users/logs',
    component: Logs
  }
]
