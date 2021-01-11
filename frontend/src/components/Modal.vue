<template>
    <div class="background" v-if=visible>
        <div class="screen">
            <h1>{{ title }}</h1>
            <div class="slot">
                <slot />
            </div>
        </div>
    </div>
</template>

<script>
export default {
    props: {
        visible: Boolean,
        title: String,
    }
}
</script>

<style lang="scss" scoped>
    @import '../variables';
    .background {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background: rgba(0, 0, 0, .5);
    
        .screen {
            position: relative;
            width: 50%;
            top: 50%;
            left: 50%;
            transform: translate(-50%, -50%);
            background: $dark_blue_color;
            padding: 2em;
            border-radius: 16px;
            box-shadow: 4px 4px 4px rgba(0, 0, 0, .25);

            h1 {
                width: 100%;
                text-align: center;
            }

            .slot {
                width: 80%;
                margin: 0 auto;
            }

            .kickable {
                padding-top: 1em;
            }

        }
    }
</style>