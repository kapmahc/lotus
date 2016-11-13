import SignIn from './users/SignIn'
import SignUp from './users/SignUp'
import Confirm from './users/Confirm'
import ForgotPassword from './users/ForgotPassword'
import ResetPassword from './users/ResetPassword'
import Unlock from './users/Unlock'
import Logs from './users/Logs'
import ChangePassword from './users/ChangePassword'
import UserInfo from './users/Info'

import AdminAuthor from './admin/Author'
import AdminBase from './admin/Base'
import AdminI18n from './admin/I18n'
import AdminSeo from './admin/Seo'
import AdminSmtp from './admin/Smtp'
import AdminStatus from './admin/Status'
import AdminUsers from './admin/Users'

import LeavewordsNew from './leavewords/New'
import LeavewordsIndex from './leavewords/Index'

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
  },

  {
    name: 'admin.author',
    path: '/admin/author',
    component: AdminAuthor
  },
  {
    name: 'admin.base',
    path: '/admin/base',
    component: AdminBase
  },
  {
    name: 'admin.i18n',
    path: '/admin/i18n',
    component: AdminI18n
  },
  {
    name: 'admin.seo',
    path: '/admin/seo',
    component: AdminSeo
  },
  {
    name: 'admin.smtp',
    path: '/admin/smtp',
    component: AdminSmtp
  },
  {
    name: 'admin.status',
    path: '/admin/status',
    component: AdminStatus
  },
  {
    name: 'admin.users',
    path: '/admin/users',
    component: AdminUsers
  },

  {
    name: 'leavewords.new',
    path: '/leavewords/new',
    component: LeavewordsNew
  },
  {
    name: 'leavewords.index',
    path: '/leavewords/index',
    component: LeavewordsIndex
  }
]
