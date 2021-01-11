<template>
    <div class="bandsview">
        <div class="bands">
            <h2>Bands</h2>
            <div v-for="band in bands" :key="band.id" class="card">
                <band-list-item :band=band></band-list-item>
            </div>
        </div>
        <div class="actions">
            <Button class="green" @Click="openJoinModal">Join a band</Button>
            <Button class="red" @Click="openCreationModal">Create a new band</Button>
        </div>
        <Modal :visible="creationModalVisibility" title="Create Band">
            <h2 class="error">{{ errorMessage }}</h2>
            <CreateBandForm ref="createBand" @close="closeCreationModal()" @accept="createBand" />
        </Modal>
        <Modal :visible="joinModalVisibility" title="Join a band">
            <label>
                Band Code <input id="joinBandInput" type="text" v-model="joinBandCode" />
            </label>
            <div class="modal__actions">
                <Button class="red medium" @Click="closeJoinModal()">Cancel</Button>
                <Button class="green medium" @Click="sendJoinRequest()">Continue</Button>
            </div>
        </Modal>
    </div>
</template>

<script>
import Button from '../components/Button.vue'
import Modal from '../components/Modal.vue'
import CreateBandForm from '../components/CreateBandForm.vue'
import BandListItem from '../components/BandListItem.vue'
import axios from '../services/axios'

export default {
    components: {
        Button,
        Modal,
        CreateBandForm,
        BandListItem,
    },
    data: function() {
        return {
            creationModalVisibility: false,
            errorMessage: "",
            joinModalVisibility: false,
            joinBandCode: "",
            bands: []
        }
    },
    created() {
        this.getBands();
    },
    methods: {
        closeCreationModal() {
            this.creationModalVisibility = false
        },
        openCreationModal() {
            this.creationModalVisibility = true
        },
        closeJoinModal(){
            this.joinModalVisibility = false
            this.joinBandCode = ""
        },
        openJoinModal(){
            this.joinModalVisibility = true
        },
        async createBand(formData) {
            if (formData.valid) {
                const response = await this.getAxios().post(`/users/${this.$store.state.user.MusicianID}/bands`, formData.band)
                if (response == undefined) {
                    this.errorMessage = "Looks like something went wrong on our end. Try again later"
                    return
                }
                if (response.status == 200) {
                    this.$router.push({name: "Band", params: {bandId: response.data}})
                } else {
                    this.errorMessage = "Looks like something went wrong on our end. Try again later"
                }
            } else {
                console.log("Invalid Band")
            }
        },
        async getBands() {
            let response = await this.getAxios().get(`/users/${this.$store.state.user.MusicianID}/bands`)
            this.bands = response.data

            response = await this.getAxios().get(`/users/${this.$store.state.user.MusicianID}/rehearsals`)
            this.bands.forEach(band => {
                if (response.data.includes(band.ID)) {
                    band.session = true
                }
            });
        },
        async sendJoinRequest() {
            this.getAxios().post(`/bands/${this.joinBandCode}/request`).then(() => {
                this.closeJoinModal()
                alert("Request send succesfully!")
            }).catch((err) => {
                alert("Something went wrong, please check the code you're trying to join (it should be only numbers). If the code is correct, then please try again later.")
                console.err(err)
            })
        },
        getAxios() {
            return axios()
        }
    }
}
</script>

<style lang="scss" scoped>
    @import '../variables';
    .bandsview {
        font-family: 'Rubik', sans-serif;
        color: $light_color;
        width: 100%;
    }

    .card {
        display: block;
        background: $light_blue_color;
        border-radius: 16px;
        box-shadow: 4px 4px 4px rgba(0, 0, 0, .25);
        margin: 1.5em 0;
        width: 100%;

    }

    h2.error {
        color: $red_color;
        font-size: 24px;
        text-align: center;
        margin: 0;
    }

    .actions {
        width: 100%;
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-column-gap: 1em;
    }

</style>
