<template>
  <v-card
    v-bind:color="getColor()"
    class="working-day-card"
    :class="{ 'working-day-off': !workingDay.isWorking }"
  >
    <v-card-subtitle>
      <v-container>
        <v-row>
          {{ formatDate() }}
        </v-row>
        <v-row>
          {{ getDay() }}
        </v-row>
      </v-container>
    </v-card-subtitle>
    <v-card-text>
      {{ workingDay.hours }}
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import moment from "moment";
import WorkingDay from "../models/workingDay";
import Vue from "vue";
import { Component, Prop } from "vue-property-decorator";

@Component
export default class WorkingDayComponent extends Vue {
  @Prop()
  public workingDay: WorkingDay;

  public formatDate(): string {
    return moment(this.workingDay.date).format("DD.MM.YYYY");
  }
  public getDay() {
    switch (moment(this.workingDay.date).isoWeekday()) {
      case 1:
        return "Monday";
      case 2:
        return "Tuesday";
      case 3:
        return "Wednesday";
      case 4:
        return "Thursday";
      case 5:
        return "Friday";
      case 6:
        return "Saturday";
      case 7:
        return "Sunday";
      default:
        return "??";
    }
  }

  public getColor(): string {
    switch (moment(this.workingDay.date).isoWeekday()) {
      case 6:
        return "#FFF2F2";
      case 7:
        return "#FFE5E5";
      default:
        return "white";
    }
  }
}
</script>

<style>
.working-day-card {
  /* min-width: 200px; */
}
.working-day-off {
  opacity: 0.5;
}
</style>
