<template>
  <CalendarComponent
    v-bind:generatedWorkingDays="generatedWorkingDays"
    v-bind:currentYear="generatorRequest.year"
    v-bind:currentMonth="generatorRequest.month - 1"
    class="generator-calendar"
  />
</template>

<script lang="ts">
import apiClient from "../../services/apiClient";
import CalendarComponent from "@/components/Calendar/Calendar.vue";
import { Resolve } from "@dvolper/tsdi";
import Vue from "vue";
import { Component } from "vue-property-decorator";
import { NotificationService } from "../../services/notificationService";
import WorkingDay from "@/models/workingDay";
import GeneratorRequest from "@/models/generatorRequest";

@Component({
  components: { CalendarComponent }
})
export default class GeneratorView extends Vue {
  @Resolve
  public readonly NotificationService: NotificationService;

  generatedWorkingDays: WorkingDay[] = [];
  generatorRequest: GeneratorRequest = {
    totalHours: +this.$route.params.hours,
    month: +this.$route.params.month,
    year: +this.$route.params.year
  };

  async created(): Promise<void> {
    this.generatedWorkingDays = await apiClient.generate(this.generatorRequest);
  }
}
</script>

<style></style>
