import { test, expect, describe } from "vitest";
import {BellIcon} from "~/components/core/icon-svg/bell/bell";

describe('BellIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(BellIcon).toBeDefined();
        expect(typeof BellIcon).toBe('function');
    });
});
