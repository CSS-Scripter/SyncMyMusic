import { mount } from '@vue/test-utils'
import mockAxios from '../../src/__mocks__/services/axios'
import flushPromises from 'flush-promises'
import SongListView from '@/components/SongsListView.vue'
import { nextTick } from 'vue'

const mockSongs = [
  {ID: 1, name: "foo"},
  {ID: 2, name: "bar"},
  {ID: 3, name: "foobar"},
]
describe('SongListView.vue', () => {
  it('Should render an listElement for every song', async () => {
    // Create a listener for the mock instance of Axios
    const mockListener = jest.spyOn(mockAxios(), 'get')
    mockListener.mockImplementationOnce(() => { 
        return Promise.resolve({data: mockSongs})
      }
    )

    // Create the component and flush all awaiting promises
    const wrapper = mount(SongListView)
    await flushPromises()

    // Overwrite an http function with the mock function
    // Call the mock function
    // await DOM changes
    wrapper.vm.getSongPromise = mockAxios().get
    await wrapper.vm.setupSongs()
    await nextTick()

    // Assert
    // Songs have been requested
    expect(mockListener).toHaveBeenCalledTimes(1)

    // Listelements have been rendered
    const renderedSongs = wrapper.findAll('div.listElement')
    expect(renderedSongs.length).toBe(3)
    for (let i = 0; i < mockSongs.length; i++) {
      expect(renderedSongs[i].find('div.name').text()).toMatch(mockSongs[i].name)
    }
  })
})