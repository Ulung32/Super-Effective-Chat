import Account from "./components/Account"
import ListAlgo from "./components/ListAlgo"
import ListHistory from "./components/ListHistory"
import ListMessage from "./components/ListMessage"
import MessageWriter from "./components/MessageWriter"

function Chat() {
  return (
    <div className="w-full h-full flex font-mplus text-lg font-regular">
      <div className="w-1/5 p-8 bg-black text-slate-100">
        <ListHistory />
        <ListAlgo />
        <Account/>
      </div>
      <div className="w-4/5 bg-emerald-300">
        <ListMessage />
        <MessageWriter />
      </div>
    </div>
  )
}

export default Chat