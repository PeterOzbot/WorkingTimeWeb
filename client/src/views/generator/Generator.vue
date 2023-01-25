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
import { useRoute } from "vue-router";
import CreateRequest from "@/models/createRequest";
import type EditableWorkingDay from "@/models/editableWorkingDay";
import Group from "@/models/group";

const route = useRoute()

let generatedWorkingDays = ref(new Array<WorkingDay>);


let generatorRequest: GeneratorRequest = {
  groups: [new Group("A", +route.params.a_hours), new Group("B", +route.params.b_hours)],
  month: +route.params.month,
  year: +route.params.year
};

onMounted(async () => {
  generatedWorkingDays.value = await apiClient.generate(generatorRequest);
});

async function download(workingDays: EditableWorkingDay[]) {
  await downloadExcel(workingDays, "A");
  await downloadExcel(workingDays, "B");
}

async function downloadExcel(workingDays: EditableWorkingDay[], groupId: string) {
  const createRequest = new CreateRequest();
  createRequest.days = workingDays.filter(day => day.hours > 0 && day.groupId == groupId);
  createRequest.groupId = groupId;
  const result = await apiClient.create(createRequest);
}

</script>

<style>

</style>
