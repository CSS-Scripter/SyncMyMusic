import { mount } from '@vue/test-utils'
import LeaveBand from '@/components/LeaveBand.vue'

describe('LeaveBand.vue', () => {
    it('Should render without the modal being visible', () => {
        const wrapper = mount(LeaveBand)
        expect(wrapper.findAll('h1').length).toBe(0)
    })
    it('Should open a modal upon clicking the icon', async () => {
        const wrapper = mount(LeaveBand)
        await wrapper.find('div.top-right').trigger('click')
        expect(wrapper.findAll('h1').length).toBe(1)
    })
    it('Should emit an \'leave\' event upon accepting the modal', async () => {
        const wrapper = mount(LeaveBand)
        await wrapper.find('div.top-right').trigger('click')
        await wrapper.find('button.green.medium').trigger('click')
        const leaveEvent = wrapper.emitted().leave
        expect(leaveEvent).toBeTruthy
        expect(leaveEvent.length).toBe(1)
    })
})