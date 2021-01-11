import { mount } from "@vue/test-utils";
import Header from '@/components/Header.vue'
import { createStore } from "vuex";

const createVuexStore = (credentials:String) => {
  return createStore({
    getters: {
      isLoggedIn: state => {
        return state.credentials != "" ? true : false
      },
    },
    state: {
      credentials: credentials
    },
    mutations: {
      clearCredentials(state) {
        state.credentials = ""
      }
    }
  })
}

function factory(credentials:String) {
  const store = createVuexStore(credentials)
  return mount(Header, {
    global: {
      plugins: [store]
    }
  })
}

describe('Header.vue', () => {
  it('Should say log in', () => {
    const wrapper = factory("")

    expect(wrapper.text()).toContain("Log in")
    expect(wrapper.text()).not.toContain("Log out")
  })

  it('Should say log out', () => {
    const wrapper = factory("ingelogd")

    expect(wrapper.text()).not.toContain("Log in")
    expect(wrapper.text()).toContain("Account")
  })
});

