import { test, expect, describe } from "vitest";
import {AcademicIcon} from "~/components/core/icon-svg/academic/academic";

describe('AcademicIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(AcademicIcon).toBeDefined();
        expect(typeof AcademicIcon).toBe('function');
    });
});
