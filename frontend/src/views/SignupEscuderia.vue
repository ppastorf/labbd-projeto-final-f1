<template>
  <form @submit.prevent="onSubmit">
    <!-- First and Last Name Row -->
    <div class="page-header mb-4">
                <h3 class="page-heading">Cadastrar Escuderia</h3>
              </div>
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Nome:</label><input class="form-control" placeholder="Nome" type="text" v-model="v$.form.nome.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.nome.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">URL:</label><input class="form-control" placeholder="URL" type="text" v-model="v$.form.url.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.url.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>
    </div>


    <!-- Date Row -->
    <div class="row">
      <div class="col-sm-6">
        <label for=""> Nascionalidade </label><input class="form-control" placeholder="Nascionalidade" type="Nascionalidade" v-model="v$.form.nascionalidade.$model">
        <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.nascionalidade.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
      </div>
    </div>

    <!-- Password and Confirm Password Row -->
    

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
import Cookie from '../assets/cookie.js'

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
        nome: '',
        url: '',
        nascionalidade: ''
      }
    }
  },

  validations() {
    return {
      form: {
        nome: { 
          required, name_validation: {
            $validator: validName,
            $message: 'Invalid Name. Valid name only contain letters, dashes (-) and spaces'
          } 
        },
        url: { 
          required, name_validation: {
            $validator: validName,
            $message: 'Invalid Name. Valid name only contain letters, dashes (-) and spaces'
          } 
        },
        nascionalidade: { required, name_validation: {
            $validator: validName,
            $message: 'Nascionalidade inválida. Contém caracteres não permitidos'
          }  }
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
        },
        getCookie() {
                this.userid = Cache.get("userid");
                this.tipo = Cache.get("tipo");
            }
    }
}
</script>