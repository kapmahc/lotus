import SignIn from './SignIn'
import SignUp from './SignUp'

export default {
  state: {
    count: 0
  },
  mutations: {
    increment: state => state.count++,
    decrement: state => state.count--
  },
  routes: [
    {
      name: 'users.sign-in',
      path: '/users/sign-in',
      component: SignIn
    },
    {
      name: 'users.sign-up',
      path: '/users/sign-up',
      component: SignUp
    }
  ]
}
