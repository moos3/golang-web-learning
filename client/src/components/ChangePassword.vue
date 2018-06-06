<template>
  <div class="col-sm-4 col-sm-offset-4">
    <h2>Sign Up</h2>
    <div class="alert alert-danger" v-if="error">
      <p>{{ error }}</p>
    </div>
    <div class="form-group">
      <input
        type="text"
        class="form-control"
        placeholder="Enter your username"
        v-model="credentials.username"
        v-on:keyup.enter="submit"
      >
    </div>
    <div class="form-group">
      <input
        type="password"
        class="form-control"
        placeholder="Enter your existing password"
        v-model="credentials.password"
        v-on:keyup.enter="submit"
      >
    </div>
    <div class="form-group">
      <input
        type="password"
        class="form-control"
        placeholder="Enter your new password"
        v-model="credentials.NewPassword"
        v-on:keyup.enter="submit"
      >
    </div>
    <button class="btn btn-primary" @click="submit">Change Password</button>
  </div>
</template>

<script>
import auth from '../auth'

export default {

  data () {
    return {
      credentials: {
        username: '',
        password: '',
        NewPassword: '',
      },
      error: ''
    }
  },
  methods: {
    submit () {
      var credentials = {
        username: this.credentials.username,
        password: this.credentials.password,
        newPassword: this.credentials.NewPassword
      }
      // console.log('changepassword: credentials=', credentials)
      auth.changePassword(this, credentials, 'userinfo')
    }
  }
}
</script>
