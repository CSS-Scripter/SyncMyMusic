import { mount } from '@vue/test-utils'
import Account from '@/views/Account.vue'
import { createStore } from 'vuex'
import { nextTick } from 'vue'

const createVuexStore = () => {
    return createStore({
      state: {
        user: {
          ID: 1,
          username: "Test",
          Email: "test@email.nl"
        }
      }
    })
  }

describe('Account.vue', () => {
    // Test if all elements are on the page.
    it('Has the h1 page title', () => {
        const store = createVuexStore();
        const wrapper = mount(Account, {
          global: {
            plugins: [store]
          }
        })
      expect(wrapper.findAll('h1').length).toBe(1);
    });
  
    it('Has 3 pair of labels + corresponding inputboxes', () => {
        const store = createVuexStore();
        const wrapper = mount(Account, {
          global: {
            plugins: [store]
          }
        })
      expect(wrapper.findAll('input').length).toBe(3);
      expect(wrapper.findAll('label').length).toBe(3);
    });
  
    it('Has 3 buttons with appropriate text', () => {
        const store = createVuexStore();
        const wrapper = mount(Account, {
          global: {
            plugins: [store]
          }
        })
      let buttons = wrapper.findAll('button');
      expect(buttons.length).toBe(3);
      expect(buttons[0].text()).toBe('Back');
      expect(buttons[1].text()).toBe('Update');
      expect(buttons[2].text()).toBe('Log Out');
    });

    // Other Tests
    it('Has username pre-entered in first inputbox', async () => {
        const store = createVuexStore();
        const wrapper = mount(Account, {
          global: {
            plugins: [store]
          }
        })
        console.log(wrapper.vm.$store.state.user.username)
        wrapper.vm.setName()
        await nextTick()
        const input = wrapper.find('#name')
        expect((<HTMLInputElement>input.element).value).toBe('Test')
    })
})