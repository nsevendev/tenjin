import { test, expect, describe } from "vitest";
import {WarningIcon} from "~/components/core/icon-svg/warning/warning";

describe('WarningIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(WarningIcon).toBeDefined();
        expect(typeof WarningIcon).toBe('function');
    });
});
