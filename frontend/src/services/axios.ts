import axios from 'axios'
import store from '../store/index'

export default () => {
  return axios.create({
    baseURL: "http://" + process.env.VUE_APP_BACKENDURL,
    headers: { 'Authorization': `Basic ${store.getters.getCredentials}` }
  })
}

// import axios from 'axios'
// import store from '../store/index'

// export default axios.create({
//     baseURL: "http://" + process.env.VUE_APP_BACKENDURL,
//     headers: { 'Authorization': `Basic ${store.getters.getCredentials}` }
//   })
