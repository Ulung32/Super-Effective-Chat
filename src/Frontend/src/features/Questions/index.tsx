import { FiEdit } from "react-icons/fi"
import { IoTrashBinOutline } from "react-icons/io5"
import AddQuestionDialog from "./components/AddQuestionDialog"
import { Link } from "react-router-dom"

function Questions() {
  return (
    <div className="p-8 bg-zinc-900 min-h-[100vh] font-sono">
      <div className="w-full pt-8 pb-4 text-lg rounded-3xl bg-black text-white">
        <div className="flex justify-evenly items-center bg-indigo-600 rounded-2xl mx-4 py-4 font-bold">
            <div className="text-left grow-[2] flex justify-center">
              question
            </div>
            <div className="text-left grow-[2] flex justify-center">
              answer
            </div>
            <div className="text-right grow">
              <AddQuestionDialog trigger={
                <Link to="/chat">
                  <button className="p-4 rounded-xl bg-indigo-500">
                    Back
                  </button>
                </Link>
              }/>
            </div>
            <div className="pr-4 text-right grow">
              <AddQuestionDialog trigger={
                <button className="p-4 rounded-xl bg-indigo-500">
                  Create
                </button>
              }/>
            </div>
        </div>
        <div className="mt-4">
          <div className="flex">
            <div className="grow-[2] text-center">eirgfhf weiuhdfe</div>
            <div className="grow-[2] text-center">ewuihf weidewio dewoidj</div>
            <div className="grow text-center">
              <span className="block w-9 h-9 rounded-xl hover:bg-yellow-200 hover:text-black flex justify-center items-center">
                <IoTrashBinOutline size={24}/>
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Questions