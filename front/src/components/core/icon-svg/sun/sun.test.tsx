import { expect, test } from "vitest";
import { render, screen } from "@testing-library/react";
import { SunIcon } from "./sun";

test("SunIcon renders correctly", () => {
    render(<SunIcon />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toBeInTheDocument();
});

test("SunIcon applies custom size", () => {
    render(<SunIcon size={32} />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toHaveAttribute("width", "32");
    expect(svgElement).toHaveAttribute("height", "32");
});

test("SunIcon applies custom color", () => {
    render(<SunIcon color="primary" />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toHaveAttribute("stroke", "#0089ff");
});