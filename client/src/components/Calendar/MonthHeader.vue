<template>
  <div class="calendar-date-indicator">{{ formatCurrentMonth }}</div>
</template>

<script lang="ts">
import Vue from "vue";
import { Component, Prop } from "vue-property-decorator";
import moment from "moment";

@Component
export default class MonthHeaderComponent extends Vue {
  @Prop()
  public currentMonth: Date;

  get formatCurrentMonth(): string {
    return moment(this.currentMonth).format("MMMM YYYY");
  }
}
</script>

<style scoped>
.calendar-date-indicator {
  font-size: 24px;
}
</style>
