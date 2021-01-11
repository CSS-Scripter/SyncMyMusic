import { mount } from '@vue/test-utils'
import Rehearsal from '@/views/Rehearsal.vue'
import SelectSong from '@/components/SelectSong.vue'
import SelectInstrument from '@/components/SelectInstrument.vue'
import { createStore } from 'vuex'

const createVuexStore = () => {
  return createStore({
    state: {
      user: {
        MusicianID: 1
      }
    },
  })
}

function createWrapper() {
  const store = createVuexStore();
  return mount(Rehearsal, {
    props: { bandId: 1 },
    global: {
      plugins: [store]
    }
  })
}

describe('Rehearsal.vue', () => {
  it('Should call created', async () => {
    const spyCreated = spyOn(Rehearsal, 'created')
    const wrapper = createWrapper();

    await wrapper.vm.$nextTick()
    expect(spyCreated).toHaveBeenCalled()
  })

  it('Should render end session button', async () => {
    const wrapper = createWrapper();
    wrapper.vm.ownerId = 1
    await wrapper.vm.$nextTick()

    expect(wrapper.find('button').text()).toBe("End session")
    expect(wrapper.vm.isSessionOwner()).toBeTruthy()
  })

  it("should call endSession", async () => {
    const wrapper = createWrapper();
    wrapper.vm.endSession = jest.fn()
    wrapper.vm.ownerId = 1
    await wrapper.vm.$nextTick()

    const endSessionbutton = wrapper.find('button')

    endSessionbutton.trigger('click')
    await wrapper.vm.$nextTick()

    expect(wrapper.vm.endSession).toHaveBeenCalled()
  })

  it('Should render SelectSong component', async () => {
    const wrapper = createWrapper();
    wrapper.vm.ownerId = 1
    await wrapper.vm.$nextTick()

    expect(wrapper.findComponent(SelectSong)).not.toBeNull()
  })
  
  it('Should render leave session button', async () => {
    const wrapper = createWrapper();
    wrapper.vm.ownerId = 55
    await wrapper.vm.$nextTick()

    expect(wrapper.find('button').text()).toBe("Leave session")
    expect(wrapper.vm.isSessionOwner()).toBeFalsy()
  })

  it("should call leaveSession", async () => {
    const wrapper = createWrapper();
    wrapper.vm.leaveSession = jest.fn()
    wrapper.vm.ownerId = 1
    await wrapper.vm.$nextTick()

    const leaveSessionButton = wrapper.find('button')

    leaveSessionButton.trigger('click')
    await wrapper.vm.$nextTick()

    expect(wrapper.vm.leaveSession).toHaveBeenCalled()
  })

  it("Should render 'Nog geen liedje geselecteerd'", async () => {
    const wrapper = createWrapper();
    expect(wrapper.text()).toContain("Nog geen liedje geselecteerd")
  })

  it("Should render 'Dit liedje is niet beschikbaar voor jou gekozen instrument'", async () => {
    const wrapper = createWrapper();
    wrapper.vm.songUnavailable = true
    wrapper.vm.pdfsource = false
    await wrapper.vm.$nextTick()

    expect(wrapper.text()).toContain("Dit liedje is niet beschikbaar voor jou gekozen instrument")
  })

  it("Should call selectSong", async () => {
    const wrapper = createWrapper();
    wrapper.vm.selectSong = jest.fn()
    wrapper.vm.ownerId = 1
    await wrapper.vm.$nextTick()
    const selectSong = wrapper.findComponent(SelectSong)

    selectSong.vm.$emit('song-changed')
    await wrapper.vm.$nextTick()

    expect(wrapper.vm.selectSong).toHaveBeenCalled()
  })

  xit("Should call selectInstrument", async () => {
    const wrapper = createWrapper();
    wrapper.vm.selectInstrument = jest.fn()
    await wrapper.vm.$nextTick()

    const selectInstrumentComponent = wrapper.findComponent(SelectInstrument)

    selectInstrumentComponent.vm.$emit('instrument-changed')
    await wrapper.vm.$nextTick()

    expect(wrapper.vm.selectInstrument).toHaveBeenCalled()
  })
})