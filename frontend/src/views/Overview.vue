<template>
  <div class="page-holder bg-gray-100">

        <div class="container-fluid px-lg-4 px-xl-5 contentDiv">
              <!-- Page Header-->
              <div class="page-header mb-4">
                <h1 class="page-heading">Profile</h1>
              </div>
          <section>
            <div class="row">
              <div class="col-lg-4">
                <form class="card mb-4">
                  <div class="card-header">
                    <h4 v-if="this.tipo==='escuderia'" class="card-heading">Escuderia</h4>
                    <h4 v-if="this.tipo==='admin'" class="card-heading">Admin</h4>
                    <h4 v-if="this.tipo==='piloto'" class="card-heading">Piloto</h4>
                  </div>
                  <div class="card-body">
                    <div class="row mb-3">
                      <div class="col-auto d-flex align-items-center"><img class="avatar avatar-lg p-1" src="https://therichpost.com/wp-content/uploads/2021/03/avatar2.png" alt="Avatar"></div>
                      <div class="card-footer text-end">
                    <button v-if="this.tipo==='admin'" id="btn1" class="btn btn-primary" v-on:click="submitPiloto()">Cadastrar Piloto</button>
                    <button v-if="this.tipo==='admin'" id="btn2" class="btn btn-primary" v-on:click="submitEscuderia()">Cadastrar Escuderia</button>
                  </div>
                    </div>
                  </div>
                </form>
              </div>
                <div class="col-lg-8">
                    <div v-if="this.tipo==='escuderia'" class="card-header"><div class="input-group"><input class="form-control" type="text" placeholder="Procure pelo nome do Piloto" v-model="search"><button v-on:click="submitSearch()" class="btn btn-outline-secondary" type="button">search<i class="fa fa-paper-plane"></i></button></div></div>
                    <div v-if="isHidden" id="text">
                      <table class="table table-bordered">
                        <thead>
                            <tr>
                                <th> Nome</th>
                                <th> Data de nascimento</th>
                                <th> Nascionalidade</th>
                            </tr>
                        </thead>
                        <tr v-for="loan in resultSearch">
                            <td>{{loan.nome}}</td>
                            <td>{{loan.data_nascimento}}</td>
                            <td>{{loan.nascionalidade}}</td>
                        </tr>


                        </table>
                    </div>
                    
                    <div class="card-header"><h4 class="card-heading">Overview</h4></div>
                    <div class="list-group-item border-start-0 border-end-0 py-5 border-top-0">
                        <div v-if="this.tipo==='admin'" class="admin">
                          <h6>Quantidade de pilotos cadastrados: </h6> <label class="form-label">{{ info.data.victories}}</label>
                          <h6>quantidade de escuderias cadastradas: </h6> <label class="form-label">{{ info.data.firt_year}}</label>
                          <h6>quantidade de corridas cadastradas: </h6> <label class="form-label">{{ info.data.last_year}}</label>
                          <h6>quantidade de temporadas: </h6> <label class="form-label">{{ info.data.victories}}</label>
                        </div>
                        <div v-if="this.tipo==='escuderia'" class="escuderia">
                          <h6>quantidade de vitórias da escuderia: </h6> <label class="form-label">{{ info.data.firt_year}}</label>
                          <h6>quantidade de pilotos diferentes que já correram pela escuderia: </h6> <label class="form-label">{{ info.data.last_year}}</label>
                          <h6>primeiro e ultimo ano em que ha dados da escuderia na base: </h6> <label class="form-label">{{ info.data.victories}}</label>
                        </div>
                        <div v-if="this.tipo==='escuderia'" class="escuderia">
                          <h6>quantidade de vitorias do piloto: </h6> <label class="form-label">{{ info.data.firt_year}}</label>
                          <h6>primeiro e ultimo ano em que ha dados do piloto na base : </h6> <label class="form-label">{{ info.data.last_year}}</label>
                        </div>
                    </div> 
                    
                    <div class="card-header"><h4 class="card-heading">Relatórios</h4></div>
                    <h5 class="card-heading">Relatório 5</h5>
                     <div class="list-group-item border-start-0 border-end-0 py-2 border-top-0">
                      <div id="textExample">
                        <table class="table table-bordered">
                          <thead>
                              <tr>
                                  <th>Status</th>
                                  <th>Count</th>
                              </tr>
                          </thead>
                          <tr v-for="loan in rel5.data.data">
                              <td>{{loan.status}}</td>
                              <td>{{loan.count}}</td>
                          </tr>
                        </table>
                      </div>
                    </div>  
                    <h5 class="card-heading">Relatório 6</h5>
                     <div class="list-group-item border-start-0 border-end-0 py-2 border-top-0">
                      <div id="textExample">
                        <table class="table table-bordered" v-if="this.tipo==='piloto'">
                          <thead>
                              <tr>
                                  <th>Ano</th>
                                  <th>Corridas</th>
                                  <th>Vitorias</th>

                              </tr>
                          </thead>
                          <tr v-for="loan in rel5.data.data">
                              <td>{{loan.ano}}</td>
                              <td>{{loan.corridas}}</td>
                              <td>{{loan.vitorias}}</td>
                              
                          </tr>
                        </table>
                        <table class="table table-bordered" v-if="this.tipo==='admin'">
                          <thead>
                              <tr>
                                  <th>Cidade</th>
                                  <th>Codigo iata</th>
                                  <th>Aeroporto</th>
                                  <th>Aeroporto cidade</th>
                                  <th>Tipo aeroporto</th>
                                  <th>Distância</th>
                              </tr>
                          </thead>
                          <tr v-for="loan in rel5.data.data">
                              <td>{{loan.cidade}}</td>
                              <td>{{loan.codigo}}</td> 
                              <td>{{loan.aeroporto}}</td>
                              <td>{{loan.aeroporto_cidade}}</td>   
                              <td>{{loan.distancia}}</td>                              
                          </tr>
                        </table>

                        <table class="table table-bordered" v-if="this.tipo==='escuderia'">
                          <thead>
                              <tr>
                                  <th>Nome</th>
                                  <th>Vitorias</th>

                              </tr>
                          </thead>
                          <tr v-for="loan in rel5.data.data">
                              <td>{{loan.nome}}</td>
                              <td>{{loan.vitorias}}</td>
                              
                          </tr>
                        </table>
                      </div>
                    </div>  
                </div>
            </div>
          </section>
        </div>
      </div>
    
</template>
<script>
//importing bootstrap 5 Modules
import "bootstrap/dist/css/bootstrap.min.css";
import axios from 'axios'
import Cookie from '../assets/cookie.js'

export default {
  data () {
    return {
      info: {},
      rel5: [],
      rel6: [],
      search: null,
      resultSearch: [],
      isHidden: false,
      username: '',
      userid:'',
      tipo:''
    }
  },
  methods: {
    async getDataOverview() {
      axios
      .get('http://localhost:5000/overviewpiloto')
      .then(response => (this.info = response))
    },
    async getDataRel5() {
      axios
      .get('http://localhost:5000/rel5')
      .then(response => (this.rel5 = response))
    },
    submitSearch(){
    axios.get('http://localhost:5000/pilotosearch?'+ this.search)
        .then(function( response ){
            console.log(response.data);
            this.resultSearch = response.data;
            this.isHidden = !this.isHidden;
              // if(response.data){
              //     if(response.data.user_type == "admin"){
              //       this.$router.push({ name: 'User' })
              //       console.log("redirecionando");
              //     }
              // }
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
  },
  created() {
    this.getDataOverview();
    this.getDataRel5();
  },mounted(){
      getCookie();
  }
}
</script>
<style scoped>
  #btn2 {
    margin-left: 10px;
  }
  #btn1 {
    margin-left: ;
  }
</style>