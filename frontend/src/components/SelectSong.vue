<template>
  <vue-select
    inputId="select-song"
    :options="options"
    label-by="name"
    :track-by="(option) => option.name"
    v-model="selectedSong"
    @selected="handleSelect"
  ></vue-select>
</template>

<script>
import axios from "../services/axios";
import VueSelect from "vue-next-select";
import "vue-next-select/dist/index.css";

export default {
  components: { VueSelect },
  props: ["bandId", "selectedSongId"],
  data() {
    return {
      options: [],
      selectedSong: null,
    };
  },
  methods: {
    handleSelect(e) {
      this.$emit("song-changed", e);
    },
    setSelected() {
      this.selectedSong = this.options.find(
        (song) => song.ID === this.selectedSongId
      );
    },
    getAxios() {
      return axios()
    },
    setupSongs() {
      this.getSongs().then((response) => {
        this.options = response.data;
        this.setSelected();
      })
    },
    getSongs() {
      return this.getAxios()
        .get(`/bands/${this.bandId}/songs`)
    }
  },
  created() {
    this.setupSongs()
  },
};
</script>

<style></style>
