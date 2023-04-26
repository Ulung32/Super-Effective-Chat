type MessageItemProps = {
  isBot: boolean
  time: string
  message: string
}

function MessageItem({isBot, time, message}: MessageItemProps) {
  return (
    <div className={`w-full py-2 px-12 flex justify-start items-end ${isBot ? "flex-row" : "flex-row-reverse"}`}>
      <span className={`max-w-[70%] py-2 px-4 text-white rounded-bl-xl ${ isBot ? "bg-zinc-950 rounded-r-xl" : "bg-purple-400 rounded-t-xl"}`}>{message}</span>
      <span className="mx-2 text-xs text-white">{time}</span>
    </div>
  )
}

export default MessageItem