import {Stack, Text} from "@chakra-ui/react";
import {useQuery} from "@tanstack/react-query";
import {BASE_URL_V1} from "../App";
import TodoItem from "./TodoItem";


export type Todo = {
    id: number;
    body: string;
    completed: boolean;
};

function TodoList() {
    const {data: todos, isLoading} = useQuery<Todo[]>({
        queryKey: ["todos"],
        queryFn: async () => {
            try {
                const response = await fetch(BASE_URL_V1 + '/todo');
                const data = await response.json();

                if (!response.ok) {
                    throw new Error(data.error || "Something went wrong!");
                }

                return data || [];

            } catch (error) {
                throw new Error(`${error}`);
            }
        },
    });

    return (
        <>
            <Text
                fontSize={"4xl"}
                // textTransform={"uppercase"}
                fontWeight={"bold"}
                textAlign={"center"}
                my={2}
                bgGradient={"linear(to-l, #0b85f8, #00ffff)"}
                bgClip={"text"}
            >
                Today's Tasks
            </Text>

            {!isLoading && todos?.length === 0 && (
                <Stack alignItems={"center"} gap={3}>
                    <Text fontSize={"xl"} color={"gray.500"}>
                        All tasks completed!
                    </Text>
                </Stack>
            )}

            <Stack gap={3}>
                {todos?.map((todo) => (
                    <TodoItem key={todo.id} todo={todo}/>
                ))}
            </Stack>
        </>
    );
}

export default TodoList;
