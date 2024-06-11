import {useMutation, useQueryClient} from "@tanstack/react-query";
import {Todo} from "./TodoList";
import {BASE_URL_V1} from "../App";
import {Badge, Box, Flex, Spinner, Text, useColorModeValue} from "@chakra-ui/react";
import {FaCheckCircle} from "react-icons/fa";
import {MdDelete} from "react-icons/md";

function TodoItem({todo}: { todo: Todo }) {
    const queryClient = useQueryClient();

    const {mutate: updateTodo, isPending: isUpdating} = useMutation({
        mutationKey: ["updateTodo"],
        mutationFn: async () => {
            if (todo.completed) {
                return alert("Todo is already completed!");
            }
            try {
                const response = await fetch(BASE_URL_V1 + `/todo/${todo.id}`, {
                    method: "PATCH",
                })

                const data = await response.json();

                if (!response.ok) {
                    throw new Error(data.error || "Something went wrong!");
                }

                return data;
            } catch (error) {
                throw new Error(`${error}`);
            }
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]});
        },
    });

    const {mutate: deleteTodo, isPending: isDeleting} = useMutation({
        mutationKey: ['deleteTodo'],
        mutationFn: async () => {
            const response = await fetch(BASE_URL_V1 + `/todo/${todo.id}`, {
                method: "DELETE",
            })

            const data = await response.json();

            if (!response.ok) {
                throw new Error(data.error || "Something went wrong!");
            }

            return data;
        },
        onSuccess: () => {
            queryClient.invalidateQueries({queryKey: ["todos"]})
        }
    })

    return (
        <Box
            bg={useColorModeValue("gray.400", "gray.700")}
            p={2}
            borderRadius={"5"}
        >
            <Flex gap={2} alignItems={'center'}>
                <Flex
                    flex={1}
                    alignItems={"center"}
                    border={1}
                    borderColor={"gray.600"}
                    p={2}
                    borderRadius={"lg"}
                    justifyContent={"space-between"}
                >
                    <Text
                        color={todo.completed ? "green.200" : "yellow.100"}
                        decoration={todo.completed ? "line-through" : "none"}
                    >
                        {todo.body}
                    </Text>
                    {
                        todo.completed && (
                            <Badge ml={'1'} colorScheme={'green'}>Done</Badge>
                        )
                    }

                    {
                        !todo.completed && (
                            <Badge ml={'1'} colorScheme={'yellow'}>In progress</Badge>
                        )
                    }
                </Flex>
                <Flex gap={2}>
                    <Box color={'green'} cursor={'pointer'} onClick={() => updateTodo()}>
                        {!isUpdating && <FaCheckCircle size={22}/>}
                        {isUpdating && <Spinner size={"sm"}/>}
                    </Box>
                    <Box color={"red.500"} cursor={"pointer"} onClick={() => deleteTodo()}>
                        {!isDeleting && <MdDelete size={25}/>}
                        {isDeleting && <Spinner size={"sm"}/>}
                    </Box>
                </Flex>
            </Flex>
        </Box>

    );
}

export default TodoItem;
