import { mount } from "@vue/test-utils"
import MembersOverlay from '@/components/MembersOverlay.vue';
import BandMemberListItem from '@/components/BandMemberListItem.vue';
import mockAxios from "@/__mocks__/services/axios";
import { nextTick } from "vue";
import { createStore } from "vuex";

const mockMembers = [
  { Musician: { MusicianID: 1, Name: 'hello' } },
  { Musician: { MusicianID: 2, Name: 'hello' } },
  { Musician: { MusicianID: 3, Name: 'hello' } }
]

const mockListener = jest.spyOn(mockAxios(), 'get')
mockListener.mockImplementation(() => {
  return Promise.resolve({data: mockMembers})
})

const createVuexStore = () => {
  return createStore({
    state: {
      user: {
        MusicianID: 1
      }
    },
  })
}

describe('MembersOverlay', () => {
  it('Should show a list of Members', async () => {
    const store = createVuexStore()
    const wrapper = mount(MembersOverlay, {
      global: {
        plugins: [store]
      }
    })

    wrapper.vm.getMembers = mockAxios().get
    await wrapper.vm.setupBandmembers()
    await nextTick()

    expect(mockListener).toHaveBeenCalled()
    expect(wrapper.findAllComponents(BandMemberListItem).length).toBe(3)
  })
})