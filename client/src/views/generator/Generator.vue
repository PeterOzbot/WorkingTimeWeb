<template>
  <CalendarComponent :generatedWorkingDays="generatedWorkingDays" :currentYear="generatorRequest.year"
    :currentMonth="generatorRequest.month - 1" @download="download" />
</template>

<script lang="ts" setup>

import apiClient from "../../services/apiClient";
import CalendarComponent from "@/components/calendar/Calendar.vue";
import { onMounted, reactive, ref } from "vue";
import type WorkingDay from "@/models/workingDay";
import type GeneratorRequest from "@/models/generatorRequest";
import { useRoute, useRouter } from "vue-router";
import CreateRequest from "@/models/createRequest";
import type EditableWorkingDay from "@/models/editableWorkingDay";
import { useGeneratorStore } from "@/stores/generator"

const route = useRoute();
const router = useRouter();
const generatorStore = useGeneratorStore();

let generatedWorkingDays = ref(new Array<WorkingDay>);

let generatorRequest: GeneratorRequest = {
  groups: [...generatorStore.groups],
  month: +route.params.month,
  year: +route.params.year
};

onMounted(async () => {
  generatedWorkingDays.value = await apiClient.generate(generatorRequest);
  if (generatedWorkingDays.value.length == 0) {
    router.push({ path: `/generator` })
  }
});

async function download(workingDays: EditableWorkingDay[]) {

  for (const group of generatorRequest.groups) {
    const groupWorkingDays = workingDays.filter(day => day.hours > 0 && day.groupId == group.groupId);

    const createRequest = new CreateRequest(group.groupId, groupWorkingDays);
    const result = await apiClient.create(createRequest);
  }
}

</script>

<style>

</style>
