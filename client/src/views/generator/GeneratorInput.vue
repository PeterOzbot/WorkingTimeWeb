<template>
  <div class="generator-input-container">
    <v-card class="input-container" elevation="6">
      <div class="hours-container">
        <div class="column">
          <div class="input-label">Total Hours:</div>
        </div>
        <div class="column">
          <InputComponent class="hours-input" v-model="generatorRequest.totalHours" after="h" />
        </div>
      </div>
      <div class="year-container">
        <div class="column">
          <div class="input-label">Year:</div>
        </div>
        <div class="column">
          <InputComponent class="year-input" v-model="generatorRequest.year" />
        </div>
      </div>
      <MonthSelectorComponent v-model="generatorRequest.month" />
    </v-card>

    <v-card class="generate-button" @click="generate()" elevation="6">
      Generate
    </v-card>
  </div>
</template>

<script lang="ts" setup>

import InputComponent from "@/components/controls/Input.vue";
import MonthSelectorComponent from "@/components/month-selector/MonthSelector.vue"
import type GeneratorRequest from "@/models/generatorRequest";
import { useRouter } from "vue-router";

let generatorRequest: GeneratorRequest = {
  totalHours: 0,
  month: 0,
  year: new Date().getFullYear()
};

const router = useRouter();

function generate() {
  router.push({ path: `/generator/${generatorRequest.totalHours}/${generatorRequest.month}/${generatorRequest.year}` })
}


</script>

<style scoped lang="scss">
.generator-input-container {
  position: absolute;
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.input-container {
  width: 50vw;

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

    background-color: #67b567;

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

.generate-button {
  margin: 50px;
  padding: 15px;
  background-color: #67b567;
  font-size: 30px;
}
</style>
