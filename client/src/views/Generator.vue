<template>
  <div>
    <v-container grid-list-md fluid class="grey lighten-4">
      <v-row>
        <v-col
          v-for="(workingDay, index) in workingDays"
          v-bind:key="index"
          md4
        >
          <WorkingDayComponent v-bind:workingDay="workingDay" />
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
import { Component, Prop } from "vue-property-decorator";
import { NotificationService } from "../services/notificationService";

@Component({
  components: { WorkingDayComponent }
})
export default class GeneratorView extends Vue {
  @Resolve
  public readonly NotificationService: NotificationService;
  @Prop()
  public workingDays: [];

  async created(): Promise<void> {
    this.workingDays = await apiClient.generate({ totalHours: 120 });
  }
}
</script>

<style></style>
