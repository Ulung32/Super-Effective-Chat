import { useEffect } from "react"
import { IoTrashBinOutline } from "react-icons/io5"
import AddQuestionDialog from "./components/AddQuestionDialog"
import { Link } from "react-router-dom"
import ConfirmationDialog from "../../component/ConfirmationDialog"
import { deleteQnA, getAllQnA } from "./api"
import { useQuery } from "@tanstack/react-query"

function Questions() {
  const { data } = useQuery({
    queryKey: ["get-all-QnA"],
    queryFn: getAllQnA
  })

  useEffect(() => {
    console.log(data?.data)
  }, [data])
  

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
          {
            data?.data.map((qna: any) => (
              <div className="flex">
                <div className="w-2/5 text-ellipsis text-center">{qna.Question}</div>
                <div className="w-2/5 text-ellipsis text-center">{qna.Answer}</div>
                <div className="w-1/5 text-center">
                  <ConfirmationDialog 
                    id={qna._id}
                    trigger={
                      <span className="block w-9 h-9 rounded-xl hover:bg-yellow-200 hover:text-black flex justify-center items-center cursor-pointer">
                        <IoTrashBinOutline size={24}/>
                      </span>
                    }
                    onConfirm={deleteQnA}
                    label="Apakah anda yakin untuk menghapus QnA ini"
                  />
                </div>
              </div>
            ))
          }
        </div>
      </div>
    </div>
  )
}

export default Questions