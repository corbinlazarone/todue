import { getSession } from "@/utils/session";
import { redirect } from "next/navigation";
import { LogoutButton } from "@/features/auth";
import Image from "next/image";

export default async function DashboardPage() {
  const session = await getSession();

  if (!session.isLoggedIn) {
    redirect("/");
  }

  return (
    <div>
      <div>
        <h1>Welcome, {session.userData.fistName}!</h1>
        <p>{session.userData.email}</p>
        {session.userData.picture && (
          <Image
            alt="Profile Picture"
            src={session.userData.picture}
            loading="eager"
            width={100}
            height={100}
          />
        )}
      </div>
      <div>
        <LogoutButton />
      </div>
    </div>
  );
}
