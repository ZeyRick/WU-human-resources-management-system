import { createI18n } from 'vue-i18n'
import en from '../locales/en.json'

export default defineNuxtPlugin((app) => {
    const i18n = createI18n({
        legacy: false,
        globalInjection: true,
        allowComposition: true,
        locale: 'en',
        messages: {
            en,
        },
    })

    app.vueApp.use(i18n)
})
