import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import store from '../store/index'
import HomePage from '../views/HomePage.vue'
import Register from '../views/Register.vue'
import Login from '../views/Login.vue'
import Account from '../views/Account.vue'
import Bands from '../views/Bands.vue'
import Band from '../views/Band.vue'
import Rehearsal from '../views/Rehearsal.vue'
import MembersOverlay from '../components/MembersOverlay.vue'
import InstrumentOverlay from '../components/InstrumentsOverlay.vue'
import SongsListView from '../components/SongsListView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: HomePage
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/account',
    name: 'Account',
    component: Account
  },
  {
    path: '/bands',
    name: 'Bands',
    component: Bands
  },
  {
    path: '/bands/:bandId',
    name: 'Band',
    component: Band,
    props: true,
    children: [{
      path: 'members',
      name:'Members',
      component: MembersOverlay,
      props: true
    },
    {
      path: 'instruments',
      name: 'Instruments',
      component: InstrumentOverlay,
      props: true
    },
    {
      path: 'songs',
      name: 'Songs',
      component: SongsListView,
      props: true
    }]
  },
  {
    path: '/bands/:bandId/rehearsal/:userId?',
    name: 'Rehearsal',
    component: Rehearsal,
    props:true
  }
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  if (to.name !== 'Login' && to.name !== 'Register' && to.name !== 'Home' && !store.getters.isLoggedIn) next({ name: 'Login' })
  else next();
})

export default router
