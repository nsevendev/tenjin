import { component$ } from "@builder.io/qwik";
import { InputPropsType, inputStyles } from "~/components/core/input/input-variant";
import { cn } from "~/utils/classe-name/cn";

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
    ...props
}) => {
    const inputId = props.id || `input-${Math.random().toString(36).substr(2, 9)}`;
    
    return (
        <div class={cn(
            inputStyles.container.base,
            fullWidth && inputStyles.container.fullWidth
        )}>
            {/* Label */}
            {label && (
                <label
                    for={inputId}
                    class={cn(
                        inputStyles.label.base,
                        required && inputStyles.label.required,
                        state === 'error' && inputStyles.label.error,
                        state === 'disabled' && inputStyles.label.disabled
                    )}
                >
                    {label}
                    {required && <span class="text-red-500 ml-1">*</span>}
                </label>
            )}
            
            {/* Input Wrapper */}
            <div class={cn(
                inputStyles.wrapper.base,
                inputStyles.wrapper.variants[state === "disabled" ? "defaultDisabled" : variant],
                inputStyles.wrapper.sizes[size],
                inputStyles.wrapper.states[state],
                leftIcon && inputStyles.wrapper.withLeftIcon,
                rightIcon && inputStyles.wrapper.withRightIcon,
                className
            )}>
                {/* Left Icon */}
                {leftIcon && (
                    <div class={inputStyles.icon.left}>
                        {leftIcon}
                    </div>
                )}
                
                {/* Input Element */}
                <input
                    {...props}
                    id={inputId}
                    disabled={state === 'disabled'}
                    class={cn(
                        inputStyles.input.base,
                        inputStyles.input.sizes[size],
                        leftIcon && inputStyles.input.withLeftIcon,
                        rightIcon && inputStyles.input.withRightIcon
                    )}
                />
                
                {/* Right Icon */}
                {rightIcon && (
                    <div class={inputStyles.icon.right}>
                        {rightIcon}
                    </div>
                )}
            </div>
            
            {/* Helper Text or Error */}
            {(helper || error) && (
                <div class={cn(
                    inputStyles.message.base,
                    error ? inputStyles.message.error : inputStyles.message.helper
                )}>
                    {error || helper}
                </div>
            )}
        </div>
    );
});
