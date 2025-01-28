import { delayedGreeting } from '../src/a3_Delayed_Greeting.js'

describe("Delayed Greeting", function () {
    beforeEach(function () {
        jasmine.clock().install();
    });
    afterEach(function () {
        jasmine.clock().uninstall();
    });

    it("should resolve correctly after the specified delay", function (done) {
        const name = "Vicky";
        const delay = 1000;

        const promise = delayedGreeting(name, delay);

        jasmine.clock().tick(delay);

        promise.then((greeting) => {
            expect(greeting).toBe(`Hello, ${name}!`);
            done();
        });
    });

    describe("delayedGreeting", function () {
        it("should not resolve before the specified delay", function () {
            const name = "Vicky";
            const delay = 1000;
            let resolved = false;

            const promise = delayedGreeting(name, delay);
            promise.then(() => {
                resolved = true;
            });

            jasmine.clock().tick(delay - 1);
            expect(resolved).toBe(false);
        });

        it("should handle multiple requests", function (done) {
            const delay1 = 500;
            const delay2 = 1000;

            const promise1 = delayedGreeting("User1", delay1);
            const promise2 = delayedGreeting("User2", delay2);

            promise1.then((greeting) => {
                expect(greeting).toBe("Hello, User1!");
            });

            promise2.then((greeting) => {
                expect(greeting).toBe("Hello, User2!");
            });
            done();
        });

        it("should resolve when delay is zero", function (done) {
            const name = "Vicky";
            const delay = 0;
            const promise = delayedGreeting(name, delay);

            jasmine.clock().tick(0);
            promise.then((greeting) => {
                expect(greeting).toBe(`Hello, ${name}!`);
                done();
            });
        });
    });
});
