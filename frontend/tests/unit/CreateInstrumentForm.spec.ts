import { mount } from '@vue/test-utils'
import CreateInstrumentForm from "@/components/CreateInstrumentForm.vue"

describe("CreateInstrumentForm.vue", () => {

    let wrapper = mount(CreateInstrumentForm, {
        props: { instrument: { } }
    })

    // Test if all elements are on the page.
    it('Has 1 label + corresponding inputbox', () => {
        expect(wrapper.findAll('label').length).toBe(1);
        expect(wrapper.findAll('input').length).toBe(1);
    });

    it('Has 2 buttons with appropriate text', () => {
        let buttons = wrapper.findAll('button');
        expect(buttons.length).toBe(2);
        expect(buttons[0].text()).toBe('Cancel');
        expect(buttons[1].text()).toBe('Continue');
    });

    // Test conditions without instrument name.
    it('Textfield is empty when instrumentname is undefined', () => {
        const value = wrapper.find('input').element.value
        expect(value).toBe("")
    })
    
    it('Shows NO Delete button when instrumentname is undefined', () => {
        let buttons = wrapper.findAll('button');
        expect(buttons.length).toBe(2);
    })

    // Test conditions with a given instrument name.
    const instrument = 'TestInstrument';
    let wrapperWithInstrument = mount(CreateInstrumentForm, {
        props: { instrument: { Name: instrument, BandID: 1 } }
    })

    it('Textfield contains instrumentname when it is given', () => {
        const value = wrapperWithInstrument.find('input').element.value
        expect(value).toBe(instrument)
    })

    it('Shows Delete button when instrumentname is defined', () => {
        let buttons = wrapperWithInstrument.findAll('button');
        expect(buttons.length).toBe(3);
    })
})