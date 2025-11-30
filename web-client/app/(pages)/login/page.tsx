import { LoginForm } from "@/features/auth";

export default function LoginPage() {
  return (
    <div className="bg-muted flex min-h-svh flex-col items-center justify-center gap-6 p-6 md:p-10">
      <div className="flex w-full max-w-sm flex-col gap-6">
        <a
          href="/"
          className="flex items-center gap-2 self-center font-medium text-lg"
        >
          <div className="flex size-12 items-center justify-center rounded-lg">
            <img src="../icon.ico" alt="Todue" className="size-8" />
          </div>
          Todue
        </a>
        <LoginForm />
      </div>
    </div>
  );
}
