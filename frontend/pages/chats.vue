<script setup>


</script>

<template>
  <main class="duration-300 p-5 w-full bg-dark h-fit rounded-xl flex flex-col items-center shadow-lg">

    <section class="text-white w-full">
      <div class="header flex relative items-center w-full justify-center">
        <h1 class="title">Chats</h1>
        <button class="button absolute right-0 w-fit font-bold">
          <p>New chat</p>
          <Icon @click="logout" name="uil:bolt-alt" class=" text-2xl duration-300 cursor-pointer"/>
        </button>
      </div>
    </section>
  </main>
</template>


<style>

</style>