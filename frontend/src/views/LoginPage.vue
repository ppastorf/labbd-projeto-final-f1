<template>
  <form @submit.prevent="onSubmit">
    <!-- First and Last Name Row -->

    <!-- Email Row -->
    <div class="form-group"><div class="col-sm-6">
      <label for=""> Username</label><input class="form-control" placeholder="Enter username" type="name" v-model="v$.form.username.$model">
      <!-- Error Message -->
        <div class="input-errors" v-for="(error, index) of v$.form.username.$errors" :key="index">
          <div class="error-msg">{{ error.$message }}</div>
        </div>
        </div>
    </div>


    <!-- Password and Confirm Password Row -->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Password</label><input class="form-control" placeholder="Password" type="password" v-model="v$.form.password.$model">
          <div class="pre-icon os-icon os-icon-fingerprint"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.password.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Submit Button -->
    <div class="buttons-w">
      <button v-on:click="submit()" class="btn btn-primary" :disabled="v$.form.$invalid" style="margin-left:120px">Login</button>
    </div>
    
  </form>
</template>
<script>
import axios from 'axios'
import useVuelidate from '@vuelidate/core'
import { required, email, minLength, sameAs } from '@vuelidate/validators'
import Cookie from '../assets/cookie.js'

axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';



export default {

  setup () {
    return { v$: useVuelidate() }
  },

  data() {
    return {
      form: {
        username: '',
        password: '',
      }
    }
  },

  validations() {
    return {
      form: {
        username: { required },
        password: { required, min: minLength(5) }
      },
    }
  }, methods: {
        submit(){
            axios.post('http://localhost:5000/login', this.form)
                .then(function( response ){
                    console.log(response.data);
                     if(response.data){
                         if(response.data.tipo == "admin"){
                            this.$router.push({ name: 'Admin' , params: { username: this.form.username }})
                            console.log("redirecionando");
                         }else if(response.data.tipo == "piloto"){
                            this.$router.push({ name: 'User' , params: { username: this.form.username }})
                            console.log(this.form.username);
                         }
                     }
                    //     if(response.data.user_type == "admin"){
                    //         console.log("admin");
                    //     }else{
                    //         console.log(response.)
                    //     }
                    // }else{
                    //     console.log("piru")
                    // }
                }.bind(this));
        },
        getCookie() {
                this.userid = Cache.get("userid");
                this.tipo = Cache.get("tipo");
            }
    }
}
</script>