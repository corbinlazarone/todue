import { Button } from "@/shared/ui/button";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogClose,
} from "@/shared/ui/dialog";
import { DialogDescription } from "@radix-ui/react-dialog";
import { ReactNode } from "react";

export function ConfirmDialog({
  descrip,
  trigger,
  onConfirm,
}: {
  descrip: string;
  trigger: ReactNode;
  onConfirm(e: React.FormEvent): void;
}) {
  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    onConfirm(e);
  }

  return (
    <Dialog>
      <DialogTrigger asChild>{trigger}</DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Confirm</DialogTitle>
          <DialogDescription>{descrip}</DialogDescription>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline" type="button">
                Cancel
              </Button>
            </DialogClose>
            <DialogClose asChild>
              <Button type="submit">Confirm</Button>
            </DialogClose>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
