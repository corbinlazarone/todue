import PageHeader from "@/features/dashboard/components/page-header";
import { EditAssignmentDialog } from "@/features/dashboard/components/upload/edit-assignment-dialog";

const assignment = {
  id: 1,
  name: "Programming Assignment 1: Hello World",
  description:
    "Write a simple program that prints 'Hello, World!' in Python. Submit via GitHub.",
  all_day: false,
  due_date: "2025-02-05",
  color: "#3B82F6",
  start_time: "09:00",
  end_time: "17:00",
  reminder: 24,
};

export default function HistoryPage() {
  return (
    <>
      <PageHeader
        title="History"
        // message="You upload history will appear here."
      />
      <EditAssignmentDialog
        AssignmentData={assignment}
        // onConfirm={() => onEdit(assignment.id)}
        // trigger={
        //   <Button
        //     variant="ghost"
        //     size="icon-sm"
        //     className="text-gray-400 hover:text-red-600 h-7 w-7"
        //   >
        //     <Edit2 className="h-3 w-3" />
        //   </Button>
        // }
      />
    </>
  );
}
