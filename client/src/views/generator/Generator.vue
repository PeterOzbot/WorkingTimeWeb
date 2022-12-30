<template>
  <CalendarComponent :generatedWorkingDays="generatedWorkingDays" :currentYear="generatorRequest.year"
    :currentMonth="generatorRequest.month - 1" />
</template>

<script lang="ts" setup>

import apiClient from "../../services/apiClient";
import CalendarComponent from "@/components/calendar/Calendar.vue";
import { onMounted, reactive, ref } from "vue";
import type WorkingDay from "@/models/workingDay";
import type GeneratorRequest from "@/models/generatorRequest";
import { useRoute } from "vue-router";

const route = useRoute()

let generatedWorkingDays = ref(new Array<WorkingDay>);
let generatorRequest: GeneratorRequest = {
  totalHours: +route.params.hours,
  month: +route.params.month,
  year: +route.params.year
};

onMounted(async () => {
  generatedWorkingDays.value = await apiClient.generate(generatorRequest);
});

</script>

<style>

</style>
