import SignIn from './users/SignIn'
import SignUp from './users/SignUp'
import Confirm from './users/Confirm'
import ForgotPassword from './users/ForgotPassword'
import ResetPassword from './users/ResetPassword'
import Unlock from './users/Unlock'
import Logs from './users/Logs'
import ChangePassword from './users/ChangePassword'
import UserInfo from './users/Info'

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
    name: 'users.info',
    path: '/users/info',
    component: UserInfo
  },
  {
    name: 'users.change-password',
    path: '/users/change-password',
    component: ChangePassword
  },
  {
    name: 'users.logs',
    path: '/users/logs',
    component: Logs
  }
]
