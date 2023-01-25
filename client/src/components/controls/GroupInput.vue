<template>
    <div class="hours-input">
        <input ref="inputElement" class="solo-input" type="text" :value="modelValue" @input="onInput($event)"
            @focus="selectAll($event)" />
    </div>
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

.hours-input:focus-within {
    border-bottom: solid 1px #5a5aff;
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
