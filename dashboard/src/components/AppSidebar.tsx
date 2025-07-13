import {
  Activity,
  Calendar,
  ChartNoAxesColumn,
  Code,
  Flag,
  Globe,
  Home,
  Inbox,
  Search,
  Settings,
  Shield,
  Users,
} from "lucide-react";

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar";
import logo from "@/logo-dark.svg";
import { cn } from "@/lib/utils";
const applicationItems = [
  {
    title: "Flags",
    url: "#",
    icon: Flag,
    disabled: false,
  },
  {
    title: "Environments",
    url: "#",
    icon: Globe,
    disabled: false,
  },
  {
    title: "Analytics",
    url: "#",
    icon: ChartNoAxesColumn,
    disabled: true,
  },
  {
    title: "Targeting",
    url: "#",
    icon: Users,
    disabled: true,
  },
];

const managmentItems = [
  {
    title: "Audit Logs",
    url: "#",
    icon: Activity,
    disabled: false,
  },
  {
    title: "API Keys",
    url: "#",
    icon: Code,
    disabled: true,
  },
  {
    title: "Permissions",
    url: "#",
    icon: Shield,
    disabled: true,
  },
  {
    title: "Settings",
    url: "#",
    icon: Settings,
    disabled: true,
  },
];

export function AppSidebar() {
  return (
    <Sidebar>
      <SidebarHeader>
        <div className="flex items-center">
          <img src={logo} alt="flagly logo" className="h-14 w-auto" />
          <h1 className="text-2xl font-semibold text-gray-900">
            Flag<span className="text-gray-400">ly</span>
          </h1>
        </div>
      </SidebarHeader>
      <SidebarContent>
        <SidebarGroup>
          <SidebarGroupLabel>Navigation</SidebarGroupLabel>
          <SidebarGroupContent>
            <SidebarMenu>
              {applicationItems.map((item) => (
                <SidebarMenuItem
                  key={item.title}
                  className={cn(
                    "",
                    item.disabled &&
                      " text-gray-400 opacity-50 pointer-events-none cursor-not-allowed"
                  )}
                >
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
        <SidebarGroup>
          <SidebarGroupLabel>Managment</SidebarGroupLabel>
          <SidebarMenu>
            {managmentItems.map((item) => (
              <SidebarMenuItem
                key={item.title}
                className={cn(
                  "",
                  item.disabled &&
                    " text-gray-400 opacity-50 pointer-events-none cursor-not-allowed"
                )}
              >
                <SidebarMenuButton asChild>
                  <a href={item.url}>
                    <item.icon />
                    <span>{item.title}</span>
                  </a>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>An Orbiq Product.</SidebarFooter>
    </Sidebar>
  );
}
