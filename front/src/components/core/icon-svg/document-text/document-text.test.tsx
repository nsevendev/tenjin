import { test, expect, describe } from "vitest";
import {DocumentTextIcon} from "~/components/core/icon-svg/document-text/document-text";

describe('DocumentTextIcon', () => {
    test('le composant est défini et exporté', () => {
        expect(DocumentTextIcon).toBeDefined();
        expect(typeof DocumentTextIcon).toBe('function');
    });
});
