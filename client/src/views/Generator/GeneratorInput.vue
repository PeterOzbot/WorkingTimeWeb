<template>
  <v-card class="input-container" elevation="6">
    <v-card-title>Inputs</v-card-title>
    <v-card-text>
      month:
    </v-card-text>
    <v-card-actions>
      Save
    </v-card-actions>
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "vue-property-decorator";

@Component({
  components: {}
})
export default class GeneratorInputView extends Vue {}
</script>

<style>
.input-container {
  position: absolute;
  top: 50%;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, -50%);
  width: 50vw;
  height: 50vh;
}
</style>
