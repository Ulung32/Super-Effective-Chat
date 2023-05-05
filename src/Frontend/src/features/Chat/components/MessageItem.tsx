type MessageItemProps = {
  isBot: boolean
  message: string
}

function MessageItem({isBot, message}: MessageItemProps) {
  return (
    <div className={`w-full py-2 px-12 flex justify-start items-end ${isBot ? "flex-row" : "flex-row-reverse"}`}>
      <span className={`max-w-[70%] p-4 text-white rounded-bl-xl ${ isBot ? "bg-zinc-950 rounded-r-xl" : "bg-purple-400 rounded-t-xl"}`}>{message.split("\n").map(e => 
      <>
        <span>{e}</span>
        <br/>
      </>)}</span>
    </div>
  )
}

export default MessageItem