import { useAuthStore, useChatStore } from "../../store"
import Account from "./components/Account"
import ListAlgo from "./components/ListAlgo"
import ListHistory from "./components/ListHistory"
import ListMessage from "./components/ListMessage"
import MessageWriter from "./components/MessageWriter"

function Chat() {
  const idUser = useAuthStore((state) => state.id)
  const idHistorySelected = useChatStore((state) => state.idHistorySelected)

  return (
    <div className="w-[100wh] h-[100vh] flex font-sono font-regular overflow-y-hidden">
      <div className="w-1/5 bg-black text-slate-100">
        <ListHistory idUSer={idUser}/>
        <ListAlgo />
        <Account/>
      </div>
      <div className="w-4/5 bg-zinc-900">
        <ListMessage idHistory={idHistorySelected}/>
        <MessageWriter />
      </div>
    </div>
  )
}

export default Chat