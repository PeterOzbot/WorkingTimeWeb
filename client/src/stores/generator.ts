import { ref, computed } from "vue";
import { defineStore } from "pinia";
import type Group from "@/models/group";

export const useGeneratorStore = defineStore("generator", {
  state: () => ({
    groups: new Array<Group>()
  }),
  getters: {
    getGroups: (state) => state.groups
  },
  actions: {
    setGroups(newGroups: Array<Group>) {
      this.groups = newGroups;
    }
  }
});
