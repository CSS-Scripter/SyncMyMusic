<template>
  <div class="song-list-view">
    <div class="list">
      <input
        type="text"
        class="filter"
        v-model="filterInput"
        placeholder="Filter..."/>
      <div class="listElement" v-for="song in songs" :key="song.ID">
        <div class="card" v-if="song.name.indexOf(filterInput) > -1">
          <div class="name">
            {{ song.name }}
          </div>
          <div class="edit">
            <Button class="green small" @Click="openEditSongModal(song.ID)">
              Edit</Button>
          </div>
        </div>
      </div>
      <Button class="small green" @Click="openCreateSongModal" v-if="isOwner">
      Add a new song
    </Button>
    </div>
    <Modal
      :visible="songModalVisibility"
      :title="songModalTitle">
      <CreateSongForm :song="prefilledSong" ref="createSong" @accept="createSong" @close="closeSongModal()" @delete="deleteSong">
        <SongInstrumentsListView
          v-bind:instruments="instruments"
          :songId="prefilledSong.ID"
          v-if="prefilledSong.ID != undefined"
          @cclick="openUploadPDFModal" />
      </CreateSongForm>
    </Modal>
    <Modal
      :visible="pdfUploadVisibility"
      :title="pdfUploadTitle">
      <h2 class="upload-error" v-if="uploadError != ''">{{ uploadError }}</h2>
      <div class="input-container">
        <input
          type="file"
          id="file"
          ref="file"
          v-bind:class="{ valid: pdfFile != undefined }"
          v-on:change="handleFileUpload()"/>
      </div>
      <div class="modal__actions">
          <Button class="red medium" @Click="closeUploadModal()">Cancel</Button>
          <Button class="green medium" @Click="uploadPDF()">Continue</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import Button from "./Button.vue";
import axios from "../services/axios";
import Modal from "../components/Modal";
import CreateSongForm from "../components/CreateSongForm";
import SongInstrumentsListView from "../components/SongInstrumentsListView";

export default {
  name: "SongsListView",
  components: {
    Button,
    Modal,
    CreateSongForm,
    SongInstrumentsListView,
  },
  props: {
    creationMessage: String,
    "band-id": Number,
  },
  data() {
    return {
      filterInput: "",
      songs: null,
      instruments: null,
      prefilledSong: {},
      songModalTitle: "",
      songModalVisibility: false,
      pdfUploadVisibility: false,
      pdfUploadTitle: "",
      pdfAction: "",
      pdfModalInstrumentID: 0,
      pdfFile: undefined,
      uploadError: "",
    };
  },
  computed: {
    isOwner() {
      const route = this.$route || {params:{isOwner: false}}
      return route.params.isOwner === "true" ? true : false;
    }
  },
  methods: {
    openCreateSongModal() {
      this.prefilledSong = {};
      this.songModalTitle = "Create a new song";
      this.songModalVisibility = true;
    },
    closeSongModal() {
      this.prefilledSong = {};
      this.songModalTitle = "";
      this.songModalVisibility = false;
      this.instruments.map((ins) => (ins.hasPDF = false));
    },
    createSong(formData) {
      if (JSON.stringify(formData) != JSON.stringify(this.prefilledSong)) {
        if (formData.ID == undefined) {
          this.getAxios()
            .post(`/bands/${this.bandId}/songs`, formData)
            .then((response) => {
              if (response.status == 200) {
                this.closeSongModal();
                formData.ID = response.data;
                console.log(formData)
                this.songs = [...this.songs, formData];
              }
            });
        } else {
          this.getAxios()
            .put(`/bands/${this.bandId}/songs/${formData.ID}`, formData)
            .then((response) => {
              if (response.status == 200) {
                this.closeSongModal();
                this.songs = [
                  ...this.songs.filter((song) => song.ID != formData.ID),
                  formData,
                ];
              }
            });
        }
      }
      this.closeSongModal();
    },
    deleteSong(formData) {
      this.getAxios().delete(`/songs/${formData.ID}`).then(() => {
        this.songs = this.songs.filter((song) => song.ID != formData.ID)
        this.closeSongModal()
      }).catch((err) => {
        alert("An error occured, please try again later")
        console.error(err)
      })
    },
    openEditSongModal(event) {
      this.prefilledSong = this.songs.filter((song) => song.ID == event)[0];
      this.songModalTitle = `Edit ${this.prefilledSong.name}`;
      this.songModalVisibility = true;
    },
    openUploadPDFModal(event) {
      this.pdfUploadVisibility = true;
      this.pdfAction = event.action;
      this.pdfModalInstrumentID = event.instrumentId;
      const eventAction =
        event.action.charAt(0).toUpperCase() + event.action.slice(1);
      this.pdfUploadTitle =
        eventAction + " PDF for instrument " + event.instrumentName;
    },
    closeUploadModal() {
      this.pdfUploadVisibility = false;
      this.pdfFile = undefined;
      this.pdfAction = "";
      this.pdfModalInstrumentID = 0;
      this.pdfUploadTitle = "";
    },
    handleFileUpload() {
      this.pdfFile = this.$refs.file.files[0];
      this.uploadError = "";
    },
    uploadPDF() {
      if (this.pdfFile == undefined) {
        this.uploadError = "There was no file selected!";
        return;
      }
      const formData = new FormData();
      formData.append("document", this.pdfFile);
      const url = `/songs/${this.prefilledSong.ID}/instruments/${this.pdfModalInstrumentID}`;
      if (this.pdfAction == "add") {
        this.getAxios()
          .post(url, formData, {
            headers: { "Content-Type": "multipart/form-data" },
          })
          .then((response) => {
            if (response.status == 200) {
              this.closeUploadModal();
            }
          });
      }
      if (this.pdfAction == "edit") {
        this.getAxios()
          .put(url, formData, {
            headers: { "Content-Type": "multipart/form-data" },
          })
          .then((response) => {
            if (response.status == 200) {
              this.closeUploadModal();
            }
          });
      }
    },
    setupSongs() {
      this.getSongPromise().then((response) => {
          this.songs = response.data || [];
        });
    },
    setupInstruments() {
      this.getInstrumentPromise().then((response) => {
          this.instruments = response.data || [];
        });
    },
    getSongPromise() {
      return this.getAxios().get(`/bands/${this.bandId}/songs`)
    },
    getInstrumentPromise() {
      return this.getAxios().get(`/bands/${this.bandId}/instruments`)
    },
    getAxios() {
      return axios()
    }
  },
  created() {
    this.setupSongs()
    this.setupInstruments()
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";
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
