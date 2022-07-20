<template>
  <div class="admin-container">
    <main class="admin-header">
      <div class="admin-container01">
        <img
          alt="image"
          src="/playground_assets/f1-1200h.png"
          loading="eager"
          class="admin-f1-logo"
        />
        <span class="admin-user-name heading1"><span>admin</span></span>
      </div>
    </main>
    <div class="admin-overview">
      <h1 class="admin-titulo">
        <span>over</span>
        <span></span>
        <span>view</span>
      </h1>
      <div class="admin-cards-overview">
        <div class="admin-overview-card">
          <h2 class="admin-texto"><span id="overview-pilotos">pilotos</span></h2>
          <img
            alt="image"
            src="/playground_assets/driver_report-200h.png"
            class="admin-icone"
          />
        </div>
        <div class="admin-overview-card1">
          <h2 class="admin-texto1">
            <span id="overview-escuderias">escuderias</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/escuderias-200h.jpg"
            class="admin-icone1"
          />
        </div>
        <div class="admin-overview-card2">
          <h2 class="admin-texto2">
            <span id="overview-corridas">corridas</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/circuitos-200h.png"
            class="admin-icone2"
          />
        </div>
        <div class="admin-overview-card3">
          <h2 class="admin-texto3"><span id="overview-temporadas">temporadas</span></h2>
          <img
            alt="image"
            src="/playground_assets/temporadas-200h.jpg"
            class="admin-icone3"
          />
        </div>
      </div>
    </div>
    <div class="admin-relatorios section-container">
      <div class="admin-container02 max-content-container">
        <h1 class="admin-titulo1">relatorios</h1>
        <div class="admin-container03">
          <div class="admin-card-relatorio">
            <div class="admin-container04">
              <h1 class="admin-titulo2">aeroportos</h1>
              <div class="admin-container05">
                <img
                  alt="profile"
                  src="/playground_assets/airports-200h.jpg"
                  class="admin-imagem-relatorio"
                />
                <form class="admin-form">
                  <div class="admin-container06">
                    <span class="admin-text10">cidade</span>
                    <input
                      type="text"
                      id="input-relatorio-aeroportos"
                      name="AiportsReportCity"
                      required
                      placeholder="URL"
                      class="admin-textinput input"
                    />
                  </div>
                  <button id="button-relatorio-aeroportos" class="admin-button button">
                    <span class="admin-text11"><span>CONSULTAR</span></span>
                  </button>
                </form>
              </div>
            </div>
            <div class="admin-container07">
              <h1 class="admin-titulo3"><span>resultados</span></h1>
              <div class="admin-container08">
                <img
                  alt="profile"
                  src="/playground_assets/status-200h.jpg"
                  class="admin-imagem-relatorio1"
                />
                <form class="admin-form1">
                  <button id="button-relatorio-status" class="admin-button1 button">
                    <span class="admin-text14"><span>CONSULTAR</span></span>
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="admin-cadastros section-container">
      <div class="admin-container09">
        <h1 class="admin-titulo4"><span>CADASTRAR</span></h1>
        <div class="admin-card-relatorio1">
          <div class="admin-cadastro-escuderia">
            <h1 class="admin-titulo5">
              <span></span>
              <span>ESCUDERIA</span>
            </h1>
            <form class="admin-form2">
              <div class="admin-container10">
                <span class="admin-text19"><span>constructor ref</span></span>
                <input
                  type="text"
                  id="constructor_ref_input"
                  name="ConstructorRef"
                  required
                  placeholder="ConstructorRef"
                  class="admin-textinput01 input"
                />
              </div>
              <div class="admin-container11">
                <span class="admin-text21"><span>nome</span></span>
                <input
                  type="text"
                  id="constructor_name_input"
                  name="ConstructorName"
                  required
                  placeholder="Nome"
                  class="admin-textinput02 input"
                />
              </div>
              <div class="admin-container12">
                <span class="admin-text23"><span>Nacionalidade</span></span>
                <input
                  type="text"
                  id="constructor_nationality_input"
                  name="ConstructorNationality"
                  required
                  placeholder="Nacionalidade"
                  class="admin-textinput03 input"
                />
              </div>
              <div class="admin-container13">
                <span class="admin-text25"><span>url</span></span>
                <input
                  type="text"
                  id="constructor_url_input"
                  name="ConstructorURL"
                  required
                  placeholder="URL"
                  class="admin-textinput04 input"
                />
              </div>
              <button class="admin-button2 button">CADASTRAR</button>
            </form>
          </div>
          <div class="admin-cadastro-piloto">
            <h1 class="admin-titulo6"><span>PILOTO</span></h1>
            <form class="admin-form3">
              <div class="admin-container14">
                <span class="admin-text28"><span>driver ref</span></span>
                <input
                  type="text"
                  id="driver_ref_input"
                  name="DriverRef"
                  required
                  placeholder="DriverRef"
                  class="admin-textinput05 input"
                />
              </div>
              <div class="admin-container15">
                <span class="admin-text30"><span>NUMERO</span></span>
                <input
                  type="number"
                  id="driver_number_input"
                  name="DriverNumber"
                  required
                  placeholder="Numero"
                  class="admin-textinput06 input"
                />
              </div>
              <div class="admin-container16">
                <span class="admin-text32"><span>CODIGO</span></span>
                <input
                  type="text"
                  id="driver_code_input"
                  name="DriverCode"
                  enctype="Codigo"
                  required
                  placeholder="Codigo"
                  class="admin-textinput07 input"
                />
              </div>
              <div class="admin-container17">
                <span class="admin-text34">PRIMEIRO NOME</span>
                <input
                  type="text"
                  id="driver_forename_input"
                  name="DriverForename"
                  required
                  placeholder="Primeiro nome"
                  class="admin-textinput08 input"
                />
              </div>
              <div class="admin-container18">
                <span class="admin-text35"><span>SOBRENOME</span></span>
                <input
                  type="text"
                  id="driver_surname_input"
                  name="DriverSurname"
                  required
                  placeholder="Sobrenome"
                  class="admin-textinput09 input"
                />
              </div>
              <div class="admin-container19">
                <span class="admin-text37"><span>DATA DE NASCIMENTO</span></span>
                <input
                  type="date"
                  id="driver_dob_iput"
                  name="DriveDateOfBirth"
                  required
                  class="admin-textinput10 input"
                />
              </div>
              <div class="admin-container20">
                <span class="admin-text39"><span>NACIONALIDADE</span></span>
                <input
                  type="text"
                  id="diver_nationality_input"
                  name="DriverNationality"
                  required
                  placeholder="Nacionalidade"
                  class="admin-textinput11 input"
                />
              </div>
              <button class="admin-button3 button">CADASTRAR</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

window.onload=function() {
  var cookies = parseCookie(document.cookie)
  
  if (window.location.pathname == "/admin/" && cookies.tipo == "admin") {
    axios.get('http://localhost:8080/'+cookies.tipo+'/overview')
      .then(res=>setOverviewAdmin(res.data))
      .catch(err=>console.log(err))
    document.getElementById("button-relatorio-aeroportos").addEventListener("click", getRelatorioAeroportos, false);
    document.getElementById("button-relatorio-status").addEventListener("click", getRelatorioStatus, false);
  }
}

function setOverviewAdmin(data) {
  document.getElementById("overview-pilotos").innerHTML = String(data.Pilotos).concat(" pilotos");
  document.getElementById("overview-escuderias").innerHTML = String(data.Escuderias).concat(" escuderias");
  document.getElementById("overview-corridas").innerHTML = String(data.Corridas).concat(" corridas");
  document.getElementById("overview-temporadas").innerHTML = String(data.Temporadas).concat(" temporadas");
}

function download(content, filename, contentType) {
  if(!contentType) contentType = 'application/octet-stream';
  var a = document.createElement('a');
  var blob = new Blob([content], {'type':contentType});
  a.href = window.URL.createObjectURL(blob);
  a.download = filename;
  a.click();
}

const parseCookie = str =>
  str
  .split(';')
  .map(v => v.split('='))
  .reduce((acc, v) => {
    acc[decodeURIComponent(v[0].trim())] = decodeURIComponent(v[1].trim());
    return acc;
  }, {});

function getRelatorioAeroportos() {
  console.log("get relatorio aeroportos")

  const url='http://localhost:8080/admin/aeroportos-report'
  inputValue = document.getElementById("input-relatorio-aeroportos").value
  
  axios.get(url, {params: {cidade: inputValue} })
    .then(res=>download(res.data, "relatorio-aeroportos-admin.json", "application/json"))
    .catch(err=>console.log(err))
}

function getRelatorioStatus() {
  console.log("get relatorio status")
  const Url='http://localhost:8080/admin/status-report'
  axios.get(Url)
    .then(res=>download(res.data, "relatorio-status-admin.json", "application/json"))
    .catch(err=>console.log(err))
}

export default {
  name: 'Admin',

  data() {
    return {
      rawbcob: ' ',
    }
  },

  metaInfo: {
    title: 'Admin - Sistema F1 WEB',
    meta: [
      {
        property: 'og:title',
        content: 'Admin - Sistema F1 WEB',
      },
    ],
  },
}
</script>

<style scoped>
.admin-container {
  width: 100%;
  display: flex;
  position: relative;
  min-height: 100vh;
  overflow-x: hidden;
  align-items: center;
  flex-direction: column;
  background-size: cover;
  background-image: url("https://images.unsplash.com/photo-1586076100131-32505c71d0d2?ixid=Mnw5MTMyMXwwfDF8c2VhcmNofDEzfHxwYXBlcnxlbnwwfHx8fDE2NTgyODEzNDY&ixlib=rb-1.2.1&w=1500");
}
.admin-header {
  flex: 0 0 auto;
  width: 100%;
  display: flex;
  position: relative;
  align-items: center;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
  background-size: cover;
  background-image: url("https://images.unsplash.com/photo-1558879912-38d7a8936871?ixid=Mnw5MTMyMXwwfDF8c2VhcmNofDEwfHxhc3BoYWx0fGVufDB8fHx8MTY1ODI0NjMxNA&ixlib=rb-1.2.1&w=1500");
}
.admin-container01 {
  flex: 1;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-small);
  display: flex;
  position: relative;
  align-items: flex-start;
  justify-content: flex-start;
}
.admin-f1-logo {
  align-self: flex-start;
  max-height: var(--dl-size-size-maxwidth);
  object-fit: cover;
}
.admin-user-name {
  color: #bf0303;
  width: auto;
  align-self: center;
  text-align: left;
  margin-left: var(--dl-space-space-twounits);
  border-color: #000000;
  border-width: 1px;
  padding-left: var(--dl-space-space-unit);
  border-radius: var(--dl-radius-radius-radius8);
  padding-right: var(--dl-space-space-unit);
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.admin-overview {
  width: var(--dl-size-size-maxwidth);
  display: flex;
  flex-wrap: wrap;
  max-width: auto;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
}
.admin-titulo {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-cards-overview {
  flex: 0 0 auto;
  width: 100%;
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  margin-top: var(--dl-space-space-halfunit);
  align-items: flex-start;
  grid-template-columns: 1fr 1fr 1fr 1fr;
}
.admin-overview-card {
  width: 100%;
  display: flex;
  padding: var(--dl-space-space-unit);
  position: relative;
  max-width: var(--dl-size-size-maxwidth);
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  transition: 0.3s;
  align-items: flex-start;
  border-radius: var(--dl-radius-radius-radius8);
  flex-direction: column;
  justify-content: flex-start;
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.admin-overview-card:hover {
  transform: scale(1.02);
}
.admin-texto {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.admin-icone {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.admin-overview-card1 {
  width: 100%;
  display: flex;
  padding: var(--dl-space-space-unit);
  max-width: var(--dl-size-size-maxwidth);
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  transition: 0.3s;
  align-items: flex-start;
  border-radius: var(--dl-radius-radius-radius8);
  flex-direction: column;
  justify-content: flex-start;
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.admin-overview-card1:hover {
  transform: scale(1.02);
}
.admin-texto1 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.admin-icone1 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.admin-overview-card2 {
  width: 100%;
  display: flex;
  padding: var(--dl-space-space-unit);
  max-width: var(--dl-size-size-maxwidth);
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  transition: 0.3s;
  align-items: flex-start;
  border-radius: var(--dl-radius-radius-radius8);
  flex-direction: column;
  justify-content: flex-start;
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.admin-overview-card2:hover {
  transform: scale(1.02);
}
.admin-texto2 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.admin-icone2 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.admin-overview-card3 {
  width: auto;
  display: flex;
  padding: var(--dl-space-space-unit);
  max-width: var(--dl-size-size-maxwidth);
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  transition: 0.3s;
  align-items: flex-start;
  border-radius: var(--dl-radius-radius-radius8);
  flex-direction: column;
  justify-content: flex-start;
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.admin-overview-card3:hover {
  transform: scale(1.02);
}
.admin-texto3 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.admin-icone3 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.admin-relatorios {
  width: var(--dl-size-size-maxwidth);
  display: flex;
  align-items: center;
  padding-top: 0px;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
  padding-bottom: 0px;
  justify-content: flex-start;
}
.admin-container02 {
  flex-direction: column;
}
.admin-titulo1 {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-container03 {
  flex: 2;
  width: var(--dl-size-size-maxwidth);
  height: auto;
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.admin-card-relatorio {
  width: 100%;
  display: flex;
  max-width: var(--dl-size-size-maxwidth);
  align-self: flex-end;
  background: #fff;
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  align-items: center;
  border-color: #3a0d0d;
  border-width: 4px;
  background-size: cover;
  justify-content: space-between;
  background-image: url("https://images.unsplash.com/photo-1586076100131-32505c71d0d2?ixid=Mnw5MTMyMXwwfDF8c2VhcmNofDEzfHxwYXBlcnxlbnwwfHx8fDE2NTgyODEzNDY&ixlib=rb-1.2.1&w=1200");
}
.admin-container04 {
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  padding-top: var(--dl-space-space-unit);
  padding-left: var(--dl-space-space-unit);
  padding-right: var(--dl-space-space-unit);
  flex-direction: column;
  padding-bottom: var(--dl-space-space-unit);
  justify-content: flex-start;
}
.admin-titulo2 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-container05 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.admin-imagem-relatorio {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.admin-form {
  width: 100%;
  position: relative;
  align-self: stretch;
}
.admin-container06 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text10 {
  color: #000000;
  font-size: 20px;
  font-style: normal;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-button {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.admin-text11 {
  color: #000000;
  font-style: normal;
  font-weight: 700;
}
.admin-container07 {
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.admin-titulo3 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-container08 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.admin-imagem-relatorio1 {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.admin-form1 {
  width: 100%;
  position: relative;
  align-self: center;
}
.admin-button1 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.admin-text14 {
  font-style: normal;
  font-weight: 700;
}
.admin-cadastros {
  width: var(--dl-size-size-maxwidth);
  height: 100%;
  display: flex;
  align-items: stretch;
  padding-top: 0px;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
  padding-bottom: 0px;
  justify-content: flex-start;
}
.admin-container09 {
  flex: 0;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-large);
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.admin-titulo4 {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-card-relatorio1 {
  flex: 1;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-xsmall);
  display: flex;
  position: relative;
  max-width: var(--dl-size-size-maxwidth);
  align-self: flex-start;
  background: #fff;
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  align-items: flex-start;
  padding-top: var(--dl-space-space-halfunit);
  border-color: #3a0d0d;
  border-width: 4px;
  padding-bottom: var(--dl-space-space-halfunit);
  background-size: cover;
  justify-content: flex-start;
  background-image: url('https://images.unsplash.com/photo-1586076100131-32505c71d0d2?ixid=Mnw5MTMyMXwwfDF8c2VhcmNofDEzfHxwYXBlcnxlbnwwfHx8fDE2NTgyODEzNDY&ixlib=rb-1.2.1&w=1200');
}
.admin-cadastro-escuderia {
  flex: 1;
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: stretch;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.admin-titulo5 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-form2 {
  height: auto;
  align-self: center;
  margin-top: var(--dl-space-space-unit);
}
.admin-container10 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text19 {
  color: #000000;
  font-size: 20px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput01 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container11 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text21 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput02 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container12 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text23 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput03 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container13 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text25 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput04 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-button2 {
  width: var(--dl-size-size-xxlarge);
  font-style: normal;
  margin-top: var(--dl-space-space-halfunit);
  font-weight: 700;
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.admin-cadastro-piloto {
  flex: 1;
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: stretch;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.admin-titulo6 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.admin-form3 {
  height: auto;
  align-self: center;
  margin-top: var(--dl-space-space-unit);
}
.admin-container14 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text28 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput05 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container15 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text30 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput06 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container16 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text32 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput07 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container17 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text34 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput08 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container18 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text35 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput09 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container19 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text37 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput10 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-container20 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.admin-text39 {
  color: #000000;
  font-size: 16px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.admin-textinput11 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.admin-button3 {
  width: var(--dl-size-size-xxlarge);
  font-style: normal;
  margin-top: var(--dl-space-space-halfunit);
  font-weight: 700;
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
</style>
