import {component$} from "@builder.io/qwik";
import {IconProps} from "~/components/core/icon-svg/icon-svg-type";
import {IconSvgBase} from "~/components/core/icon-svg/icon-svg-base";

export const LockCloseIcon = component$<IconProps>((props) => {
    return (
        <IconSvgBase {...props}>
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M16.5 10.5V6.75a4.5 4.5 0 1 0-9 0v3.75m-.75 11.25h10.5a2.25 2.25 0 0 0 2.25-2.25v-6.75a2.25 2.25 0 0 0-2.25-2.25H6.75a2.25 2.25 0 0 0-2.25 2.25v6.75a2.25 2.25 0 0 0 2.25 2.25Z"/>
        </IconSvgBase>
    )
})
