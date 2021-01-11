<template>
  <div class="list">
    <input
      type="text"
      class="filter"
      v-model="filterInput"
      placeholder="Filter..."/>
    <div class="listElement" v-for="instrument in instruments" :key="instrument.ID">
      <div class="card" v-if="instrument.Name.indexOf(filterInput) > -1">
          <div class="name">
            {{ instrument.Name }}
          </div>
          <div class="edit">
            <Button class="green small" @Click="openEditInstrumentModal(instrument.ID)">
              Edit</Button>
          </div>
      </div>
    </div>
    <Button @Click="openCreationModal" class="green small button-margin" v-if="isOwner"
      >Add a new instrument</Button
    >
    <Modal
      :visible="creationModalVisibility"
      :title="this.modalTitle">
      <h2 class="error">{{ errorMessage }}</h2>
      <CreateInstrumentForm
      ref="createInstrument"
      :instrument="prefilledInstrument"
      :deleteButtonContent="deleteMessage"
        @close="closeCreationModal()"
        @accept="createInstrument"
        @delete="deleteInstrument()"/>
    </Modal>
  </div>
</template>

<script>
import axios from "../services/axios";
import Button from "../components/Button";
import Modal from "../components/Modal.vue";
import CreateInstrumentForm from "../components/CreateInstrumentForm.vue";

export default {
  name: "InstrumentOverlay",
  components: {
    Button,
    Modal,
    CreateInstrumentForm,
  },
  props: ["with-add-action", "band-id"],
  data() {
    return {
      filterInput: "",
      prefilledInstrument: {},
      instruments: [],
      creationModalVisibility: false,
      modalTitle: "",
      errorMessage: "",
      deleteMessage: null,
    };
  },
  created() {
    this.getInstruments();
  },
  computed: {
    isOwner() {
      const route = this.$route || {params:{isOwner: false}}
      return route.params.isOwner === "true" ? true : false;
    }
  },
  methods: {
    closeCreationModal() {
      this.creationModalVisibility = false;
      this.prefilledInstrument = {};
      this.deleteMessage = null;
      this.errorMessage = "";
    },
    openCreationModal() {
      this.modalTitle = "Create Instrument";
      this.creationModalVisibility = true;
    },
    instrumentSelected(instrument) {
      this.$emit("instrument-selected", instrument);
    },
    getInstruments() {
      this.getInstrumentPromise()
        .then((response) => {
          this.instruments = response.data;
        });
    },
    getInstrumentPromise() {
      return this.getAxios().get(`/bands/${this.bandId}/instruments`)
    },
    getAxios() {
      return axios()
    },
    async createInstrument(formData) {
      if (formData.valid) {
        const instrument = {
          Name: formData.instrument.name,
          BandID: parseInt(this.bandId),
        };
        const response = await axios().post(
          `/bands/${this.bandId}/instruments`,
          instrument
        );
        if (response == undefined) {
          this.errorMessage =
            "Looks like something went wrong on our end. Try again later";
          return;
        }
        if (response.status == 200) {
          this.creationModalVisibility = false;
          this.getInstruments();
        } else {
          this.errorMessage =
            "Looks like something went wrong on our end. Try again later";
        }
      } else {
        this.errorMessage = "Invalid instrument";
      }
    },
    openEditInstrumentModal(event) {
      this.prefilledInstrument = this.instruments.filter((instrument) => instrument.ID == event)[0];
      this.modalTitle = `Edit ${this.prefilledInstrument.Name}`;
      this.creationModalVisibility = true;
    },
    deleteInstrument() {
      axios()
      .delete(`/bands/${this.bandId}/instruments/${this.prefilledInstrument.ID}`)
      .then((response) => {
        if (response.status == 200) {
          this.closeCreationModal();
          this.getInstruments();
        }
      });
    }
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";

.button-margin {
  margin-bottom: 1.5em;
}

.list {
  width: 100%;
  font-family: "Rubik", sans-serif;

  .filter {
    background: none;
    border: 5px solid $red_color;
    border-radius: 16px;
    padding: 0.5em 1em;
    width: calc(100% - 2em - 10px);
    color: $light_color;
    font-size: 18px;
    margin-bottom: 1em;
    outline: none;
  }

  .card {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    background: darken($light_blue_color, 2);
    box-shadow: 4px 4px 4px rgba(0, 0, 0, 0.25);
    border-radius: 16px;
    margin-bottom: 1em;

    .name {
      display: flex;
      flex-direction: column;
      justify-content: center;
      color: $light_color;
      font-size: 18px;
      padding-left: 2em;
    }
  }
}
</style>
