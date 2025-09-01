import {component$} from "@builder.io/qwik";
import {IconProps} from "~/components/core/icon-svg/icon-svg-type";
import {IconSvgBase} from "~/components/core/icon-svg/icon-svg-base";

export const UserIcon = component$<IconProps>((props) => {
    return (
        <IconSvgBase {...props}>
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z"/>
        </IconSvgBase>
    )
})
