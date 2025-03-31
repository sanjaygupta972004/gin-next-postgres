import { ButtonHTMLAttributes, ReactNode } from 'react';
import cn from 'classnames';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  customClass?: string;
  isPrimary?: boolean;
  children?: ReactNode | Array<ReactNode>
}

export function Button({
  children,
  customClass = "",
  isPrimary = true,
  type,
  ...rest
}: ButtonProps) {
  return (
    <button
      type={type}
      className={cn(
        customClass,
        "py-2 px-4 font-semibold rounded-lg cursor-pointer transition-all duration-500",
        {
          "bg-zinc-100 hover:bg-gray-300 text-zinc-950": isPrimary,
          "bg-zinc-950 text-white border border-solid border-zinc-800 hover:border-zinc-500": !isPrimary,
        }
      )}
      {...rest}
    >
      {children}
    </button>
  )
}