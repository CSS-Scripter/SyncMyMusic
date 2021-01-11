import { mount } from '@vue/test-utils'
import mockAxios from '../../src/__mocks__/services/axios'
import flushPromises from 'flush-promises'
import InstrumentsOverlay from '@/components/InstrumentsOverlay.vue'
import { nextTick } from 'vue'

const mockInstruments = [
  {ID: 1, Name: "foo"},
  {ID: 2, Name: "bar"},
  {ID: 3, Name: "foobar"},
]
describe('InstrumentOverlay.vue', () => {
  it('Should render a listElement for each instrument', async () => {
    // Create a listener for the mock instance of Axios
    const mockListener = jest.spyOn(mockAxios(), 'get')
    mockListener.mockImplementationOnce(() => { 
        return Promise.resolve({data: mockInstruments})
      }
    )

    // Create the component and flush all awaiting promises
    const wrapper = mount(InstrumentsOverlay)
    await flushPromises()

    // Overwrite an http function with the mock function
    // Call the mock function
    // await DOM changes
    wrapper.vm.getInstrumentPromise = mockAxios().get
    await wrapper.vm.getInstruments()
    await nextTick()

    // Assert
    // Instruments have been requested
    expect(mockListener).toHaveBeenCalledTimes(1)

    // Listelements have been rendered
    const renderedInstruments = wrapper.findAll('div.listElement')
    expect(renderedInstruments.length).toBe(3)
    for (let i = 0; i < mockInstruments.length; i++) {
      expect(renderedInstruments[i].find('div.name').text()).toMatch(mockInstruments[i].Name)
    }
  })
})