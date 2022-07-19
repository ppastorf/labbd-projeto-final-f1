<template>
  <div class="driver-container">
    <main class="driver-header">
      <div class="driver-container1">
        <img
          alt="image"
          src="/playground_assets/f1-1200h.png"
          loading="eager"
          class="driver-f1-logo"
        />
        <span class="driver-user-name heading1">
          <span>&#123;driver</span>
          <span>&#125;</span>
        </span>
      </div>
    </main>
    <div class="driver-overview">
      <h1 class="driver-titulo"><span>overview</span></h1>
      <div class="driver-cards-overview">
        <div class="driver-overview-card">
          <h2 class="driver-texto">
            <span>
              vitorias:
              <span v-html="rawqwy0"></span>
            </span>
            <br />
            <span>&#123;&#125;</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/trofeu-200h.webp"
            class="driver-icone"
          />
        </div>
        <div class="driver-overview-card1">
          <h2 class="driver-texto1">
            <span>primeiro ano:</span>
            <br />
            <span>&#123;&#125;</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/temporadas-200h.jpg"
            class="driver-icone1"
          />
        </div>
        <div class="driver-overview-card2">
          <h2 class="driver-texto2">
            <span>ultimo ano:</span>
            <br />
            <span>&#123;&#125;</span>
          </h2>
          <img
            alt="image"
            src="/playground_assets/temporadas-200h.jpg"
            class="driver-icone2"
          />
        </div>
      </div>
    </div>
    <div class="driver-relatorios section-container">
      <div class="driver-container2 max-content-container">
        <h1 class="driver-titulo1">relatorios</h1>
        <div class="driver-container3">
          <div class="driver-card-relatorio">
            <div class="driver-container4">
              <h1 class="driver-titulo2"><span>vitorias</span></h1>
              <div class="driver-container5">
                <img
                  alt="profile"
                  src="/playground_assets/trofeu-200h.webp"
                  class="driver-imagem-relatorio"
                />
                <form class="driver-form">
                  <button class="driver-button button">
                    <span class="driver-text13"><span>consultar</span></span>
                  </button>
                </form>
              </div>
            </div>
            <div class="driver-container6">
              <h1 class="driver-titulo3">resultados</h1>
              <div class="driver-container7">
                <img
                  alt="profile"
                  src="/playground_assets/status-200h.jpg"
                  class="driver-imagem-relatorio1"
                />
                <form class="driver-form1">
                  <button class="driver-button1 button">
                    <span class="driver-text15"><span>consultar</span></span>
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Driver',

  data() {
    return {
      rawqwy0: ' ',
    }
  },

  metaInfo: {
    title: 'Driver - Sistema F1 WEB',
    meta: [
      {
        property: 'og:title',
        content: 'Driver - Sistema F1 WEB',
      },
    ],
  },
}
</script>

<style scoped>
.driver-container {
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
.driver-header {
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
.driver-container1 {
  flex: 1;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-small);
  display: flex;
  position: relative;
  align-items: flex-start;
  justify-content: flex-start;
}
.driver-f1-logo {
  align-self: flex-start;
  max-height: var(--dl-size-size-maxwidth);
  object-fit: cover;
}
.driver-user-name {
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
.driver-overview {
  display: flex;
  position: relative;
  flex-wrap: wrap;
  max-width: auto;
  align-self: center;
  align-items: center;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
}
.driver-titulo {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.driver-cards-overview {
  flex: 0 0 auto;
  display: grid;
  grid-gap: var(--dl-space-space-twounits);
  align-self: stretch;
  margin-top: var(--dl-space-space-halfunit);
  align-items: flex-start;
  grid-template-columns: 1fr 1fr 1fr 1fr;
}
.driver-overview-card {
  width: 100%;
  display: flex;
  padding: var(--dl-space-space-unit);
  position: relative;
  max-width: var(--dl-size-size-maxwidth);
  box-shadow: 5px 5px 10px 0px rgba(18, 18, 18, 0.1);
  transition: 0.3s;
  align-items: stretch;
  border-radius: var(--dl-radius-radius-radius8);
  flex-direction: column;
  justify-content: flex-start;
  background-image: linear-gradient(225deg, rgb(189, 195, 199) 3.00%,rgba(0, 0, 0, 0.9) 31.00%);
}
.driver-overview-card:hover {
  transform: scale(1.02);
}
.driver-texto {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.driver-icone {
  height: var(--dl-size-size-medium);
  align-self: flex-start;
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.driver-overview-card1 {
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
.driver-overview-card1:hover {
  transform: scale(1.02);
}
.driver-texto1 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.driver-icone1 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.driver-overview-card2 {
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
.driver-overview-card2:hover {
  transform: scale(1.02);
}
.driver-texto2 {
  color: #cec70c;
  margin-bottom: var(--dl-space-space-twounits);
}
.driver-icone2 {
  width: var(--dl-size-size-large);
  height: var(--dl-size-size-medium);
  object-fit: cover;
  border-radius: var(--dl-radius-radius-radius4);
}
.driver-relatorios {
  width: var(--dl-size-size-maxwidth);
  display: flex;
  align-items: center;
  padding-top: 0px;
  margin-bottom: var(--dl-space-space-unit);
  flex-direction: column;
  padding-bottom: 0px;
  justify-content: flex-start;
}
.driver-container2 {
  flex-direction: column;
}
.driver-titulo1 {
  color: #bf0303;
  width: var(--dl-size-size-maxwidth);
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.driver-container3 {
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
.driver-card-relatorio {
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
.driver-container4 {
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
.driver-titulo2 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.driver-container5 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.driver-imagem-relatorio {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.driver-form {
  width: 100%;
  position: relative;
  align-self: center;
}
.driver-button {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.driver-text13 {
  color: #000000;
  font-style: normal;
  font-weight: 700;
}
.driver-container6 {
  width: 50%;
  height: auto;
  display: flex;
  position: relative;
  align-self: center;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.driver-titulo3 {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.driver-container7 {
  flex: 1;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: flex-start;
}
.driver-imagem-relatorio1 {
  flex: 0;
  width: auto;
  height: auto;
  margin: var(--dl-space-space-unit);
  object-fit: cover;
  flex-shrink: 0;
  border-radius: var(--dl-radius-radius-round);
}
.driver-form1 {
  width: 100%;
  position: relative;
  align-self: center;
}
.driver-button1 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  border-color: #0a0a0a;
  border-width: 2px;
  border-radius: var(--dl-radius-radius-radius4);
  margin-bottom: var(--dl-space-space-halfunit);
  background-color: #e4e406;
}
.driver-text15 {
  font-style: normal;
  font-weight: 700;
}
</style>
