<template>
    <v-dialog v-model="visible" :title="title" :width="width" :show-close="showClose" :fullscreen="fullscreen">
        <slot></slot>
    </v-dialog>
</template>

<script setup>
const props = defineProps({
    title: {
        type: String,
    },
    visible: {
        type: Boolean,
        default: false,
    },
    width: {
        type: String,
        default: "60%",
    },
    showClose: {
        type: Boolean,
        default: false,
    },
    fullscreen: {
        type: Boolean,
        default: false,
    },
})
</script>