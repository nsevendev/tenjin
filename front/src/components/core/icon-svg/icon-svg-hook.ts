import {colorSvg, IconProps, iconSizes, strokeWidths} from "~/components/core/icon-svg/icon-svg-type";

export const getIconSvgProps = ({
    size = iconSizes.lg,
    color = colorSvg.currentColor,
    strokeWidth = strokeWidths.normal
}: Partial<IconProps>) => {
    const sizeValue = typeof size === 'string' && size in iconSizes
        ? iconSizes[size as keyof typeof iconSizes]
        : size;
    
    const strokeValue = typeof strokeWidth === 'string' && strokeWidth in strokeWidths
        ? strokeWidths[strokeWidth as keyof typeof strokeWidths]
        : strokeWidth;
    
    const colorValue = typeof color === 'string' && color in colorSvg
        ? colorSvg[color as keyof typeof colorSvg]
        : color;
    
    return { sizeValue, strokeValue, colorValue };
};
