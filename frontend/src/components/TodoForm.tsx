import { Button, Flex, Input, Spinner } from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { FormEvent, useState } from "react";
import { IoAdd } from "react-icons/io5";
import { BASE_URL_V1 } from "../App";

function TodoForm() {
  const [newTodo, setNewTodo] = useState("");

  const queryClient = useQueryClient();

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    // Handle fetching APIs
    mutationKey: ["createTodo"],
    mutationFn: async (e: FormEvent) => {
      e.preventDefault();
      try {
        const response = await fetch(BASE_URL_V1 + "/todo", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            body: newTodo,
          }),
        });

        const data = await response.json();

        if (!response.ok) {
          throw new Error(data.error || "Something went wrong!");
        }
        return data;
      } catch (error) {
        throw new Error(`Fetch Todo Error: ${error}`);
      }
    },
    // Handle success and error
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
      setNewTodo("");
    },
    onError: (error: any) => {
      alert(error.message);
    },
  });

  return (
    <>
      <form onSubmit={createTodo}>
        <Flex gap={2}>
          <Input
            type="text"
            value={newTodo}
            onChange={(e) => setNewTodo(e.target.value)}
            ref={(input) => input && input.focus()}
          />
          <Button mx={2} type="submit" _active={{ transform: "scale(.97)" }}>
            {isCreating ? <Spinner size={"xs"} /> : <IoAdd size={20} />}
          </Button>
        </Flex>
      </form>
    </>
  );
}

export default TodoForm;
