// https://nuxt.com/docs/api/configuration/nuxt-config


export default defineNuxtConfig({
  build: {
    transpile: ['@vuepic/vue-datepicker']
  },
  devtools: { enabled: true },
  css: ['~/assets/css/main.css'],
  modules: ['nuxt-icon', "@nuxt/image", "@nuxt/ui"],
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
  runtimeConfig: {
    public:{
      backendUrl: 'http://localhost:3001'
    }
  },
})