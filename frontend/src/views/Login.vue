<template>
  <form id="login-form" @submit.prevent="submitForm">
    <div class="screen">
      <h1>Log In</h1>
      <h2 v-if="credentialsIncorrect" id="incorrect-credentials">Wachtwoord of emailadres niet juist</h2>
      <div class="form-control">
        <label for="email">Email</label>
        <input class="inputbox" id="email" name="email" type="email" v-model="email" />
      </div>

      <div class="form-control">
        <label for="password">Password</label>
        <input class="inputbox" id="password" name="password" type="password" v-model="password" />
      </div>

      <div class="modal__actions">
        <Button type="button" class="red smallbutton" @click="returnHome">Cancel</Button>
        <Button type="submit" class="green smallbutton">Log In</Button>
      </div>
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
      email: "",
      password: "",
      credentialsIncorrect: false,
    };
  },
  methods: {
    submitForm() {
      const user = {
        email: this.email,
        password: this.password,
      };
      this.$store.commit("setCredentials", user);
      this.login()
        .then((response) => {
          this.$store.commit("setUser", response.data);
          if (!response.test) {
            router.push({ name: "Bands" });
          }
        })
        .catch(() => {
          this.activateCredentialsIncorrect();
          this.$store.commit("clearCredentials");
        });
    },
    login() {
      return this.getAxios().get(`/login/${this.email}`)
    },
    returnHome() {
      router.push({ name: "Home" });
    },
    activateCredentialsIncorrect() {
      this.credentialsIncorrect = true;
    },
    getAxios() {
      return axios()
    }
  },
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
        padding-bottom: 3%;
    }
}

.smallbutton {
  padding: 1em;
  width: 45%;
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

.form-control {
  margin: 0.5rem 0;
  text-align: center;
  padding-bottom: 3%;
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

input,
select {
  display: block;
  width: 100%;
  font: inherit;
  margin-top: 0.5rem;
}

select {
  width: auto;
}

input[type="checkbox"],
input[type="radio"] {
  display: inline-block;
  width: auto;
  margin-right: 1rem;
}

input[type="checkbox"] + label,
input[type="radio"] + label {
  font-weight: normal;
}
</style>
