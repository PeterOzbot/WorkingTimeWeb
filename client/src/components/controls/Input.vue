<template>
  <div class="hours-input">
    <div :class="{ 'first-column': after != undefined }">
      <input :class="{ 'after-input': after != undefined, 'solo-input': after == undefined }" type="number"
        v-model="value" @keyup="$emit('update:value', value)" @focus="selectAll($event)" />
    </div>
    <div class="second-column" v-if="after != undefined">
      <div class="after-content">{{ after }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>

const props = defineProps({
  value: Number,
  after: String
})

function selectAll(event: FocusEvent) {
  (<HTMLInputElement>event.target)?.select();
}

</script>

<style scoped lang="scss">
.hours-input {
  margin: 20px;
  border-bottom: solid 1px transparent;

  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}

.hours-input input {
  outline-width: 0;
}

.hours-input .solo-input {
  text-align: center;
}

.hours-input .after-input {
  text-align: right;
}

.hours-input:focus-within {
  border-bottom: solid 1px #5a5aff;
}

.first-column {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.second-column {
  display: flex;
  flex-direction: column;
  flex: 1;
}

.hours-input input::-moz-selection {
  background: #94cb94;
}

.hours-input input::selection {
  background: #94cb94;
}

.after-content {
  margin-left: 10px;
}

input {
  width: 100%;
  min-width: 0;
}
</style>
