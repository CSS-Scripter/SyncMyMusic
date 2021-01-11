import { flushPromises, mount } from '@vue/test-utils'
import SelectInstrument from '@/components/SelectInstrument.vue'
import mockAxios from '../../src/__mocks__/services/axios'
import { nextTick } from 'vue'

const mockInstruments = [
  {ID: 1, Name: "foo"},
  {ID: 2, Name: "bar"},
  {ID: 3, Name: "foobar"},
]
const mockListener = jest.spyOn(mockAxios(), 'get')
  mockListener.mockImplementationOnce(() => { 
    return Promise.resolve({data: mockInstruments})
  }
)

describe('SelectInstrument.vue', () => {

  it('Should call axios.get, fill options with mocksongs', async () => {
    const wrapper = mount(SelectInstrument, {
      props: {
        selectedInstrumentId: 2
      }
    })
    await wrapper.vm.$nextTick()
    await flushPromises()

    wrapper.vm.getInstruments = mockAxios().get
    await wrapper.vm.setupInstruments()
    await nextTick()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.vm.options).toEqual(mockInstruments)
    expect(wrapper.vm.selectedInstrument.ID).toEqual(2)
  })

})