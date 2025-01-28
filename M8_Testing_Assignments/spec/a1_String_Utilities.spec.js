import { capitalize, reverseString } from '../src/a1_String_Utilities.js';

describe("String Utilities", function () {
    describe("capitalize", function () {
        it("should capitalize first letter of the word", function () {
            expect(capitalize("vicky")).toBe("Vicky");
        });

        it("should return an empty string when input is empty string", function () {
            expect(capitalize("")).toBe("");
        });

        it("should handle single-character words", function () {
            expect(capitalize("v")).toBe("V");
        });

        it("should not change already capitalized words", function () {
            expect(capitalize("Vicky")).toBe("Vicky");
        });

        it("should work with number and symbol at start", function () {
            expect(capitalize("123word")).toBe("123word");
            expect(capitalize("@hello")).toBe("@hello");
        });
    });

    describe("reverseString", function () {
        it("should reverse string", function () {
            expect(reverseString("vicky")).toBe("ykciv");
        });

        it("should return an empty string when input is empty string", function () {
            expect(reverseString("")).toBe("");
        });

        it("should handle palindromes", function () {
            expect(reverseString("madam")).toBe("madam");
        });

        it("should reverse a single-character string", function () {
            expect(reverseString("v")).toBe("v");
        });

        it("should reverse strings with numbers & special characters", function () {
            expect(reverseString("123@abc!")).toBe("!cba@321");
        });
    });
});

