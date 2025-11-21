import { getSession } from "@/utils/session";
import { redirect } from "next/navigation";
import { SidebarInset, SidebarProvider } from "@/shared/ui/sidebar";
import { AppSidebar } from "@/features/dashboard";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const session = await getSession();

  if (!session.isLoggedIn) {
    redirect("/");
  }

  return (
    <SidebarProvider>
      <AppSidebar variant="inset" userData={session.userData} />
      <SidebarInset>{children}</SidebarInset>
    </SidebarProvider>
  );
}
