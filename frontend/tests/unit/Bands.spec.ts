import Bands from '@/views/Bands.vue'
import {flushPromises, mount} from '@vue/test-utils'
import { nextTick } from 'vue'
import mockAxios from '../../src/__mocks__/services/axios'

const mockBands = [
    {
        ID: 1,
        name: 'hello',
    },
    {
        ID: 2,
        name: 'World'
    },
    {
        ID: 3,
        name: 'hello World',
    }
]

const mockRehearsals = [1,3]

const mockStore = {
    state: {
        user: {
            MusicianID: 2
        }
    }
}

describe('Bands.Vue', () => {
    beforeEach(() => {
        jest.clearAllMocks()
        jest.spyOn(window, 'alert').mockReturnValue()
    })
    it('Should render the bands correctly', async () => {
        const mockListener = jest.spyOn(mockAxios(), 'get')
        mockListener.mockImplementation((url, ...extra) => {
            if (url.indexOf('bands') > -1) {
                return Promise.resolve({data: mockBands})
            } else {
                return Promise.resolve({data: mockRehearsals})
            }
          }
        )

        const mockProperties = {
            global: {
                mocks: { $store: mockStore }
            }
        }

        const wrapper = mount(Bands, mockProperties)
        wrapper.vm.getAxios = () => mockAxios()
        await flushPromises()
        wrapper.vm.getBands()
        await nextTick()
        await nextTick()

        const bandCards = wrapper.findAll('.card')
        expect(bandCards.length).toBe(mockBands.length)
        expect(mockListener).toBeCalledTimes(2)
    })

    it('Should send a correct join request', async () => {
        const mockListener = jest.spyOn(mockAxios(), 'post')
        mockListener.mockImplementation(() => {
                return Promise.resolve({data: "ok"})
          }
        )

        const mockProperties = {
            global: {
                mocks: { $store: mockStore }
            }
        }

        const wrapper = mount(Bands, mockProperties)
        wrapper.vm.getAxios = () => mockAxios()
        await flushPromises()
        await wrapper.find('.actions button.green').trigger('click')

        const modalRoot = wrapper.find('.background')
        expect(modalRoot.find('h1').text()).toBe("Join a band")

        const joinBandInput = wrapper.find('#joinBandInput')
        await joinBandInput.setValue('123')

        wrapper.vm.sendJoinRequest()
        await nextTick()

        expect(mockListener).toBeCalledTimes(1)
        expect(mockListener).toBeCalledWith('/bands/123/request')
    })
    xit('Should send a valid create band request', async () => {
        const mockBandID = 12
        
        const mockListener = jest.spyOn(mockAxios(), 'post')
        mockListener.mockImplementation(() => {
                return Promise.resolve({data: mockBandID})
          }
        )
       
        const mockRouter = {
            push: jest.fn()
        }

        const mockProperties = {
            global: {
                mocks: { $store: mockStore, $router: mockRouter }
            }
        }
        
        const wrapper = mount(Bands, mockProperties)
        await flushPromises()
        wrapper.vm.getAxios = mockAxios()
        await wrapper.vm.createBand({valid: true, band: {name: "Hel"}})
        expect(mockListener).toBeCalledTimes(1)
        expect(mockListener).toBeCalledWith(`/users/${mockStore.state.user.MusicianID}/bands`, {name: "Hel"})

        expect(mockRouter.push).toBeCalledTimes(1)
        expect(mockRouter.push).toBeCalledWith({name: "Band", params: {bandId: mockBandID}})
    })
})