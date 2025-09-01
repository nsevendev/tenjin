import { test, expect, describe } from "vitest";
import {EnvelopeIcon} from "~/components/core/icon-svg/envelope/envelope";

describe('EnvelopeIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(EnvelopeIcon).toBeDefined();
        expect(typeof EnvelopeIcon).toBe('function');
    });
});
