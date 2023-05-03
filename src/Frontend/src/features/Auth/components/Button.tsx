import { ButtonHTMLAttributes } from "react"
import { CgSpinner } from "react-icons/cg"
import clsx from "clsx"

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
    label: string
    className?: string,
    loading?: boolean
}

function Button({
    label,
    className,
    type,
    disabled,
    loading,
    ...props
}: ButtonProps,) {
  return (
    <button className={clsx("rounded-2xl min-w-[80px] box-border p-3", className, disabled && "opacity-40", loading && "bg-white" )} type={type} disabled={disabled || loading} {...props}>{
        loading ? <span className="block flex justify-center items-center animate-spin"><CgSpinner size={24}/></span> :
        label
    }</button>
  )
}

export default Button