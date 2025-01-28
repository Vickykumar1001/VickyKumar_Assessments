import { toggleVisibility } from '../src/a5_DOM_Manipulation.js';

describe("DOM Manipulation Tests", () => {
    let mockElement;
    let spyOnDisplay;

    beforeEach(() => {
        let displayValue = "block";
        mockElement = {
            style: {
                get display() {
                    return displayValue;
                },
                set display(value) {
                    displayValue = value;
                },
            },
        };
        spyOnDisplay = spyOnProperty(mockElement.style, "display", "set").and.callThrough();
    });

    it("should set display to 'none' if the element is visible", () => {
        toggleVisibility(mockElement);

        expect(mockElement.style.display).toEqual("none");
        expect(spyOnDisplay).toHaveBeenCalledWith("none");
        expect(spyOnDisplay).toHaveBeenCalledTimes(1);
    });

    it("should set display to 'block' if the element is hidden", () => {
        mockElement.style.display = "none";
        expect(spyOnDisplay).toHaveBeenCalledTimes(1);

        spyOnDisplay.calls.reset();
        toggleVisibility(mockElement);

        expect(mockElement.style.display).toEqual("block");
        expect(spyOnDisplay).toHaveBeenCalledTimes(1);
        expect(spyOnDisplay).toHaveBeenCalledWith("block");
    });

    it("should toggle visibility twice properly", () => {
        toggleVisibility(mockElement);
        expect(mockElement.style.display).toBe("none");
        expect(spyOnDisplay).toHaveBeenCalledTimes(1);
        expect(spyOnDisplay).toHaveBeenCalledWith("none");

        spyOnDisplay.calls.reset();

        toggleVisibility(mockElement);
        expect(mockElement.style.display).toBe("block");
        expect(spyOnDisplay).toHaveBeenCalledTimes(1);
        expect(spyOnDisplay).toHaveBeenCalledWith("block");
    });
});
