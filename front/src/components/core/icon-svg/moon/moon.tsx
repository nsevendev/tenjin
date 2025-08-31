import {component$} from "@builder.io/qwik";
import {IconProps} from "~/components/core/icon-svg/icon-svg-type";
import {IconSvgBase} from "~/components/core/icon-svg/icon-svg-base";

export const MoonIcon = component$<IconProps>((props) => {
    return (
        <IconSvgBase {...props}>
            <path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/>
        </IconSvgBase>
    )
})