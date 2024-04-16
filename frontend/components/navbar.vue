<script setup>

</script>

<template>
  <aside class="bg-dark p-5 w-56 items-center rounded-xl text-center flex flex-col gap-2">
    <IconLogo />
    <NuxtLink to="/home" class="button  justify-start"><Icon name="iconamoon:home" class="text-xl"/><p>Home</p></NuxtLink>
    <NuxtLink to="/profile" class="button  justify-start"><Icon name="fa-regular:user" class="text-xl"/><p>Profile</p></NuxtLink>
    <NuxtLink to="/chats" class="button  justify-start"><Icon name="iconamoon:heart" class="text-xl"/><p>Chats</p></NuxtLink>
  </aside>
</template>

<style>

</style>