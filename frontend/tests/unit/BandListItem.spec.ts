import BandListItem from '@/components/BandListItem.vue'
import { mount } from '@vue/test-utils'
import { nextTick } from 'vue'
import router from '../../src/router/index'

jest.mock('../../src/router/index.ts')

describe('BandListItem.vue', () => {
    afterEach(() => {
        jest.clearAllMocks()
    })
    it('Should render the elements correctly', () => {
        const wrapper = mount(BandListItem, {props: {band: {name: "Testers"}}})
        expect(wrapper.find('h3').text()).toBe("Testers")
    })
    it('Shouldn\'t render a join session button when there is no session', () => {
        const wrapper = mount(BandListItem, {props: {band: {name: "Testers"}}})
        expect(wrapper.findAll('button').length).toBe(1)
    })
    it('Should render a join session button when there is a session', () => {
        const wrapper = mount(BandListItem, {props: {band: {name: "Testers", session: {}}}})
        expect(wrapper.findAll('button').length).toBe(2)
        expect(wrapper.findAll('button')[1].text()).toBe('Join Session')
    })
    it('Should push the band route upon pressing View Band', async () => {
        
        const mockBand = {
            ID: 1, 
            name: "Testers", 
            session: {}
        }
        const mockStore = {
            commit: jest.fn((name, value) => {})
        }

        const wrapperProperties = {
            props: { band: mockBand },
            mocks: { router },
            global: { mocks: { $store: mockStore } }
        }

        const wrapper = mount(BandListItem, wrapperProperties)
        await wrapper.find('button').trigger('click')
        await nextTick()
        expect(router.push).toHaveBeenCalledTimes(1)
        expect(router.push).toHaveBeenCalledWith({
            name: 'Band', params: {bandId: mockBand.ID}
        })
        expect(mockStore.commit).toHaveBeenCalledTimes(1)
        expect(mockStore.commit).toHaveBeenCalledWith('setBand', mockBand)
    })
    it('Should push the rehearsal route upon pressing View Band', async () => {
        const mockBand = {
            ID: 1, 
            name: "Testers", 
            session: {}
        }

        const wrapperProperties = {
            props: { band: mockBand },
            mocks: { router },
        }

        const wrapper = mount(BandListItem, wrapperProperties)
        await wrapper.find('button.red').trigger('click')
        await nextTick()
        expect(router.push).toHaveBeenCalledTimes(1)
        expect(router.push).toHaveBeenCalledWith({
            name: 'Rehearsal', params: {bandId: mockBand.ID}
        })
    })
})