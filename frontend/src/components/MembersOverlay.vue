<template>
  <div>
    <input
      type="text"
      class="filter"
      v-model="filterInput"
      placeholder="Filter..."
    />
    <div v-for="(request, id) in joinRequests" :key="id" class="requestList">
      <div class="jrcard">
        <div class="align-vertical">
          <div class="name">{{ request }}</div>
        </div>
        <div class="actions">
          <Button class="red small" @Click="openRequestModal(id, 'deny')">Deny</Button>
          <Button class="green small" @Click="openRequestModal(id, 'accept')">Accept</Button>
        </div>
      </div>
    </div>
    <div v-for="member in bandmembers" :key="member.MusicianID">
      <BandMemberListItem
        :member="member"
        :isOwner="isOwner"
        @member-updated="getMembers"
        @delete="deleteMemberFromBand"
      ></BandMemberListItem>
    </div>
    <Button v-if="isOwner"
      @Click="openCreationModal" 
      class="green small button-margin">Invite a member
    </Button>
    <Modal
      :visible="creationModalVisibility"
      title="Invite Member">
      <h2 class="error">{{ errorMessage }}</h2>
      <AddMemberForm ref="addMember" 
        @close="closeCreationModal()"
        @accept="addMemberToBand" />
    </Modal>
    <Modal :visible="requestModalVisibility" :title="requestModalTitle">
      <div class="modal__actions">
          <Button class="red medium" @Click="closeRequestModal()">Cancel</Button>
          <Button class="green medium" @Click="processRequest()">Continue</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import axios from "../services/axios";
import Button from "../components/Button";
import Modal from "../components/Modal.vue";
import AddMemberForm from "../components/AddMemberForm.vue";
import BandMemberListItem from "./BandMemberListItem.vue";

export default {
  components: {
    Button,
    Modal,
    AddMemberForm,
    BandMemberListItem,
  },
  props: ["band-id"],
  data() {
    return {
      filterInput: "",
      bandmembers: [],
      creationModalVisibility: false,
      errorMessage: "",
      requestModalVisibility: false,
      requestModalTitle: "",
      requestID: 0,
      requestAction: "",
      joinRequests: {}
    };
  },
  created() {
    this.setupBandmembers()
    this.setupJoinRequests()
  },
  computed: {
    isOwner() {
      const route = this.$route || {params:{isOwner: false}}
      return route.params.isOwner === "true" ? true : false;
    }
  },
  methods: {
    getAxios() {
      return axios()
    },
    getMembers() {
      return this.getAxios().get(`/bands/${this.bandId}/members`)
    },
    getJoinRequests() {
      return this.getAxios().get(`/bands/${this.bandId}/request`)
    },
    setupBandmembers() {
      this.getMembers().then((response) => {
        this.bandmembers = response.data
      });
    },
    setupJoinRequests() {
      this.getJoinRequests().then((response) => {
        this.joinRequests = response.data
        console.log(this.joinRequests)
      })
    },
    closeCreationModal() {
      this.creationModalVisibility = false;
    },
    openCreationModal() {
      this.creationModalVisibility = true;
    },
    async addMemberToBand(formData) {
      if (formData.valid) {
        const response = await this.getAxios().post(`/bands/${this.bandId}/invite`, {
          email: formData.member.email,
        });
        if (response == undefined) {
          this.errorMessage =
            "Looks like something went wrong on our end. Try again later";
          return;
        }
        if (response.status == 200 || response.status == 201) {
          this.creationModalVisibility = false;
        } else {
          this.errorMessage =
            "Looks like something went wrong on our end. Try again later";
        }
      } else {
        console.log("Invalid Membername");
      }
    },
    async deleteMemberFromBand(musicianId) {
      this.getAxios()
      .delete(`/bands/${this.bandId}/members/${musicianId}`)      
      .then((response) => {
        if (response.status == 200) {
          this.closeCreationModal();
          this.setupBandmembers();
        }
      });
    },
    openRequestModal(id, action) {
      this.requestAction = action
      this.requestID = id
      this.requestModalTitle = `Weet je zeker dat je ${this.joinRequests[id]} wilt ${action == 'deny' ? 'weigeren' : 'accepteren'}?`
      this.requestModalVisibility = true
    },
    closeRequestModal() {
      this.requestAction = ''
      this.requestID = 0
      this.requestModalTitle = ''
      this.requestModalVisibility = false
    },
    processRequest() {
      this.getAxios().put(`/bands/${this.bandId}/request/${this.requestID}/${this.requestAction}`).then(() => {
        this.setupJoinRequests()
        this.setupBandmembers()
        this.closeRequestModal()
      })
    }
  },
};
</script>

<style lang="scss" scoped>
@import "../variables";

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

  .jrcard {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    background-color: $light_blue_color;
    margin: 1em 0;
    padding: 1em;
    border-radius: 16px;
    box-shadow: 4px 4px 8px rgba(0, 0, 0, .15);

    .name {
      color: $light_color;
      font-weight: bold;
    } 

   .actions {
      display: flex;
      flex-direction: row;

      Button {
        width: 100px;

        &:nth-child(1) {
          border-radius: 16px 0 0 16px;
        }
        &:nth-child(2) {
          border-radius: 0 16px 16px 0;
        }
      }
    }
  }
</style>
