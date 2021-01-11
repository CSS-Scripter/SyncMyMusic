import { mount } from '@vue/test-utils'
import AddMemberForm from "@/components/AddMemberForm.vue"

describe("CreateInstrumentForm.vue", () => {

    let wrapper = mount(AddMemberForm)

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

    // Tests covering Errormessage visiblity.
    it('Shouldn\'t render errormessage initially with empty email', async () => {
        expect(wrapper.findAll('p').length).toBe(0);
    });

    it('Should render errormessage if email is empty after change', async () => {
        const inputWrapper = wrapper.find('#memberemail')
        await inputWrapper.setValue('t');
        await inputWrapper.trigger('keyup', {
            key: 't'
        })
        await inputWrapper.setValue('');
        expect((<HTMLInputElement>inputWrapper.element).value).toBe('')
        expect(wrapper.findAll('p').length).toBe(1);
    });

    it('Should render errormessage if email is too short', async () => {
        const inputWrapper = wrapper.find('#memberemail')
        await inputWrapper.setValue('t');
        await inputWrapper.trigger('keyup', {
            key: 't'
        })
        expect((<HTMLInputElement>inputWrapper.element).value).toBe('t')
        expect(wrapper.findAll('p').length).toBe(1);
    });

    it('Shouldn\'t render errormessage if email is long enough', async () => {
        const inputWrapper = wrapper.find('#memberemail')
        await inputWrapper.setValue('lol');
        await inputWrapper.trigger('keyup', {
            key: 'l'
        })
        await inputWrapper.trigger('keyup', {
            key: 'o'
        })
        await inputWrapper.trigger('keyup', {
            key: 'l'
        })
        expect((<HTMLInputElement>inputWrapper.element).value).toBe('lol')
        expect(wrapper.findAll('p').length).toBe(0);
    });
})