import { mount } from '@vue/test-utils'
import Button from '@/components/Button.vue'

describe('Button.vue', () => {
  it('Should render \'button\' as button text', () => {
    const buttonText = "button"
    const wrapper = mount(Button, {
      slots: { default: buttonText }
    });
    const button = wrapper.find('button')
    expect(button.text().trim()).toBe(buttonText)
  })

})
