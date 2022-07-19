<template>
  <div class="login-container">
    <main class="login-header">
      <div class="login-container1">
        <img
          alt="image"
          src="/playground_assets/f1-1200h.png"
          loading="eager"
          class="login-f1-logo"
        />
        <span class="login-user-name heading1"><span>sistema f1 web</span></span>
      </div>
    </main>
    <div class="login-body section-container">
      <div class="login-login-form">
        <h1 class="login-titulo"><span>login</span></h1>
        <form class="login-form">
          <div class="login-container2">
            <span class="login-text2"><span>login</span></span>
            <input
              type="text"
              id="login_input"
              name="LoginInput"
              required
              placeholder="login"
              class="login-textinput input"
            />
          </div>
          <div class="login-container3">
            <span class="login-text4">
              <span>senha</span>
              <span class="login-text6"></span>
            </span>
            <input
              type="password"
              id="password_input"
              name="Password"
              required
              placeholder="********"
              class="login-textinput1 input"
            />
          </div>
          <button class="login-button button">LOGIN</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Login',
  metaInfo: {
    title: 'Sistema F1 WEB',
    meta: [
      {
        property: 'og:title',
        content: 'Sistema F1 WEB',
      },
    ],
  },
}
</script>

<style scoped>
.login-container {
  width: 100%;
  display: flex;
  position: relative;
  min-height: 100vh;
  overflow-x: hidden;
  align-items: center;
  flex-direction: column;
  background-size: cover;
  justify-content: flex-start;
  background-image: url("https://images.unsplash.com/photo-1586076100131-32505c71d0d2?ixid=Mnw5MTMyMXwwfDF8c2VhcmNofDEzfHxwYXBlcnxlbnwwfHx8fDE2NTgyODEzNDY&ixlib=rb-1.2.1&w=1500");
}
.login-header {
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
.login-container1 {
  flex: 1;
  width: var(--dl-size-size-maxwidth);
  height: var(--dl-size-size-small);
  display: flex;
  position: relative;
  align-items: flex-start;
  justify-content: flex-start;
}
.login-f1-logo {
  align-self: flex-start;
  max-height: var(--dl-size-size-maxwidth);
  object-fit: cover;
}
.login-user-name {
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
.login-body {
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
.login-login-form {
  flex: 1;
  width: 100%;
  height: var(--dl-size-size-maxwidth);
  display: flex;
  position: relative;
  align-self: stretch;
  align-items: flex-start;
  flex-direction: column;
  justify-content: flex-start;
}
.login-titulo {
  color: #bf0303;
  width: auto;
  font-size: 3rem;
  align-self: center;
  text-align: center;
  border-radius: var(--dl-radius-radius-radius8);
}
.login-form {
  height: auto;
  align-self: center;
  margin-top: var(--dl-space-space-unit);
}
.login-container2 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.login-text2 {
  color: #000000;
  font-size: 20px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.login-textinput {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.login-container3 {
  flex: 0 0 auto;
  width: 200px;
  display: flex;
  align-items: flex-start;
  flex-direction: column;
}
.login-text4 {
  color: #000000;
  font-size: 18px;
  font-family: IBM Plex Sans Condensed;
  font-weight: 400;
  line-height: 1.5;
  text-decoration: none;
}
.login-text6 {
  font-style: normal;
  font-weight: 700;
}
.login-textinput1 {
  width: var(--dl-size-size-xxlarge);
  margin-top: var(--dl-space-space-halfunit);
  margin-bottom: var(--dl-space-space-halfunit);
}
.login-button {
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
