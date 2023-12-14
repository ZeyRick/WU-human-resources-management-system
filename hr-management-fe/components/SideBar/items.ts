import { type MenuOption } from "naive-ui";
import {
  BookOutline as BookIcon,
  PersonOutline as PersonIcon,
  WineOutline as WineIcon,
} from "@vicons/ionicons5";
import { UserRoute } from "~/constants/routes";

const Devider: MenuOption = {
  key: "divider-1",
  type: "divider",
};

export const menuOptions: MenuOption[] = [
  Devider,
  {
    label: "Management",
    key: "Dance Dance Dance",
    icon: renderIcon(BookIcon),
    children: [
      {
        label: renderRoute(UserRoute.path, UserRoute.label),
        key: UserRoute.key,
        icon: renderIcon(UserRoute.icon),
      },
      {
        label: "Sheep Man",
        key: "sheep-man",
        icon: renderIcon(PersonIcon),
      },
    ],
  },
];
