<template>
    <div>
        <div class="top-right" @click="openModal">
            <img src="../assets/trash.svg" alt="Trashcan">
        </div>
        <Modal :visible="modalVisible" title="Are you sure you want to delete your band?">
            <div class="centered">This action cannot be reversed</div>
            <div class="modal__actions">
                <Button class="red medium" @Click="closeModal">Cancel</Button>
                <Button class="green medium" @Click="$emit('delete')">Delete</Button>
            </div>
        </Modal>
    </div>
</template>

<script>
import Modal from "./Modal.vue"
import Button from "./Button.vue"
export default {
    name: "DeleteBand",
    components: {
        Modal,
        Button
    },
    data() {
        return {
            modalVisible: false,
        }
    },
    methods: {
        openModal() {
            this.modalVisible = true
        },
        closeModal() {
            this.modalVisible = false
        }
    }    
}
</script>

<style lang="scss" scoped>
    .top-right {
        position: absolute;
        top: 0;
        right: 0;
        width: 50px;
        height: 50px;
        padding: 10px;

        &:hover {
            cursor: pointer;
        }
    }
    .centered {
        text-align: center;
        padding-bottom: 50px;
        font-size: 24px;
    }
</style>