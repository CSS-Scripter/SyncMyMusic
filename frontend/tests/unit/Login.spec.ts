import { flushPromises, mount, VueWrapper } from '@vue/test-utils'
import Login from '@/views/Login.vue'
import mockAxios from '@/__mocks__/services/axios'
import { nextTick } from 'vue'
import { createStore } from 'vuex'

const user = {
  ID: 1,
  Name: "Test",
  Email: "test@email.nl"
}

let wrapper: VueWrapper<any>;

beforeEach(() => {
  wrapper = mount(Login)
});

const createVuexStore = () => {
  return createStore({
    mutations: {
      setCredentials(user) { return true },
      setUser(user) { return true },
      clearCredentials() { return true }
    }
  })
}

describe('Login.vue', () => {
  // Test if all elements are on the page.
  it('Has the h1 page title', () => {
    expect(wrapper.findAll('h1').length).toBe(1);
  });

  it('Has 1 pair of labels + corresponding inputboxes', () => {
    expect(wrapper.findAll('.inputbox').length).toBe(2);
    expect(wrapper.findAll('label').length).toBe(2);
  });

  it('Has 2 buttons with appropriate text', () => {
    let buttons = wrapper.findAll('button');
    expect(buttons.length).toBe(2);
    expect(buttons[0].text()).toBe('Cancel');
    expect(buttons[1].text()).toBe('Log In');
  });

  // Test Login(related) functionality
  it('Should log user in', async () => {
    const store = createVuexStore();
    const wrapper = mount(Login, {
      global: {
        plugins: [store]
      }
    })

    const mockListener = jest.spyOn(mockAxios(), 'get')
    mockListener.mockImplementationOnce(() => {
      return Promise.resolve({ data: user, test: true })
    })

    wrapper.vm.login = mockAxios().get

    wrapper.vm.email = "test@email.nl"
    wrapper.vm.password = "geheim"

    wrapper.find('#login-form').trigger('submit')
    await nextTick()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.find('#incorrect-credentials').exists()).toBeFalsy()
  })

  xit('Should show incorrect credentials message', async () => {
    const store = createVuexStore();
    const wrapper = mount(Login, {
      global: {
        plugins: [store]
      }
    })

    const mockListener = jest.spyOn(mockAxios(), 'get')
    mockListener.mockRejectedValue('')

    wrapper.vm.login = mockAxios().get

    wrapper.find('#login-form').trigger('submit')
    await flushPromises()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.find('#incorrect-credentials').exists()).toBeTruthy()
  })
})
