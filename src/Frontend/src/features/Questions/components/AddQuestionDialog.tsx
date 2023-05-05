import * as Dialog from "@radix-ui/react-dialog"
import { ReactNode } from "react"
import { QnARequest, createQnA } from "../api"
import { useForm } from "react-hook-form"
import Button from "../../../component/Button"
import Textbox from "./Textbox"
import { useMutation } from "@tanstack/react-query"
import { AiOutlineClose } from "react-icons/ai"
import { useChatAction, useChatStore } from "../../../store"
import toast from "react-hot-toast"

type AddQuestionDialogProps = {
  trigger: ReactNode
} 

function AddQuestionDialog({trigger}: AddQuestionDialogProps) {
  const qnas = useChatStore((state) => state.qnas)
  const setQnas = useChatAction().setQnas
  const {
    formState: { isDirty, isSubmitting },
    handleSubmit,
    register,
  } = useForm<QnARequest>()

  const mutation = useMutation({
    mutationFn: createQnA
  })

  const onSubmit = async (req: QnARequest) => {
    try {
      const res = await mutation.mutateAsync(req)
      setQnas([...qnas, {
        answer: res.data.Answer,
        question: res.data.Question,
        id: res.data._id
      }])
      toast.success("Berhasil membuat QnA")
    } catch (err) {
      toast.error("Gagl membuat QnA")
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
                  Add QnA
                </div>
                <Dialog.Close asChild>
                  <span className="font-bold cursor-pointer">
                    <AiOutlineClose/>
                  </span>
                </Dialog.Close>
              </div>
              
              <form className="w-[400px] max-w-[80%]" onSubmit={handleSubmit(onSubmit)}>
                <Textbox label="question" className="my-4 p-4 focus:outline-white" name="question" register={register} required={true}/>
                <Textbox label="answer" className="my-4 focus:outline-white" name="answer" register={register} required={true}/>
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

export default AddQuestionDialog