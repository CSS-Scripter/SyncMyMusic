import { flushPromises, mount } from '@vue/test-utils'
import SongInstrumentsListView from '@/components/SongInstrumentsListView.vue'
import { nextTick } from 'vue'
import mockAxios from '@/__mocks__/services/axios'

const mockInstruments = [
    {ID: 1, Name: "T"},
    {ID: 2, Name: "Tr"},
    {ID: 3, Name: "Tru"},
    {ID: 4, Name: "Trum"}
]
const mockInstrumentsWithPDF = [
    {InstrumentID: 1, Name: "T", PDFLocation: "t.pdf"},
    {InstrumentID: 2, Name: "Tr", PDFLocation: "t.pdf"},
    {InstrumentID: 4, Name: "Trum", PDFLocation: "t.pdf"}
]

describe('SongInstrumentsListView.vue', () => {
    it('Should render instruments', async () => {
        const wrapper = getWrapper()
        const mockAxiosListener = await mockAxiosOnWrapper(wrapper, 'get', [])
        wrapper.vm.setupInstruments()
        await nextTick()
        await nextTick()
        
        const instrumentElements = wrapper.findAll('.element')
        for (let i = 0; i < mockInstruments.length; i++) {
            const instrumentElement = instrumentElements[i]
            expect(instrumentElement.find('.name').text()).toBe(mockInstruments[i].Name)
        }
    })

    it('Should order to list to start with instruments without a PDF', async () => {
        const wrapper = getWrapper()
        await mockAxiosOnWrapper(wrapper, 'get', mockInstrumentsWithPDF)
        wrapper.vm.setupInstruments()
        await nextTick()
        await nextTick()

        const instrumentElements = wrapper.findAll('.element')
        
        expect(instrumentElements[0].find('button').text()).toBe('Add')
        expect(instrumentElements[1].find('button').text()).toBe('Edit')
        expect(instrumentElements[2].find('button').text()).toBe('Edit')
        expect(instrumentElements[3].find('button').text()).toBe('Edit')
    })
})

function getWrapper() {
    return mount(SongInstrumentsListView, {props:{instruments: mockInstruments}})
}

async function mockAxiosOnWrapper(wrapper: any, method: 'get'| 'post', mockData: any) {
    const mockListener = getMockAxios(method, mockData)
    wrapper.vm.getAxios = () => mockAxios()
    await flushPromises()
    return mockListener
}

function getMockAxios(method: 'get' | 'post', mockData: any) {
    const mockListener = jest.spyOn(mockAxios(), method)
    mockListener.mockImplementationOnce(() => {
        return Promise.resolve({data: mockData})
    })
    return mockListener
}