<template>
  <form id="login-form" @submit.prevent="submitForm">
    <div class="form-control">
      <label for="user-name">Username</label>
      <input class="box" id="user-name" name="user-name" type="text" v-model="userName" required />
    </div>
    <div class="form-control">
      <label for="email">Email</label>
      <input  class="box" id="email" name="email" type="email" v-model="email" required />
    </div>

    <div class="form-control">
      <label for="password">Password</label>
      <input  class="box" id="password" name="password" type="password" v-model="password" required />
    </div>
    <div class="form-control">
      <label for="repeatpassword">Repeat Password</label>
      <input class="box" id="repeatpassword" name="repeatpassword" type="password" v-model="repeatedPassword" required />
    </div>

    <div>
      <button>Cancel</button>
      <button id="signup" :disabled="!formValid">Sign Up</button>
    </div>
  </form>
</template>

<script>
import axios from '../services/axios';
import { defineComponent } from 'vue';
import router from '../router'

export default defineComponent({
  name: 'SignUp',
  data() {
    return {
      userName: '',
      email: '',
      password: '',
      repeatedPassword: ''
    };
  },
  computed: {
    formValid() {
      // ...
      return true;
    }
  },
  methods: {
    submitForm() {
      if (this.password === this.repeatedPassword)
      {
        const user = {
            username: this.userName,
            email: this.email,
            password: this.password
        }

        this.userName = '';
        this.email = '';
        this.password = '';
        this.repeatedPassword = '';

        axios().post('/register', user).then((result) => {
          router.push("/login")
        });
      } else {
        console.log('WARNING: Entered passwords aren\'t equal!');
      }
    },
  },
});
</script>


<style lang="scss" scoped>
@import "../variables";

  .box {
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

form {
  margin: 2rem auto;
  max-width: 40rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.26);
  padding: 2rem;
  background-color: #ffffff;
}

.form-control {
  margin: 0.5rem 0;
}

label {
  font-weight: bold;
}

h2 {
  font-size: 1rem;
  margin: 0.5rem 0;
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

input[type='checkbox'],
input[type='radio'] {
  display: inline-block;
  width: auto;
  margin-right: 1rem;
}

input[type='checkbox'] + label,
input[type='radio'] + label {
  font-weight: normal;
}

button {
  font: inherit;
  border: 1px solid #0076bb;
  background-color: #0076bb;
  color: white;
  cursor: pointer;
  padding: 0.75rem 2rem;
  border-radius: 30px;
}

button:hover,
button:active {
  border-color: #002350;
  background-color: #002350;
}
</style>