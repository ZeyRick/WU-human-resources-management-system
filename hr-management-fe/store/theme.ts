import { defineStore } from "pinia";

export const useDarkThemeStore = defineStore({
  id: "isDarkTheme",
  state: () => ({
    isDarkTheme: ref(false),
  }),
  actions: {
    setDarkTheme(value: boolean) {
      this.isDarkTheme = value;
    },
  },
});
