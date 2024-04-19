<script setup>

import { ref } from 'vue'
import { getCookie, refreshTokens } from '@/utils/helpers'
import axios from 'axios'

const code = ref(['', '', '', '', '', ''])

const runtimeConfig = useRuntimeConfig()
const router = useRouter()
const checkCode = async (event, i) => {
  code.value[i] = code.value[i].toUpperCase()
  if (code.value[i].length != 1 || !code.value[i].match(/[A-Z0-9]/)) {
    return
  }
  if (i < code.value.length - 1) {
    event.target.nextElementSibling.focus()
  }
  for (let j = 0; j < code.value.length; j++) {
    if (code.value[j] == '') {
      return
    }
  }
  
  try {
    console.log("Auth: ")
    console.log(getCookie("Authorization"))
    await axios.post(runtimeConfig.public.backendUrl + "/users/verify/email", {
      code: code.value.join('')
    },{
      headers:{
        'Authorization': getCookie("Authorization")
      }
    })
    router.push('/home')
  } catch (error) {
    console.log(error)
    if (error?.response.status === 401) {
      await refreshTokens()
      checkCode
      return
    }
  }
}


</script>

<template>
  <section class="w-screen h-screen gradient-background flex justify-center items-center">
    <div class="flex bg-dark rounded-xl w-fit flex-col items-center p-5 text-white gap-2">
      <h1 class="title">Enter verification code:</h1>
      <div class="flex gap-2">
        <input v-for="(v, i) in code" :key="i" v-model="code[i]" @input="checkCode($event, i)" type="text"
          class="p-2 bg-white text-dark rounded-xl font-mono w-16 uppercase text-center text-6xl" maxlength="1">
      </div>
    </div>
  </section>
</template>


<style></style>