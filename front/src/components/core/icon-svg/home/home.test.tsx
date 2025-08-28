import { test, expect, describe } from "vitest";
import {HomeIcon} from "~/components/core/icon-svg/home/home";

describe('HomeIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(HomeIcon).toBeDefined();
        expect(typeof HomeIcon).toBe('function');
    });
});
