import { ReactNode } from "react";
import cn from 'classnames';
interface SectionProps {
  label?: string;
  className?: string;
  children: ReactNode | Array<ReactNode>
};

function Section({
  label = "",
  className = "",
  children
}: SectionProps) {
  return (
    <div className={cn(
      "relative flex flex-col gap-4 transition-all duration-500 z-1",
      "border border-solid border-zinc-800 focus-within:border-zinc-700 rounded-lg p-8",
      className
    )}>
      <p className="absolute top-0 -translate-y-[50%] px-2 bg-zinc-950 text-[16px] font-bold ">{label}</p>
      {children}
    </div>
  )
}

export default Section;