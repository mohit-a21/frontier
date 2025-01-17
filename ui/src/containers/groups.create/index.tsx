import { zodResolver } from "@hookform/resolvers/zod";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";

import {
  Form,
  FormControl,
  FormField,
  FormLabel,
  FormSubmit,
} from "@radix-ui/react-form";
import { Button, Flex, Sheet, Text } from "@raystack/apsara";
import * as z from "zod";

import { useCallback, useState } from "react";
import { FormProvider, useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { update } from "~/api";
import { CustomFieldName } from "~/components/CustomField";
import { SheetFooter } from "~/components/sheet/footer";
import { SheetHeader } from "~/components/sheet/header";
import { Organisation } from "~/types/organisation";
import { fetcher } from "~/utils/helper";

const GroupSchema = z.object({
  name: z
    .string()
    .trim()
    .min(2, { message: "Must be 2 or more characters long" }),
  slug: z
    .string()
    .trim()
    .toLowerCase()
    .min(2, { message: "Must be 2 or more characters long" }),
  orgId: z.string().trim(),
});
export type GroupForm = z.infer<typeof GroupSchema>;

export default function NewGroup() {
  const [organisation, setOrganisation] = useState();
  const navigate = useNavigate();
  const { data, error } = useSWR("/v1beta1/admin/organizations", fetcher);
  const { trigger } = useSWRMutation(
    `/v1beta1/organizations/${organisation}/groups`,
    update,
    {}
  );
  const { organizations = [] } = data || { organizations: [] };

  const methods = useForm<GroupForm>({
    resolver: zodResolver(GroupSchema),
    defaultValues: {},
  });

  const onOpenChange = useCallback(() => {
    navigate("/console/groups");
  }, []);

  const onSubmit = async (data: any) => {
    await trigger(data);
    navigate("/console/groups");
    navigate(0);
  };

  const onChange = (e: any) => {
    setOrganisation(e.target.value);
  };

  return (
    <Sheet open={true}>
      <Sheet.Content
        side="right"
        style={{
          width: "30vw",
          padding: 0,
          borderRadius: "var(--pd-8)",
          boxShadow: "var(--shadow-sm)",
        }}
        close={false}
      >
        <FormProvider {...methods}>
          <Form onSubmit={methods.handleSubmit(onSubmit)}>
            <SheetHeader
              title="Add new group"
              onClick={onOpenChange}
            ></SheetHeader>
            <Flex direction="column" gap="large" style={styles.main}>
              <CustomFieldName
                name="name"
                register={methods.register}
                control={methods.control}
              />
              <CustomFieldName
                name="slug"
                register={methods.register}
                control={methods.control}
              />
              <FormField name="orgId" style={styles.formfield}>
                <Flex
                  style={{
                    marginBottom: "$1",
                    alignItems: "baseline",
                    justifyContent: "space-between",
                  }}
                >
                  <FormLabel
                    style={{
                      fontSize: "11px",
                      color: "#6F6F6F",
                      lineHeight: "16px",
                    }}
                  >
                    Organisation Id
                  </FormLabel>
                </Flex>
                <FormControl asChild>
                  <select
                    {...methods.register("orgId")}
                    style={styles.select}
                    onChange={onChange}
                  >
                    {organizations.map((org: Organisation) => (
                      <option value={org.id} key={org.id}>
                        {org.name}
                      </option>
                    ))}
                  </select>
                </FormControl>
              </FormField>
            </Flex>
            <SheetFooter>
              <FormSubmit asChild>
                <Button variant="primary" style={{ height: "inherit" }}>
                  <Text size={4}>Add group</Text>
                </Button>
              </FormSubmit>
            </SheetFooter>
          </Form>
        </FormProvider>
      </Sheet.Content>
    </Sheet>
  );
}

const styles = {
  main: {
    padding: "32px",
    margin: 0,
    height: "calc(100vh - 125px)",
    overflow: "auto",
  },
  formfield: {
    width: "80%",
    marginBottom: "40px",
  },
  select: {
    height: "32px",
    borderRadius: "8px",
    padding: "8px",
    border: "none",
    backgroundColor: "transparent",
    "&:active,&:focus": {
      border: "none",
      outline: "none",
      boxShadow: "none",
    },
  },
};
