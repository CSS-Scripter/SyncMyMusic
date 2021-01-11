<template>
  <form @submit.prevent="submitForm">
    <div class="screen">
      <h1 class="big-bottom-space">Account</h1>
      <h2 v-if="nameincorrect">Naam voldoet niet</h2>
      <h2 v-if="passwordsincorrect">Wachtwoorden voldoen niet</h2>
      <div class="form-control big-bottom-space">
        <label for="name">Name</label>
        <input class="inputbox" id="name" name="name" type="text" v-model="name" />
      </div>

      <div class="form-control">
        <label for="password">New password</label>
        <input class="inputbox" id="password" name="password" type="password" v-model="password" />
      </div>

      <div class="form-control">
        <label for="repeatpassword">Repeat new password</label>
        <input class="inputbox" id="repeatpassword" name="repeatpassword" type="password" v-model="repeatpassword" />
      </div>

      <div class="modal__actions">
        <Button class="red smallbutton" @click="returnHome">Back</Button>
        <Button class="green smallbutton">Update</Button>
      </div>
      <Button class="red widebutton" @click="logOut">Log Out</Button>
    </div>
  </form>
</template>

<script>
import axios from "../services/axios";
import router from "../router/index";
import Button from '../components/Button.vue'

export default {
  components: {
        Button,
  },
  data() {
    return {
      name: "",
      password: "",
      repeatpassword: "",
      nameincorrect: false,
      passwordsincorrect: false,
    };
  },
  computed: {
      isLoggedIn: function() {
          return this.$store.getters.isLoggedIn;
      }
  },
  methods: {
    submitForm() {
      this.nameincorrect = false;
      this.passwordsincorrect = false;
      if (this.name.length < 3) {
        this.nameincorrect = true;
        return;
      }      
      if (this.password !== this.repeatpassword || this.password.length < 3) {
        this.passwordsincorrect = true;
        return;
      }

      const user = {
          username: this.name,
          email: this.$store.state.user.email,
          password: this.password
      }

      axios().put(`/users/${this.$store.state.user.MusicianID}`, user).then((response) => {
          if (response.status == 200) {
            this.logOut();
          }
          else {
            console.log(response);
          }
        });
    },
    returnHome() {
      if (this.isLoggedIn) {
        router.push({ name: "Bands" });
        return;
      }
      router.push({ name: "Home" });
    },
    logOut: function() {
        router.push({ name: "Home" });
        this.$store.commit("clearCredentials")
    },
    setName() {
      this.name = this.$store.state.user.username
    }
  },
  created() {
    this.setName()
  }
};
</script>

<style lang="scss" scoped>
@import "../variables";

form {
  margin: 2rem auto;
  padding: 2rem;
}

.screen {
    position: absolute;
    width: 40%;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    padding: 2rem;
    margin: auto;

    h1 {
        width: 100%;
        font-family: 'Rubik', sans-serif;
        text-align: center;
        color: $light_color;
    }
}

.widebutton {
  padding: 1em;
  width: 100%;
  margin-top: 20px;
}

.smallbutton {
  padding: 1em;
  width: 45%;
  margin-bottom: 20px;
}

.inputbox {
  background: none;
  border: 5px solid $green_color;
  border-radius: 16px;
  padding: 0.5em 1em;
  width: calc(100% - 2em - 10px);
  color: $light_color;
  font-size: 18px;
  font-weight: bold;
  text-align: center;
  margin-bottom: 1em;
  outline: none;
}

.big-bottom-space {
  padding-bottom: 3%;
}

.form-control {
  margin: 0.5rem 0;
  text-align: center;
}

label {
  font-family: 'Rubik', sans-serif;
  font-weight: bold;
  color: $light_color;
}

h2 {
  font-size: 1rem;
  margin: 0.5rem 0;
  color: $red_color;
}
</style>
