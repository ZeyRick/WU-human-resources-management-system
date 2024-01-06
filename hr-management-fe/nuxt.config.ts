export default defineNuxtConfig({
    ssr: false,
    devtools: { enabled: true },
    css: ['~/assets/css/main.css'],
    runtimeConfig: {
        public: {
            baseURL: process.env.BASE_URL,
        },
    },
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
    build: {
        transpile:
            process.env.NODE_ENV === 'production'
                ? ['naive-ui', 'vueuc', '@css-render/vue3-ssr', '@juggle/resize-observer']
                : ['@juggle/resize-observer'],
    },
    vite: {
        optimizeDeps: {
            include:
                process.env.NODE_ENV === 'development' ? ['naive-ui', 'vueuc', 'date-fns-tz/formatInTimeZone'] : [],
        },
    },
    modules: ['@pinia/nuxt'],
})
