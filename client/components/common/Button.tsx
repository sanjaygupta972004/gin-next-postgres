import cn from 'classnames';

interface CustomizedButtonProps {
  label?: string;
  className?: string;
  isPrimary?: boolean;
  type?: "submit" | "button";
}

export function CustomizedButton({
  label = "",
  className = "",
  isPrimary = true,
  type,
}: CustomizedButtonProps) {
  return (
    <button
      type={type}
      className={cn(
        className,
        "py-3 px-4 font-semibold rounded-lg cursor-pointer transition-all duration-500",
        {
          "bg-zinc-100 hover:bg-zinc-400 text-zinc-950": isPrimary,
          "bg-zinc-950 text-white border border-solid border-zinc-800 hover:border-zinc-500": !isPrimary,
        }
      )}
    >
      {label}
    </button>
  )
}