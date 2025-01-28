export function delayedGreeting(name, delay) {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve(`Hello, ${name}!`);
        }, delay);
    });
}