import {component$} from "@builder.io/qwik";
import {containerVariants, inputFieldVariants, InputPropsType, labelVariants, messageVariants, wrapperVariants} from "~/components/core/input/input-variant";
import {cn} from "~/utils/classe-name/cn";

export const Input = component$<InputPropsType>(({
    variant = 'default',
    size = 'base',
    state = 'default',
    fullWidth = false,
    class: className,
    label,
    helper,
    error,
    required,
    leftIcon,
    rightIcon,
    id,
    ...props
}) => {
    // id stable : si fourni on l'utilise, sinon fallback (ok en SSR)
    const inputId = id ?? `in-${Math.random().toString(36).slice(2, 9)}`;
    const hasLeft = !!leftIcon;
    const hasRight = !!rightIcon;
    const effectiveVariant = state === 'disabled' && variant === 'default' ? 'defaultDisabled' : variant;
    
    return (
        <div class={containerVariants({ fullWidth })}>
            {/* Label */}
            {label && (
                <label
                    for={inputId}
                    class={labelVariants({ state: state === 'disabled' ? 'disabled' : state === 'error' ? 'error' : 'default' })}
                >
                    {label}
                    {required && <span class="text-red-500 ml-1">*</span>}
                </label>
            )}
            
            {/* Wrapper */}
            <div
                class={cn(
                    wrapperVariants({
                        variant: effectiveVariant,
                        size,
                        state,
                        leftIcon: hasLeft,
                        rightIcon: hasRight,
                    }),
                    className
                )}
            >
                {/* Left Icon */}
                {hasLeft && <div class="absolute left-3 flex items-center justify-center text-gray-400 pointer-events-none">{leftIcon}</div>}
                
                {/* Input */}
                <input
                    id={inputId}
                    disabled={state === 'disabled'}
                    class={inputFieldVariants({ size, leftIcon: hasLeft, rightIcon: hasRight })}
                    {...props}
                />
                
                {/* Right Icon */}
                {hasRight && <div class="absolute right-3 flex items-center justify-center text-gray-400">{rightIcon}</div>}
            </div>
            
            {/* Helper / Error */}
            {(helper || error) && (
                <div class={messageVariants({ kind: error ? 'error' : 'helper' })}>
                    {error || helper}
                </div>
            )}
        </div>
    );
});
