<template>
  <vue-select
    :options="options"
    label-by="Name"
    :track-by="(option) => option.Name"
    v-model="selectedInstrument"
    @selected="handleSelect"
  ></vue-select>
</template>

<script>
import axios from "../services/axios";
import VueSelect from "vue-next-select";
import "vue-next-select/dist/index.css";

export default {
  components: { VueSelect },
  props: ["bandId", "selectedInstrumentId"],
  data() {
    return {
      options: [],
      selectedInstrument: null,
    };
  },
  methods: {
    handleSelect(e) {
      this.$emit("instrument-changed", e);
    },
    setSelected() {
      this.selectedInstrument = this.options.find(
        (instrument) => instrument.ID === this.selectedInstrumentId
      );
    },
    getInstruments() {
      return this.getAxios().get(`/bands/${this.bandId}/instruments`)
    },
    setupInstruments() {
      this.getInstruments().then((response) => {
        this.options = response.data;
        this.setSelected();
      })
    },
    getAxios() {
      return axios()
    }
  },
  created() {
    this.setupInstruments()
  },
};
</script>

<style></style>
