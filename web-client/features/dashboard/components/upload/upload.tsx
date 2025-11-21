import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";
import { Label } from "@/shared/ui/label";

export function Upload() {
  return (
    <div className="flex w-full max-w-sm items-center gap-2 mt-15">
      <Label htmlFor="syllabus">Syllabus</Label>
      <Input id="syllabus" type="file" />
      <Button>Extract</Button>
    </div>
  );
}
