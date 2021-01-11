import { mount, VueWrapper } from '@vue/test-utils'
import Band from '@/views/Band.vue'
import { createStore } from 'vuex'

const createVuexStore = () => {
    return createStore({
      state: {
        band: {
          ID: 88,
          Name: 'MyTestBand',
          OwnerID: 22
        },
        user: {
          MusicianID: 1,
          username: "Test",
          Email: "test@email.nl"
        }
      }
    })
  }

let wrapper: VueWrapper<any>;

beforeEach(() => {
    const store = createVuexStore();
    wrapper = mount(Band, {
      global: {
        plugins: [store]
      }
    })
});

describe('Band.vue', () => {
    // Test if all elements are on the page.
    it('Has the h1 band title', () => {
      expect(wrapper.findAll('h1').length).toBe(1);
    });

    it('Shows the joincode on page', () => {
        const joincode = wrapper.find('p');
        expect(joincode.exists()).toBe(true);
    });
  
    it('Has 4 buttons with appropriate text', () => {
      let buttons = wrapper.findAll('button');
      expect(buttons.length).toBe(4);
      expect(buttons[0].text()).toBe('Repertoire');
      expect(buttons[1].text()).toBe('Members');
      expect(buttons[2].text()).toBe('Instruments');
      expect(buttons[3].text()).toBe('Start Jamming');
    });

    it('Overlay buttons should trigger function', async () => {
        let ButtonClickMock = jest.fn();
        wrapper.vm.showOverlay = ButtonClickMock;
        let buttons = wrapper.findAll('button');
        await buttons[0].trigger('click');
        await buttons[1].trigger('click');
        await buttons[2].trigger('click');

        expect(ButtonClickMock).toHaveBeenCalledTimes(3);
    });
})