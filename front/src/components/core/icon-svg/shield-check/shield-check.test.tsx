import { test, expect, describe } from "vitest";
import {ShieldCheckIcon} from "~/components/core/icon-svg/shield-check/shield-check";

describe('ShieldCheckIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(ShieldCheckIcon).toBeDefined();
        expect(typeof ShieldCheckIcon).toBe('function');
    });
});
