import { test, expect, describe } from "vitest";
import {UserIcon} from "~/components/core/icon-svg/user/user";

describe('UserIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(UserIcon).toBeDefined();
        expect(typeof UserIcon).toBe('function');
    });
});
