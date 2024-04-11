<script setup>
  import {ref, onMounted} from 'vue'
  import { getCookie, refreshTokens } from '@/utils/helpers';
  
  const runtimeConfig = useRuntimeConfig()

  const layout = ref('unauthorized')



  onMounted(() => {
    if (getCookie('Refresh')) {
      refreshTokens()
      layout.value = 'default'
    } else{
      layout.value = 'unauthorized'
    }
  })
</script>

<template>
  <div class="wrapper w-full hm-screen bg-light flex justify-center">
    <NuxtLayout :name="layout">
      <NuxtPage/>
    </NuxtLayout>
  </div>
</template>
