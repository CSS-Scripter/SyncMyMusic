<template>
    <div class="form">
        <form class="modal">
            <label for="songname">Name</label>
            <input ref="name" type="text" id="songname" class="valid" :value=song.name>
            
            <label for="artist">Artist</label>
            <input ref="artist" type="text" id="artist" class="valid" :value=song.artist>    
        </form>
        <slot />
        <Button v-if="song.ID != undefined" class="red medium" @Click="$emit('delete', submit())">Delete</Button>
        <div class="modal__actions">
            <Button class="red medium" @Click="$emit('close')">Cancel</Button>
            <Button class="green medium" @Click="$emit('accept', submit())">Continue</Button>
        </div>
    </div>
</template>

<script>
import Button from './Button'
export default {
    name: "CreateSongForm",
    components: {Button},
    props: ['song'],
    methods: {
        submit() {
            return {
                ...this.song,
                name: this.$refs.name.value,
                artist: this.$refs.artist.value
            }
        }
    }
}
</script>
