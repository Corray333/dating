<script setup>
import { ref, onBeforeMount, watch } from 'vue'
import { getCookie, refreshTokens, clearCookies } from '@/utils/helpers';
import { useRouter } from 'vue-router'
import axios from 'axios'

const router = useRouter()

onBeforeMount(() => {
  if (!document.cookie) {
    router.push('/login')
  }
})

const runtimeConfig = useRuntimeConfig()

const user = ref(null)
const changed = ref(false)



const getUser = async () => {
  try {
    const jwt = getCookie('Authorization') // get your JWT
    const payload = jwt.split('.')[1] // get the payload
    const claims = JSON.parse(atob(payload))   // decode and parse the payload
    const { data } = await axios.get(`${runtimeConfig.public.backendUrl}/users/${claims.id}`, {
      headers: {
        Authorization: getCookie('Authorization')
      },
    })
    user.value = data.user
    watch(user, () => {
      changed.value = true
    }, { deep: true })
  } catch (error) {
    console.error(error)
    if (error?.response.status === 401) {
      await refreshTokens()
      getUser()
      return
    }
  }
}

const avatarUrl = ref('')
const file = ref(null)

const handleFileUpload = (event) => {
  console.log('test')
  if (event.target.files[0].size > 500 * 1024) {
    fileMsg.value = "File is too large"
    return
  }
  file.value = event.target.files[0]
  const reader = new FileReader()

  reader.onload = (e) => {
    avatarUrl.value = e.target.result
  }
  reader.readAsDataURL(event.target.files[0])
}

onBeforeMount(() => {
  getUser()
})

const saveChanges = async () => {
  const formData = new FormData()
  if (file.value != null) formData.append('avatar', file.value)
  formData.append('user', JSON.stringify(user.value))

  try {
    const jwt = getCookie('Authorization') 
    const payload = jwt.split('.')[1] 
    const claims = JSON.parse(atob(payload))   
    let url = `${runtimeConfig.public.backendUrl}/users/${claims.id}`
    await axios.put(url, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        'Authorization': getCookie("Authorization")
      },
    })
  } catch (error) {
    console.log(error)
  }
}

const logout = ()=>{
  clearCookies()
  router.push('/login')
}

</script>

<template>
  <main class="duration-300 p-5 w-full bg-dark h-fit rounded-xl flex flex-col items-center shadow-lg">
    <section v-if="user" class="flex w-full flex-col items-center gap-5 text-white">
      <h1 class="title text-white flex items-center gap-5">Profile <Icon @click="logout" name="mdi:logout" class="hover:text-primary duration-300 cursor-pointer"/></h1>
      <div class="content w-full grid gap-5 ">
        <div class="avatar w-48 h-48 overflow-hidden rounded-full relative">
          <input @input="changed = true" type="file" id="fileInput" class="hidden" @change="handleFileUpload" />
          <label for="fileInput"
            class="text-center absolute mx-auto bg-gray-900 bg-opacity-80 h-full w-full rounded-full flex items-center justify-center text-5xl text-green-400 opacity-0 duration-300 cursor-pointer border-green-400 border-8 hover:opacity-100">
            <Icon name="mdi:camera" />
          </label>
          <img :src="file ? avatarUrl : user.avatar" alt="Avatar" class="object-cover  w-full h-full">
        </div>
        <div class=" w-full flex flex-col gap-2">
          <NuxtLink v-if="!user.emailVerified" to="/verify/email" class="button">Verify your email</NuxtLink>
          <div class="flex gap-2">
            <div class="w-full">
              <p>Name:</p>
              <input type="text" class="text-input w-full" v-model="user.name">
            </div>
            <div class="w-full">
              <p>Surname:</p>
              <input type="text" class="text-input w-full" v-model="user.surname">
            </div>
            <div class="w-full">
              <p>Patronymic:</p>
              <input type="text" class="text-input w-full" v-model="user.patronymic">
            </div>
          </div>
          <div class="flex gap-2">
            <div class="w-full">
              <p>Username:</p>
              <input type="text" class="text-input w-full" v-model="user.username" disabled>
            </div>
            <div class="w-full">
              <p>Email:</p>
              <input type="text" class="text-input w-full" v-model="user.email" disabled>
            </div>
          </div>
          <div class="flex gap-2">
            <div class="w-full">
              <p>Phone:</p>
              <input type="text" class="text-input w-full" v-model="user.phone">
            </div>
            <div class="w-full">
              <p>City:</p>
              <UInput icon="i-heroicons-magnifying-glass-20-solid" size="lg" color="white" placeholder="Search..."
                model-value="" v-model="user.city" class="border-none" />
            </div>
          </div>
          <div class="flex gap-2">
            <div class="w-full">
              <p>Bio:</p>
              <textarea type="text" class="text-input w-full" v-model="user.bio"></textarea>
            </div>
          </div>
          <button @click="saveChanges" class="button w-fit px-5" :class="changed ? '' : 'disabled'">Save</button>
        </div>
      </div>
    </section>
  </main>
</template>


<style scoped>
/* .avatar{
  width: 200px;
  height: 200px;
  margin: 0;
  padding: 0;
} */

.content {
  grid-template-columns: 12rem 1fr;
}

textarea {
  min-height: 150px;
  max-height: 300px;
}
</style>