import { mount } from '@vue/test-utils'
import Modal from '@/components/Modal.vue'

describe('Modal.vue', () => {
  it('Shouldn\'t render when passed an visibility value of False', () => {
    const title = "Hello World"
    const wrapper = mount(Modal, {
      props: { visible: false, title: title }
    })
    const titleWrapper = wrapper.find('h1')
    expect(titleWrapper.exists()).toBe(false)
  })

  it('Should render when passed an visibility value of True', () => {
    const title = "Hello World"
    const wrapper = mount(Modal, {
      props: { visible: true, title: title }
    })
    const titleWrapper = wrapper.find('h1')
    expect(titleWrapper.exists()).toBe(true)
    expect(titleWrapper.text()).toMatch(title)
  })
})
