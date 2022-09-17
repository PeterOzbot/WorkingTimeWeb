<template>
  <v-card class="input-container" elevation="6">
    <div class="hours-container">
      <div class="column">
        <div class="input-label">Total Hours:</div>
      </div>
      <div class="column">
        <InputComponent
          class="hours-input"
          :model.sync="generatorRequest.totalHours"
          after="h"
        />
      </div>
    </div>
    <div class="year-container">
      <div class="column">
        <div class="input-label">Year:</div>
      </div>
      <div class="column">
        <InputComponent
          class="year-input"
          :model.sync="generatorRequest.year"
        />
      </div>
    </div>
    <div class="month-container">
      Months
    </div>
    
  </v-card>
</template>

<script lang="ts">
import Vue from "vue";
import { Component } from "vue-property-decorator";
import InputComponent from "@/components/Controls/Input.vue";
import GeneratorRequest from "@/models/generatorRequest";

@Component({
  components: { InputComponent }
})
export default class GeneratorInputView extends Vue {
  generatorRequest: GeneratorRequest = {
    totalHours: 0,
    month: 0,
    year: new Date().getFullYear()
  };
}
</script>

<style scoped lang="scss">
.input-container {
  position: absolute;
  top: 50%;
  left: 50%;
  margin-right: -50%;
  transform: translate(-50%, -50%);
  width: 50vw;
  height: 50vh;

  display: flex;
  flex-flow: column;

  .column {
    display: flex;
    flex-direction: column;
    flex-basis: 100%;
    flex: 1;
  }

  .hours-container {
    flex: 0 0 auto;

    width: 100%;
    flex-wrap: wrap;
    display: flex;
    flex-flow: row;

    background-color:#67b567;

    .input-label {
      margin: 20px;
      margin-right: 50px;
      font-size: 30px;
      text-align: right;
    }
    .hours-input {
      font-size: 30px;
      margin-left: 60px;
      margin-right: 60px;
    }
  }

  .month-container {
    //background-color: green;
    flex: 1 0 auto;
  }

  .year-container {
    width: 100%;
    flex: 0 0 auto;
    flex-wrap: wrap;

    display: flex;
    flex-flow: row;

    border-bottom: 1px solid #9d9d9d;

    .input-label {
      margin: 20px;
      margin-top: 18px;
      margin-right: 50px;
      font-size: 30px;
      text-align: right;
    }
    .year-input {
      font-size: 30px;
      margin-left: 60px;
      margin-right: 60px;
    }
  }
}
</style>
