import type { ColumnDef } from "@tanstack/react-table";
import { createColumnHelper } from "@tanstack/react-table";
import { Link } from "react-router-dom";
import type { User } from "~/types/user";

const columnHelper = createColumnHelper<User>();
export const getColumns: (users: User[]) => ColumnDef<User, any>[] = (
  users: User[]
) => {
  return [
    columnHelper.accessor("id", {
      header: "ID",
      //@ts-ignore
      filterVariant: "text",
      cell: ({ row, getValue }) => {
        return <Link to={`${row.getValue("id")}`}>{getValue()}</Link>;
      },
    }),
    {
      header: "Name",
      accessorKey: "name",
      filterVariant: "text",
      cell: (info) => info.getValue(),
    },
    {
      header: "Email",
      accessorKey: "email",
      filterVariant: "text",
      cell: (info) => info.getValue(),
      footer: (props) => props.column.id,
    },
    {
      header: "Create At",
      accessorKey: "createdAt",
      meta: {
        headerFilter: false,
      },
      cell: (info) =>
        new Date(info.getValue() as Date).toLocaleString("en", {
          month: "long",
          day: "numeric",
          year: "numeric",
        }),

      footer: (props) => props.column.id,
    },
  ];
};