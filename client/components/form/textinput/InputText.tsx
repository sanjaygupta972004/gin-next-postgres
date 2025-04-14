import cn from 'classnames';
import { InputHTMLAttributes, ReactNode } from 'react';
import { Control, Controller, FieldValues, Path } from 'react-hook-form';

interface FormInputTextProps<TFieldValues extends FieldValues> extends InputHTMLAttributes<HTMLInputElement> {
  label?: ReactNode | Array<ReactNode>;
  name: Path<TFieldValues>;
  control: Control<TFieldValues>;
  customClass?: string;
}

const FormInputText = <TFieldValues extends Record<string, unknown>>({
  label = '',
  name,
  control,
  customClass,
  ...rest
}: FormInputTextProps<TFieldValues>) => {

  return (
    <Controller
      control={control}
      name={name}
      render={({ field, fieldState: { error } }) => {
        return (
          <div className="flex flex-col flex-1 w-full">
            <p className='flex items-center gap-2 font-bold text-sm ml-1'>{label}</p>
            <input
              {...rest}
              {...field}
              name={name}
              className={cn(
                'mt-2 py-3 px-4',
                'text-sm disabled:text-zinc-500',
                'transition-all duration-700',
                'bg-transparent border border-solid border-zinc-800 outline-none focus-within:border-zinc-500 hover:border-zinc-500 rounded-lg',
                customClass,
                {
                  'shadow-rose-600 shadow-sm border-none': !!error?.message,
                },
              )}
              value={(field.value ?? '') as string}
            />
            {error?.message && <span className="mt-1.5 ml-1 text-rose-600 font-semibold text-[14px] antialiased"> {error?.message}</span>}
          </div>
        );
      }}
    />
  );
};
export default FormInputText;
