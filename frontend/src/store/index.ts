import {createStore} from 'vuex';
import axios from 'axios';

const store = createStore({
  state: {
    band: null,
    user: {},
    credentials: "",
  },
  getters: {
    isLoggedIn: state => {
      return state.credentials != "" ? true : false
    },
    getCredentials: state => {
      return state.credentials
    },
    axios: state => {
      return axios.create({
        baseURL: "http://" + process.env.VUE_APP_BACKENDURL,
        headers: { 'Authorization': `Basic ${state.credentials}` }
      })
    }
  },
  mutations: {
    setBand(state, band) {
      state.band = band
    },
    setUser(state, user) {
      delete user.password
      state.user = user
    },
    setCredentials(state, { email, password }) {
      state.credentials = btoa(email + ':' + password)
    },
    clearCredentials(state) {
      state.credentials = ""
    }
  }
})

export default store;