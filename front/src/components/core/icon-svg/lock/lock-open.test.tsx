import { test, expect, describe } from "vitest";
import {LockOpenIcon} from "~/components/core/icon-svg/lock/lock-open";

describe('LockOpenIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(LockOpenIcon).toBeDefined();
        expect(typeof LockOpenIcon).toBe('function');
    });
});
