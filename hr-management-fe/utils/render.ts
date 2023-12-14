import { NIcon } from "naive-ui";
import { RouterLink } from "vue-router";

export function renderIcon(icon: Component) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

export function renderRoute(path: string, label: string) {
  return () =>
    h(
      RouterLink,
      {
        to: {
          name: path || "home",
          params: {
            lang: "en-US",
          },
        },
      },
      { default: () => label || "opsy... 😛" }
    );
}
