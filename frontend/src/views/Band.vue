<template>
  <div class="band">
    <h1 v-if="band">{{ band.name }}</h1>
    <p>Join Code: {{ band.ID }}</p>
    <div class="band-content">
      <div class="left-container">
        <Button class="green small button-margin" @click="showOverlay('Songs')">
          Repertoire</Button>
        <Button
          class="green small button-margin"
          @click="showOverlay('Members')">
            Members</Button>
        <Button
          class="green small button-margin"
          @click="showOverlay('Instruments')">
            Instruments
        </Button>

        <span class="spacer"></span>

        <div v-if="band.session">
          <Button class="red small button-margin">Join Jamming</Button>
        </div>
        <div v-else>
          <h2>There is currently no session active</h2>
          <Button id="start-jamming" class="green small button-margin" @click="startRehearsal"
            >Start Jamming</Button
          >
        </div>
      </div>
      <div class="list-view">
        <router-view></router-view>
      </div>
    </div>
    <LeaveBand v-if="!isOwner()" @leave="leaveBand"></LeaveBand>
    <DeleteBand @delete="deleteBand" v-else></DeleteBand>
  </div>
</template>

<script>
import Button from "../components/Button";
import router from "../router";
import axios from "../services/axios";
import DeleteBand from "../components/DeleteBand.vue";
import LeaveBand from "../components/LeaveBand.vue";

export default {
  components: {
    Button,
    DeleteBand,
    LeaveBand
  },
  props: ["bandId"],
  data() {
    return {
      band: {},
    };
  },
  computed: {
    user() {
      return this.$store.state.user
    }
  },
  methods: {
    showOverlay(overlay) {
      router.push({ name: overlay, params: { bandId: this.bandId, isOwner: this.isOwner() } });
    },
    isOwner() {
      return this.band.OwnerID === this.user.MusicianID
    },
    startRehearsal() {
      router.push({ name: "Rehearsal", params: { bandId: this.bandId } });
    },
    setupBand() {
      axios()
        .get(`/bands/${this.bandId}`)
        .then((response) => {
          this.band = response.data;
        });
    },
    deleteBand() {
      axios()
        .delete(`/bands/${this.bandId}`)
        .then((response) => {
          if (response.status == 200) {
            this.$router.push({ name: "Bands" });
          }
        });
    },
    leaveBand() {
      axios()
        .delete(`/bands/${this.bandId}/members/${this.user.MusicianID}`)
        .then((response) => {
            if (response.status == 200) {
              this.$router.push({ name: "Bands" });
            }
        });
    }
  },
  created: function() {
    this.setupBand();
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";

.band {
  color: $light_color;
  font-family: "Rubik", sans-serif;
}

h1 {
  font-family: "Saira", sans-serif;
  font-weight: bold;
  font-style: normal;
  font-size: 36px;
  color: $light_color;
  margin: 0;
  padding: 0;
}

h2 {
  font-family: "Rubik", sans-serif;
  color: $light_color;
  width: 100%;
}

.button-margin {
  margin-bottom: 1.5em;
}

.band-content {
  display: flex;
  flex-direction: row;
  height: 100%;
}

.left-container {
  display: flex;
  flex-direction: column;
  width: 30%;
  height: 100%;

  .spacer {
    flex: 1;
  }
}

.list-view {
  flex: 1;
  padding-left: 1em;
}

.upload-error {
  text-align: center;
  color: $red_color;
}

.input-container {
  margin: 0 auto;
  display: flex;
  flex-direction: row;
  justify-content: center;
}

#file {
  border: none;
  padding-right: 2em;
  font-size: 20px;
  outline: none;
  cursor: pointer;

  &::-webkit-file-upload-button {
    background: none;
    border-radius: 16px;
    border: 5px solid $light_color;
    padding: 1em;
    color: $light_color;
    font-weight: bold;
    margin-right: 1em;
  }

  &.valid {
    color: $green_color;

    &::-webkit-file-upload-button {
      border-color: $green_color;
    }
  }
}
.right-container {
  width: 50%;
  margin-left: 5%;
}
</style>
