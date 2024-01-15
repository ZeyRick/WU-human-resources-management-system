import { PersonOutline } from "@vicons/ionicons5";

export type Route = {
    label: string
    key: string
}

export const UserRoute: Route = {
    label: 'clock_management',
    key: 'admin-clock',
}

export const Routes: Route[] = [UserRoute];
