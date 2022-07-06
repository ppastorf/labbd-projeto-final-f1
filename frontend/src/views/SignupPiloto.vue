<template>
  <form @submit.prevent="onSubmit" v-if="this.tipo==='admin'">
    <!-- First and Last Name Row -->
    <div class="page-header mb-4">
                <h3 class="page-heading">Cadastrar Piloto</h3>
              </div>
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Primeiro nome:</label><input class="form-control" placeholder="Primeiro nome" type="text" v-model="v$.form.firstName.$model">
          <div class="pre-icon os-icon os-icon-user-male-circle"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.firstName.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Sobrenome:</label><input class="form-control" placeholder="Sobrenome" type="text" v-model="v$.form.lastName.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.lastName.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>
    </div>


    <!-- Date Row -->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Data de nascimento:</label><input class="form-control" type="date" v-model="v$.form.date.$model">
          <div class="pre-icon os-icon os-icon-email-2-at2"></div>
          <!-- Error Message -->
            <div class="input-errors" v-for="(error, index) of v$.form.date.$errors" :key="index">
              <div class="error-msg">{{ error.$message }}</div>
            </div>
        </div>
      </div>
      <div class="col-sm-6">
        <label for=""> Nascionalidade:</label><input class="form-control" placeholder="Nascionalidade" type="Nascionalidade" v-model="v$.form.nascionalidade.$model">
        <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.nascionalidade.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
      </div>
    </div>

    <!-- Password and Confirm Password Row -->
    <div class="row">
      <div class="col-sm-6">
        <div class="form-group">
          <label for=""> Numero:</label><input class="form-control" placeholder="Numero" type="number" v-model="v$.form.numero.$model">
          <div class="pre-icon os-icon os-icon-fingerprint"></div>
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.numero.$errors" :key="index">
            <div class="error-msg">{{ error.$message }}</div>
          </div>
        </div>
      </div>

      <div class="col-sm-6">
        <div class="form-group">
          <label for="">Código:</label><input class="form-control" placeholder="Código" type="codigo" v-model="v$.form.codigo.$model">
          <!-- Error Message -->
          <div class="input-errors" v-for="(error, index) of v$.form.codigo.$errors" :key="index">
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
  <div v-else>
  Você não é administrador
  </div>
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
        firstName: '',
        lastName: '',
        numero: '',
        codigo: '',
        date: '',
        nascionalidade: ''
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
        numero: { required, Number },
        codigo: { required },
        date: { required },
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
    ,mounted(){
      getCookie();
  }
}
</script>