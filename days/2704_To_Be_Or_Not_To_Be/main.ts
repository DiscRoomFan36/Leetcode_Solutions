
type ToBeOrNotToBe = {
    toBe: (val: any) => boolean;
    notToBe: (val: any) => boolean;
};

function expect(val: any): ToBeOrNotToBe {
    const result: ToBeOrNotToBe = {
        toBe:    x => {
            if (x === val) return true;
            throw new Error("Not Equal");
        },
        notToBe: x => {
            if (x !== val) return true;
            throw new Error("Equal");
        },
    };
    return result;

};

/**
 * expect(5).toBe(5); // true
 * expect(5).notToBe(5); // throws "Equal"
 */