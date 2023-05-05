import { BiPaperPlane } from "react-icons/bi"
import { useForm } from "react-hook-form"
import { useMutation } from "@tanstack/react-query"
import { getAnswer, ChatRequest } from "../api"
import { useChatAction, useChatStore } from "../../../store"
import { CgSpinner } from "react-icons/cg"

type MessageRequest = {
  message: string
}

function MessageWriter() {
  const {
    formState: { isSubmitting },
    handleSubmit,
    register,
    reset
  } = useForm<MessageRequest>({
    defaultValues: {
      message: ""
    }
  })
  const algo = useChatStore((state) => state.algo)
  const idHistory = useChatStore((state) => state.idHistorySelected)
  const setChats = useChatAction().setChats
  const chats = useChatStore((state) => state.chats)

  const mutation = useMutation({
    mutationFn: getAnswer
  })

  const onSubmit = async (input: MessageRequest) => {
    const data: ChatRequest = {
      query: input.message,
      method: algo,
      historyId: idHistory
    }
    setChats([...chats, {
      id: "1",
      idHistory: idHistory,
      message: input.message,
      isBot: false
    }])
    const res = await mutation.mutateAsync(data)
    setChats([...chats, {
      id: "1",
      idHistory: idHistory,
      message: res.data,
      isBot: true
    }])
    reset({message: ""})
  }

  return (
    <div className="w-[80vw] h-24 absolute bottom-0 flex items-center justify-center">
        <div className="w-[65%] p-4 bg-zinc-950 flex justify-between items-center rounded-3xl">
          <form className="w-full flex justify-between" onSubmit={handleSubmit(onSubmit)}>
            <input type="text" className="w-[90%] bg-zinc-950 text-white outline-none" {...register("message")}/>
            <button type="submit" className="w-8 h-8 rounded-xl bg-yellow-200 flex items-center justify-center">
              {
                isSubmitting ? 
                <span className="block animate-spin">
                  <CgSpinner size="24"/>
                </span> 
                :
                <BiPaperPlane size="24" color="black"/>
              }
            </button>
          </form>
        </div>
    </div>
  )
}

export default MessageWriter