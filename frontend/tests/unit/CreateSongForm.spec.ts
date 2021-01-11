import { mount } from '@vue/test-utils'
import CreateSongForm from '@/components/CreateSongForm.vue'
import { nextTick } from 'vue'

describe('CreateSongForm.Vue', () => {
    it('Should have empty fields upon receiving an empty song', async () => {
        const wrapper = mount(CreateSongForm, {props: {song: {}}})
        await nextTick()
        const inputFields = wrapper.findAll('input')
        for(let input of inputFields) {
            expect((<HTMLInputElement>wrapper.find('input#songname').element).value.trim().length).toBe(0)
        }
    })
    it('Should have prefilled fields upon receiving a song', async () => {
        const song = {ID: 1, name: "Clouds", artist: "Paper Idols"}
        const wrapper = mount(CreateSongForm, {props: {song}})
        await nextTick()
        expect((<HTMLInputElement>wrapper.find('input#songname').element).value).toBe(song.name)
        expect((<HTMLInputElement>wrapper.find('input#artist').element).value).toBe(song.artist)
    })
    it('Shouldn\'t load the delete button upon receiving an empty song', async () => {
        const wrapper = mount(CreateSongForm, {props: {song: {}}})
        await nextTick()
        const buttons = wrapper.findAll('button')
        expect(buttons.length).toBe(2)
        for (let button of buttons) {
            expect(button.text()).not.toBe('Delete')
        }
    })
    it('Should load the delete button upon receiving a filled song', async () => {
        const song = {ID: 1, name: "Clouds", artist: "Paper Idols"}
        const wrapper = mount(CreateSongForm, {props: {song}})
        await nextTick()
        const buttons = wrapper.findAll('button')
        expect(buttons.length).toBe(3)
        expect(buttons[0].text()).toBe('Delete')
    })
    it('Should emit an delete event upon pressing the delete button', async () => {
        const song = {ID: 1, name: "Clouds", artist: "Paper Idols"}
        const wrapper = mount(CreateSongForm, {props: {song}})
        await nextTick()
        const buttons = wrapper.findAll('button')
        buttons[0].trigger('click')
        await nextTick()
        const deleteEvent = wrapper.emitted().delete
        expect(deleteEvent).toBeTruthy()
        expect(deleteEvent.length).toBe(1)
        expect(deleteEvent[0]).toStrictEqual([song])
    })
    it('Should emit an accept event upon pressing the continue button', async () => {
        const song = {ID: 1, name: "Clouds", artist: "Paper Idols"}
        const wrapper = mount(CreateSongForm, {props: {song}})
        await nextTick()
        const buttons = wrapper.findAll('button')
        buttons[2].trigger('click')
        await nextTick()
        const acceptEvent = wrapper.emitted().accept
        expect(acceptEvent).toBeTruthy()
        expect(acceptEvent.length).toBe(1)
        expect(acceptEvent[0]).toStrictEqual([song])
    })
    it('Should emit an close event upon pressing the close button', async () => {
        const song = {ID: 1, name: "Clouds", artist: "Paper Idols"}
        const wrapper = mount(CreateSongForm, {props: {song}})
        await nextTick()
        const buttons = wrapper.findAll('button')
        buttons[1].trigger('click')
        await nextTick()
        const closeEvent = wrapper.emitted().close
        expect(closeEvent).toBeTruthy()
        expect(closeEvent.length).toBe(1)
    })
})