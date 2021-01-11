<template>
  <div class="list">
    <div
      class="element"
      v-for="instrument of instrumentsWithPDF"
      :key="instrument.ID"
    >
      <div class="name">{{ instrument.Name }}</div>
      <div class="actions">
        <Button
          class="small red"
          v-if="instrument.hasPDF"
          @click="
            $emit('cclick', {
              action: 'edit',
              instrumentName: instrument.Name,
              instrumentId: instrument.ID,
            })
          "
          >Edit</Button
        >
        <Button
          class="small green"
          v-else
          @click="
            $emit('cclick', {
              action: 'add',
              instrumentName: instrument.Name,
              instrumentId: instrument.ID,
            })
          "
          >Add</Button
        >
      </div>
    </div>
  </div>
</template>

<script>
import Button from "./Button";
import axios from "../services/axios";

export default {
  name: "SongInstrumentsListView",
  components: {
    Button,
  },
  props: ["songId", "instruments"],
  data() {
    return {
      elements: [],
      instrumentsWithPDF: [],
    };
  },
  methods: {
    getAxios() {
      return axios()
    },
    setupInstruments() {
      this.getAxios()
        .get(`songs/${this.songId}/instruments`)
        .then((response) => {
          this.instrumentsWithPDF = [...this.instruments];
          for (const instrumentWithPDF of response.data) {
            const foundElement = this.instrumentsWithPDF.filter(
              (instrument) => instrument.ID == instrumentWithPDF.InstrumentID
            )[0];
            if (foundElement != undefined) {
              foundElement.hasPDF = true;
            }
          }

          const noPDFInstruments = this.instrumentsWithPDF.filter(
            (instrument) => !instrument.hasPDF
          );
          const PDFInstruments = this.instrumentsWithPDF.filter(
            (instrument) => instrument.hasPDF
          );
          this.instrumentsWithPDF = [...noPDFInstruments, ...PDFInstruments];
        });
    }
  },
  created: function() {
    this.setupInstruments()
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";

.list {
  margin-top: 2em;
  max-height: 300px;
  overflow-y: auto;

  .element {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    background-color: darken($light_blue_color, 2);
    border-radius: 16px;
    box-shadow: 4px 4px 4px rgba(0, 0, 0, 0.25);
    margin-bottom: 1em;

    .name {
      margin-top: 15px;
      padding-left: 2em;
    }

    .actions {
      width: 60px;
    }
  }
}
</style>
