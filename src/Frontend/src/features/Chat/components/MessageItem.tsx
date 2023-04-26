type MessageItemProps = {
  isBot: boolean
  time: string
  message: string
}

function MessageItem({isBot, time, message}: MessageItemProps) {
  return (
    <div className={`w-[70%] flex ${isBot ? "flex-row" : "flex-row-reverse"}`}>
      <span className="p-2 rounded-lg bg-slate-300">{message}</span>
      <span className="text-xs text-slate-200">{time}</span>
    </div>
  )
}

export default MessageItem