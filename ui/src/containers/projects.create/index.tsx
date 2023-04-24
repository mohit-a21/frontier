import { zodResolver } from "@hookform/resolvers/zod";
import useSWR from "swr";
import useSWRMutation from "swr/mutation";

import {
  Button,
  Container,
  Flex,
  Sheet,
  SheetContent,
  Text,
  TextField,
} from "@odpf/apsara";
import {
  Form,
  FormControl,
  FormField,
  FormLabel,
  FormMessage,
  FormSubmit,
} from "@radix-ui/react-form";
import * as z from "zod";

import { useCallback } from "react";
import { FormProvider, useForm, UseFormRegister } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { fetcher, update } from "~/api";
import { SheetFooter } from "~/components/sheet/footer";
import { SheetHeader } from "~/components/sheet/header";
import { Organisation } from "~/types/organisation";
import { capitalizeFirstLetter } from "~/utils/helper";

const ProjectSchema = z.object({
  name: z
    .string()
    .trim()
    .min(3, { message: "Must be 3 or more characters long" }),
  slug: z
    .string()
    .trim()
    .toLowerCase()
    .min(3, { message: "Must be 3 or more characters long" }),
  orgId: z.string().trim(),
});
export type ProjectForm = z.infer<typeof ProjectSchema>;

export default function NewProject() {
  const navigate = useNavigate();
  const { data, error } = useSWR("/admin/v1beta1/organizations", fetcher);
  const { trigger } = useSWRMutation("/admin/v1beta1/projects", update, {});
  const { organizations = [] } = data || { organizations: [] };

  const methods = useForm<ProjectForm>({
    resolver: zodResolver(ProjectSchema),
    defaultValues: {},
  });

  const onOpenChange = useCallback(() => {
    navigate("/projects");
  }, []);

  const onSubmit = async (data: any) => {
    await trigger(data);
    navigate("/projects");
  };

  return (
    <Sheet open={true}>
      <SheetContent
        side="right"
        css={{
          width: "30vw",
          borderRadius: "$3",
          backgroundColor: "$gray1",
          boxShadow: "0px 0px 6px 1px #E2E2E2",
        }}
        close={false}
      >
        <FormProvider {...methods}>
          <Form onSubmit={methods.handleSubmit(onSubmit)}>
            <SheetHeader
              title="Add new project"
              onClick={onOpenChange}
            ></SheetHeader>
            <Container css={styles.main}>
              <CustomFieldName name="name" register={methods.register} />
              <CustomFieldName name="slug" register={methods.register} />
              <FormField name="orgId" style={styles.formfield}>
                <Flex
                  css={{
                    marginBottom: "$1",
                    alignItems: "baseline",
                    justifyContent: "space-between",
                  }}
                >
                  <FormLabel>Organisation Id</FormLabel>
                </Flex>
                <FormControl asChild>
                  <select {...methods.register("orgId")}>
                    {organizations.map((org: Organisation) => (
                      <option value={org.id}>{org.name}</option>
                    ))}
                  </select>
                </FormControl>
              </FormField>
            </Container>
            <SheetFooter>
              <FormSubmit asChild>
                <Button variant="primary" css={{ height: "inherit" }}>
                  <Text
                    css={{
                      fontSize: "14px",
                      color: "white",
                      fontWeight: "normal",
                      lineHeight: "20px",
                      py: "$2",
                    }}
                  >
                    Add project
                  </Text>
                </Button>
              </FormSubmit>
            </SheetFooter>
          </Form>
        </FormProvider>
      </SheetContent>
    </Sheet>
  );
}

type CustomFieldNameProps = {
  name: string;
  register: UseFormRegister<ProjectForm>;
};

const CustomFieldName = ({ name, register }: CustomFieldNameProps) => {
  return (
    <FormField name={name} style={styles.formfield}>
      <Flex
        css={{
          marginBottom: "$1",
          alignItems: "baseline",
          justifyContent: "space-between",
        }}
      >
        <FormLabel
          style={{ fontSize: "11px", color: "#6F6F6F", lineHeight: "16px" }}
        >
          {capitalizeFirstLetter(name)}
        </FormLabel>
        <FormMessage match="valueMissing">Please enter your {name}</FormMessage>
        <FormMessage match="typeMismatch">
          Please provide a valid {name}
        </FormMessage>
      </Flex>
      <FormControl asChild>
        <TextField
          css={{
            height: "32px",
            color: "$grass12",
            borderRadius: "$3",
            padding: "$2",
          }}
          type="name"
          {...register(name as any)}
          required
          placeholder={`Enter your ${name}`}
        />
      </FormControl>
    </FormField>
  );
};

const styles = {
  main: { padding: "32px", width: "80%", margin: 0 },
  formfield: {
    marginBottom: "40px",
  },
};
