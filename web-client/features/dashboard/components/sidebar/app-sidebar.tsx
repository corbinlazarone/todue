"use client";

import Link from "next/link";
import Image from "next/image";
import { usePathname } from "next/navigation";
import { SessionData } from "@/utils/session";
import {
  UploadIcon,
  GalleryHorizontalEnd,
  LayoutDashboard,
} from "lucide-react";
import {
  Sidebar,
  SidebarContent,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuItem,
  SidebarMenuButton,
  SidebarFooter,
  SidebarGroup,
} from "@/shared/ui/sidebar";
import { NavUser } from "./nav-user";
import { useDashboard } from "../../context";

const navItems = [
  {
    title: "Dashboard",
    url: "/dashboard/",
    icon: LayoutDashboard,
  },
  {
    title: "Upload",
    url: "/dashboard/upload",
    icon: UploadIcon,
  },
  {
    title: "History",
    url: "/dashboard/history",
    icon: GalleryHorizontalEnd,
  },
];

interface AppSidebarProps extends React.ComponentProps<typeof Sidebar> {
  userData: SessionData["userData"];
}

export function AppSidebar({ userData, ...props }: AppSidebarProps) {
  const pathname = usePathname();
  const { sidebarDisabled } = useDashboard();

  return (
    <Sidebar
      collapsible="offcanvas"
      className={sidebarDisabled ? "pointer-events-none opacity-50" : ""}
      {...props}
    >
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton asChild size="lg">
              <Link href="/">
                <Image
                  src="/icon.ico"
                  width={20}
                  height={20}
                  className="size-5 group-data-[collapsible=icon]:size-6"
                  alt="Todue"
                />
                <span className="text-base font-semibold">Todue</span>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent className="flex flex-col gap-2">
        <SidebarGroup>
          <SidebarMenu>
            {navItems.map((item) => (
              <SidebarMenuItem key={item.url}>
                <SidebarMenuButton asChild isActive={pathname === item.url}>
                  <Link href={item.url}>
                    <item.icon />
                    <span>{item.title}</span>
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={userData} />
      </SidebarFooter>
    </Sidebar>
  );
}
