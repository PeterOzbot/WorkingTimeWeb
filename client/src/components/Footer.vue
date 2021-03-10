<template>
  <v-snackbar
    v-model="NotificationService.show"
    :timeout="NotificationService.timeout"
  >
    {{ NotificationService.currentMessage }}

    <template v-slot:action="{ attrs }">
      <v-btn
        color="blue"
        text
        v-bind="attrs"
        @click="NotificationService.close()"
      >
        Close
      </v-btn>
    </template>
  </v-snackbar>
</template>

<script lang="ts">
import { Resolve } from "@dvolper/tsdi";
import Vue from "vue";
import { Component } from "vue-property-decorator";
import { NotificationService } from "../services/notificationService";

@Component
export default class FooterComponent extends Vue {
  @Resolve
  public readonly NotificationService: NotificationService;
}
</script>
