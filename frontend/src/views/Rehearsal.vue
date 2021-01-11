<template>
  <div id="rehearsal" style="display:grid;grid-template-columns:1fr 3fr; height:100%">
    <div>
      <h1 id="welcome">Welkom bij de repetitie van band: {{ bandId }}</h1>
      <SelectInstrument
        @instrument-changed="selectInstrument"
        :bandId="bandId"
        :selectedInstrumentId="instrumentId"
      />
      <SelectSong
        v-if="isSessionOwner()"
        @song-changed="selectSong"
        :bandId="bandId"
        :selectedSongId="selectedSong ? selectedSong.ID : null"
        placeholder="Select song"
      />
      <Button class="red medium" @click="endSession" v-if="isSessionOwner()">End session</Button>
      <Button class="red medium" @click="leaveSession" v-else>Leave session</Button>
    </div>
    <div>
      <h1 id="not-available" v-if="songUnavailable">
        Dit liedje is niet beschikbaar voor jou gekozen instrument!
      </h1>
      <h1 id="no-song" v-if="!pdfsource && !songUnavailable">
        Nog geen liedje geselecteerd!
      </h1>
      <embed
        id="pdf"
        v-if="pdfsource"
        :src="pdfsource"
        :key="pdfsource"
        width="100%"
        height="100%"
        toolbar="0"
        type="application/pdf"
      />
    </div>
  </div>
</template>

<script>
import axios from "../services/axios";
import router from "../router"
import SelectInstrument from "../components/SelectInstrument.vue";
import SelectSong from "../components/SelectSong.vue";
import Button from"../components/Button.vue";

export default {
  components: { SelectInstrument, SelectSong, Button },
  props: ["bandId"],
  data() {
    return {
      message: "Rehearsal not started!",
      websocket: null,
      selectedSong: null,
      pdfsource: null,
      songUnavailable: false,
      instrumentId: 1,
      sessionId: null,
      ownerId: null,
    };
  },
  methods: {
    isSessionOwner() {
      const userId = this.$store.state.user.MusicianID;
      return userId == this.ownerId
    },
    startRehearsal() {
      this.websocket = new WebSocket(
        `ws://${process.env.VUE_APP_BACKENDURL}/ws/${this.bandId}?userID=${this.$store.state.user.MusicianID}`
      );
      this.websocket.addEventListener("message", this.receiveMessage);
    },
    leaveSession() {
      this.websocket.close();
      router.go(-1);
    },
    endSession() {
      this.leaveSession();
    },
    receiveMessage(event) {
      const dataObject = JSON.parse(event.data);
      if (!this.sessionId && !this.ownerId) {
        this.sessionId = dataObject.sessionId;
        this.ownerId = dataObject.ownerId;
      }
      if (dataObject.type === 2) this.selectedSong = { ID: dataObject.body };
      if (this.selectedSong) {
        this.getSong();
      }
      if (dataObject.type == 1 && this.selectedSong)
        this.selectSong(this.selectedSong);
      if (dataObject.type == -1)
        this.leaveSession();
    },
    selectSong(song) {
      this.websocket.send(song.ID);
    },
    selectInstrument(instrument) {
      this.instrumentId = instrument.ID;
      this.getSong();
    },
    getSong() {
      axios()
        .get(
          `/files/${this.bandId}/${this.selectedSong.ID}/${this.instrumentId}.pdf`
        )
        .then((result) => {
          const pdfSoureString = result.request.responseURL + "#toolbar=0";
          this.songUnavailable = false;
          this.pdfsource = pdfSoureString;
        })
        .catch(() => {
          this.pdfsource = null;
          this.songUnavailable = true;
        });
    },
  },
  created() {
    this.startRehearsal();
  },
};
</script>

<style></style>
