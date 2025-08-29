import { test, expect, describe } from "vitest";
import {EyeIcon} from "~/components/core/icon-svg/eye/eye";

describe('EyeIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(EyeIcon).toBeDefined();
        expect(typeof EyeIcon).toBe('function');
    });
});
