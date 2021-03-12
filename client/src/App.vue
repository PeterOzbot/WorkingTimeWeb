<template>
  <v-app>
    <router-view></router-view>
    <!-- <FooterComponent></FooterComponent> -->
  </v-app>
</template>

<script>
//import FooterComponent from "./components/Footer.vue";

export default {
  name: "App"
  //components: { FooterComponent }
};
</script>

<style>
/* 
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

body {
  font-family: sans-serif;
  font-weight: 100;
  --grey-100: #e4e9f0;
  --grey-200: #cfd7e3;
  --grey-300: #b5c0cd;
  --grey-800: #3e4e63;
  --grid-gap: 1px;
  --day-label-size: 20px;
} */
</style>
