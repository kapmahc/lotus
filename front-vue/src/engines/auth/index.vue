<script>
import SignIn from './SignIn'
import SignUp from './SignUp'

export default {
  routes: [
    {path: '/users/sign-in', component: SignIn},
    {path: '/users/sign-up', component: SignUp}
  ]
}
</script>
