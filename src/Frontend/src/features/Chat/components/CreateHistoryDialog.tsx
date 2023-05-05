import * as Dialog from "@radix-ui/react-dialog"
import { ReactNode } from "react"
import { useForm } from "react-hook-form"
import Button from "../../../component/Button"
import { useMutation } from "@tanstack/react-query"
import { AiOutlineClose } from "react-icons/ai"
import { HistoryRequest, createHistory } from "../api"
import { useAuthStore, useChatAction, useChatStore } from "../../../store"
import Textbox from "./Textbox"
import { History } from "@/models"
import toast from "react-hot-toast"

type CreateHistoryDialogProps = {
  trigger: ReactNode
} 

function CreateHistoryDialog({trigger}: CreateHistoryDialogProps) {
  const idUser = useAuthStore((state) => state.id)
  const histories = useChatStore((state) => state.histories)
  const setHistories = useChatAction().setHistories

  const {
    formState: { isDirty, isSubmitting },
    handleSubmit,
    register,
  } = useForm<HistoryRequest>({
    defaultValues: {
        userId: idUser,
        name: ""
    }
  })

  const mutation = useMutation({
    mutationFn: createHistory
  })

  const onSubmit = async (req: HistoryRequest) => {
    try {
      const res = await mutation.mutateAsync(req)
      const historyTemp: History = {
        name: res.data.name,
        idUser: res.data.userid,
        id: res.data._id
      }
      setHistories([...histories, historyTemp])
      toast.success("Berhasil membuat history")
    } catch (err) {
      toast.error("Gagl membuat history")
    }
  }

  return (
    <Dialog.Root>
      <Dialog.Trigger asChild>
        {trigger}
      </Dialog.Trigger>
      <Dialog.Portal>
        <Dialog.Overlay/>
        <div className="fixed inset-0 z-10 flex items-center bg-white/60 font-sono">
          <Dialog.Content className="w-full flex justify-center">
            <div className="w-full max-w-xl p-6 bg-zinc-900 rounded-2xl flex flex-col items-center">
              <div className="w-full flex justify-between text-white">
                <div className="text-xl font-bold">
                  Create History
                </div>
                <Dialog.Close asChild>
                  <span className="font-bold cursor-pointer">
                    <AiOutlineClose/>
                  </span>
                </Dialog.Close>
              </div>
              
              <form className="w-[400px] max-w-[80%]" onSubmit={handleSubmit(onSubmit)}>
                <Textbox label="id-user" className="my-4 p-4 focus:outline-white hidden" name="userId" register={register} required={true}/>
                <Textbox label="name" className="my-4 focus:outline-white" name="name" register={register} required={true}/>
                <div className="flex justify-center">
                  <Button label="submit" type="submit" className="bg-yellow-200 mx-4 cursor-pointer" disabled={!isDirty} loading={isSubmitting}/>
                </div>
              </form>
            </div>
          </Dialog.Content>
        </div>
      </Dialog.Portal>
    </Dialog.Root>
  )
}

export default CreateHistoryDialog