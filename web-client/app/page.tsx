import { Landing } from "@/features/landing";
import { getSession } from "@/utils/session";

export default async function Page() {
  const session = await getSession();
  return <Landing isLoggedIn={session.isLoggedIn} />;
}
