<script setup>

import {useRoute} from 'vue-router'

const route = useRoute()

console.log(route)



</script>

<template>
    <section class="flex w-full relative justify-center p-5">
        <section v-if="route.path=='/login'">
            <slot />
        </section>
        <section v-else class="relative wm w-full">
            <navbar v-if="route.path != '/login'" class="absolute -left-60"/>
            <slot />
        </section>
    </section>
</template>