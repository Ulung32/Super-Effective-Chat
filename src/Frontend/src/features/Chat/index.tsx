import Account from "./components/Account"
import ListAlgo from "./components/ListAlgo"
import ListHistory from "./components/ListHistory"
import ListMessage from "./components/ListMessage"
import MessageWriter from "./components/MessageWriter"

function Chat() {
  return (
    <div className="w-[100wh] h-[100vh] flex font-mplus font-regular overflow-y-hidden">
      <div className="w-1/5 bg-black text-slate-100">
        <ListHistory />
        <ListAlgo />
        <Account/>
      </div>
      <div className="w-4/5 bg-zinc-900">
        <ListMessage />
        <MessageWriter />
      </div>
    </div>
  )
}

export default Chat