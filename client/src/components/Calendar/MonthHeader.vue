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
  public currentMonth: number;
  @Prop()
  public currentYear: number;

  get formatCurrentMonth(): string {
    return moment(new Date(this.currentYear, this.currentMonth, 1)).format(
      "MMMM YYYY"
    );
  }
}
</script>

<style scoped>
.calendar-date-indicator {
  font-size: 28px;
  padding: 10px;
}
</style>
