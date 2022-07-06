<template>
  <form @submit.prevent="onSubmit">
    <!-- First and Last Name Row -->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> First Name:</label><input class="form-control" placeholder="Enter first name" type="text" v-model="v$.form.firstName.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.firstName.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Last Name:</label><input class="form-control" placeholder="Enter last name" type="text" v-model="v$.form.lastName.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.lastName.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>
    </div>


    <!-- Email Row -->
    <div class="form-group">
      <label for=""> Email address</label><input class="form-control" placeholder="Enter email" type="email" v-model="v$.form.email.$model">
      <div class="pre-icon os-icon os-icon-email-2-at2"></div>
      <!-- Error Message -->
        <div class="input-errors" v-for="(error, index) of v$.form.email.$errors" :key="index">
          <div class="error-msg">{{ error.$message }}</div>
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

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Confirm Password</label><input @input="checkPassword()" class="form-control" placeholder="Confirm Password" type="password" v-model="v$.form.confirmPassword.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.confirmPassword.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Submit Button -->
    <div class="buttons-w">
      <button v-on:click="submit()" class="btn btn-primary" :disabled="v$.form.$invalid" style="margin-left:120px">Signup</button>
    </div>
    
  </form>
</template>
<script>
import axios from 'axios'
import useVuelidate from '@vuelidate/core'
import { required, email, minLength, sameAs } from '@vuelidate/validators'

axios.defaults.headers.common['Access-Control-Allow-Origin'] = '*';
export function validName(name) {
  let validNamePattern = new RegExp("^[a-zA-Z]+(?:[-'\\s][a-zA-Z]+)*$");
  if (validNamePattern.test(name)){
    return true;
  }
  return false;
}


export default {

  setup () {
    return { v$: useVuelidate() }
  },

  data() {
    return {
      form: {
        firstName: '',
        lastName: '',
        email: '',
        password: '',
        confirmPassword: '',
      }
    }
  },

  validations() {
    return {
      form: {
        firstName: { 
          required, name_validation: {
            $validator: validName,
            $message: 'Invalid Name. Valid name only contain letters, dashes (-) and spaces'
          } 
        },
        lastName: { 
          required, name_validation: {
            $validator: validName,
            $message: 'Invalid Name. Valid name only contain letters, dashes (-) and spaces'
          } 
        },
        email: { required, email },
        password: { required, min: minLength(6) },
        confirmPassword: { required }
      },
    }
  }, methods: {
        submit(){
            axios.post('https://eool0umaj3oyst0.m.pipedream.net', this.form)
                .then(function( response ){
                    console.log(response.data);
                     if(response.data){
                         if(response.data.user_type == "admin"){
                            this.$router.push({ name: 'User' , params: { email: this.form.email }})
                            console.log("redirecionando");
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
        }
    }
}
</script>