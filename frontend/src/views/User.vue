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
                    <h4 class="card-heading">Piloto</h4>
                  </div>
                  <div class="card-body">
                    <div class="row mb-3">
                      <div class="col-auto d-flex align-items-center"><img class="avatar avatar-lg p-1" src="https://therichpost.com/wp-content/uploads/2021/03/avatar2.png" alt="Avatar"></div>
                      <div class="col">
                        <label class="form-label">Name </label><br>
                        <label class="form-label">{{username}}</label>
                      </div>
                    </div>
                  </div>
                  
                </form>
              </div>
                <div class="col-lg-8">
                    <div class="card-header"><h4 class="card-heading">Overview</h4></div>
                    <div class="list-group-item border-start-0 border-end-0 py-5 border-top-0">
                        <h6>Número de vitórias: </h6> <label class="form-label">{{ info.data.victories}}</label>
                        <h6>Primeiro ano de entrada: </h6> <label class="form-label">{{ info.data.firt_year}}</label>
                        <h6>Último ano de entrada: </h6> <label class="form-label">{{ info.data.last_year}}</label>
                    </div> 
                    <div class="card-header"><h4 class="card-heading">Relatórios</h4></div>
                    <div class="list-group-item border-start-0 border-end-0 py-2 border-top-0">
                        <h5 class="card-heading">Relatório 5</h5>
                        <div id="textExample">
                            <table class="table table-bordered">
                                <thead>
                                    <tr>
                                        <th> Ano</th>
                                        <th> Corridas</th>
                                        <th> Vitorias</th>
                                    </tr>
                                </thead>
                                <tr v-for="loan in rel5.data">
                                    <td>{{loan.year}}</td>
                                    <td>{{loan.corridas}}</td>
                                    <td>{{loan.vitorias}}</td>
                                </tr>


                            </table>
                        </div>
                        <h5 class="card-heading">Relatório 6</h5>
                        <div id="textExample">
                            <table class="table table-bordered">
                                <thead>
                                    <tr>
                                        <th> Status</th>
                                        <th> Count</th>
                                    </tr>
                                </thead>
                                <tr v-for="item in rel6.data">
                                    <td>{{item.Status}}</td>
                                    <td>{{item.Count}}</td>
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
      username: ''
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
    async getDataRel6() {
      axios
      .get('http://localhost:5000/rel6')
      .then(response => (this.rel6 = response))
    }
  },
  created() {
    this.getDataOverview();
    this.getDataRel5();
    this.getDataRel6();
  },
  mounted(){
      this.username = this.$route.params.username
      console.log(this.$route.params)
  }


}
</script>