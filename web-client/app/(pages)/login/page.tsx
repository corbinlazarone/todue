import Link from "next/link";
import Image from "next/image";
import { LoginForm } from "@/features/auth";

export default function LoginPage() {
  return (
    <div className="bg-muted flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
      <div className="flex w-full max-w-sm flex-col gap-6">
        <Link
          href="/"
          className="flex items-center gap-2 self-center font-medium text-lg"
        >
          <div className="flex size-12 items-center justify-center rounded-lg">
            <Image src="/icon.ico" alt="Todue" width={32} height={32} className="size-8" />
          </div>
          Todue
        </Link>
        <LoginForm />
      </div>
    </div>
  );
}
