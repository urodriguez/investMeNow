<template>
  <div class="app flex-row align-items-center">
    <div class="container">
      <b-row class="justify-content-center">
        <b-col md="8">
          <b-card-group>
            <b-card no-body class="p-4">
              <b-card-body>
                <b-form>
                  <div class="text-center"><img class="navbar-brand-full" src="img/brand/logo.png" style="max-width:250px; max-height:250px;"></div>
                  <p class="text-muted">Sign In to your account</p>
                  <b-input-group class="mb-3">
                    <b-input-group-prepend><b-input-group-text><i class="icon-user"></i></b-input-group-text></b-input-group-prepend>
                    <b-form-input type="text" class="form-control" placeholder="Username" autocomplete="username email" />
                  </b-input-group>
                  <b-input-group class="mb-4">
                    <b-input-group-prepend><b-input-group-text><i class="icon-lock"></i></b-input-group-text></b-input-group-prepend>
                    <b-form-input type="password" class="form-control" placeholder="Password" autocomplete="current-password"/>
                  </b-input-group>
                  <b-row>
                    <b-col cols="6">
                      <b-button variant="primary" v-on:click="login" class="">Login <i class="fas fa-angle-right"></i></b-button>
                    </b-col>
                    <b-col cols="6" class="text-right">
                      <b-button variant="link" class="px-0">Forgot password?</b-button>
                    </b-col>
                  </b-row>
                </b-form>
              </b-card-body>
            </b-card>
            <b-card no-body class="text-white bg-primary py-5 d-md-down-none" style="width:44%">
              <b-card-body class="text-center">
                <div class="text-center">
                  <img class="navbar-brand-full" src="img/brand/logo-inver.png" style="max-width:250px; max-height:250px;">
                  </div>
                <div>
                  <h2>Sign up</h2>
                  <b-button variant="primary" class="active mt-3">Register Now <i class="fas fa-sign-in-alt"></i></b-button>
                </div>
              </b-card-body>
            </b-card>
          </b-card-group>
        </b-col>
      </b-row>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: 'Login',
  methods: {
    login: function () {
      axios.post('http://localhost:3000/login', {UserName: "user", Password: "user"},
      { 
        headers: { 'Content-type': 'application/x-www-form-urlencoded', }
      })
      .then(function (response) {
        this.$router.push('/round')
      }.bind(this))
      .catch(function (error) {
        alert("Invalid login")
      });
    }
  }
}
</script>
