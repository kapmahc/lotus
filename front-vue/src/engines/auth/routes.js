import SignIn from './SignIn'
import SignUp from './SignUp'
import Confirm from './Confirm'
import ForgotPassword from './ForgotPassword'
import Unlock from './Unlock'

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
    name: 'users.confirm',
    path: '/users/confirm',
    component: Confirm
  },
  {
    name: 'users.unlock',
    path: '/users/unlock',
    component: Unlock
  }
]
