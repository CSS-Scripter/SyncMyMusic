import { mount } from '@vue/test-utils'
import SignUp from '@/components/SignUp.vue'

describe('SignUp.vue', () => {

    let wrapper = mount(SignUp)

    // Test if all elements are on the page.
    it('Has 4 pair of labels with predefined text + corresponding inputboxes', () => {
      let labels = wrapper.findAll('label');
      expect(labels.length).toBe(4);
      expect(labels[0].text()).toBe('Username');
      expect(labels[1].text()).toBe('Email');
      expect(labels[2].text()).toBe('Password');
      expect(labels[3].text()).toBe('Repeat Password');
      expect(wrapper.findAll('input').length).toBe(4);
    });
  
    it('Has 2 buttons with appropriate text', () => {
      let buttons = wrapper.findAll('button');
      expect(buttons.length).toBe(2);
      expect(buttons[0].text()).toBe('Cancel');
      expect(buttons[1].text()).toBe('Sign Up');
    });

    // Other Tests
    it('Fires click event for Sign Up', async () => {
        let ButtonClickMock = jest.fn();
        wrapper.vm.submitForm = ButtonClickMock;
        await wrapper.find('#login-form').trigger('submit')

        expect(ButtonClickMock).toHaveBeenCalled();
    })
})