import { mount } from "@vue/test-utils"
import BandMemberListItem from '@/components/BandMemberListItem.vue';
import { createStore } from "vuex";

const mockMember = {
  Musician: {
    MusicianID: 1,
    Name: 'hello'
  }
}

const createVuexStore = () => {
  return createStore({
    state: {
      user: {
        MusicianID: 1
      }
    },
  })
}

describe('BandMemberListItem.vue', () => {
  it('Should render edit button', async () => {
    const store = createVuexStore()
    const wrapper = mount(BandMemberListItem, {
      props: {
        member: mockMember,
        isOwner: false
      },
      global: {
        plugins: [store]
      }
    })

    expect(wrapper.find('#edit-button')).toBeTruthy()
  })
})