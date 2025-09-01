import { test, expect, describe } from "vitest";
import {PhoneIcon} from "~/components/core/icon-svg/phone/phone";

describe('PhoneIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(PhoneIcon).toBeDefined();
        expect(typeof PhoneIcon).toBe('function');
    });
});
