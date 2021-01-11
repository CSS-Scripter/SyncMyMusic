<template>
  <div class="member-list-item">
    <div class="card-content">
      <h3>{{ member.Musician.username }}, {{member.Instrument ? member.Instrument.Name : "no instrument selected"}}</h3>
      <Button @click="toggleEditUserModal" class="small red" v-if="allowEdit" id="edit-button">Edit</Button>
    </div>
    <Modal
      :visible=showModal
      :title="modalTitle">
      <h2 class="error">{{ errorMessage }}</h2>
      <SelectInstrument
        @instrument-changed="changeSelectedInstrument"
        :bandId="member.Band.ID"
        :selectedInstrumentId="member.Instrument ? member.Instrument.ID : null"
      />
      <Button v-if="allowEdit" class="red medium widebutton" @Click="$emit('delete', this.member.Musician.MusicianID)">Kick User</Button>
      <div class="modal__actions">
        <Button class="red medium" @Click="toggleEditUserModal">Cancel</Button>
        <Button class="green medium" @Click="editMember">Continue</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
// import router from '../router'
import axios from "../services/axios";
import Button from "./Button.vue";
import Modal from "./Modal.vue";
import SelectInstrument from "./SelectInstrument.vue";

export default {
  components: { Button, Modal, SelectInstrument },
  props: ["member", "isOwner"],
  data() {
    return {
      showModal: false,
      errorMessage: "",
      modalTitle: `Edit ${this.member.Musician.username}`,
      selectedInstrument: null,
    };
  },
  computed: {
    allowEdit() {
      return (this.isOwner || this.$store.state.user.MusicianID == this.member.Musician.MusicianID)
    }
  },
  methods: {
    changeSelectedInstrument(e) {
      this.selectedInstrument = e;
    },
    toggleEditUserModal() {
      this.showModal = !this.showModal;
      if (!this.showModal) {
        this.selectedInstrument = null;
      }
    },
    editMember() {
      axios()
        .put(
          `/bands/${this.member.Band.ID}/members/${this.member.Musician.MusicianID}`,
          {
            InstrumentID: this.selectedInstrument.ID,
          }
        )
        .then(() => {
          this.toggleEditUserModal();
          this.$emit("member-updated");
        });
    },
    kickMember() {
      axios()
      .delete(`/bands/${this.member.Band.ID}/members/${this.member.Musician.MusicianID}`)
      .then((response) => {
        if (response.status === 200) {
          this.toggleEditUserModal();
        }
      });
    }
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";
.card-content {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  padding: 1.5em 4em;
  background-color: $light_blue_color;
}

h3 {
  margin: 0;
  font-size: 20px;
  color: $light_color;
}

button {
  margin-top: 1em;
  width: 180px;
}

.widebutton {
  width: 100%;
}

.song {
  font-weight: lighter;
  font-size: 16px;
  font-style: italic;
}
</style>
