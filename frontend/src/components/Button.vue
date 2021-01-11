<template>
    <button><slot /></button>
</template>

<script>
export default {
    name: "Button"
}
</script>

<style lang="scss" scoped>
    @import '../variables';
    button {
        width: 100%;
        font-size: 24px;
        font-weight: bold;
        outline: none;
        border: none;
        padding: 1.5em;
        border-radius: 16px;
        user-select: none;
    
        &.green {
            background-color: $green_color;
        }
        &.red {
            background-color: $red_color;
        }

        &.small {
            font-size: 16px;
            padding: 1em;
        }
        &.medium {
            padding: 1em;
        }

        &:hover {
            cursor: pointer;
        }
    }
</style>