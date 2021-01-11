<template>
    <div class="benefit">
        <div class="icon-container">
            <img :src="require('@/assets/' + src)" :alt="alt">
        </div>
        <div class="desc"><slot /></div>
    </div>
</template>

<script>
export default {
    name: "Benefit",
    props: {
      src: String,
      alt: String,
    }
}
</script>

<style lang="scss" scoped>
    @import '../variables';
    .benefit {
      color: $light_color;
      font-weight: bold;
      font-size: 24px;

      display: grid;
      grid-template-columns: 50px 100px;
      grid-column-gap: 10px;

      .icon-container {
        width: 100%;
        height: 100%;
      }

      .desc {
        margin: calc(50px - 24px / 2) 0;
      }
    }
</style>