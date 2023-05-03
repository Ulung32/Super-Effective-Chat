import { BiPaperPlane } from "react-icons/bi"
import { useState } from "react"
import { useMutation } from "@tanstack/react-query"
import { getAnswer, ChatRequest } from "../api"
import { useChatAction, useChatStore } from "../../../store"

function MessageWriter() {
  const [message, setMessage] = useState<string>("")
  const algo = useChatStore((state) => state.algo)
  const idHistory = useChatStore((state) => state.idHistorySelected)
  const setChats = useChatAction().setChats
  const chats = useChatStore((state) => state.chats)

  const mutation = useMutation({
    mutationFn: getAnswer
  })

  const onSubmit = async () => {
    const data: ChatRequest = {
      query: message,
      method: algo,
      historyId: idHistory
    }
    setChats([...chats, {
      id: "1",
      idHistory: idHistory,
      message: message,
      isBot: false
    }])
    const res = await mutation.mutateAsync(data)
    console.log(res.data);
    setChats([...chats, {
      id: "1",
      idHistory: idHistory,
      message: res.data,
      isBot: true
    }])
  }

  return (
    <div className="w-[80vw] h-24 absolute bottom-0 flex items-center justify-center">
        <div className="w-[65%] p-4 bg-zinc-950 flex justify-between items-center rounded-3xl">
            <input type="text" className="w-[90%] bg-zinc-950 text-white outline-none" onChange={(e) => setMessage(e.target.value)}/>
            <button className="w-8 h-8 rounded-xl bg-yellow-200 flex items-center justify-center" onClick={() => onSubmit()}>
                <BiPaperPlane size="24" color="black"/>
            </button>
        </div>
    </div>
  )
}

export default MessageWriter