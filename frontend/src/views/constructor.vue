<template>
  <div class="constructor-container">
    <main class="constructor-header">
      <div class="constructor-container01">
        <img
          alt="image"
          src="/playground_assets/f1-1200h.png"
          loading="eager"
          class="constructor-f1-logo"
        />
        <span class="constructor-user-name heading1">
          <span>&#123;constructor&#125;</span>
        </span>
      </div>
    </main>
    <div class="constructor-overview">
      <h1 class="constructor-titulo"><span>overview</span></h1>
      <div class="constructor-cards-overview">
        <div class="constructor-overview-card">
          <h2 class="constructor-texto">
            <span id="overview-pilotos">vitorias</span>
            <br />
            <span>&#123;&#125;</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/trofeu-200h.webp"
            class="constructor-icone"
          />
        </div>
        <div class="constructor-overview-card1">
          <h2 class="constructor-texto1">
            <span id="overview-pilotos">pilotos</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/driver_report-200h.png"
            class="constructor-icone1"
          />
        </div>
        <div class="constructor-overview-card2">
          <h2 class="constructor-texto2">
            <span>primeiro ano:</span>
            <br />
            <span id="overview-primeiro-ano"></span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/temporadas-200h.jpg"
            class="constructor-icone2"
          />
        </div>
        <div class="constructor-overview-card3">
          <h2 class="constructor-texto3">
            <span>primeiro ano:</span>
            <br />
            <span id="overview-ultimo-ano"></span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/temporadas-200h.jpg"
            class="constructor-icone3"
          />
        </div>
      </div>
    </div>
    <div class="constructor-relatorios section-container">
      <div class="constructor-container02 max-content-container">
        <h1 class="constructor-titulo1">relatorios</h1>
        <div class="constructor-container03">
          <div class="constructor-card-relatorio">
            <div class="constructor-container04">
              <h1 class="constructor-titulo2"><span>pilotos</span></h1>
              <div class="constructor-container05">
                <img
                  alt="profile"
                  src="/playground_assets/driver_report-200h.png"
                  class="constructor-imagem-relatorio"
                />
                <form class="constructor-form">
                  <button class="constructor-button button">
                    <span class="constructor-text13"><span>CONSULTAR</span></span>
                  </button>
                </form>
              </div>
            </div>
            <div class="constructor-container06">
              <h1 class="constructor-titulo3">resultados</h1>
              <div class="constructor-container07">
                <img
                  alt="profile"
                  src="/playground_assets/status-200h.jpg"
                  class="constructor-imagem-relatorio1"
                />
                <form class="constructor-form1">
                  <button class="constructor-button1 button">
                    <span class="constructor-text15"><span>CONSULTAR</span></span>
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="constructor-consultas section-container">
      <div class="constructor-container08">
        <h1 class="constructor-titulo4"><span>CONSULTAR</span></h1>
        <div class="constructor-card-relatorio1">
          <div class="constructor-cadastro-piloto">
            <h1 class="constructor-titulo5"><span>PILOTO</span></h1>
            <form class="constructor-form2">
              <div class="constructor-container09">
                <span class="constructor-text19">PRIMEIRO NOME</span>
                <input
                  type="text"
                  id="driver_forename_input"
                  name="DriverForename"
                  required
                  placeholder="Primeiro nome"
                  class="constructor-textinput input"
                />
              </div>
              <button class="constructor-button2 button">CONSULTAR</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

function download(content, filename, contentType) {
  if(!contentType) contentType = 'application/octet-stream';
  var a = document.createElement('a');
  var blob = new Blob([content], {'type':contentType});
  a.href = window.URL.createObjectURL(blob);
  a.download = filename;
  a.click();
}

function parseCookie(str) {
  str
  .split(';')
  .map(v => v.split('='))
  .reduce((acc, v) => {
    acc[decodeURIComponent(v[0].trim())] = decodeURIComponent(v[1].trim());
    return acc;
  }, {})
}

window.onload=function() {
  var cookies = parseCookie(document.cookie)
  
  if (window.location.pathname == "/escuderia/" && cookies.tipo == "escuderia") {
    axios.get('http://localhost:8080/'+cookies.tipo+'/overview')
      .then(res=>setOverviewEscuderia(res.data))
      .catch(err=>console.log(err))
    
    document.getElementById("button-relatorio-escuderia").addEventListener("click", getRelatorioEscuderia, false);
    document.getElementById("button-relatorio-status").addEventListener("click", getRelatorioStatusEscuderia, false);
  }
}

function setOverviewEscuderia(data) {
  document.getElementById("overview-vitorias").innerHTML = String(data.Vitorias) + " vitorias";
  document.getElementById("overview-pilotos").innerHTML = String(data.PilotosUnicos) + " pilotos";
  document.getElementById("overview-primeiro-ano").innerHTML = "primeiro ano: " + String(data.PrimeiroAno);
  document.getElementById("overview-ultimo-ano").innerHTML = "primeiro ano: " + String(data.UltimoAno);
}

export default {
  name: 'Constructor',

  data() {
    return {
      rawcs2h: ' ',
    }
  },

 methods: {
    getRelatorioConstructor: function() {
    },
    getRelatorioStatusConstructor: function() {
    }
  },

  metaInfo: {
    title: 'Constructor - Sistema F1 WEB',
    meta: [
      {
        property: 'og:title',
        content: 'Constructor - Sistema F1 WEB',
      },
    ],
  },
}
</script>

<style scoped>
.constructor-container {
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
.constructor-header {
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
.constructor-container01 {
  flex: 1;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-small);
  display: flex;
  position: relative;
  align-items: flex-start;
  justify-content: flex-start;
}
.constructor-f1-logo {
  align-self: flex-start;
  max-height: var(--dl-size-size-maxwidth);
  object-fit: cover;
}
.constructor-user-name {
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
.constructor-overview {
  width: var(--dl-size-size-maxwidth);
  display: flex;
  flex-wrap: wrap;
  max-width: auto;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
}
.constructor-titulo {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-cards-overview {
  flex: 0 0 auto;
  width: 100%;
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  margin-top: var(--dl-space-space-halfunit);
  align-items: flex-start;
  grid-template-columns: 1fr 1fr 1fr 1fr;
}
.constructor-overview-card {
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
.constructor-overview-card:hover {
  transform: scale(1.02);
}
.constructor-texto {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.constructor-icone {
  height: var(--dl-size-size-medium);
  align-self: flex-start;
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.constructor-overview-card1 {
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
.constructor-overview-card1:hover {
  transform: scale(1.02);
}
.constructor-texto1 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.constructor-icone1 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.constructor-overview-card2 {
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
.constructor-overview-card2:hover {
  transform: scale(1.02);
}
.constructor-texto2 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.constructor-icone2 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.constructor-overview-card3 {
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
.constructor-overview-card3:hover {
  transform: scale(1.02);
}
.constructor-texto3 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.constructor-icone3 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.constructor-relatorios {
  width: var(--dl-size-size-maxwidth);
  display: flex;
  align-items: center;
  padding-top: 0px;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
  padding-bottom: 0px;
  justify-content: flex-start;
}
.constructor-container02 {
  flex-direction: column;
}
.constructor-titulo1 {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-container03 {
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
.constructor-card-relatorio {
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
.constructor-container04 {
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
.constructor-titulo2 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-container05 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.constructor-imagem-relatorio {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.constructor-form {
  width: 100%;
  position: relative;
  align-self: center;
}
.constructor-button {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.constructor-text13 {
  color: #000000;
  font-style: normal;
  font-weight: 700;
}
.constructor-container06 {
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.constructor-titulo3 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-container07 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.constructor-imagem-relatorio1 {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.constructor-form1 {
  width: 100%;
  position: relative;
  align-self: center;
}
.constructor-button1 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.constructor-text15 {
  font-style: normal;
  font-weight: 700;
}
.constructor-consultas {
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
.constructor-container08 {
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
.constructor-titulo4 {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-card-relatorio1 {
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
.constructor-cadastro-piloto {
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
.constructor-titulo5 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.constructor-form2 {
  height: auto;
  align-self: center;
  margin-top: var(--dl-space-space-unit);
}
.constructor-container09 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.constructor-text19 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.constructor-textinput {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.constructor-button2 {
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
