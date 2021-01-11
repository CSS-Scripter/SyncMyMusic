<template>
  <div class="page">
    <div class="banner">
      <div class="vertical-align">
        <div class="banner-text">
          <span class="highlighted">Stream</span> your band's <span class="highlighted">sheet music</span> easily to your band members, without any of the hassle
        </div>
      </div>
      <div class="banner-image">
        <img src="../assets/banner-image.svg" alt="Banner image (Note list by undraw)">
      </div>
    </div>
    <div class="pros">
      <Benefit src="zap.svg" alt="Lightning bolt (Feathericons zap.svg)">Fast</Benefit>
      <Benefit src="smile.svg" alt="Smile (Feathericons smile.svg)">Simple</Benefit>
      <Benefit src="no-money.svg" alt="Free (Feathericons dollar-sign.svg and slash.svg)">Free</Benefit>
    </div>
    <div class="bottom-banner">
      <div class="description">
        Always having trouble getting to the right song in your sheetmusic? Well no more! Someone picks a song, and you get the sheet music, easy as that! Still not convinced? Then just try it, it’s free! We promise not to sell your data.
      </div>
      <div class="get-started vertical-align">
        <Button @click="getStarted" class="green">Get Started</Button>
      </div>
    </div>
  </div>
</template>

<script>
  import Benefit from '../components/Benefit.vue';
  import Button from '../components/Button.vue';
  export default {
    name: 'Homepage',
    components: {
      Button,
      Benefit,
    },
    methods: {
      getStarted() {
        this.$router.push('register')
      }
    }
  }
</script>

<style lang="scss" scoped>
  @import '../variables';
  .vertical-align {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  span.highlighted {
    color: $red_color;
  }

  .page {
    font-family: "Rubik", sans-serif;
  }
  
  .banner {
    display: grid;
    grid-template-columns: 1fr 350px;
    grid-column-gap: 80px;
    padding: 4em;
    margin: 0 auto;

    .banner-text {
      font-weight: bold;
      font-style: normal;
      font-size: 24px;
      color: $light_color;
    }

    .banner-image {
      width: 350px;
      height: 311px;
    }
  }

  .pros {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    width: 620px;
    margin: 0 auto;
  }

  .bottom-banner {
    display: grid;
    grid-template-columns: 350px 350px;
    grid-column-gap: 200px;
    padding-top: 4em;

    .description {
      text-align: center;
      color: $light_color;
      font-weight: lighter;
      line-height: 24px;
      font-size: 18px;
    }
  }
</style>