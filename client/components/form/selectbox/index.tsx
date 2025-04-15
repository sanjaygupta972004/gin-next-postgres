import React, { ReactNode } from 'react';
import { Controller, Control, FieldValues, Path } from 'react-hook-form';
import { SelectHTMLAttributes } from 'react';
import cn from "classnames";

interface FormSelectBoxOption {
  value: string | number;
  label: string;
}

interface FormSelectBoxProps<TFieldValues extends FieldValues> extends SelectHTMLAttributes<HTMLSelectElement> {
  label: ReactNode | Array<ReactNode>;
  options: FormSelectBoxOption[];
  control: Control<TFieldValues>;
  name: Path<TFieldValues>;
}

const FormSelectBox = <TFieldValues extends Record<string, unknown>>({
  label,
  name,
  options,
  control,
  ...rest
}: FormSelectBoxProps<TFieldValues>) => {

  return (
    <Controller
      name={name}
      control={control}
      render={({ field, fieldState: { error } }) => (
        <div className="flex flex-col flex-1 w-full">
          <p className='flex items-center gap-2 font-bold text-sm ml-1'>{label}</p>
          <select
            {...rest}
            {...field}
            value={`${field.value}` || ''}
            id={name}
            className={cn(
              'appearance-none mt-2 py-3 px-4',
              'text-sm disabled:text-zinc-500',
              'transition-all duration-700',
              'bg-transparent border border-solid border-zinc-800 outline-none focus-within:border-zinc-500 hover:border-zinc-500 rounded-lg',
              {
                'shadow-rose-600 shadow-sm border-none': !!error?.message,
              },
            )}
          >
            {options.map((option) => (
              <option className='bg-zinc-900' key={option.value} value={option.value}>
                {option.label}
              </option>
            ))}
          </select>
          {error?.message && <span className="mt-1.5 ml-1 text-rose-600 font-semibold text-[14px] antialiased"> {error?.message}</span>}
        </div>
      )}
    />
  );
};

export default FormSelectBox;