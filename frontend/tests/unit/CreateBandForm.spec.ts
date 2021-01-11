import { flushPromises, mount } from '@vue/test-utils'
import CreateBandForm from '@/components/CreateBandForm.vue'
import { nextTick } from 'vue'

describe('CreateBandForm.vue', () => {
    it('Should return invalid for an invalid band name', async () => {
        const wrapper = mount(CreateBandForm)
        wrapper.find('input').setValue('He')
        const {valid} = wrapper.vm.submit()
        expect(valid).toBe(false)
    })
    it('Should return valid for an barely valid band name', async () => {
        const wrapper = mount(CreateBandForm)
        wrapper.find('input').setValue('Hell')
        const {valid, band} = wrapper.vm.submit()
        expect(valid).toBe(true)
        expect(band.name).toBe('Hell')
    })
    it('Should return valid for an valid band name', async () => {
        const wrapper = mount(CreateBandForm)
        wrapper.find('input').setValue('Hello World')
        const {valid, band} = wrapper.vm.submit()
        expect(valid).toBe(true)
        expect(band.name).toBe('Hello World')
    })
})