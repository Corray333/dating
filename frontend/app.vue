<script setup>
  import {ref, onMounted} from 'vue'
  import { getCookie, refreshTokens } from '@/utils/helpers';
  




  onMounted(() => {
    if (getCookie('Refresh')) {
      refreshTokens()
    } 
  })
</script>

<template>
  <div class="wrapper w-full hm-screen bg-light flex justify-center">
    <NuxtLayout>
      <NuxtPage/>
    </NuxtLayout>
  </div>
</template>


<style>

.page-enter-active,
.page-leave-active {
  transition: all 0.3s;
}
.page-enter-from,
.page-leave-to {
  transform:scale(0);
}

</style>