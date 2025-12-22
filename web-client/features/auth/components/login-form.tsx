"use client";

import { cn } from "@/lib/utils";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/shared/ui/card";
import { Field, FieldDescription, FieldGroup } from "@/shared/ui/field";
import { GoogleLogin } from "@react-oauth/google";
import { login } from "../api/actions";
import { useRouter } from "next/navigation";
import { toast } from "sonner";

export function LoginForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const router = useRouter();

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card>
        <CardHeader className="text-center">
          <CardTitle className="text-xl">Welcome back</CardTitle>
          <CardDescription>Login with your Google account</CardDescription>
        </CardHeader>
        <CardContent>
          <form>
            <FieldGroup>
              <Field>
                <GoogleLogin
                  onSuccess={async (credentialResponse) => {
                    const jwt = credentialResponse.credential;
                    try {
                      await login(jwt);
                      router.push("/dashboard");
                    } catch {
                      toast.error("Failed to login. Try again or contact us.");
                    }
                  }}
                  onError={() => {
                    toast.error("Failed to login. Try again or contact us.");
                  }}
                  theme="outline"
                  size="large"
                />
              </Field>
            </FieldGroup>
          </form>
        </CardContent>
      </Card>
      <FieldDescription className="px-6 text-center">
        By clicking creating an account, you agree to our{" "}
        <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.
      </FieldDescription>
    </div>
  );
}
