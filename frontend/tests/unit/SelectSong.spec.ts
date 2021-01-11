import { mount } from '@vue/test-utils'
import SelectSong from '@/components/SelectSong.vue'
import mockAxios from '../../src/__mocks__/services/axios'
import { nextTick } from 'vue'

const mockSongs = [
  {ID: 1, name: "foo"},
  {ID: 2, name: "bar"},
  {ID: 3, name: "foobar"},
]
const mockListener = jest.spyOn(mockAxios(), 'get')
  mockListener.mockImplementation(() => { 
    return Promise.resolve({data: mockSongs})
  }
)

describe('SelectSong.vue', () => {

  it('Should call axios.get, fill options with mocksongs', async () => {
    const wrapper = mount(SelectSong)

    wrapper.vm.getSongs = mockAxios().get
    await wrapper.vm.setupSongs()
    await nextTick()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.vm.options).toEqual(mockSongs)
    expect(wrapper.vm.selectedSong).toBeNull()
  })

  it('Should call axios.get, Select song with id 2', async () => {
    const wrapper = mount(SelectSong, {
      props: {
        selectedSongId: 2
      }
    })

    wrapper.vm.getSongs = mockAxios().get
    await wrapper.vm.setupSongs()
    await nextTick()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.vm.selectedSong.ID).toEqual(2)
  })

})