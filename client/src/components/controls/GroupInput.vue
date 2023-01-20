<template>
    <input ref="inputElement" type="text" :value="modelValue" @input="onInput($event)" />
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const emit = defineEmits(['update:modelValue'])
const props = defineProps({
    modelValue: String,
    after: String,
})

const inputElement = ref<HTMLInputElement>()
defineExpose({
    focus: () => { inputElement.value?.focus(); }
})


function onInput(event: Event) {
    const htmlInputElement = <HTMLInputElement>event?.target;
    if (htmlInputElement) {
        emit('update:modelValue', htmlInputElement.value)
    }
}

</script>

<style scoped lang="scss">
input {
    width: 100%;
    min-width: 0;
}
</style>
