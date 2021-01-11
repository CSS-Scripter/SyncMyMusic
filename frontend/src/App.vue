<template>
  <div class="site">
    <Header />
    <div class="content">
      <router-view />
    </div>
    <span class="spacer" />
    <Footer />
  </div>
</template>

<script>
import Header from "./components/Header";
import Footer from "./components/Footer";

export default {
  components: {
    Header,
    Footer,
  },
};
</script>

<style lang="scss">
@import "./variables";
img {
  width: 100%;
  height: 100%;
  user-select: none;
}

form.modal {
  font-size: 18px;
  color: $light_color;

  label {
    padding-left: 0.5em;
    padding-bottom: 0.2em;
  }

  input,
  label {
    display: block;
    width: 100%;
  }

  input {
    border-radius: 16px;
    border: 5px solid $light_color;

    font-size: 18px;
    color: $light_color;
    background: none;
    padding: 0.2em 0.4em;
    outline: none;

    &.invalid {
      border-color: $red_color;
    }

    &.valid {
      border-color: $green_color;
    }
  }

  p.error {
    margin-top: 0;
    color: $red_color;
  }
}


.modal__actions {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    padding-top: 1em;

    Button {
        width: 45%;
    }
}

.align-vertical {
  display: flex;
  flex-direction: column;
  justify-content: center;
}
</style>

<style lang="scss" scoped>
.site {
  height: 100%;
  width: 100%;
  display: flex;
  flex-direction: column;
}

.content {
  width: 100%;
  max-width: 1000px;
  margin: 0 auto;
  height: 100%;
}

.spacer {
  flex: 1;
}
</style>
