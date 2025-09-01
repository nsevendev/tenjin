import { test, expect, describe } from "vitest";
import {ChartBarIcon} from "~/components/core/icon-svg/chart-bar/chart-bar";

describe('ChartBarIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(ChartBarIcon).toBeDefined();
        expect(typeof ChartBarIcon).toBe('function');
    });
});
