<template>
  <div>
    <v-container>
      <v-row v-for="(week, weekIndex) in workingDays" v-bind:key="weekIndex">
        <v-col v-for="(day, dayIndex) in week" v-bind:key="dayIndex" md4>
          <WorkingDayComponent v-bind:workingDay="day" />
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>

<script lang="ts">
import apiClient from "../services/apiClient";
import WorkingDayComponent from "@/components/WorkingDay.vue";
import { Resolve } from "@dvolper/tsdi";
import Vue from "vue";
import { Component } from "vue-property-decorator";
import { NotificationService } from "../services/notificationService";
import WorkingDay from "@/models/workingDay";
import moment from "moment";

@Component({
  components: { WorkingDayComponent }
})
export default class GeneratorView extends Vue {
  @Resolve
  public readonly NotificationService: NotificationService;
  workingDays: WorkingDay[][] = [];

  async created(): Promise<void> {
    const workingDays = await apiClient.generate({ totalHours: 120 });

    let week: WorkingDay[] = [];
    for (const day of workingDays) {
      week.push(day);
      if (moment(day.date).isoWeekday() === 7) {
        this.workingDays.push(week);
        week = [];
      }
    }
    if (week.length > 0) {
      this.workingDays.push(week);
    }
  }
}
</script>

<style></style>
