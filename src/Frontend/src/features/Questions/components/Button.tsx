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
    <button className={clsx("rounded-2xl p-3", className, disabled && "opacity-40", loading && "bg-white" )} type={type} disabled={disabled || loading} {...props}>{
        loading ? <span className="animate-spin"><CgSpinner/></span> :
        label
    }</button>
  )
}

export default Button