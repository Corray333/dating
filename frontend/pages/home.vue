<script setup>

import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

// onMounted(() => {
//   if (!document.cookie) {
//     router.push('/login')
//   }
// })

</script>

<template>
  <main class="duration-300 p-5 w-full bg-dark h-fit rounded-xl flex flex-col items-center shadow-lg">
    <section>
      <h1 class="title text-white">Home</h1>
    </section>
  </main>
</template>


<style>

</style>