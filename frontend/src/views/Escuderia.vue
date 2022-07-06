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
                    <h4 class="card-heading">Escuderia</h4>
                  </div>
                  <div class="card-body">
                    <div class="row mb-3">
                      <div class="col-auto d-flex align-items-center"><img class="avatar avatar-lg p-1" src="https://therichpost.com/wp-content/uploads/2021/03/avatar2.png" alt="Avatar"></div>
                      <div class="col">
                        <label class="form-label">Name</label>
                        <input class="form-control" placeholder="Your name">
                      </div>
                    </div>
                    <div class="mb-3"> 
                      <label class="form-label">Bio</label>
                      <textarea class="form-control" rows="8">Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.</textarea>
                    </div>
                    <div class="mb-3"> 
                      <label class="form-label">Email</label>
                      <input class="form-control" placeholder="you@domain.com">
                    </div>
                    <label class="form-label">Password</label>
                    <input class="form-control" type="password" value="password">
                  </div>
                  <div class="card-footer text-end">
                    <button class="btn btn-primary">Save</button>
                  </div>
                </form>
              </div>
                <div class="col-lg-8">
                    <div class="card-header"><div class="input-group"><input class="form-control" type="text" placeholder="Procure pelo nome do Piloto" v-model="search"><button v-on:click="submitSearch()" class="btn btn-outline-secondary" type="button">search<i class="fa fa-paper-plane"></i></button></div></div>
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
                        <h6>Número de vitórias: </h6> <label class="form-label">{{ info.data.victories}}</label>
                        <h6>Primeiro ano de entrada: </h6> <label class="form-label">{{ info.data.firt_year}}</label>
                        <h6>Último ano de entrada: </h6> <label class="form-label">{{ info.data.last_year}}</label>
                    </div> 
                    <div class="card-header"><h4 class="card-heading">Relatórios</h4></div>
                     <div class="list-group-item border-start-0 border-end-0 py-2 border-top-0">
                      <div id="textExample">
                        <table class="table table-bordered">
                          <thead>
                              <tr>
                                  <th> Ano</th>
                                  <th> Corridas</th>
                                  <th> Vitorias</th>
                              </tr>
                          </thead>
                          <tr v-for="loan in rel5.data.data">
                              <td>{{loan.year}}</td>
                              <td>{{loan.corridas}}</td>
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

export default {
  data () {
    return {
      info: {},
      rel5: [],
      rel6: [],
      search: null,
      resultSearch: [],
      isHidden: false
    }
  },
  methods: {
    async getDataOverview() {
      axios
      .get('https://eox56vpp94quv3j.m.pipedream.net')
      .then(response => (this.info = response))
    },
    async getDataRel5() {
      axios
      .get('https://eov5ar2pl5ebk1h.m.pipedream.net')
      .then(response => (this.rel5 = response))
    },
    submitSearch(){
    axios.get('https://eoc5dvu9sek9v4m.m.pipedream.net?name='+ this.search)
        .then(function( response ){
            console.log(response.data);
            this.resultSearch = response.data.data;
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
    }
  },
  created() {
    this.getDataOverview();
    this.getDataRel5();
  },
}
</script>