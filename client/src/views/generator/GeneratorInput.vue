<template>
  <div class="generator-input-container">
    <v-card class="input-container" elevation="6">
      <GroupHours v-model="hourGroups" />
      <div class="year-container">
        <div class="column">
          <div class="input-label">Year:</div>
        </div>
        <div class="column">
          <InputComponent class="year-input" v-model="year" />
        </div>
      </div>
      <MonthSelectorComponent v-model="month" />
    </v-card>

    <v-card class="generate-button" @click="generate()" elevation="6">
      Generate
    </v-card>
  </div>
</template>

<script lang="ts" setup>

import InputComponent from "@/components/controls/HoursInput.vue";
import MonthSelectorComponent from "@/components/input/MonthSelector.vue"
import GroupHours from "@/components/input/GroupHours.vue"
import { useRouter } from "vue-router";
import { ref } from "vue";

let hourGroups = ref(new Array<number>());
let month = ref(0);
let year = ref(new Date().getFullYear());

const router = useRouter();

function generate() {
  router.push({ path: `/generator/${hourGroups.value[0]}/${hourGroups.value[1]}/${month.value}/${year.value}` })
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
