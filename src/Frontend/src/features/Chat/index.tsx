import { useAuthStore, useChatAction, useChatStore } from "../../store"
import Account from "./components/Account"
import ListAlgo from "./components/ListAlgo"
import ListHistory from "./components/ListHistory"
import ListMessage from "./components/ListMessage"
import MessageWriter from "./components/MessageWriter"
import Logo from "./components/Logo"
import clsx from "clsx"
import { BiMenu } from "react-icons/bi"

function Chat() {
  const idUser = useAuthStore((state) => state.id)
  const idHistorySelected = useChatStore((state) => state.idHistorySelected)
  const isOpen = useChatStore((state) => state.isOpenSidebar)
  const setIsOPen = useChatAction().setIsOpenSidebar


  return (
    <div className="w-[100wh] h-[100vh] flex font-sono font-regular overflow-y-hidden">
      <div className={clsx("w-64 absolute md:static md:w-1/5 bg-black text-slate-100 z-10 duration-300 overflow-hidden", {"w-0" : !isOpen})}>
        <ListHistory idUSer={idUser}/>
        <ListAlgo />
        <Account/>
      </div>
      <span className={clsx("absolute z-20 md:hidden top-2 text-white left-2", { "left-[16.5rem]" : isOpen })} onClick={() => setIsOPen(!isOpen)}>
        <BiMenu size="48"/>
      </span>
      <div className="w-full md:w-4/5 bg-zinc-900">
        {
          idHistorySelected == "" ? 
          <Logo/>
          :
          <ListMessage idHistory={idHistorySelected}/>
        }
        <MessageWriter />
      </div>
    </div>
  )
}

export default Chat