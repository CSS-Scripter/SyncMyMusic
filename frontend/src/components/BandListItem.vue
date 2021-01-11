<template>
  <div class="card-content">
      <div class="left-content">
          <h3>{{band.name}}</h3>
          <Button class="green small" @click="toBandPage">View Band</Button>
      </div>
      <div v-if=band.session class="right-content">
          <h3>Currently Playing</h3>
          <!-- <div class="song">{{band.session.currentSong}} - {{band.session.artist}}</div> -->
          <Button class="red small" @click="joinSession">Join Session</Button>
      </div>
  </div>
</template>

<script>
import router from '../router'
import Button from './Button.vue'

export default {
  components: {Button},
  props: [
    'band'
  ],
  methods: {
    toBandPage() {
      this.$store.commit('setBand', this.band)
      router.push({ name: 'Band', params: { bandId: this.band.ID }})
    },
    joinSession() {
      router.push({ name: 'Rehearsal', params: { bandId: this.band.ID } })
    }
  }
}
</script>

<style lang="scss" scoped>
.card-content {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding: 1.5em 4em;

    .left-content {
        display: flex;
        flex-direction: column;
        justify-content: space-between;
    }
    
    h3 {
        margin: 0;
        font-size: 20px;
    }

    Button {
        margin-top: 1em;
        width: 180px;
    }

    .song {
        font-weight: lighter;
        font-size: 16px;
        font-style: italic;
    }
}

</style>