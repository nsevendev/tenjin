import { test, expect, describe } from "vitest";
import {CogIcon} from "~/components/core/icon-svg/cog/cog";

describe('CogIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(CogIcon).toBeDefined();
        expect(typeof CogIcon).toBe('function');
    });
});
