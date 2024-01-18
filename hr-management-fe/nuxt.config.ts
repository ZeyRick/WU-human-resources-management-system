import { resolve, dirname } from 'node:path'
import { fileURLToPath } from 'url'
import VueI18nVitePlugin from '@intlify/unplugin-vue-i18n/vite'

export default defineNuxtConfig({
    ssr: false,
    devtools: { enabled: true },
    css: ['~/assets/css/main.css'],
    runtimeConfig: {
        public: {
            apiURL: process.env.API_URL,
        },
    },
    postcss: {
        plugins: {
            tailwindcss: {},
            autoprefixer: {},
        },
    },
    build: {
        transpile: [
            ...(process.env.NODE_ENV === 'production'
                ? ['naive-ui', 'vueuc', '@css-render/vue3-ssr', '@juggle/resize-observer']
                : ['@juggle/resize-observer']),
        ],
    },
    vite: {
        plugins: [
            VueI18nVitePlugin({
                include: [resolve(dirname(fileURLToPath(import.meta.url)), './locales/*.json')],
            }),
        ],
        optimizeDeps: {
            include:
                process.env.NODE_ENV === 'development' ? ['naive-ui', 'vueuc', 'date-fns-tz/formatInTimeZone'] : [],
        },
    },
    modules: ['@pinia/nuxt'],
})
