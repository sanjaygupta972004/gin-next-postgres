import cn from 'classnames';
import { InputHTMLAttributes, ReactNode } from 'react';

interface InputTextProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  labelIcon?: ReactNode;
  customClass?: string;
}

const InputText = ({
  label = '',
  labelIcon,
  customClass,
  ...rest
}: InputTextProps) => {

  return (
    <div className="flex flex-col gap-2 flex-1 w-full">
      <p className="flex items-center gap-2 font-semibold">{labelIcon}{label}</p>
      <input
        {...rest}
        className={cn(
          'py-3 px-4',
          'text-sm disabled:text-zinc-500',
          'transition-all duration-700',
          'bg-transparent border border-solid border-zinc-800 outline-none focus-within:border-zinc-500 hover:border-zinc-500 rounded-lg',
          customClass,
        )}
      />
    </div>
  );
};
export default InputText;
