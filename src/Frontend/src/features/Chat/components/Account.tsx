import DropDownAccount from "./DropDownAccount"

function Account() {
  return (
    <div className="flex h-[10vh] bg-zinc-800 items-center justify-between w-full px-4">
      <div className="flex justify-start items-center">
        <div className="w-8 h-8 bg-black rounded-full"/>
        <span className="ml-3">ernfiuer erkjg</span>
      </div>
      <DropDownAccount trigger={<button>
        ...
      </button>}/>
    </div>
  )
}

export default Account