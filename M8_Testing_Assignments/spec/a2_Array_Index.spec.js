import { getElement } from '../src/a2_Array_Index.js';
describe("Array Index Handling", function () {
    describe("getElement", function () {
        const testArray = [1, 2, 3, 4, 5, 6, 7];

        it("should return the correct element for a valid index", function () {
            expect(getElement(testArray, 0)).toBe(1);
            expect(getElement(testArray, 2)).toBe(3);
            expect(getElement(testArray, 6)).toBe(7);
        });

        it("should throw error for negative index", function () {
            expect(function () {
                getElement(testArray, -1);
            }).toThrowError("Index out of bounds");
        });

        it("should throw error for index greater than or equal to array length", function () {
            expect(function () {
                getElement(testArray, 8);
            }).toThrowError("Index out of bounds");
            expect(function () {
                getElement(testArray, 10);
            }).toThrowError("Index out of bounds");
        });

        it("should handle empty arrays", function () {
            const emptyArray = [];
            expect(function () {
                getElement(emptyArray, 0);
            }).toThrowError("Index out of bounds");
            expect(function () {
                getElement(emptyArray, -1);
            }).toThrowError("Index out of bounds");
        });
    });
});
