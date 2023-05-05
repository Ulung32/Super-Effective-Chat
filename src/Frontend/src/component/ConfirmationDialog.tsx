import * as AlertDialog from "@radix-ui/react-alert-dialog"
import { ReactNode, useState } from "react"
import Button from "./Button"
import toast from "react-hot-toast"

type ConfirmationDialogProps = {
  id: string
  trigger: ReactNode
  label: string
  onConfirm: (id: string) => Promise<unknown>
}

function ConfirmationDialog({trigger, label, onConfirm, id}: ConfirmationDialogProps) {
  const [open, setOpen] = useState(false)

  return (
    <AlertDialog.Root open={open} onOpenChange={setOpen}>
        <AlertDialog.Trigger asChild>{trigger}</AlertDialog.Trigger>
        <AlertDialog.Portal>
          <div className="fixed inset-0 z-10 flex items-center bg-white/60 font-sono">
            <AlertDialog.Content className="w-full flex justify-center">
              <div className="w-full max-w-lg p-6 bg-zinc-900 rounded-2xl flex flex-col items-center">
                <div className="w-full text-white">
                  <div className="text-lg font-bold mb-4">
                    {label}
                  </div>
                  <div className="w-full flex justify-center">
                    <AlertDialog.Action asChild>
                      <Button
                        label="Delete"
                        className="bg-red-700 mr-4"
                        onClick={() =>
                          onConfirm(id).then(() => {
                            setOpen(false)
                            toast.success("Berhasil menghapus data")
                          }).catch(() => {
                            toast.error("Gagal menghapus data")
                          })
                        }
                      />
                    </AlertDialog.Action>
                    <Button
                      label="Cancel"
                      onClick={() => setOpen(false)}
                    />
                  </div>
                </div>
              </div>
            </AlertDialog.Content>
          </div>
        </AlertDialog.Portal>
      </AlertDialog.Root>
  )
}

export default ConfirmationDialog