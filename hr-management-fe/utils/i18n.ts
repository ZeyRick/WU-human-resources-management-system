import { createI18n } from 'vue-i18n'
import en from '../locales/en.json'

const i18n = createI18n({
    legacy: false,
    globalInjection: true,
    allowComposition: true,
    locale: 'en',
    messages: {
        en,
    },
})

export default i18n
