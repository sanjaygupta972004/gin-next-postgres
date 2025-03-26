import cn from 'classnames';

interface CustomizedButtonProps {
  label?: string;
  className?: string;
  type?: "primary" | "secondary";
}

export function CustomizedButton({
  label = "",
  className = "",
  type = "primary"
}: CustomizedButtonProps) {
  return (
    <button
      type="button"
      className={cn(
        className,
        "py-3 px-4 font-semibold rounded-lg cursor-pointer transition-all duration-500",
        {
          "bg-zinc-100 hover:bg-zinc-400 text-zinc-950": type === "primary",
          "bg-zinc-950 text-white border border-solid border-zinc-800 hover:border-zinc-500": type === "secondary"
        }
      )}
    >
      {label}
    </button>
  )
}