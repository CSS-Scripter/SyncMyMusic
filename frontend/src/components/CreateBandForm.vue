<template>
    <div class="form">
        <form class="modal">
            <label for="bandname">Name</label>
            <input type="text" id="bandname" class="invalid" v-model="name" @keyup="validate" @change="validate">
            <p class="error" v-if="errorMessage != ''">{{ errorMessage }}</p>
        
        </form>
        <div class="modal__actions">
                <Button class="red medium" @Click="$emit('close')">Cancel</Button>
                <Button class="green medium" @Click="$emit('accept', submit())">Continue</Button>
            </div>
        </div>
</template>

<script>
import Button from './Button'
export default {
    name: "CreateBandForm",
    components: {Button},
    data() {
        return {
            name: "",
            errorMessage: ""
        }
    },
    methods: {
        validate(event) {
            if (event.target.value.length < 4) {
                event.target.className = "invalid"
                this.errorMessage = "Bandname should be atleast 4 characters long"
            } else {
                event.srcElement.className = "valid"
                this.errorMessage = ""
            }
        },
        submit() {
            if (this.name.trim().length < 4) {
                return {
                    valid: false
                }
            }
            return {
                valid: true,
                band: {
                    name: this.name
                }
            }
        }
    }
}
</script>
