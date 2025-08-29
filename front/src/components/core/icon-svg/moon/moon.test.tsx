import { expect, test } from "vitest";
import { render, screen } from "@testing-library/react";
import { MoonIcon } from "./moon";

test("MoonIcon renders correctly", () => {
    render(<MoonIcon />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toBeInTheDocument();
});

test("MoonIcon applies custom size", () => {
    render(<MoonIcon size={32} />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toHaveAttribute("width", "32");
    expect(svgElement).toHaveAttribute("height", "32");
});

test("MoonIcon applies custom color", () => {
    render(<MoonIcon color="primary" />);
    const svgElement = screen.getByRole("img", { hidden: true });
    expect(svgElement).toHaveAttribute("stroke", "#0089ff");
});