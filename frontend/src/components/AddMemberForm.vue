<template>
    <div class="form">
        <form class="modal">
            <label for="memberemail">Email</label>
            <input type="email" id="memberemail" class="invalid" v-model="email" @keyup="validate">
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
    name: "AddMemberForm",
    components: {Button},
    data() {
        return {
            email: "",
            errorMessage: ""
        }
    },
    methods: {
        validate(event) {
            if (event.target.value.length < 3) {
                event.target.className = "invalid"
                this.errorMessage = "Membername should be atleast 3 characters long"
            } else {
                event.srcElement.className = "valid"
                this.errorMessage = ""
            }
        },
        submit() {
            if (this.email.trim().length < 3) {
                return {
                    valid: false
                }
            }
            return {
                valid: true,
                member: {
                    email: this.email
                }
            }
        }
    }
}
</script>
