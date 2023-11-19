import { PersonOutline } from "@vicons/ionicons5";

export type Route = {
  label: string;
  key: string;
  path: string;
  icon: Component;
};

export const UserRoute: Route = {
  label: "User",
  key: "user",
  path: "user",
  icon: PersonOutline,
};

export const Routes: Route[] = [UserRoute];
