import { test, expect, describe } from "vitest";
import {CalendarIcon} from "~/components/core/icon-svg/calendar/calendar";

describe('CalendarIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(CalendarIcon).toBeDefined();
        expect(typeof CalendarIcon).toBe('function');
    });
});
