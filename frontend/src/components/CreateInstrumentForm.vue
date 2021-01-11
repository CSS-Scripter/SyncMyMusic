<template>
    <div class="form">
        <form class="modal">
            <label for="instrumentname">Name</label>
            <input type="text" id="instrumentname" class="invalid" :value=instrument.Name @keyup="updateName($event)">
            <p class="error" v-if="errorMessage != ''">{{ errorMessage }}</p>
        </form>
        <Button v-if="instrument.Name !== undefined" class="red medium" @Click="$emit('delete')">Delete instrument</Button>
        <div class="modal__actions">
            <Button class="red medium" @Click="$emit('close')">Cancel</Button>
            <Button class="green medium" @Click="$emit('accept', submit())">Continue</Button>
        </div>
    </div>
</template>

<script>
import Button from './Button'
export default {
    name: "CreateInstrumentForm",
    components: {Button},
    props: ['instrument'],
    data() {
        return {
            name: "",
            errorMessage: "",
            deleteButtonContent: null,
        }
    },
    methods: {
        updateName(event) {
            this.name = event.target.value
        },
        submit() {
            if (this.name.trim().length < 3) {
                return {
                    valid: false
                }
            }
            return {
                valid: true,
                instrument: {
                    name: this.name
                }
            }
        }
    }
}
</script>
