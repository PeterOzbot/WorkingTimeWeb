<template>
  <div class="hours-input">
    <div :class="{ 'first-column': after != undefined }">
      <input ref="inputElement" :class="{ 'after-input': after != undefined, 'solo-input': after == undefined }"
        type="number" :value="modelValue" @input="onInput($event)" @focus="selectAll($event)" />
    </div>
    <div class="second-column" v-if="after != undefined">
      <div class="after-content">{{ after }}</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
  modelValue: Number,
  after: String,
})

const inputElement = ref<HTMLInputElement>()
defineExpose({
  focus: () => { inputElement.value?.focus(); }
})


function onInput(event: Event) {
  const htmlInputElement = <HTMLInputElement>event?.target;
  if (htmlInputElement) {
    let newValue = Number(htmlInputElement.value);
    htmlInputElement.value = newValue.toString();
    emit('update:modelValue', newValue)
  }
}

function selectAll(event: FocusEvent) {
  (<HTMLInputElement>event.target)?.select();
}

</script>

<style scoped lang="scss">
.hours-input {
  border-bottom: solid 1px transparent;
  align-content: center;
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
